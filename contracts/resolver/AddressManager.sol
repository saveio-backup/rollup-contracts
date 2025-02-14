// SPDX-License-Identifier: GPL-v3
pragma solidity ^0.8.0;

import "../interfaces/IAddressManager.sol";
import "./AddressName.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "../interfaces/IAddressResolver.sol";

contract AddressManager is IAddressManager, IAddressResolver, OwnableUpgradeable {
    mapping(bytes32 => address) public getAddrByHash;

    function initialize() public initializer {
        __Ownable_init();
    }

    function setAddress(string memory _name, address _addr) public onlyOwner {
        bytes32 _hash = hash(_name);
        address _old = _setAddress(_hash, _addr);
        emit AddressSet(_name, _old, _addr);
    }

    function _setAddress(bytes32 _hash, address _addr) internal returns (address) {
        require(_addr != address(0), "empty addr");
        address _old = getAddrByHash[_hash];
        getAddrByHash[_hash] = _addr;
        return _old;
    }

    function setAddressBatch(string[] calldata _names, address[] calldata _addrs) public onlyOwner {
        uint256 _len = _names.length;
        require(_len == _addrs.length, "length mismatch");
        for (uint256 i = 0; i < _len; i++) {
            string calldata _name = _names[i];
            address _addr = _addrs[i];
            bytes32 _hash = hash(_name);
            address _old = _setAddress(_hash, _addr);
            emit AddressSet(_name, _old, _addr);
        }
    }

    function getAddr(string memory _name) public view returns (address) {
        return getAddrByHash[hash(_name)];
    }

    function resolve(string memory _name) public view returns (address) {
        address _addr = this.getAddr(_name);
        require(_addr != address(0), "no name saved");
        return _addr;
    }

    function dao() public view returns (IDAO) {
        return IDAO(resolve(AddressName.DAO));
    }

    function rollupInputChain() public view returns (IRollupInputChain) {
        return IRollupInputChain(resolve(AddressName.ROLLUP_INPUT_CHAIN));
    }

    function rollupInputChainContainer() public view returns (IChainStorageContainer) {
        return IChainStorageContainer(resolve(AddressName.ROLLUP_INPUT_CHAIN_CONTAINER));
    }

    function rollupStateChain() public view returns (IRollupStateChain) {
        return IRollupStateChain(resolve(AddressName.ROLLUP_STATE_CHAIN));
    }

    function rollupStateChainContainer() public view returns (IChainStorageContainer) {
        return IChainStorageContainer(resolve(AddressName.ROLLUP_STATE_CHAIN_CONTAINER));
    }

    function stakingManager() public view returns (IStakingManager) {
        return IStakingManager(resolve(AddressName.STAKING_MANAGER));
    }

    function challengeFactory() public view returns (IChallengeFactory) {
        return IChallengeFactory(resolve(AddressName.CHALLENGE_FACTORY));
    }

    function l1CrossLayerWitness() public view returns (IL1CrossLayerWitness) {
        return IL1CrossLayerWitness(resolve(AddressName.L1_CROSS_LAYER_WITNESS));
    }

    function l2CrossLayerWitness() public view returns (IL2CrossLayerWitness) {
        return IL2CrossLayerWitness(resolve(AddressName.L2_CROSS_LAYER_WITNESS));
    }

    function stateTransition() public view returns (IStateTransition) {
        return IStateTransition(resolve(AddressName.STATE_TRANSITION));
    }

    function hash(string memory _name) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(_name));
    }
}
