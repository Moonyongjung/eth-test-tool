NODEMODDIR="./node_modules"

if [ ! -d "$NODEMODDIR" ]; then
  npm install ./lib/ethers/ 
fi

go run tool.go