package tx

import (	
	"github.com/ethereum/go-ethereum/common"

	eclient "github.com/Moonyongjung/eth-test-tool/client"	
	"github.com/Moonyongjung/eth-test-tool/util"
)

func GetTransaction(txhash string) {
	client := eclient.NewRpcClient()

	commonTxHash := common.HexToHash(txhash)
	tx, isPending, err := client.Client.TransactionByHash(client.Ctx, commonTxHash)
	if err != nil {
		util.LogRpcClient(err)
		return
	} 

	receipt, err := client.Client.TransactionReceipt(client.Ctx, commonTxHash)
	if err != nil {
		util.LogTool(err)
		return
	}

	json, err := tx.MarshalJSON()
	if err != nil {
		util.LogTool(err)
		return
	}		
	
	rjson, err := receipt.MarshalJSON()
	if err != nil {
		util.LogTool(err)
		return
	}
	
	v, r, s := tx.RawSignatureValues()
	util.LogRpcClient("Get transaction JSON : ", string(json))
	util.LogRpcClient("Nonce : ", tx.Nonce())
	util.LogRpcClient("Gas Price : ", tx.GasPrice())
	util.LogRpcClient("Gas Limit : ", tx.Gas())
	util.LogRpcClient("To : ", tx.To())
	util.LogRpcClient("Chain ID : ", tx.ChainId())
	util.LogRpcClient("is pending? : ", isPending)
	util.LogRpcClient("ECDSA signature output (v, r) : ", v, ", ", r)
	util.LogRpcClient("Recovery ID for Publid key (s) : ", s)
	util.LogRpcClient("Tx hash : ", tx.Hash())
	
	util.LogRpcClient("=======================================================================")
	util.LogRpcClient("Get transaction receipt : ", string(rjson))
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
	
}