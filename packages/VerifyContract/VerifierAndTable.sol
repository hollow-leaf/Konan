// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.22;

import "./Blobstream.sol";
import "./lib/tree/binary/BinaryMerkleProof.sol";
import "./lib/tree/namespace/NamespaceMerkleMultiproof.sol";
import "./lib/tree/Types.sol";
import "openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol";
import "./IDAOracle.sol";

import "./Constants.sol";
import "./DataRootTuple.sol";
import "./lib/verifier/DAVerifier.sol";

struct VerifyResult {
    uint256 blockHeight;
    bool verified;
}

contract VerifierAndTable {
    address bridge;
    mapping(bytes28 => VerifyResult) public VerifyTable;

    constructor(address _bridge) {
        bridge = _bridge;
    }

    function IsNamespaceVerified(bytes28 namespace) public view returns (bool) {
        VerifyResult memory res = VerifyTable[namespace];
        return res.verified;
    }

    function VerifySharesToDataRootTupleRoot(SharesProof memory sharesProof) public returns (bool, DAVerifier.ErrorCodes){

        Namespace memory _namespace;
        _namespace = sharesProof.namespace;

        bytes[] memory _data;
        _data = sharesProof.data;

        bytes32 _datahash = keccak256(abi.encode(_data));


        bool res;
        DAVerifier.ErrorCodes err;

        (res, err) = DAVerifier.verifySharesToDataRootTupleRoot(Blobstream(bridge), sharesProof);

        if(res) {
            VerifyTable[_namespace.id] = VerifyResult(sharesProof.attestationProof.tuple.height, res);
        }

        return (res, err);
    }

    function bytes32ToBytes(bytes32 data) public pure returns (bytes28) {
        bytes28 result;
        bytes28 x = bytes28(0xff000000000000000000000000000000000000000000000000000000);
        uint i = 0;
        while (i < 28) {
            result ^= (x & data[i]) >> 8*i;
            ++i;
        }
        return result;
    }

}