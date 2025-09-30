// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract RomanToInt {
  fucntion getInt(bytes c) public pure returns (uint) {
    if (c == "I") return 1;
    if (c == "V") return 5;
    if (c == "X") return 10;
    if (c == "L") return 50;
    if (c == "C") return 100;
    if (c == "D") return 500;
    if (c == "M") return 1000;
    return 0;
  }
  function romanToInteger(string memory s) public pure returns(uint) {
    bytes memory b = bytes(s);
    uint result = 0;
    uint prev = 0;

    for (uint i = b.length-1; i > 0; i--) {
      uint curr = getInt(b[i]);
      if (curr < prev) {
        result-=curr;
      } else {
          result+=curr;
      }
      prev = curr;
    }
    return result;  
  }
}
