// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract SimpleERC20 {
    // 代币基本信息
    string public name;
    string public symbol;
    uint8 public decimals;
    uint256 public totalSupply;

    // 合约所有者
    address public owner;

    // mapping
    mapping (address=>uint256) public balanceOf;
    mapping (address=>mapping(address=>uint256)) public allowance;

    // event
    event Transfer(address indexed from,address indexed to,uint256 value);
    event Approval(address indexed owner,address indexed spender,uint256 value);
    event Mint(address indexed to,uint256 value);

    // 修饰器：只有所有者可以调用
    modifier onlyOwner() {
        require(msg.sender == owner,"Only owner can call this function");
        _;
    }

    constructor(
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 _initialSupply
    ){
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
        totalSupply = _initialSupply * 10 ** uint256(_decimals);
        owner = msg.sender;
        
        balanceOf[msg.sender] = totalSupply;
        emit Transfer(address(0),msg.sender,totalSupply);
    }

    // 转账功能
    function transfer(address _to, uint256 _value) public returns(bool) {
        require(_to != address(0),"Invalid address");
        require(balanceOf[msg.sender]>= _value,"Insufficiant balance");

        balanceOf[msg.sender]-=_value;
        balanceOf[_to]+= _value;

        emit Transfer(msg.sender,_to,_value);
        return true;
    }

    // 授权功能
    function approve(address _spender,uint256 _value) public returns(bool) {
        allowance[msg.sender][_spender]=_value;
        emit Approval(msg.sender,_spender,_value);
        return true;
    }

    // 代扣转账功能
    function transferFrom(
        address _from,
        address _to,
        uint256 _value
    ) public returns (bool) {
        require(_from != address(0),"Invalid address");
        require(_to != address(0),"Invalid address");
        require(balanceOf[_from] >= _value,"Insufficiant balance");
        require(allowance[_from][msg.sender] >= _value,"Allowance exceeded");

        balanceOf[_from]-=_value;
        balanceOf[_to]+=_value;
        allowance[_from][msg.sender]-=_value;

        emit Transfer(_from, _to, _value);
        return true;
    }

    // 增发代币
    function mint(address _to,uint256 _value) public onlyOwner {
        require(_to != address(0),"Mint to invalid address");
        require(_value > 0,"Mint value must be greater than zero");

        uint256 mintAmount = _value * 10 ** (decimals);
        totalSupply += mintAmount;
        balanceOf[_to] += mintAmount;

        emit Transfer(address(0), _to,  mintAmount);
        emit Mint(_to,mintAmount);
    }

    // 转移所有权
    function TransferOwnership(address _newOwner) public onlyOwner {
        require(_newOwner != address(0),"New owner is invalid address");
        owner = _newOwner;
    }
}
