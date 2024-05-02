// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IBlackStoneEvent {
    event Drawn(uint256 indexed tokenId, uint256 drawId, uint256 drawIndex);
    event DrawPropertySet(uint256 drawId, uint256[] probability);
}
