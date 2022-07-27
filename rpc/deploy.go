package rpc

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Moonyongjung/eth-test-tool/client"
	"github.com/Moonyongjung/eth-test-tool/manage"
	"github.com/Moonyongjung/eth-test-tool/util"
)

type contractBytecode struct {
	contractBytecode string
}

func DeployEthTestToolContract() {
	//-- New RPC client
	bcClient := client.NewRpcClient()

	//-- Get config	
	address := util.GetConfig().Get("AccountAddress")
	priKey := util.GetConfig().Get("PrivateKey")
	gasLimitStr := util.GetConfig().Get("gasLimit")
	gasLimit, err := util.ToInt(gasLimitStr)
	if err != nil {
		util.LogRpcClient("Gas limit setting error")
		return 
	}
	_, _, chainId, gasPrice := util.GetConfigRpc().Get()

	accountNonce := manage.NonceMng().NowNonce()
	util.LogTool("Deploy smart contract")
	util.LogTool("Account Nonce : ", accountNonce)
	util.LogTool("Chain ID : ", chainId)	
	util.LogTool("Suggested gas price (wei) : ", gasPrice, "[usually 10gwei")		
	
	//-- User address, prikey
	byte20Address := util.AddressStringToByte20(address)
	util.LogTool("Account Address : ", byte20Address.Hex())

	//-- Remove 0x
	priKey = priKey[2:]
	priKeyCrypto, _ := crypto.HexToECDSA(priKey)

	//-- Generate Transactor and set transactor's values
	//-- ING... locl eth test chain id = 15
	// chainId = big.NewInt(15)
	contractAuth, _ := bind.NewKeyedTransactorWithChainID(priKeyCrypto, chainId)

	//-- ING... The below function not using chain ID
	// contractAuth := bind.NewKeyedTransactor(priKeyCrypto)
	contractAuth.Nonce = big.NewInt(int64(accountNonce))
	contractAuth.Value = big.NewInt(0)
	contractAuth.GasLimit = uint64(gasLimit)
	contractAuth.GasPrice = gasPrice
	util.LogTool("Contract Authentication info : ", contractAuth)

	//-- Deploy contract
	ctrtAddr, tx, _, err := DeployEthTestTool(contractAuth, bcClient.Client)
	if err != nil {
		util.LogTool(err)
		return
	}	

	util.LogTool("Deploy Complete")
	util.LogTool("Contract address : ", ctrtAddr)
	util.LogTool("Deploy Transaction : ", tx.Hash())	
	util.LogTool("tx size : ", tx.Size())
	util.LogTool("Gas Limit from chain : ", tx.Gas())
	util.LogTool("Gas Price from chain : ", tx.GasPrice())	
	
	//-- Auto create contract address json file
	util.StoreContractAddress(ctrtAddr)

	manage.NonceMng().AddNonce()
}
