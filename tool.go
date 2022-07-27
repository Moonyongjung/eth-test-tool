package main

import (
	"os"
	"os/signal"
	"syscall"

	eclient "github.com/Moonyongjung/eth-test-tool/client"
	"github.com/Moonyongjung/eth-test-tool/util"
	tool "github.com/Moonyongjung/eth-test-tool/T"
)

var configPath = "./config/config.json"
var configKeyPath = "./config/configKey.json"
var configContractPath = "./config/configContract.json"

func init() {
	util.GetConfig().Read(configPath)
	util.GetConfig().Read(configKeyPath)
	util.GetConfig().Read(configContractPath)
}

func main() {
	
	eclient.InitRpcClient()

	tool.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	util.LogTool("Shutting down the server...")
	util.LogTool("Server gracefully stopped")	
}

