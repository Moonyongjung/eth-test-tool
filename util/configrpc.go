package util

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
)

var rpconce sync.Once
var rpcInstance *ConfigRpc

type ConfigRpc struct {
	AccountNonce    uint64
	LatestBlockNum  *types.Block
	ChainId         *big.Int
	SuggestGasPrice *big.Int
}

func GetConfigRpc() *ConfigRpc {
	rpconce.Do(func() {
		rpcInstance = &ConfigRpc{}
	})
	return rpcInstance
}

func (c *ConfigRpc) Set(
	latestBlockNum *types.Block,
	accountNonce uint64,
	chainId *big.Int,
	suggestGasPrice *big.Int) {

	c.LatestBlockNum = latestBlockNum
	c.AccountNonce = accountNonce
	c.ChainId = chainId
	c.SuggestGasPrice = suggestGasPrice
}

func (c *ConfigRpc) Get() (*types.Block, uint64, *big.Int, *big.Int) {
	return c.LatestBlockNum, c.AccountNonce, c.ChainId, c.SuggestGasPrice
}
