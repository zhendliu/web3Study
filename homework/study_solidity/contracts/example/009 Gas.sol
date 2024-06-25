// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity >=0.8.24 <0.9.0;
contract Gas {
    uint256 public i = 0;
    function forever() public {
        while (true) {
            i += 1;
        }
    }
}