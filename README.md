# Eth-test-tool

## Introduction
EVM based chain test tool.
There is a simple solidity contract to test(store and inquiry number).

## Prerequisites
- Go, Geth, nodejs
- Set config file
  - `./config/config.json`
    ```yaml
  {
    "targetUrl":"http://localhost:8545",
    "gasLimit":"58360"  
  }
  ```
  - `targetUrl` : Target chain information
  - `gasLimit` : Transaction gas limit
- Set contract config file
  - `./config/configContract.json`
  ```yaml
  {
    "contractDir":"./rpc/contracts/",
    "contractAbiDir":"contractAbi.json",
    "contractBytecodeDir":"contractBytecode.json",
    "contractAddressDir":"contractAddress.json"
  }
  ```
  - `contractDir` : Default directory of contract
  - `contractAbiDir`, `contractBytecodeDir` : Solidity contract ABI and bytecode file to deploy. (Default test solidity contract ABI & bytecode are set)
  - `contractAddressDir` : After deploying contract, contract address saved file directory automatically
- Set key config file
  - `./config/configKey.json`
  ```yaml
  {
    "mnemonic":""       
  }
  ```
  - `mnemonic` : mnemonic words to crate eth test tool's address and private key

## Usage
1. Set config file
2. ./start.sh
3. If account has no coin, send coin to ETH account address in order to pay gas fee
4. Test

## Functionalities
1. Contract deploy
2. Send invoke transaction
3. Call transaction
4. Get transaction information (by hash)
5. Get block information (by hash or height)
6. Account info
