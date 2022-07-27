package tx

import (
	"math/big"
	"strconv"
	"strings"
	
	"github.com/ethereum/go-ethereum/common"
	
	eclient "github.com/Moonyongjung/eth-test-tool/client"
	"github.com/Moonyongjung/eth-test-tool/util"
)

func GetBlock(blockVal string) {
	client := eclient.NewRpcClient()

	if strings.Contains(blockVal, "0x") {
		blockHash := blockVal		
		commonBlockHash := common.HexToHash(blockHash)
		block, err := client.Client.BlockByHash(client.Ctx, commonBlockHash)
		if err != nil {
			util.LogRpcClient(err)
			return
		}

		blockHeader, err := block.Header().MarshalJSON()
		if err != nil {
			util.LogRpcClient(err)
			return
		}
		util.LogRpcClient("Block header : ", string(blockHeader))		
		util.LogRpcClient("Block hash : ", block.Hash())
		util.LogRpcClient("Block Number : ", block.Number())
		util.LogRpcClient("Block gas limit : ", block.GasLimit())
		util.LogRpcClient("Block gas used : ", block.GasUsed())
		util.LogRpcClient("Timestamp : ", block.Time())

		txs := block.Body().Transactions
		if len(txs) == 0 {
			util.LogRpcClient("The block has no transaction")
		} else {
			for i, val := range(txs) {
				util.LogRpcClient("Transactions in the block -",i+1,": ", val.Hash())	
			}
		}
		

		uncles := block.Body().Uncles
		if len(uncles) == 0 {
			util.LogRpcClient("The block has no uncle blocks")
		} else {
			for i, val := range(uncles) {
				util.LogRpcClient("Uncles in the block -",i+1,": ", val.Hash())	
			}
		}
			

	} else {
		blockNumber, err := strconv.Atoi(blockVal)
		if err != nil {
			util.LogRpcClient("Input correct number")
			return
		}	

		block, err := client.Client.BlockByNumber(client.Ctx, big.NewInt(int64(blockNumber)))
		if err != nil {
			util.LogRpcClient(err)
			return
		}

		blockHeader, err := block.Header().MarshalJSON()
		if err != nil {
			util.LogRpcClient(err)
			return
		}
		util.LogRpcClient("Block header : ", string(blockHeader))		
		util.LogRpcClient("Block hash : ", block.Hash())
		util.LogRpcClient("Block Number : ", block.Number())
		util.LogRpcClient("Block gas limit : ", block.GasLimit())
		util.LogRpcClient("Block gas used : ", block.GasUsed())
		util.LogRpcClient("Timestamp : ", block.Time())

		txs := block.Body().Transactions
		if len(txs) == 0 {
			util.LogRpcClient("The block has no transaction")
		} else {
			for i, val := range(txs) {
				util.LogRpcClient("Transactions in the block -",i+1,": ", val.Hash())	
			}
		}
		

		uncles := block.Body().Uncles
		if len(uncles) == 0 {
			util.LogRpcClient("The block has no uncle blocks")
		} else {
			for i, val := range(uncles) {
				util.LogRpcClient("Uncles in the block -",i+1,": ", val.Hash())	
			}
		}
	}
}