// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract Mapping{
    //mapping的key是内置数据类型
    mapping(address => uint256) public balance;

    function get(address  _add) public view returns (uint256){
        return balance[_add];
    }

    function set(address _add, uint256 _amount) public{
        balance[_add] = _amount;
    }

    function remove(address _addr) public{
        delete balance[_addr];
    }
}

contract NestedMapping{
    mapping(address => mapping(uint256 => bool)) public nested;

    function get(address _add, uint256 _amount) public view returns (bool){
        return nested[_add][_amount];
    }

    function set(address _addr, uint256 _amount, bool _b) public{
        nested[_addr][_amount] = _b;
    }

    function remove(address _addr, uint256 _amount) public{
        delete nested[_addr][_amount];
    }
}