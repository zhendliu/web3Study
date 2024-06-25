// SPDX-License-Identifier: MIT
// compiler version must be greater than or equal to 0.8.24 and less than 0.9.0
pragma solidity >=0.8.24 <0.9.0;

contract Enum {

    enum Status {
        Pending,
        Shipped,
        Accepted,
        Rejected,
        Canceled
    }


    Status public status;

    
    function get() public view returns (Status) {
        return status;
    }


    function set(Status _status) public {
        status = _status;
    }


    function cancel() public {
        status = Status.Canceled;
    }

    
    function reset() public {
        delete status;
    }
}