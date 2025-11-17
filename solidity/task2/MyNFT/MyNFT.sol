// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol"; // 导入标准的ERC721合约实现
import "@openzeppelin/contracts/utils/Counters.sol"; // 提供安全的计数器功能，用于生成tokenId 防止溢出和其他数学错误
import "@openzeppelin/contracts/access/Ownable.sol"; // 提供所有权管理功能，只有合约所有者可以执行某些操作

contract MyNFT is ERC721,Ownable {  // 多重继承 继承ERC721标准和Ownable功能 继承顺序会影响构造函数调用的顺序
    using Counters for Counters.Counter; // 为Counter类型添加扩展方法 using A for B: 将库 A 的函数附加到类型 B 上
    Counters.Counter private _tokenIds; // 私有的计数器，用于生成唯一的tokenId

    // 基础 token URI
    string private _baseTokenURI; // 存储基础URI，用于构建完整的tokenURI

    // 事件：NFT 铸造事件
    event NFTMinted(uint256 indexed tokenId,address indexed recipient,string tokenURI);// 记录NFT铸造的详细信息

    // 构造函数：设置NFT名称和符号
    constructor(
        string memory name,  // NFT集合名称
        string memory symbol, // NFT符号
        string memory baseTokenURI // 基础元数据URI
    ) ERC721(name, symbol) Ownable(msg.sender) { // 调用父类构造函数  设置合约部署者为所有者
        _baseTokenURI = baseTokenURI;   
    }
    
    // 设置基础token URI
    function _baseURI() internal view virtual override returns (string memory) { // 只读不修改 父类开许可（virtual），子类来执行（override）允许重写（继承并有个性化功能）
        return _baseTokenURI;
    }  // 完整的tokenURI=baseURI+tokenId

    // 铸造NFT函数
    function mintNFT(address recipient,string memory tokenURI) public onlyOwner returns (uint256) { // 接受NFT的地址  NFT的元数据URI 只有所有者可以调用  值类型不需要指定
        _tokenIds.increment(); // 记录已铸造的NFT数量 自增函数：计数器+1
        uint256 newTokenId = _tokenIds.current(); // 获取当前计数器值作为tokenId

        _mint(recipient,newTokenId);   // 调用父类ERC721的_mint函数  指定NFT的接收者，为这个NFT分配唯一的tokenId

        emit NFTMinted(newTokenId,recipient,tokenURI);  // 触发事件 新的NFT的ID  接收者地址 NFT的元数据URI
        return newTokenId;
    }

    // 批量铸造 NFT
    function batchMintNFT(address recipient,uint256 count) public onlyOwner returns (uint256[] memory) { 
        // 所有新 NFT 都会分配给这个地址  要铸造的 NFT 数量  返回的是整数数组  引用类型（数组、结构体、字符串需要指定存储位置）
        uint256[] memory tokenIds = new uint256[](count);   // 创建一个长度为 count 的动态数组

        for(uint256 i = 0;i < count; i++) {  // 每次铸造一个NFT
            _tokenIds.increment();  // 计数器自增，为每个NFT生成一个新的tokenId
            uint256 newTokenId = _tokenIds.current(); 
            _mint(recipient,newTokenId);
            tokenIds[i] = newTokenId; 
        }
        return tokenIds;
    }

    // 获取当前tokenId计数
    function getCurrentTokenId() public view returns (uint256) {
        return _tokenIds.current();
    }

    // 获取合约信息
    function getContractInfo() public view returns (string memory name,string memory symbol,uint256 totalMinted) {
        return (super.name(),super.symbol(),_tokenIds.current()-1);  // super.name()和super.symbol()调用父合约ERC721的查询函数
    }   
    // 返回了合约的名称、符号和已铸造的NFT总数
    // 铸造时，计算器先+1，再current()获取当前值作为tokenId  
    // 查询时，current()返回的是下一个要用的ID，不是已使用的ID
}
