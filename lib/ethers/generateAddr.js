const ethers = require("ethers")
const fs = require("fs")

function generateAdd() {
    var jsondata
    fs.readFile("./config/configKey.json", 'utf8', (err, data) => {
        if (err) {
            console.log(err)
        }       
        jsondata = JSON.parse(data)
        
        const walletMnemonic = ethers.Wallet.fromMnemonic(jsondata.mnemonic).address
        const walletPrivateKey = ethers.Wallet.fromMnemonic(jsondata.mnemonic).privateKey
        console.log(walletMnemonic)
        console.log(walletPrivateKey)
    })        
}

generateAdd()
