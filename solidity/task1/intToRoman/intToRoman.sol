// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract intToRoman {
  struct RomanMap {
    uint value;
    string symbol;
  }
    function integerToRoman(uint num) public pure returns (string memory) {
        require(num >= 1 && num < 3999,"Out of range");
        RomanMap[13] memory roman = [
        RomanMap(1000, "M"),
        RomanMap(900, "CM"),
        RomanMap(500, "D"),
        RomanMap(400, "CD"),
        RomanMap(100, "C"),
        RomanMap(90, "XC"),
        RomanMap(50, "L"),
        RomanMap(40, "XL"),
        RomanMap(10, "X"),
        RomanMap(9, "IX"),
        RomanMap(5, "V"),
        RomanMap(4, "IV"),
        RomanMap(1, "I")
        ];
        string memory result = "";
        for (uint i = 0; i < roman.length; i++) {
            while(num >= roman[i].value) {
            result = string.concat(result,roman[i].symbol);
            num -= roman[i].value;
            }
        }
    return result;
    }
}
