// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract Struct{
    struct Todo{
        string text;
        bool completed;
    }

    Todo[] public todos;

    function create1(string calldata _text) public {
        todos.push(Todo(_text, false));
    }

    function create2(string calldata _text) public{
        todos.push(Todo({text : _text, completed: false}));
    }

    function create3(string calldata _text) public{
        Todo memory _todo;
        _todo.text = _text;
        todos.push(_todo);
    }

    // update text
    function updateText(uint256 _index, string calldata _text) public {
        Todo storage todo = todos[_index];
        todo.text = _text;
    }
}