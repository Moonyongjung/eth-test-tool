package tx

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	eclient "github.com/Moonyongjung/eth-test-tool/client"
	"github.com/Moonyongjung/eth-test-tool/manage"
	"github.com/Moonyongjung/eth-test-tool/util"
	"github.com/Moonyongjung/eth-test-tool/rpc"
)

var invokeName = "store"

func SendInvokeTx(val string) {
	defaultGasLimitStr := util.GetConfig().Get("gasLimit")
	defaultGasLimit, err := util.ToInt(defaultGasLimitStr)
	if err != nil {
		util.LogRpcClient("Gas limit setting error")
		return
	}

	valNum := util.StringToBigInt(val)
	
	client := eclient.NewRpcClient()

	//-- Private key for Transaction signing
	priKey := util.GetConfig().Get("PrivateKey")			
	priKey = priKey[2:]
	priKeyCrypto, _ := crypto.HexToECDSA(priKey)
	_, _, chainId, gasPrice := util.GetConfigRpc().Get()
	gasPrice = util.StringToBigInt("10000000")
	accountNonce := manage.NonceMng().NowNonce()

	byteData := rpc.GetAbiPack(invokeName, valNum)
	contractAddress, err := rpc.GetContractAddr()
	if err != nil {
		util.LogRpcClient(err)
		return
	} else {
		util.LogRpcClient("Target Contract address : " + contractAddress.Hex())

		tx := types.NewTransaction(
			accountNonce, 
			contractAddress, 
			big.NewInt(0),
			uint64(defaultGasLimit),
			gasPrice,
			byteData)	
		
		//-- ING... locl eth test chain id = 15
		// chainId = big.NewInt(15)
		signer := types.NewEIP155Signer(chainId)
	
		//-- Also, the version of blockchain based on ethereum is able to exist which don't use signer.
		//   the below function do not support signer.
		// signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, priKeyCrypto)
		signedTx, err := types.SignTx(tx, signer, priKeyCrypto)
		if err != nil {
			util.LogRpcClient(err)
			return
		}
	
		err = client.Client.SendTransaction(client.Ctx, signedTx)
		if err != nil {
			util.LogRpcClient(err)
			return
		}
	
		util.LogRpcClient("Account nonce : " + util.Uint64ToString(accountNonce))	
		
		//-- Need to increase account nonce after sending transaction.
		manage.NonceMng().AddNonce()
	
		waitTxReceipt(client, signedTx)
	}
}

func waitTxReceipt(client *eclient.BcClient, signedTx *types.Transaction) {
	//-- Count is wait time (sec)
	count := 100
	for {
		receipt, err := client.Client.TransactionReceipt(client.Ctx, signedTx.Hash())
		if err != nil {
			count = count - 1
			if count < 0 {
				util.LogRpcClient("Block not mined, Timeout")
			}
			util.LogRpcClient("Transaction receipt is not arrived yet...")
			time.Sleep(time.Second*1)
		} else {
			util.LogRpcClient("Transaction receipt =============")
			util.LogRpcClient("Status : ", receipt.Status)
			util.LogRpcClient("Transaction Receipt : ", receipt)
			util.LogRpcClient("Transaction Hash : " , receipt.TxHash)
			util.LogRpcClient("Transaction index : ", receipt.TransactionIndex)
			util.LogRpcClient("Type : ", receipt.Type)			
			util.LogRpcClient("Gas Used : " , receipt.GasUsed)
			util.LogRpcClient("Cumulative gas used : ", receipt.CumulativeGasUsed)
			util.LogRpcClient("Block Number : ", receipt.BlockNumber)
			util.LogRpcClient("Block Hash : ", receipt.BlockHash)
			util.LogRpcClient("Receipt logs : ", receipt.Logs)
			
			break
		}
	}
}