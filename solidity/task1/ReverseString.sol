// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract ReverseString {
    function reverse(string memory _str) public pure returns (string memory) {
        bytes memory strBytes = bytes(_str);
        bytes memory reversed = new bytes(strBytes.length);
        
        for (uint i = 0; i < strBytes.length; i++) {
            reversed[i] = strBytes[strBytes.length - 1 - i];
        }
        
        return string(reversed);
    }
}
    
