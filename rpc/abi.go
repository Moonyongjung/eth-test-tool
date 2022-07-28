package rpc

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/Moonyongjung/eth-test-tool/util"
)

func GetAbiPack(callName string, val *big.Int) []byte {

	contractAbi, err := EthTestToolMetaData.GetAbi()
	if err != nil {
		util.LogTool("ABI get err : ", err)
	}	

	var abiByteData []byte
	
	if val == nil {
		abiByteData, err = contractAbi.Pack(callName)
		if err != nil {
			util.LogTool("ABI pack err : ", err)
		}
	} else {
		abiByteData, err = contractAbi.Pack(callName, val)
		if err != nil {
			util.LogTool("ABI pack err : ", err)
		}
	}	

	return abiByteData
}

func GetAbiUnpack(callName string, data []byte) []interface{} {
	contractAbi, err := EthTestToolMetaData.GetAbi()
	if err != nil {
		util.LogTool("ABI get err : ", err)
	}

	unpacked, err := contractAbi.Unpack(callName, data)
	if err != nil {
		util.LogTool("ABI unpack err : ", err)
	}

	return unpacked
}

func GetContractAddr() (common.Address, error) {
	contractDir := util.GetConfig().Get("contractDir")
	contractAddressDir := util.GetConfig().Get("contractAddressDir")

	var addressJsonStruct util.ContractAddressStruct
	addressData := util.JsonUnmarshal(addressJsonStruct, contractDir+contractAddressDir)
	address := addressData.(map[string]interface{})["ContractAddress"].(string)
	
	if address == "" {
		return common.Address{}, errors.New("No contract address, contract deploy first")
	}
	
	return util.AddressStringToByte20(address), nil
}