package tx

import (
	"math/big"	
	
	"github.com/ethereum/go-ethereum/common"
	
	eclient "github.com/Moonyongjung/eth-test-tool/client"
	"github.com/Moonyongjung/eth-test-tool/util"
)

func GetAccountInfo() {
	accountAddress := util.GetConfig().Get("AccountAddress")
	client := eclient.NewRpcClient()

	latestBlockNum, err := client.Client.BlockNumber(client.Ctx)
	if err != nil {
		util.LogTool(err)
		return
	}

	commonAccountAddress := common.HexToAddress(accountAddress)
	latestBlock := big.NewInt(int64(latestBlockNum))	
	balance , err := client.Client.BalanceAt(client.Ctx, commonAccountAddress, latestBlock)
	if err != nil {
		util.LogTool(err)
	}	

	currentNonce, err := client.Client.NonceAt(client.Ctx, commonAccountAddress, nil)
	if err != nil {
		util.LogRpcClient(err)
	}
	util.LogTool("Check block height(latest) : ", latestBlock)
	util.LogTool("Account balance : ", balance)
	util.LogTool("Account nonce : ", currentNonce)	
}