// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Voting {
    mapping (string => uint256) private votes;
    string[] private candidates;

    // 投票
    function vote(string memory candidate) public {
        if (votes[candidate] == 0) {
        candidates.push(candidate);
    }
        votes[candidate]+=1;
    }

    // 获取票数
    function getVotes(string memory candidate) public view returns (uint256) {
        return votes[candidate];
    }

    // 重置票数
    function resetVotes() public {
        uint length = candidates.length;
        for(uint i=0;i < length; i++) {
            votes[candidates[i]] = 0;
        }
    }
}
