package account

import (
	"os/exec"
	"strings"

	"github.com/Moonyongjung/eth-test-tool/util"
)

func CreateAccountFromMnemonic() {
	var strList []string
	cmd := exec.Command("node", "./lib/ethers/generateAddr.js")
	output, err := cmd.Output()
	if err != nil {
		util.LogRpcClient(err)
		return
	}
	strOutput := string(output)
	strList = strings.Split(strOutput, "\n")
	util.LogTool("Eth account address : ", strList[0])
	
	util.StoreKey(strList)	
}