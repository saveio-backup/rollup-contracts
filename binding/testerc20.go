package binding

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/laizy/web3"
	"github.com/laizy/web3/contract"
	"github.com/laizy/web3/jsonrpc"
	"github.com/laizy/web3/utils"
	"github.com/mitchellh/mapstructure"
)

var (
	_ = json.Unmarshal
	_ = big.NewInt
	_ = fmt.Printf
	_ = utils.JsonStr
	_ = mapstructure.Decode
)

// TestERC20 is a solidity contract
type TestERC20 struct {
	c *contract.Contract
}

// DeployTestERC20 deploys a new TestERC20 contract
func DeployTestERC20(provider *jsonrpc.Client, from web3.Address, name string, symbol string) *contract.Txn {
	return contract.DeployContract(provider, from, abiTestERC20, binTestERC20, name, symbol)
}

// NewTestERC20 creates a new instance of the contract at a specific address
func NewTestERC20(addr web3.Address, provider *jsonrpc.Client) *TestERC20 {
	return &TestERC20{c: contract.NewContract(addr, abiTestERC20, provider)}
}

// Contract returns the contract object
func (_a *TestERC20) Contract() *contract.Contract {
	return _a.c
}

// calls

// Allowance calls the allowance method in the solidity contract
func (_a *TestERC20) Allowance(owner web3.Address, spender web3.Address, block ...web3.BlockNumber) (retval0 *big.Int, err error) {
	var out map[string]interface{}
	_ = out // avoid not used compiler error

	out, err = _a.c.Call("allowance", web3.EncodeBlock(block...), owner, spender)
	if err != nil {
		return
	}

	// decode outputs

	if err = mapstructure.Decode(out["0"], &retval0); err != nil {
		err = fmt.Errorf("failed to encode output at index 0")
	}

	return
}

// BalanceOf calls the balanceOf method in the solidity contract
func (_a *TestERC20) BalanceOf(account web3.Address, block ...web3.BlockNumber) (retval0 *big.Int, err error) {
	var out map[string]interface{}
	_ = out // avoid not used compiler error

	out, err = _a.c.Call("balanceOf", web3.EncodeBlock(block...), account)
	if err != nil {
		return
	}

	// decode outputs

	if err = mapstructure.Decode(out["0"], &retval0); err != nil {
		err = fmt.Errorf("failed to encode output at index 0")
	}

	return
}

// Decimals calls the decimals method in the solidity contract
func (_a *TestERC20) Decimals(block ...web3.BlockNumber) (retval0 uint8, err error) {
	var out map[string]interface{}
	_ = out // avoid not used compiler error

	out, err = _a.c.Call("decimals", web3.EncodeBlock(block...))
	if err != nil {
		return
	}

	// decode outputs

	if err = mapstructure.Decode(out["0"], &retval0); err != nil {
		err = fmt.Errorf("failed to encode output at index 0")
	}

	return
}

// Name calls the name method in the solidity contract
func (_a *TestERC20) Name(block ...web3.BlockNumber) (retval0 string, err error) {
	var out map[string]interface{}
	_ = out // avoid not used compiler error

	out, err = _a.c.Call("name", web3.EncodeBlock(block...))
	if err != nil {
		return
	}

	// decode outputs

	if err = mapstructure.Decode(out["0"], &retval0); err != nil {
		err = fmt.Errorf("failed to encode output at index 0")
	}

	return
}

// Symbol calls the symbol method in the solidity contract
func (_a *TestERC20) Symbol(block ...web3.BlockNumber) (retval0 string, err error) {
	var out map[string]interface{}
	_ = out // avoid not used compiler error

	out, err = _a.c.Call("symbol", web3.EncodeBlock(block...))
	if err != nil {
		return
	}

	// decode outputs

	if err = mapstructure.Decode(out["0"], &retval0); err != nil {
		err = fmt.Errorf("failed to encode output at index 0")
	}

	return
}

