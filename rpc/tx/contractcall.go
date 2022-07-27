package tx

import (
	"math/big"	

	"github.com/ethereum/go-ethereum"	

	eclient "github.com/Moonyongjung/eth-test-tool/client"
	"github.com/Moonyongjung/eth-test-tool/util"
	"github.com/Moonyongjung/eth-test-tool/rpc"
)

var callName = "retrieve"

func SendCallTx() {
	//-- Input for Call contract is empty string
	defaultGasLimitStr := util.GetConfig().Get("gasLimit")
	defaultGasLimit, err := util.ToInt(defaultGasLimitStr)
	if err != nil {
		util.LogRpcClient("Gas limit setting error")
		return
	}
		
	client := eclient.NewRpcClient()
	fromAddress := util.GetConfig().Get("AccountAddress")			
	byte20Address := util.AddressStringToByte20(fromAddress)
	_, _, _, gasPrice := util.GetConfigRpc().Get()

	byteData := rpc.GetAbiPack(callName, nil)
	contractAddress, err := rpc.GetContractAddr()
	if err != nil {
		util.LogRpcClient(err)
		return
	} else {
		util.LogRpcClient("Call Contract address : " + contractAddress.Hex())

		msg := ethereum.CallMsg {
			From: byte20Address,
			To: &contractAddress,
			Gas: uint64(defaultGasLimit),
			GasPrice: gasPrice,
			Value: big.NewInt(0),
			Data: byteData,
		}

		res, err := client.Client.CallContract(client.Ctx, msg, nil)			
		if err != nil {
			util.LogRpcClient(err)
			return
		}

		result := rpc.GetAbiUnpack(callName, res)

		util.LogRpcClient("Contract Reponse : ", result)
	}
		
}