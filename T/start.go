package t

import (
	"fmt"
	"strconv"

	"github.com/Moonyongjung/eth-test-tool/rpc"
	"github.com/Moonyongjung/eth-test-tool/rpc/tx"
	"github.com/Moonyongjung/eth-test-tool/util"
)

func Start() {		
	for{
		var s string

		util.LogTool("=======================================================================")
		util.LogTool("Enter number")
		util.LogTool("1. Contract deploy", "2. Send invoke transaction", "3. Call transaction")
		util.LogTool("4. Get transaction (by hash)", "5. Get Block (by hash or height)       ")
		util.LogTool("6. Account info(eth-tool)", "7. Account info(retrieve account)")
		util.LogTool("8. Send coin(from eth-tool address")
		util.LogTool("=======================================================================")
		fmt.Scan(&s)

		if s == "1" {
			rpc.GolangBindings()
			rpc.DeployEthTestToolContract()	
		} else if s == "2" {
			var val string
			util.LogTool("Input to store value (number)")
			
			fmt.Scan(&val)
			_, err := strconv.Atoi(val)
			if err != nil {
				util.LogTool("Input type is not number")
			} else {
				tx.SendInvokeTx(val)
			}
			
		} else if s == "3" {
			tx.SendCallTx()
		} else if s == "4" {
			var val string
			util.LogTool("Input transaction hash")
			fmt.Scan(&val)

			tx.GetTransaction(val)

		} else if s == "5" {
			var val string
			util.LogTool("Input block hash or number")
			fmt.Scan(&val)

			tx.GetBlock(val)
		} else if s == "6" {
			tx.GetAccountInfo()
		} else if s == "7" {
			var val string
			util.LogTool("Retrieve account")
			fmt.Scan(&val)

			tx.RetrieveAccountInfo(val)
		} else if s == "8" {
			var val string
			util.LogTool("Input address to send coin")
			fmt.Scan(&val)

			var val2 string
			util.LogTool("Input amount")
			fmt.Scan(&val2)
			tx.SendCoinTx(val, val2)
		} else {
			util.LogTool("Input correct number")
		}
	}
	
}