// TotalSupply calls the totalSupply method in the solidity contract
func (_a *TestERC20) TotalSupply(block ...web3.BlockNumber) (retval0 *big.Int, err error) {
	var out map[string]interface{}
	_ = out // avoid not used compiler error

	out, err = _a.c.Call("totalSupply", web3.EncodeBlock(block...))
	if err != nil {
		return
	}

	// decode outputs

	if err = mapstructure.Decode(out["0"], &retval0); err != nil {
		err = fmt.Errorf("failed to encode output at index 0")
	}

	return
}

// txns

// Approve sends a approve transaction in the solidity contract
func (_a *TestERC20) Approve(spender web3.Address, amount *big.Int) *contract.Txn {
	return _a.c.Txn("approve", spender, amount)
}

// DecreaseAllowance sends a decreaseAllowance transaction in the solidity contract
func (_a *TestERC20) DecreaseAllowance(spender web3.Address, subtractedValue *big.Int) *contract.Txn {
	return _a.c.Txn("decreaseAllowance", spender, subtractedValue)
}

// IncreaseAllowance sends a increaseAllowance transaction in the solidity contract
func (_a *TestERC20) IncreaseAllowance(spender web3.Address, addedValue *big.Int) *contract.Txn {
	return _a.c.Txn("increaseAllowance", spender, addedValue)
}

// Transfer sends a transfer transaction in the solidity contract
func (_a *TestERC20) Transfer(to web3.Address, amount *big.Int) *contract.Txn {
	return _a.c.Txn("transfer", to, amount)
}

// TransferFrom sends a transferFrom transaction in the solidity contract
func (_a *TestERC20) TransferFrom(from web3.Address, to web3.Address, amount *big.Int) *contract.Txn {
	return _a.c.Txn("transferFrom", from, to, amount)
}

// events

func (_a *TestERC20) ApprovalTopicFilter(owner []web3.Address, spender []web3.Address) [][]web3.Hash {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	var query [][]interface{}
	query = append(query, []interface{}{ApprovalEventID}, ownerRule, spenderRule)

	topics, err := contract.MakeTopics(query...)
	utils.Ensure(err)

	return topics
}

func (_a *TestERC20) FilterApprovalEvent(owner []web3.Address, spender []web3.Address, startBlock uint64, endBlock ...uint64) ([]*ApprovalEvent, error) {
	topic := _a.ApprovalTopicFilter(owner, spender)

	logs, err := _a.c.FilterLogsWithTopic(topic, startBlock, endBlock...)
	if err != nil {
		return nil, err
	}
	res := make([]*ApprovalEvent, 0)
	evts := _a.c.Abi.Events["Approval"]
	for _, log := range logs {
		args, err := evts.ParseLog(log)
		if err != nil {
			return nil, err
		}
		var evtItem ApprovalEvent
		err = json.Unmarshal([]byte(utils.JsonStr(args)), &evtItem)
		if err != nil {
			return nil, err
		}
		evtItem.Raw = log
		res = append(res, &evtItem)
	}
	return res, nil
}

func (_a *TestERC20) TransferTopicFilter(from []web3.Address, to []web3.Address) [][]web3.Hash {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var query [][]interface{}
	query = append(query, []interface{}{TransferEventID}, fromRule, toRule)

	topics, err := contract.MakeTopics(query...)
	utils.Ensure(err)

	return topics
}

func (_a *TestERC20) FilterTransferEvent(from []web3.Address, to []web3.Address, startBlock uint64, endBlock ...uint64) ([]*TransferEvent, error) {
	topic := _a.TransferTopicFilter(from, to)

	logs, err := _a.c.FilterLogsWithTopic(topic, startBlock, endBlock...)
	if err != nil {
		return nil, err
	}
	res := make([]*TransferEvent, 0)
	evts := _a.c.Abi.Events["Transfer"]
	for _, log := range logs {
		args, err := evts.ParseLog(log)
		if err != nil {
			return nil, err
		}
		var evtItem TransferEvent
		err = json.Unmarshal([]byte(utils.JsonStr(args)), &evtItem)
		if err != nil {
			return nil, err
		}
		evtItem.Raw = log
		res = append(res, &evtItem)
	}
	return res, nil
}
