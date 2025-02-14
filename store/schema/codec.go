package schema

import (
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/laizy/web3"
	"github.com/laizy/web3/crypto"
	"github.com/laizy/web3/utils"
	"github.com/laizy/web3/utils/codec"
	"github.com/ontology-layer-2/rollup-contracts/binding"
	"github.com/ontology-layer-2/rollup-contracts/merkle"
)

type AppendedTransaction struct {
	Proposer        web3.Address
	Index           uint64
	StartQueueIndex uint64
	QueueNum        uint64
	InputHash       web3.Hash
}

func (e *AppendedTransaction) Serialization(sink *codec.ZeroCopySink) {
	sink.WriteAddress(e.Proposer)
	sink.WriteUint64(e.Index)
	sink.WriteUint64(e.StartQueueIndex)
	sink.WriteUint64(e.QueueNum)
	sink.WriteHash(e.InputHash)
}

func (e *AppendedTransaction) Deserialization(source *codec.ZeroCopySource) (err error) {
	e.Proposer, err = source.ReadAddress()
	if err != nil {
		return err
	}
	e.Index, err = source.ReadUint64()
	if err != nil {
		return err
	}
	e.StartQueueIndex, err = source.ReadUint64()
	if err != nil {
		return err
	}
	e.QueueNum, err = source.ReadUint64()
	if err != nil {
		return err
	}
	e.InputHash, err = source.ReadHash()
	return err
}

//type StateBatchAppendedEvent struct {
type RollupStateBatchInfo struct {
	Index     uint64
	Proposer  web3.Address
	Timestamp uint64
	BlockHash web3.Hash
}

func (e *RollupStateBatchInfo) Serialization(sink *codec.ZeroCopySink) {
	sink.WriteUint64(e.Index)
	sink.WriteAddress(e.Proposer)
	sink.WriteUint64(e.Timestamp)
	sink.WriteHash(e.BlockHash)
}

func (e *RollupStateBatchInfo) Deserialization(source *codec.ZeroCopySource) (err error) {
	e.Index, err = source.ReadUint64()
	if err != nil {
		return err
	}
	e.Proposer, err = source.ReadAddress()
	if err != nil {
		return err
	}
	e.Timestamp, err = source.ReadUint64()
	if err != nil {
		return err
	}
	e.BlockHash, err = source.ReadHash()
	return err
}

type EnqueuedTransaction struct {
	QueueIndex uint64
	From       web3.Address
	To         web3.Address
	RlpTx      []byte
	Timestamp  uint64
}

func (e *EnqueuedTransaction) Deserialization(source *codec.ZeroCopySource) (err error) {
	e.QueueIndex, err = source.ReadUint64()
	if err != nil {
		return err
	}
	e.From, err = source.ReadAddress()
	if err != nil {
		return err
	}
	e.To, err = source.ReadAddress()
	if err != nil {
		return err
	}
	e.RlpTx, err = source.ReadVarBytes()
	if err != nil {
		return err
	}
	e.Timestamp, err = source.ReadUint64()
	return err
}

func (e *EnqueuedTransaction) Serialization(sink *codec.ZeroCopySink) {
	sink.WriteUint64(e.QueueIndex)
	sink.WriteAddress(e.From)
	sink.WriteAddress(e.To)
	sink.WriteVarBytes(e.RlpTx)
	sink.WriteUint64(e.Timestamp)
}

func (e *EnqueuedTransaction) MustToTransaction() *types.Transaction {
	tx := new(types.Transaction)
	err := tx.UnmarshalBinary(e.RlpTx)
	utils.Ensure(err)
	return tx
}

func EnqueuedTransactionFromEvent(e *binding.TransactionEnqueuedEvent) *EnqueuedTransaction {
	return &EnqueuedTransaction{
		QueueIndex: e.QueueIndex,
		From:       e.From,
		To:         e.To,
		RlpTx:      e.RlpTx,
		Timestamp:  e.Timestamp,
	}
}

