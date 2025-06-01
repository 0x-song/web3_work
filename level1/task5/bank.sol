// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract Bank{

    address public immutable owner;

    event Deposit(address _addr, uint256 _amount);

    event Withdraw(uint256 _amount);

    receive() external payable{
        emit Deposit(msg.sender,  msg.value);
    }

    constructor(){
        owner = msg.sender;
    }

    function withdraw() external {
        require(msg.sender == owner, "not owner");
        emit Withdraw(address(this).balance);
        selfdestruct(payable(msg.sender));
    }

    function getBalance() external view returns (uint256){
        return address(this).balance;
    }
}