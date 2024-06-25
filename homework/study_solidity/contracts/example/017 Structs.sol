// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity >=0.8.24 <0.9.0;


contract Todos {

    struct Todo {
        string text;
        bool completed;
    }

     Todo[] public todos;

    function create(string calldata _text) public {
        todos.push(Todo(_text, false));

        todos.push(Todo({text: _text, completed: false}));

    
        Todo memory todo;
        todo.text = _text;

        todos.push(todo);
    }

    function getTods() public  view returns (Todo[] memory){
            return todos;
    }



  
    function get(uint256 _index)  public view  returns (string memory text, bool completed)
    {
        Todo storage todo = todos[_index];
        return (todo.text, todo.completed);
    }


    function updateText(uint256 _index, string calldata _text) public {
        Todo storage todo = todos[_index];
        todo.text = _text;
    }


    function toggleCompleted(uint256 _index) public {
        Todo storage todo = todos[_index];
        todo.completed = !todo.completed;
    }




}