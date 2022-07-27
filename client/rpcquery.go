package client

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Moonyongjung/eth-test-tool/util"
)

//-- If Queries are used only InitRpcClient, change the number of calling NewRpcClient to once
func QueryBlock() (*types.Block, error) {
	client := NewRpcClient()

	block, err := client.Client.BlockByNumber(client.Ctx, nil)
	if err != nil {
		util.LogRpcClient(err)		
	}

	return block, err
}

func QueryNonce() (uint64, error) {
	client := NewRpcClient()	
	address := util.GetConfig().Get("AccountAddress")
	byte20Address := util.AddressStringToByte20(address)

	currentNonce, err := client.Client.NonceAt(client.Ctx, byte20Address, nil)
	if err != nil {
		util.LogRpcClient(err)
	}

	return currentNonce, err
}

func QueryChainId() (*big.Int, error) {
	client := NewRpcClient()

	chainId, err := client.Client.ChainID(client.Ctx)
	if err != nil {
		util.LogRpcClient(err)
	}

	return chainId, err
}

func QuerySuggestGasPrice() (*big.Int, error) {
	client := NewRpcClient()
	
	gasPrice, err := client.Client.SuggestGasPrice(client.Ctx)
	if err != nil {
		util.LogRpcClient(err)
	}

	return gasPrice, err
}
