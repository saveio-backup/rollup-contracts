package rollup

import (
	"testing"

	"github.com/laizy/web3/utils"
	"github.com/laizy/web3/utils/common"
	"github.com/ontology-layer-2/rollup-contracts/deploy"
	"github.com/ontology-layer-2/rollup-contracts/tests/contracts"
	"github.com/stretchr/testify/assert"
)

func TestChainSize(t *testing.T) {
	chainEnv := contracts.LocalL1ChainEnv
	signer := contracts.SetupLocalSigner(chainEnv.ChainId, chainEnv.PrivKey)
	l1Chain := deploy.DeployL1Contracts(signer, chainEnv.ChainConfig)

	size, err := l1Chain.InputChainStorage.ChainSize()
	utils.Ensure(err)
	assert.Equal(t, size, uint64(0))
}

func TestAppend(t *testing.T) {
	chainEnv := contracts.LocalL1ChainEnv
	signer := contracts.SetupLocalSigner(chainEnv.ChainId, chainEnv.PrivKey)
	l1Chain := deploy.DeployL1Contracts(signer, chainEnv.ChainConfig)

	// not owner
	element := common.BytesToHash([]byte("element"))
	receipt := l1Chain.InputChainStorage.Append(element).SetGasLimit(5000000).SetGasPrice(2000).Sign(signer).SendTransaction(signer)
	utils.EnsureTrue(receipt.Status == 0)

	// change storage owner
	l1Chain.AddressManager.SetAddress("RollupInputChain", signer.Address()).Sign(signer).SendTransaction(signer)

	receipt = l1Chain.InputChainStorage.Append(element).SetGasLimit(5000000).SetGasPrice(2000).Sign(signer).SendTransaction(signer)
	utils.EnsureTrue(receipt.Status == 1)
}
