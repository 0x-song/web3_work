// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract WETH{
    string public name = "Wrapped Ether";

    string public symbol = "WETH";

    uint8 public decimals = 18;

    event Approval(address indexed src, address indexed delegateAds, uint256 amount);

    event Transfer(address indexed src, address indexed toAds, uint256 amount);

    event Deposit(address indexed src, uint256 amount);

    event Withdraw(address indexed src, uint256 amount);

    mapping(address => uint256) public balanceOf;

    mapping(address => mapping(address => uint256)) public allowance;

    //存入一定数量的 eth，weth合约中会分配对应数量的 额度
    function deposit() public payable{
        balanceOf[msg.sender] += msg.value;
        emit Deposit(msg.sender, msg.value);
    }
    //提取一定的eth，那么 weth的额度会减少，同时进行eth转账
    function withdraw(uint256 amount) public{
        require(balanceOf[msg.sender] > amount, "not enought funds");
        balanceOf[msg.sender] -= amount;
        //这行代码为什么？？？？
        payable(msg.sender).transfer(amount);
        emit Withdraw(msg.sender, amount);
    }

    //这个啥意思
    function totalSupply() public view returns (uint256){
        return address(this).balance;
    }

    function approve(address delegateAds, uint256 amount) public  returns (bool){
        allowance[msg.sender][delegateAds] = amount;
        emit Approval(msg.sender, delegateAds, amount);
        return true;
    }

    function transfer(address toAds, uint256 amount) public returns (bool){
        return transferFrom(msg.sender, toAds, amount);
    }

    function transferFrom(address src, address toAds, uint256 amount) public returns (bool){
        require(balanceOf[src] >= amount);
        if(src != msg.sender){
            //通过合约进行转账划扣
            
            require(allowance[src][msg.sender] >= amount);
            allowance[src][msg.sender] -= amount;
        }
        balanceOf[toAds] += amount;
        balanceOf[src] -= amount;
        emit Transfer(src, toAds, amount);
        return true;
    }

    fallback() external payable{
        deposit();
    }
    receive() external payable{
        deposit();
    }

}