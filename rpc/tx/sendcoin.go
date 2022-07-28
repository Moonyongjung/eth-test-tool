package tx

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	eclient "github.com/Moonyongjung/eth-test-tool/client"
	"github.com/Moonyongjung/eth-test-tool/manage"
	"github.com/Moonyongjung/eth-test-tool/util"	
)

func SendCoinTx(toAccount string, amount string) {
	defaultGasLimitStr := util.GetConfig().Get("gasLimit")
	defaultGasLimit, err := util.ToInt(defaultGasLimitStr)
	if err != nil {
		util.LogRpcClient("Gas limit setting error")
		return
	}	
	amountNumber := util.StringToBigInt(amount)
	client := eclient.NewRpcClient()

	//-- Private key for Transaction signing
	priKey := util.GetConfig().Get("PrivateKey")			
	priKey = priKey[2:]
	priKeyCrypto, _ := crypto.HexToECDSA(priKey)
	_, _, chainId, gasPrice := util.GetConfigRpc().Get()
	// gasPrice = util.StringToBigInt("10000000")
	accountNonce := manage.NonceMng().NowNonce()

	toAccountConvert := util.AddressStringToByte20(toAccount)
	
	
	util.LogRpcClient("To address : " + toAccountConvert.Hex())

	tx := types.NewTransaction(
		accountNonce, 
		toAccountConvert, 
		amountNumber,
		uint64(defaultGasLimit),
		gasPrice,
		nil)	
	
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