type InputChainInfo struct {
	PendingQueueIndex uint64
	TotalBatches      uint64
	QueueSize         uint64
}

func (i *InputChainInfo) Serialization(sink *codec.ZeroCopySink) {
	sink.WriteUint64(i.PendingQueueIndex)
	sink.WriteUint64(i.TotalBatches)
	sink.WriteUint64(i.QueueSize)
}

func (i *InputChainInfo) DeSerialization(source *codec.ZeroCopySource) error {
	reader := source.Reader()
	i.PendingQueueIndex = reader.ReadUint64()
	i.TotalBatches = reader.ReadUint64()
	i.QueueSize = reader.ReadUint64()
	return reader.Error()
}

type StateChainInfo struct {
	TotalSize      uint64
	LastEventBlock uint64
	LastEventIndex uint64
}

func (s *StateChainInfo) Serialization(sink *codec.ZeroCopySink) {
	sink.WriteUint64(s.TotalSize)
	sink.WriteUint64(s.LastEventBlock)
	sink.WriteUint64(s.LastEventIndex)
}

func (s *StateChainInfo) Deserialization(source *codec.ZeroCopySource) (err error) {
	reader := source.Reader()
	s.TotalSize = reader.ReadUint64()
	s.LastEventBlock = reader.ReadUint64()
	s.LastEventIndex = reader.ReadUint64()
	return reader.Error()
}

func CalcQueueHash(queues []*EnqueuedTransaction) web3.Hash {
	return crypto.Keccak256Hash(SerializeEnqueuedTxsInfo(queues))
}

func SerializeEnqueuedTxsInfo(queues []*EnqueuedTransaction) []byte {
	b := codec.NewZeroCopySink(nil)
	for _, queue := range queues {
		txHash := crypto.Keccak256Hash(queue.RlpTx)
		b.WriteHash(txHash).WriteUint64BE(queue.Timestamp)
	}
	return b.Bytes()
}

type CrossLayerSentMessage struct {
	BlockNumber  uint64
	MessageIndex uint64
	Target       web3.Address
	Sender       web3.Address
	MMRRoot      web3.Hash
	Message      []byte
}

func (s *CrossLayerSentMessage) Serialization(sink *codec.ZeroCopySink) {
	sink.WriteUint64(s.BlockNumber)
	sink.WriteUint64(s.MessageIndex)
	sink.WriteAddress(s.Target)
	sink.WriteAddress(s.Sender)
	sink.WriteHash(s.MMRRoot)
	sink.WriteVarBytes(s.Message)
}

func (s *CrossLayerSentMessage) Deserialization(source *codec.ZeroCopySource) (err error) {
	reader := source.Reader()
	s.BlockNumber = reader.ReadUint64()
	s.MessageIndex = reader.ReadUint64()
	s.Target = reader.ReadAddress()
	s.Sender = reader.ReadAddress()
	s.MMRRoot = reader.ReadHash()
	s.Message = reader.ReadVarBytes()
	return reader.Error()
}

type L1TokenBridgeETHEvent struct {
	From   web3.Address
	To     web3.Address
	Amount *big.Int
	Data   []byte
}

func (s *L1TokenBridgeETHEvent) Serialization(sink *codec.ZeroCopySink) {
	sink.WriteAddress(s.From)
	sink.WriteAddress(s.To)
	sink.WriteVarBytes(s.Amount.Bytes())
	sink.WriteVarBytes(s.Data)
}

func (s *L1TokenBridgeETHEvent) Deserialization(source *codec.ZeroCopySource) (err error) {
	reader := source.Reader()
	s.From = reader.ReadAddress()
	s.To = reader.ReadAddress()
	amountData := reader.ReadVarBytes()
	s.Amount = new(big.Int).SetBytes(amountData)
	s.Data = reader.ReadVarBytes()
	return reader.Error()
}

type TokenBridgeERC20Event struct {
	L1Token web3.Address
	L2Token web3.Address
	From    web3.Address
	To      web3.Address
	Amount  *big.Int
	Data    []byte
}

