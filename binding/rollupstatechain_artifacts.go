package binding

import (
	"encoding/hex"
	"fmt"

	"github.com/laizy/web3/abi"
)

var abiRollupStateChain *abi.ABI

// RollupStateChainAbi returns the abi of the RollupStateChain contract
func RollupStateChainAbi() *abi.ABI {
	return abiRollupStateChain
}

var binRollupStateChain []byte

// RollupStateChainBin returns the bin of the RollupStateChain contract
func RollupStateChainBin() []byte {
	return binRollupStateChain
}

var binRuntimeRollupStateChain []byte

// RollupStateChainBinRuntime returns the runtime bin of the RollupStateChain contract
func RollupStateChainBinRuntime() []byte {
	return binRuntimeRollupStateChain
}

func init() {
	var err error
	abiRollupStateChain, err = abi.NewABI(abiRollupStateChainStr)
	if err != nil {
		panic(fmt.Errorf("cannot parse RollupStateChain abi: %v", err))
	}
	if len(binRollupStateChainStr) != 0 {
		binRollupStateChain, err = hex.DecodeString(binRollupStateChainStr[2:])
		if err != nil {
			panic(fmt.Errorf("cannot parse RollupStateChain bin: %v", err))
		}
	}
	if len(binRuntimeRollupStateChainStr) != 0 {
		binRuntimeRollupStateChain, err = hex.DecodeString(binRuntimeRollupStateChainStr[2:])
		if err != nil {
			panic(fmt.Errorf("cannot parse RollupStateChain bin runtime: %v", err))
		}
	}
}

