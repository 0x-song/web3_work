// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.24;

contract Exceed{

    uint8 public number;

    constructor(){
        number = 0;
    }

    function reduce() public {
        number = number - 1;
    }
    
}