func (s *TokenBridgeERC20Event) Serialization(sink *codec.ZeroCopySink) {
	sink.WriteAddress(s.L1Token)
	sink.WriteAddress(s.L2Token)
	sink.WriteAddress(s.From)
	sink.WriteAddress(s.To)
	sink.WriteVarBytes(s.Amount.Bytes())
	sink.WriteVarBytes(s.Data)
}

func (s *TokenBridgeERC20Event) Deserialization(source *codec.ZeroCopySource) (err error) {
	reader := source.Reader()
	s.L1Token = reader.ReadAddress()
	s.L2Token = reader.ReadAddress()
	s.From = reader.ReadAddress()
	s.To = reader.ReadAddress()
	amountData := reader.ReadVarBytes()
	s.Amount = new(big.Int).SetBytes(amountData)
	s.Data = reader.ReadVarBytes()
	return reader.Error()
}

type L1TokenBridgeETHEvents []*L1TokenBridgeETHEvent

func (s L1TokenBridgeETHEvents) Serialization(sink *codec.ZeroCopySink) {
	sink.WriteInt64(int64(len(s)))
	for _, evt := range s {
		sink.WriteVarBytes(codec.SerializeToBytes(evt))
	}
}

func DeserializationL1TokenBridgeETHEvents(source *codec.ZeroCopySource) (s []*L1TokenBridgeETHEvent, err error) {
	length, eof := source.NextInt64()
	if eof {
		return nil, io.ErrUnexpectedEOF
	}

	for i := int64(0); i < length; i++ {
		data, err := source.ReadVarBytes()
		if err != nil {
			return nil, err
		}
		evt := &L1TokenBridgeETHEvent{}
		err = evt.Deserialization(codec.NewZeroCopySource(data))
		if err != nil {
			return nil, err
		}
		s = append(s, evt)
	}
	return s, nil
}

// this struct used to track all enqueue block
type ChainedEnqueueBlockInfo struct {
	PrevEnqueueBlock uint64 // prev l2 enqueue block
	CurrEnqueueBlock uint64 // curr l2 enqueue blcok
	TotalEnqueuedTx  uint64 // total tx in all l2 enqueue block
}

func (s *ChainedEnqueueBlockInfo) Serialization(sink *codec.ZeroCopySink) {
	sink.WriteUint64(s.PrevEnqueueBlock)
	sink.WriteUint64(s.CurrEnqueueBlock)
	sink.WriteUint64(s.TotalEnqueuedTx)
}

func (s *ChainedEnqueueBlockInfo) Deserialization(source *codec.ZeroCopySource) (err error) {
	reader := source.Reader()
	s.PrevEnqueueBlock = reader.ReadUint64()
	s.CurrEnqueueBlock = reader.ReadUint64()
	s.TotalEnqueuedTx = reader.ReadUint64()
	return reader.Error()
}

func DeserializeCompactMerkleTree(data []byte) (uint64, []web3.Hash, error) {
	source := codec.NewZeroCopySource(data)
	treeSize, err := source.ReadUint64()
	if err != nil {
		return 0, nil, err
	}
	hashCount := (len(data) - 8) / web3.HashLength
	hashes := make([]web3.Hash, 0, hashCount)
	for i := 0; i < hashCount; i++ {
		hash, err := source.ReadHash()
		if err != nil {
			return 0, nil, err
		}
		hashes = append(hashes, hash)
	}
	return treeSize, hashes, nil
}

func SerializeCompactMerkleTree(tree *merkle.CompactMerkleTree) []byte {
	treeSize := tree.TreeSize()
	hashes := tree.Hashes()
	value := codec.NewZeroCopySink(make([]byte, 0, 8+len(hashes)*web3.HashLength))
	value.WriteUint64(treeSize)
	for _, hash := range hashes {
		value.WriteHash(hash)
	}
	return value.Bytes()
}