var binRollupStateChainStr = "0x608060405234801561001057600080fd5b50611796806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063935b0d781161005b578063935b0d78146100d2578063cd6dc687146100f3578063e46020a114610106578063e9c706df1461011d57600080fd5b8063325aeae21461008257806376ef0aaa146100aa57806392927f11146100bf575b600080fd5b610095610090366004611487565b610130565b60405190151581526020015b60405180910390f35b6100bd6100b8366004611487565b610157565b005b6100956100cd366004611487565b610549565b6100da61071f565b60405167ffffffffffffffff90911681526020016100a1565b6100bd610101366004611503565b610824565b61010f60015481565b6040519081526020016100a1565b6100bd61012b36600461152f565b610917565b600042600154836040015167ffffffffffffffff1661014f9190611616565b111592915050565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635dbaf68b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101e8919061162e565b6040517fb363ff8500000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff919091169063b363ff8590602401602060405180830381865afa158015610254573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610278919061164b565b610308576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f6e6c79207065726d6974746564206279206368616c6c656e676520636f6e7460448201527f726163740000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b61031181610549565b610377576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c696420737461746520696e666f000000000000000000000000000060448201526064016102ff565b61038081610130565b156103e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f737461746520636f6e6669726d6564000000000000000000000000000000000060448201526064016102ff565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663388f2a0a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610454573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610478919061162e565b60208201516040517f5682afa900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015273ffffffffffffffffffffffffffffffffffffffff9190911690635682afa990602401600060405180830381600087803b1580156104f157600080fd5b505af1158015610505573d6000803e3d6000fd5b50508251602084015160405191935067ffffffffffffffff1691507f911ace459082010270e47b0b03415673320c53da9e7918fc7d0b0c379f80514590600090a350565b600080600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663388f2a0a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105dd919061162e565b90508073ffffffffffffffffffffffffffffffffffffffff166331fe09496040518163ffffffff1660e01b8152600401602060405180830381865afa15801561062a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064e919061166d565b67ffffffffffffffff16836020015167ffffffffffffffff1610801561071857506106788361117a565b60208401516040517fada8679800000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015273ffffffffffffffffffffffffffffffffffffffff83169063ada8679890602401602060405180830381865afa1580156106f2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610716919061168a565b145b9392505050565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663388f2a0a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561078d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107b1919061162e565b73ffffffffffffffffffffffffffffffffffffffff166331fe09496040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061081f919061166d565b905090565b60006108306001611193565b9050801561086557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8616021790556001829055801561091257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634162169f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610984573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109a8919061162e565b6040517fed4be8ae00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff919091169063ed4be8ae90602401602060405180830381865afa158015610a14573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a38919061164b565b610a9e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f6f6e6c792070726f706f7365720000000000000000000000000000000000000060448201526064016102ff565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663388f2a0a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b30919061162e565b90508073ffffffffffffffffffffffffffffffffffffffff166331fe09496040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b7d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ba1919061166d565b67ffffffffffffffff168267ffffffffffffffff1614610c1d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f737461727420706f73206d69736d61746368000000000000000000000000000060448201526064016102ff565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322828cc26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cae919061162e565b6040517f6f49712b00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff9190911690636f49712b90602401602060405180830381865afa158015610d1a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d3e919061164b565b610da4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f756e7374616b656400000000000000000000000000000000000000000000000060448201526064016102ff565b6000835111610e0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f6e6f20626c6f636b20686173686573000000000000000000000000000000000060448201526064016102ff565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166374aee6c96040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e7c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ea0919061162e565b73ffffffffffffffffffffffffffffffffffffffff1663761a26616040518163ffffffff1660e01b8152600401602060405180830381865afa158015610eea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f0e919061166d565b67ffffffffffffffff1683518273ffffffffffffffffffffffffffffffffffffffff166331fe09496040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f89919061166d565b67ffffffffffffffff16610f9d9190611616565b1115611005576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f65786365656420696e70757420636861696e206865696768740000000000000060448201526064016102ff565b604080516080810182526000808252602082018190524267ffffffffffffffff81169383019390935233606083015284905b865181101561111657868181518110611052576110526116a3565b602090810291909101810151845267ffffffffffffffff83169084015273ffffffffffffffffffffffffffffffffffffffff8516636483ec256110948561117a565b6040518263ffffffff1660e01b81526004016110b291815260200190565b6020604051808303816000875af11580156110d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110f5919061166d565b5081611100816116d2565b925050808061110e906116f9565b915050611037565b508467ffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167ffd1ab91e7c217cde3474f0c085a92f117c977c8a9c04b903d549129f00de539a858960405161116a929190611731565b60405180910390a3505050505050565b60006111858261131e565b805190602001209050919050565b60008054610100900460ff161561124a578160ff1660011480156111b65750303b155b611242576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016102ff565b506000919050565b60005460ff8084169116106112e1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016102ff565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b606081600001518260200151836040015184606001516040516020016113ad949392919093845260c092831b7fffffffffffffffff00000000000000000000000000000000000000000000000090811660208601529190921b16602883015260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016603082015260440190565b6040516020818303038152906040529050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611439576114396113c3565b604052919050565b67ffffffffffffffff8116811461145757600080fd5b50565b803561131981611441565b73ffffffffffffffffffffffffffffffffffffffff8116811461145757600080fd5b60006080828403121561149957600080fd5b6040516080810181811067ffffffffffffffff821117156114bc576114bc6113c3565b6040528235815260208301356114d181611441565b602082015260408301356114e481611441565b604082015260608301356114f781611465565b60608201529392505050565b6000806040838503121561151657600080fd5b823561152181611465565b946020939093013593505050565b6000806040838503121561154257600080fd5b823567ffffffffffffffff8082111561155a57600080fd5b818501915085601f83011261156e57600080fd5b8135602082821115611582576115826113c3565b8160051b92506115938184016113f2565b82815292840181019281810190898511156115ad57600080fd5b948201945b848610156115cb578535825294820194908201906115b2565b96506115da905087820161145a565b9450505050509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115611629576116296115e7565b500190565b60006020828403121561164057600080fd5b815161071881611465565b60006020828403121561165d57600080fd5b8151801515811461071857600080fd5b60006020828403121561167f57600080fd5b815161071881611441565b60006020828403121561169c57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600067ffffffffffffffff8083168181036116ef576116ef6115e7565b6001019392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361172a5761172a6115e7565b5060010190565b60006040820167ffffffffffffffff851683526020604081850152818551808452606086019150828701935060005b8181101561177c57845183529383019391830191600101611760565b509097965050505050505056fea164736f6c634300080d000a"

