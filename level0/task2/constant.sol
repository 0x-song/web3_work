// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract Constant{
    address public constant owner = address(1);

    //string和bytes只可以使用constant
    string public constant str = "hello,world";

}

contract Immutable{
    address public immutable owner;

//immutable既可以在声明时初始化，也可以在构造函数中初始化
    uint256 public immutable number = 1024;

    constructor(){
        owner = msg.sender;
    }
}