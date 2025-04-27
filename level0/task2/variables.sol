// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract Variables {
    // 状态变量存储在链上.
    string public text = "Hello";
    uint256 public num = 123;

    function doSomething() public {
        // 局部变量不会存储到链上
        uint256 i = 456;

        // 全局变量，可以获取区块链的信息
        uint256 timestamp = block.timestamp; // Current block timestamp
        address sender = msg.sender; // address of the caller
    }
}