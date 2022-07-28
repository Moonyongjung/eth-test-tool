package rpc

import (
	"io"
	"os"	

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/Moonyongjung/eth-test-tool/util"
)

func GolangBindings() {
	//-- Get config
	contractDir := util.GetConfig().Get("contractDir")
	contractAbiDir := util.GetConfig().Get("contractAbiDir")
	contractBytecodeDir := util.GetConfig().Get("contractBytecodeDir")
	_, _, chainId, _ := util.GetConfigRpc().Get()
	util.LogTool("Target Chain ID : ", chainId)	

	//-- Smart contract bytecode
	var bytecodeJsonStruct contractBytecode
	bytecodeData := util.JsonUnmarshal(bytecodeJsonStruct, contractDir+contractBytecodeDir)
	bytecode := bytecodeData.(map[string]interface{})["contractBytecode"].(string)	

	//-- Smart contract ABI
	ABI := util.ParsingAbi(contractDir + contractAbiDir)	
	
	util.LogTool("ABI : ", ABI)
	types := []string{"ethTestTool"}
	abi := []string{ABI}
	bytecodes := []string{bytecode}

	bind, err := bind.Bind(types, abi, bytecodes, nil, "rpc", bind.LangGo, nil, nil)
	if err != nil {
		util.LogTool(err)
	}	
	f, _ := os.Create("./rpc/form.go")
	_, err = io.WriteString(f, bind)

}
