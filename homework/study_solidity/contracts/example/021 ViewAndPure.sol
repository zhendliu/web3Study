// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity >=0.8.24 <0.9.0;

contract ViewAndPure {
    uint256 public x = 1;

    function addToX(uint256 y) public view returns (uint256) {
        return x + y;
    }

    function add(uint256 i, uint256 j) public pure returns (uint256) {
        return i + j;
    }
}