// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./ERC721A.sol";
import {IVRFCoordinatorV2Plus} from "@chainlink/contracts/src/v0.8/vrf/dev/interfaces/IVRFCoordinatorV2Plus.sol";
import {VRFConsumerBaseV2Plus} from "@chainlink/contracts/src/v0.8/vrf/dev/VRFConsumerBaseV2Plus.sol";
import {VRFV2PlusClient} from "@chainlink/contracts/src/v0.8/vrf/dev/libraries/VRFV2PlusClient.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import {IBlackStoneEvent} from "./interfaces/events/IBlackStoneEvent.sol";

contract BlackStone is ERC721A, ReentrancyGuard, VRFConsumerBaseV2Plus, IBlackStoneEvent {
    uint256 public immutable maxPerAddressDuringMint;
    uint256 public immutable amountForDevs;
    uint256 public immutable totalTypes;
    uint256[] public drawTypes;
    struct RequestStatus {
        bool exists; // whether a requestId exists
        bool fulfilled; // whether the request has been successfully fulfilled
        uint256 drawId;
        uint256 rank;
    }
    mapping(uint256 => RequestStatus) public s_requests; 
    IVRFCoordinatorV2Plus COORDINATOR;
    uint256 public lastRequestId;
    uint256 s_subscriptionId;
    uint32 callbackGasLimit = 100000;
    uint16 requestConfirmations = 3;
    uint32 numWords = 1;
    bytes32 public keyHash = 0x1770bdc7eec7771f7ba4ffd640f34260d7f095b79c92d34a5b2551d6f6cfd2be;
    address public vrfCoordinator = 0x5CE8D5A2BC84beb22a398CCA51996F7930313D61;
    struct SaleConfig {
        uint32 publicSaleStartTime;
        uint64 publicPrice;
        uint32 publicSaleKey;
    }

    struct SpecialEvent {
        bool exists;
        uint256 startTimestamp;
        uint256 endTimestamp;
        uint256[] probability;
    }

    SaleConfig public saleConfig;

    // lotteryDraw probability
    mapping(uint256 => uint256[]) public drawProbability;
    // special probability
    mapping(uint256 => SpecialEvent) public specialProbability;
    // tokenId to drawId to drawIndex
    mapping(uint256 => mapping(uint256 => uint256)) public drawStatus;
    // requestId to tokenId to drawId
    mapping(uint256 => uint256) public requestToDrawId;

    constructor(
        uint256 subscriptionId,
        uint256 maxBatchSize_,
        uint256[] memory drawtypes_,
        uint256 collectionSize_,
        uint256 amountForDevs_
    ) ERC721A("BlackStone", "BROCK", maxBatchSize_, collectionSize_) VRFConsumerBaseV2Plus(vrfCoordinator){
        totalTypes = drawtypes_.length;
        drawTypes = drawtypes_;
        s_subscriptionId = subscriptionId;
        COORDINATOR = IVRFCoordinatorV2Plus(vrfCoordinator);
        maxPerAddressDuringMint = maxBatchSize_;
        amountForDevs = amountForDevs_;
    }

    modifier callerIsUser() {
        require(tx.origin == msg.sender, "The caller is another contract");
        _;
    }

    function setDrawProperty(
        uint256 drawId,
        uint256[] memory probability
    ) external onlyOwner {
        require(
            drawTypes[drawId] == probability.length,
            "Draw property length does not match draw type"
        );
        uint256 checkSum = 0;
        for (uint i = 0; i < probability.length; i++) {
            checkSum += probability[i];
        }
        require(
            checkSum == 100,
            "total probability should be 100"
        );
        drawProbability[drawId] = probability;
        emit DrawPropertySet(drawId, probability);
    }

    function createSpecialEvent(
        uint256 startTimestamp,
        uint256 endTimestamp,
        uint256 drawId,
        uint256[] memory probability
    ) external onlyOwner {
        require(
            drawTypes[drawId] == probability.length,
            "Draw property length does not match draw type"
        );
        uint256 checkSum = 0;
        for (uint i = 0; i < probability.length; i++) {
            checkSum += probability[i];
        }
        require(
            checkSum == 100,
            "total probability should be 100"
        );
        specialProbability[drawId]= SpecialEvent({
            exists: true,
            startTimestamp: startTimestamp,
            endTimestamp: endTimestamp,
            probability: probability
        });
        emit SpecialPropertySet(drawId, probability, startTimestamp, endTimestamp);
    }

    //input your tokenId
    function drawByVRF(
        uint256 _tokenId,
        uint256 _drawId
    ) external nonReentrant {
        require(
            _drawId == 0 || drawStatus[_tokenId][0] != 0,
            "Draw not started"
        );
        require(
            drawStatus[_tokenId][_drawId] == 0,
            "Draw already completed"
        );
        require(
            _drawId < totalTypes,
            "Draw id out of range"
        );
        require(
            drawProbability[_drawId].length > 0,
            "Draw property not set"
        );
        require(
            ownerOf(_tokenId) == msg.sender,
            "not Owner"
        );
        uint256 requestId = COORDINATOR.requestRandomWords(
            VRFV2PlusClient.RandomWordsRequest({
                keyHash: keyHash,
                subId: s_subscriptionId,
                requestConfirmations: requestConfirmations,
                callbackGasLimit: callbackGasLimit,
                numWords: numWords,
                extraArgs: VRFV2PlusClient._argsToBytes(
                    VRFV2PlusClient.ExtraArgsV1({nativePayment: false})
                )
            })
        );

        s_requests[requestId] = RequestStatus({
            exists: true,
            fulfilled: false,
            drawId: _drawId,
            rank: uint256(0)
        });

        lastRequestId = requestId;
        drawStatus[_tokenId][_drawId] = requestId;
    }

    function fulfillRandomWords(
        uint256 _requestId,
        uint256[] memory _randomWords
    ) internal override {
        require(s_requests[_requestId].exists, "request not found");
        s_requests[_requestId].fulfilled = true;
        uint256 drawNumber = _randomWords[0] % 100;
        uint256 drawId = s_requests[_requestId].drawId;
        uint256[] memory probability;
        if (specialProbability[drawId].exists && block.timestamp >= specialProbability[drawId].startTimestamp && block.timestamp <= specialProbability[drawId].endTimestamp) {
            probability = specialProbability[drawId].probability;
        } else {
            probability = drawProbability[drawId];
        }
        for (uint i = 0; i < probability.length; i++) {
            if( drawNumber >= probability[i] ){
                drawNumber -= probability[i];
            }else {
                s_requests[_requestId].rank = i+1;
            }
        }
    }

    function tokenURI(
        uint256 tokenId
    ) public view virtual override returns (string memory) {
        require(
            _exists(tokenId),
            "ERC721Metadata: URI query for nonexistent token"
        );
        string memory baseURI = _baseURI();
        uint256 requestId = drawStatus[tokenId][0];
        uint256 number = s_requests[requestId].rank;
        if ( number == 0 ) {
            return 
            bytes(baseURI).length > 0
                ? string(abi.encodePacked(baseURI,"0.json"))
                : "";
        } else {
            number -= 1;
        }
        for (uint i = 1; i < drawTypes.length; i++) {
            uint256 num = drawTypes[i] + 1;
            uint256 requestIndex = drawStatus[tokenId][i];
            if (s_requests[requestIndex].fulfilled == true) {
                number *= num;
                number += s_requests[requestIndex].rank;
            } else {
                number *= num;
            }
        }
        number += 1;
        return
            bytes(baseURI).length > 0
                ? string(abi.encodePacked(baseURI, Strings.toString(number),".json"))
                : "";
    }

    function publicSaleMint(
        uint256 quantity,
        uint256 callerPublicSaleKey
    ) external payable callerIsUser {
        SaleConfig memory config = saleConfig;
        uint256 publicSaleKey = uint256(config.publicSaleKey);
        uint256 publicPrice = uint256(config.publicPrice);
        uint256 publicSaleStartTime = uint256(config.publicSaleStartTime);
        require(
            publicSaleKey == callerPublicSaleKey,
            "called with incorrect public sale key"
        );
        require(
            isPublicSaleOn(publicPrice, publicSaleKey, publicSaleStartTime),
            "public sale has not begun yet"
        );
        require(
            totalSupply() + quantity <= collectionSize,
            "reached max supply"
        );
        require(
            numberMinted(msg.sender) + quantity <= maxPerAddressDuringMint,
            "can not mint this many"
        );
        _safeMint(msg.sender, quantity);
        refundIfOver(publicPrice * quantity);
    }

    function refundIfOver(uint256 price) private {
        require(msg.value >= price, "Need to send more ETH.");
        if (msg.value > price) {
            payable(msg.sender).transfer(msg.value - price);
        }
    }

    function isPublicSaleOn(
        uint256 publicPriceWei,
        uint256 publicSaleKey,
        uint256 publicSaleStartTime
    ) public view returns (bool) {
        return
            publicPriceWei != 0 &&
            publicSaleKey != 0 &&
            block.timestamp >= publicSaleStartTime;
    }

    function setPublicSaleKey(uint32 key) external onlyOwner {
        saleConfig.publicSaleKey = key;
    }

    // For marketing etc.
    function devMint(uint256 quantity) external onlyOwner {
        require(
            totalSupply() + quantity <= amountForDevs,
            "too many already minted before dev mint"
        );
        require(
            quantity % maxBatchSize == 0,
            "can only mint a multiple of the maxBatchSize"
        );
        uint256 numChunks = quantity / maxBatchSize;
        for (uint256 i = 0; i < numChunks; i++) {
            _safeMint(msg.sender, maxBatchSize);
        }
    }

    // // metadata URI
    string private _baseTokenURI;

    function _baseURI() internal view virtual override returns (string memory) {
        return _baseTokenURI;
    }

    function setBaseURI(string calldata baseURI) external onlyOwner {
        _baseTokenURI = baseURI;
    }

    function withdrawMoney() external onlyOwner nonReentrant {
        (bool success, ) = msg.sender.call{value: address(this).balance}("");
        require(success, "Transfer failed.");
    }

    function setOwnersExplicit(
        uint256 quantity
    ) external onlyOwner nonReentrant {
        _setOwnersExplicit(quantity);
    }

    function numberMinted(address owner) public view returns (uint256) {
        return _numberMinted(owner);
    }

    function getOwnershipData(
        uint256 tokenId
    ) external view returns (TokenOwnership memory) {
        return ownershipOf(tokenId);
    }
}
