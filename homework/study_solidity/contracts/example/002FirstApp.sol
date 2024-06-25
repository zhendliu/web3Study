// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity >=0.8.24 <0.9.0;

contract Counter{
    
    uint256 public count = 18;

    function get() public view  returns(uint256){
        return count;
    }

     function inc() public {
        count += 1;
    }

    function dec() public {
        count -= 1;
    }
     
}