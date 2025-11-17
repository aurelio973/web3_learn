// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract BeggingContract {
    // 合约所有者
    address public owner;
    
    // 记录每个捐赠者的捐赠金额
    mapping(address => uint256) public donations;
    
    // 捐赠总额
    uint256 public totalDonations;
    
    // 捐赠事件
    event Donation(address indexed donor, uint256 amount, uint256 timestamp);
    
    // 提款事件
    event Withdrawal(address indexed owner, uint256 amount, uint256 timestamp);
    
    // 捐赠者信息结构体
    struct Donor {
        address donorAddress;
        uint256 amount;
    }
    
    // 所有捐赠者列表
    address[] public donorAddresses;
    
    // 捐赠时间段限制
    uint256 public startTime;
    uint256 public endTime;
    bool public timeLimitEnabled;
    
    // 修饰符：只有所有者可以调用
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }
    
    // 修饰符：检查捐赠时间限制
    modifier withinDonationTime() {
        if (timeLimitEnabled) {
            require(block.timestamp >= startTime && block.timestamp <= endTime, "Donations are not allowed at this time");
        }
        _;
    }
    
    // 构造函数
    constructor() {
        owner = msg.sender;
        timeLimitEnabled = false; // 默认不启用时间限制
    }
    
    // 捐赠函数 - 接收以太币
    function donate() external payable withinDonationTime {
        require(msg.value > 0, "Donation amount must be greater than 0");
        
        // 如果是新捐赠者，添加到列表
        if (donations[msg.sender] == 0) {
            donorAddresses.push(msg.sender);
        }
        
        // 更新捐赠记录
        donations[msg.sender] += msg.value;
        totalDonations += msg.value;
        
        // 触发捐赠事件
        emit Donation(msg.sender, msg.value, block.timestamp);
    }
    
    // 提款函数 - 只有所有者可以调用
    function withdraw() external onlyOwner {
        uint256 balance = address(this).balance;  // 合约当前的以太坊余额
        require(balance > 0, "No funds to withdraw");
        
        // 转账给所有者
        payable(owner).transfer(balance);
        
        // 触发提款事件
        emit Withdrawal(owner, balance, block.timestamp);
    }
    
    // 查询某个地址的捐赠金额
    function getDonation(address donor) external view returns (uint256) {
        return donations[donor];
    }
    
    // 获取捐赠者总数
    function getDonorCount() external view returns (uint256) {
        return donorAddresses.length;
    }
    
    // 获取捐赠排行榜（前N名）
    function getTopDonors(uint256 topN) external view returns (Donor[] memory) {
        require(topN > 0 && topN <= donorAddresses.length, "Invalid number of top donors");
        
        // 创建捐赠者数组
        Donor[] memory allDonors = new Donor[](donorAddresses.length);
        for (uint256 i = 0; i < donorAddresses.length; i++) {
            allDonors[i] = Donor(donorAddresses[i], donations[donorAddresses[i]]);
        }
        
        // 简单排序（冒泡排序，适用于小数据量）
        for (uint256 i = 0; i < allDonors.length - 1; i++) {
            for (uint256 j = 0; j < allDonors.length - i - 1; j++) {
                if (allDonors[j].amount < allDonors[j + 1].amount) {
                    Donor memory temp = allDonors[j];
                    allDonors[j] = allDonors[j + 1];
                    allDonors[j + 1] = temp;
                }
            }
        }
        
        // 返回前N名
        Donor[] memory topDonors = new Donor[](topN);
        for (uint256 i = 0; i < topN; i++) {
            topDonors[i] = allDonors[i];
        }
        
        return topDonors;
    }
    
    // 设置捐赠时间限制
    function setDonationTimeLimit(uint256 _startTime, uint256 _endTime) external onlyOwner {
        require(_startTime < _endTime, "Start time must be before end time");
        require(_endTime > block.timestamp, "End time must be in the future");
        
        startTime = _startTime;
        endTime = _endTime;
        timeLimitEnabled = true;
    }
    
    // 禁用时间限制
    function disableTimeLimit() external onlyOwner {
        timeLimitEnabled = false;
    }
    
    // 获取合约余额
    function getContractBalance() external view returns (uint256) {
        return address(this).balance;
    }
    
    // 获取当前时间
    function getCurrentTime() external view returns (uint256) {
        return block.timestamp;
    }
    
    // 接收以太币的回退函数
    receive() external payable withinDonationTime {
        require(msg.value > 0, "Donation amount must be greater than 0");
        
        if (donations[msg.sender] == 0) {
            donorAddresses.push(msg.sender);
        }
        
        donations[msg.sender] += msg.value;
        totalDonations += msg.value;
        
        emit Donation(msg.sender, msg.value, block.timestamp);
    }
}
