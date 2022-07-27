package client

import (
	"context"
	"net/http"
	"os"
	

	"github.com/ethereum/go-ethereum/ethclient"
	erpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/Moonyongjung/eth-test-tool/account"
	"github.com/Moonyongjung/eth-test-tool/manage"
	"github.com/Moonyongjung/eth-test-tool/util"
)

var keyStorePath = "./account/keystore/key.json"

type BcClient struct {
	Ctx    context.Context
	Client *ethclient.Client
}

func InitRpcClient() {
	if _, err := os.Stat(keyStorePath); err != nil {
		account.CreateAccountFromMnemonic()
	}
	util.GetConfig().Read(keyStorePath)
	accountAddress := util.GetConfig().Get("AccountAddress")
	block, _ := QueryBlock()
	currentNonce, _ := QueryNonce()
	chainId, _ := QueryChainId()
	suggestGasPrice, _ := QuerySuggestGasPrice()

	util.LogTool("Account address : ", accountAddress)
	util.LogTool("Account nonce : ", currentNonce)

	//-- Save info which are latest block number, account nonce, 
	//   Chain ID and Gas price
	util.GetConfigRpc().Set(block, currentNonce, chainId, suggestGasPrice)
	manage.NonceMng().NewNonce()
}

//-- Connect to Ethereum blockchain
//   For testing, connect to BSC based on ethereum core
func NewRpcClient() *BcClient {
	ctx, _ := context.WithCancel(context.Background())
	// defer cancel()

	//-- Target blockchain node URL
	targetUrl := util.GetConfig().Get("targetUrl")
	util.LogTool("Target URL : ", targetUrl)
	httpDefaultTransport := http.DefaultTransport
	defaultTransportPointer, ok := httpDefaultTransport.(*http.Transport)
	if !ok {
		util.LogRpcClient(ok)
	}
	defaultTransport := *defaultTransportPointer
	defaultTransport.DisableKeepAlives = true

	httpClient := &http.Client{Transport: &defaultTransport}
	rpcClient, err := erpc.DialHTTPWithClient(targetUrl, httpClient)
	if err != nil {
		util.LogRpcClient(err)
	}

	ethClient := ethclient.NewClient(rpcClient)

	return &BcClient{ctx, ethClient}
}
