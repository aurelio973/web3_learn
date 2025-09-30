// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract BinarySearch {
    function search(uint[] memory arr,uint target) public pure returns(int) {
        uint left = 0;
        uint right = arr.length;

        while(left < right) {
            uint mid = (left + right)/2;
            // 中间值
            if(arr[mid] == target) return int(mid);
            // 右半区间
            if(arr[mid] < target) left = mid + 1;
            // 左半区间
            else right = mid;
        }
        return -1;
    }
}