var binRuntimeRollupStateChainStr = "0x608060405234801561001057600080fd5b506004361061007d5760003560e01c8063935b0d781161005b578063935b0d78146100d2578063cd6dc687146100f3578063e46020a114610106578063e9c706df1461011d57600080fd5b8063325aeae21461008257806376ef0aaa146100aa57806392927f11146100bf575b600080fd5b610095610090366004611487565b610130565b60405190151581526020015b60405180910390f35b6100bd6100b8366004611487565b610157565b005b6100956100cd366004611487565b610549565b6100da61071f565b60405167ffffffffffffffff90911681526020016100a1565b6100bd610101366004611503565b610824565b61010f60015481565b6040519081526020016100a1565b6100bd61012b36600461152f565b610917565b600042600154836040015167ffffffffffffffff1661014f9190611616565b111592915050565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635dbaf68b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101e8919061162e565b6040517fb363ff8500000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff919091169063b363ff8590602401602060405180830381865afa158015610254573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610278919061164b565b610308576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f6e6c79207065726d6974746564206279206368616c6c656e676520636f6e7460448201527f726163740000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b61031181610549565b610377576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c696420737461746520696e666f000000000000000000000000000060448201526064016102ff565b61038081610130565b156103e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f737461746520636f6e6669726d6564000000000000000000000000000000000060448201526064016102ff565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663388f2a0a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610454573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610478919061162e565b60208201516040517f5682afa900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015273ffffffffffffffffffffffffffffffffffffffff9190911690635682afa990602401600060405180830381600087803b1580156104f157600080fd5b505af1158015610505573d6000803e3d6000fd5b50508251602084015160405191935067ffffffffffffffff1691507f911ace459082010270e47b0b03415673320c53da9e7918fc7d0b0c379f80514590600090a350565b600080600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663388f2a0a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105dd919061162e565b90508073ffffffffffffffffffffffffffffffffffffffff166331fe09496040518163ffffffff1660e01b8152600401602060405180830381865afa15801561062a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064e919061166d565b67ffffffffffffffff16836020015167ffffffffffffffff1610801561071857506106788361117a565b60208401516040517fada8679800000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015273ffffffffffffffffffffffffffffffffffffffff83169063ada8679890602401602060405180830381865afa1580156106f2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610716919061168a565b145b9392505050565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663388f2a0a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561078d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107b1919061162e565b73ffffffffffffffffffffffffffffffffffffffff166331fe09496040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061081f919061166d565b905090565b60006108306001611193565b9050801561086557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b600080547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8616021790556001829055801561091257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634162169f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610984573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109a8919061162e565b6040517fed4be8ae00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff919091169063ed4be8ae90602401602060405180830381865afa158015610a14573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a38919061164b565b610a9e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f6f6e6c792070726f706f7365720000000000000000000000000000000000000060448201526064016102ff565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663388f2a0a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b30919061162e565b90508073ffffffffffffffffffffffffffffffffffffffff166331fe09496040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b7d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ba1919061166d565b67ffffffffffffffff168267ffffffffffffffff1614610c1d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f737461727420706f73206d69736d61746368000000000000000000000000000060448201526064016102ff565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322828cc26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cae919061162e565b6040517f6f49712b00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff9190911690636f49712b90602401602060405180830381865afa158015610d1a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d3e919061164b565b610da4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f756e7374616b656400000000000000000000000000000000000000000000000060448201526064016102ff565b6000835111610e0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f6e6f20626c6f636b20686173686573000000000000000000000000000000000060448201526064016102ff565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166374aee6c96040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e7c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ea0919061162e565b73ffffffffffffffffffffffffffffffffffffffff1663761a26616040518163ffffffff1660e01b8152600401602060405180830381865afa158015610eea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f0e919061166d565b67ffffffffffffffff1683518273ffffffffffffffffffffffffffffffffffffffff166331fe09496040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f89919061166d565b67ffffffffffffffff16610f9d9190611616565b1115611005576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f65786365656420696e70757420636861696e206865696768740000000000000060448201526064016102ff565b604080516080810182526000808252602082018190524267ffffffffffffffff81169383019390935233606083015284905b865181101561111657868181518110611052576110526116a3565b602090810291909101810151845267ffffffffffffffff83169084015273ffffffffffffffffffffffffffffffffffffffff8516636483ec256110948561117a565b6040518263ffffffff1660e01b81526004016110b291815260200190565b6020604051808303816000875af11580156110d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110f5919061166d565b5081611100816116d2565b925050808061110e906116f9565b915050611037565b508467ffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167ffd1ab91e7c217cde3474f0c085a92f117c977c8a9c04b903d549129f00de539a858960405161116a929190611731565b60405180910390a3505050505050565b60006111858261131e565b805190602001209050919050565b60008054610100900460ff161561124a578160ff1660011480156111b65750303b155b611242576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016102ff565b506000919050565b60005460ff8084169116106112e1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016102ff565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b606081600001518260200151836040015184606001516040516020016113ad949392919093845260c092831b7fffffffffffffffff00000000000000000000000000000000000000000000000090811660208601529190921b16602883015260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016603082015260440190565b6040516020818303038152906040529050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611439576114396113c3565b604052919050565b67ffffffffffffffff8116811461145757600080fd5b50565b803561131981611441565b73ffffffffffffffffffffffffffffffffffffffff8116811461145757600080fd5b60006080828403121561149957600080fd5b6040516080810181811067ffffffffffffffff821117156114bc576114bc6113c3565b6040528235815260208301356114d181611441565b602082015260408301356114e481611441565b604082015260608301356114f781611465565b60608201529392505050565b6000806040838503121561151657600080fd5b823561152181611465565b946020939093013593505050565b6000806040838503121561154257600080fd5b823567ffffffffffffffff8082111561155a57600080fd5b818501915085601f83011261156e57600080fd5b8135602082821115611582576115826113c3565b8160051b92506115938184016113f2565b82815292840181019281810190898511156115ad57600080fd5b948201945b848610156115cb578535825294820194908201906115b2565b96506115da905087820161145a565b9450505050509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115611629576116296115e7565b500190565b60006020828403121561164057600080fd5b815161071881611465565b60006020828403121561165d57600080fd5b8151801515811461071857600080fd5b60006020828403121561167f57600080fd5b815161071881611441565b60006020828403121561169c57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600067ffffffffffffffff8083168181036116ef576116ef6115e7565b6001019392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361172a5761172a6115e7565b5060010190565b60006040820167ffffffffffffffff851683526020604081850152818551808452606086019150828701935060005b8181101561177c57845183529383019391830191600101611760565b509097965050505050505056fea164736f6c634300080d000a"

