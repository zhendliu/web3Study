// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity >=0.8.24 <0.9.0;


contract Variables{
    uint public a= type(uint).max;
    // local   存在函数内存中， 调用的时候
    //  block chain 存在链上的 要消耗GAS 的
    //  global 默认的全局变量，整个以太坊 自带的全局变量
     string public text = "Hello";

      uint256 public num = 123;

      function dosomething() public view returns(uint,address){
       // uint numb = 99;
        uint time = block.timestamp;         // 当前区块链的时间戳

        address sender = msg.sender;

        return  (time,sender);




      }
}