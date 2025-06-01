// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract MultiSigWallet{

    event success(bytes32 txHash);
    event failure(bytes32 txHash);

    //多签持有人
    address[] public owners;
    mapping(address => bool) public isOwner;
    //多签持有人数量
    uint256 public ownerCount;
    //多签执行门槛
    uint256 public threshold;
    //防止多链签名重放攻击
    uint256 public nonce;

    constructor(address[] memory _owners, uint256 _threshold){
        _setUpOwners(_owners, _threshold);
    }

    function _setUpOwners(address[] memory _owners, uint256 _threshold) internal {
        require(_threshold <= _owners.length);
        require(_threshold >= 1);
        for(uint256 i = 0; i < _owners.length; ++i){
            address owner = _owners[i];
            require(owner != address(0) && owner != address(this) && !isOwner[owner]);
            owners.push(owner);
            isOwner[owner] = true;
        }
        ownerCount = _owners.length;
        threshold = _threshold;
    }

    function execTransaction(address to, uint256 value, bytes memory data, bytes memory signatures) public payable virtual returns (bool success){
        bytes32 txHash = encodeTransactionData(to, value, data, nonce, block.chainid);
        nonce ++;
        checkSignatures(txHash, signatures);
        (success, ) = to.call{value : value}(data);
        if(success){
            emit success(txHash);
        }else {
            emit failure(txHash);
        }
    }

    function checkSignatures(bytes32 txHash, bytes memory signatures) public view {
        uint256 _threshold = threshold;
        require(_threshold > 0);
        require(signatures.length >= 65 * _threshold);
    // 通过一个循环，检查收集的签名是否有效
    // 大概思路：
    // 1. 用ecdsa先验证签名是否有效
    // 2. 利用 currentOwner > lastOwner 确定签名来自不同多签（多签地址递增）
    // 3. 利用 isOwner[currentOwner] 确定签名者为多签持有人
        address lastOwner = address(0); 
        address currentOwner;
        uint8 v;
        bytes32 r;
        bytes32 s;
        uint256 i;
        for (i = 0; i < _threshold; i++) {
            (v, r, s) = signatureSplit(signatures, i);
            // 利用ecrecover检查签名是否有效
            currentOwner = ecrecover(keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", dataHash)), v, r, s);
            require(currentOwner > lastOwner && isOwner[currentOwner], "WTF5007");
            lastOwner = currentOwner;
        }

    }

    /// 将单个签名从打包的签名分离出来
/// @param signatures 打包签名
/// @param pos 要读取的多签index.
    function signatureSplit(bytes memory signatures, uint256 pos)internal pure returns (uint8 v, bytes32 r, bytes32 s){
        // 签名的格式：{bytes32 r}{bytes32 s}{uint8 v}
        assembly {
            let signaturePos := mul(0x41, pos)
            r := mload(add(signatures, add(signaturePos, 0x20)))
            s := mload(add(signatures, add(signaturePos, 0x40)))
            v := and(mload(add(signatures, add(signaturePos, 0x41))), 0xff)
        }
    }
    /// @dev 编码交易数据
    /// @param to 目标合约地址
    /// @param value msg.value，支付的以太坊
    /// @param data calldata
    /// @param _nonce 交易的nonce.
    /// @param chainid 链id
    /// @return 交易哈希bytes.
    function encodeTransactionData(address to,uint256 value,bytes memory data,uint256 _nonce,uint256 chainid) public pure returns (bytes32) {
        bytes32 safeTxHash =keccak256(abi.encode(to,value,keccak256(data), _nonce,chainid));
        return safeTxHash;
    }



}