// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SimpleCounter {
    uint256 private count;
    
    event CountChanged(uint256 newCount, address changedBy);
    
    function getCount() public view returns (uint256) {
        return count;
    }
    
    function increment() public {
        count += 1;
        emit CountChanged(count, msg.sender);
    }
    
    function decrement() public {
        require(count > 0, "Count cannot be negative");
        count -= 1;
        emit CountChanged(count, msg.sender);
    }
    
    function setCount(uint256 _count) public {
        count = _count;
        emit CountChanged(count, msg.sender);
    }
}