var abiRollupStateChainStr = `[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint8","name":"version","type":"uint8"}],"name":"Initialized","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"_proposer","type":"address"},{"indexed":true,"internalType":"uint64","name":"_startIndex","type":"uint64"},{"indexed":false,"internalType":"uint64","name":"_timestamp","type":"uint64"},{"indexed":false,"internalType":"bytes32[]","name":"_blockHash","type":"bytes32[]"}],"name":"StateBatchAppended","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint64","name":"_stateIndex","type":"uint64"},{"indexed":true,"internalType":"bytes32","name":"_blockHash","type":"bytes32"}],"name":"StateRollbacked","type":"event"},{"inputs":[{"internalType":"bytes32[]","name":"_blockHashes","type":"bytes32[]"},{"internalType":"uint64","name":"_startAt","type":"uint64"}],"name":"appendStateBatch","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"fraudProofWindow","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_addressResolver","type":"address"},{"internalType":"uint256","name":"_fraudProofWindow","type":"uint256"}],"name":"initialize","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"components":[{"internalType":"bytes32","name":"blockHash","type":"bytes32"},{"internalType":"uint64","name":"index","type":"uint64"},{"internalType":"uint64","name":"timestamp","type":"uint64"},{"internalType":"address","name":"proposer","type":"address"}],"internalType":"struct Types.StateInfo","name":"_stateInfo","type":"tuple"}],"name":"isStateConfirmed","outputs":[{"internalType":"bool","name":"_confirmed","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"components":[{"internalType":"bytes32","name":"blockHash","type":"bytes32"},{"internalType":"uint64","name":"index","type":"uint64"},{"internalType":"uint64","name":"timestamp","type":"uint64"},{"internalType":"address","name":"proposer","type":"address"}],"internalType":"struct Types.StateInfo","name":"_stateInfo","type":"tuple"}],"name":"rollbackStateBefore","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"totalSubmittedState","outputs":[{"internalType":"uint64","name":"","type":"uint64"}],"stateMutability":"view","type":"function"},{"inputs":[{"components":[{"internalType":"bytes32","name":"blockHash","type":"bytes32"},{"internalType":"uint64","name":"index","type":"uint64"},{"internalType":"uint64","name":"timestamp","type":"uint64"},{"internalType":"address","name":"proposer","type":"address"}],"internalType":"struct Types.StateInfo","name":"_stateInfo","type":"tuple"}],"name":"verifyStateInfo","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
