// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract CrowFunding {
    address public immutable beneficiary;

    uint256 public immutable fundingGoal;

    //当前已经募集的金额
    uint256 public fundingAmount;

    mapping(address => uint256) public funders;

    mapping(address => bool) private fundersContributed;

    address[] public fundersList;

    bool public available = true;

    constructor(address _beneficiary, uint256 _fundingGoal){
        beneficiary = _beneficiary;
        fundingGoal = _fundingGoal;
    }
    
    function contribute() external payable{
        require(available, "crowFunding is closed");
        require(msg.value <= fundingGoal - fundingAmount, "the goal is complete");
        fundingAmount += msg.value;
        funders[msg.sender] += msg.value;
        if(!fundersContributed[msg.sender]){
            fundersList.push(msg.sender);  //存放所有参与者地址
            fundersContributed[msg.sender] = true;
        }
    }

    function close() external returns (bool){
        
        available = false;
        payable(beneficiary).transfer(fundingAmount);
        return true;
    }

    function fundersCount() public view returns (uint256){
        return fundersList.length;
    }

}