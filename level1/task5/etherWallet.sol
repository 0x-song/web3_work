// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract EtherWallet{

    address payable public immutable  owner;

    event Log(string operation, address from, uint256 value, bytes data);

    constructor(){
        owner = payable(msg.sender);
    }

    receive() external payable {
        emit Log("receive", msg.sender, msg.value, "");
    }

    function withdraw1(uint256 amount) external {
        require(msg.sender == owner, "You are not owner!");
        require(amount <= address(this).balance, "balance is not enough");
        payable(msg.sender).transfer(amount);
    }

    function withdraw2(uint256 amount) external {
        require(msg.sender == owner, "You are not owner!");
        require(address(this).balance >= amount,"you have no balance");
        bool result = payable(msg.sender).send(amount);
        require(result, "Send Failed");
    }
    
    function withdraw3(uint256 amount) external {
        require(msg.sender == owner, "you are not owner");
        require(address(this).balance >= amount,"you have no balance");
        (bool result, ) = msg.sender.call{value : amount}("");
        require(result, "Send Failed");
    }

    function getBalance() external view returns (uint256){
        return address(this).balance;
    }
}
