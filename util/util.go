package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/logrusorgru/aurora"
)

type ContractAddressStruct struct {
	ContractAddress string
}

type KeyStruct struct {
	AccountAddress string
	PrivateKey string
}

func ParsingAbi(jsonFilePath string) string {
	file, err := os.Open(jsonFilePath)
	if err != nil {
		LogTool(err)
	}
	defer file.Close()

	var abiStrSlice []string

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		abiStrSlice = append(abiStrSlice, str)
		if err == io.EOF {
			break
		}
	}
	str := strings.Join(abiStrSlice, "")	
	str = strings.Replace(str, " ", "", -1)

	return str
}

func JsonUnmarshal(jsonStruct interface{}, jsonFilePath string) interface{} {
	jsonData, err := os.Open(jsonFilePath)
	if err != nil {
		LogTool(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonData)
	jsonStruct = JsonUnmarshalData(jsonStruct, byteValue)

	return jsonStruct
}

func JsonUnmarshalData(jsonStruct interface{}, byteValue []byte) interface{} {
	json.Unmarshal(byteValue, &jsonStruct)

	return jsonStruct
}

func JsonMarshal(jsonData interface{}, jsonFilePath string) {
	byteData, err := JsonMarshalData(jsonData)
	if err != nil {
		LogTool(err)
	}
	err = ioutil.WriteFile(jsonFilePath, byteData, os.FileMode(0644))
	if err != nil {		
		// LogTool("json err ", err)
		path := strings.Split(jsonFilePath, "/")
		pathPop := path[:len(path)-1]
		filePath := strings.Join(pathPop, "/")		

		err := os.Mkdir(filePath, 0755)
		if err != nil {
			LogTool("mkdir err ", err)
		}
		err = ioutil.WriteFile(jsonFilePath, byteData, os.FileMode(0644))
	}
}

func JsonMarshalData(jsonData interface{}) ([]byte, error) {
	byteData, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		LogTool(err)
	}

	return byteData, err
}

func AddressStringToByte20(address string) common.Address {
	var byte20Address common.Address
	byte20Address = common.HexToAddress(address)

	return byte20Address
}

func ToString(value interface{}, defaultValue string) string {
	str := strings.TrimSpace(fmt.Sprintf("%v", value))
	if str == "" {
		return defaultValue
	} else {
		return str
	}
}

func ToInt(value interface{}) (int, error) {
	return strconv.Atoi(value.(string))
}

func Uint64ToString(input uint64) string {
	result := strconv.FormatUint(input, 10)
	return result
}

func StringToBigInt(v string) (*big.Int) {
	n, err := strconv.Atoi(v)
	if err != nil {
		LogTool("Fail to parse big.int : ", v)
	}
	return big.NewInt(int64(n))
}

func StoreContractAddress(contractAddr common.Address) {
	addr := contractAddr.Hex()	
	
	jsonData := ContractAddressStruct{
		ContractAddress: addr, 
	}
	
	f := "./rpc/contracts/contractAddress.json"
	JsonMarshal(jsonData, f)
}

func StoreKey(strList []string) {
	//-- strList[0] : account address
	//   strList[1] : Private key	
	jsonData := KeyStruct{
		AccountAddress: strList[0],
		PrivateKey: strList[1],
	}

	f := "./account/keystore/key.json"
	JsonMarshal(jsonData, f)
}

func ConvertConfigParam(str string) []string {
	var strList []string
	if strings.Contains(str, "mnemonic") {
		str = strings.Replace(str, "\r", "", -1)
		str = strings.Replace(str, "\n", "", -1)
		str = strings.Replace(str, ",", "", -1)
		str = strings.Replace(str, "\"", "", -1)
		strList = strings.Split(str, ":")
		strList[0] = strings.Replace(strList[0], " ", "", -1)
		strList[1] = strings.TrimRight(strList[1], " ")
	} else if strings.Contains(str, "http") {
		str = strings.Replace(str, " ", "", -1)
		str = strings.Replace(str, "\r", "", -1)
		str = strings.Replace(str, "\n", "", -1)
		str = strings.Replace(str, ",", "", -1)
		str = strings.Replace(str, "\"", "", -1)
		strList = strings.Split(str, ":")
		strList[1] = strings.Join(strList[1:], ":")
	} else {
		str = strings.Replace(str, " ", "", -1)
		str = strings.Replace(str, "\r", "", -1)
		str = strings.Replace(str, "\n", "", -1)
		str = strings.Replace(str, ",", "", -1)
		str = strings.Replace(str, "\"", "", -1)
		strList = strings.Split(str, ":")
	}
	
	return strList
}

func LogHttpServer(log ...interface{}) {
	str := ToString(log, "")
	fmt.Println(aurora.Blue("HTTPServer ").String() + str)
}

func LogHttpClient(log ...interface{}) {
	str := ToString(log, "")
	fmt.Println(aurora.Green("HTTPClient ").String() + str)
}

func LogRpcClient(log ...interface{}) {
	str := ToString(log, "")
	fmt.Println(aurora.Brown("RPCClient  ").String() + str)
}

func LogTool(log ...interface{}) {
	str := ToString(log, "")
	fmt.Println(aurora.White("EthTool    ").String() + str)
}