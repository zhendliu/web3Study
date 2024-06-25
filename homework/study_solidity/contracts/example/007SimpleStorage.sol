// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity >=0.8.24 <0.9.0;

contract SimpleStorage {
     uint256 public num;
     function set(uint256 _num) public {
        num = _num;
    }
    function get() public view returns (uint256) {
        return num;
    }

}