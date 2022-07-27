package manage

import (
	"sync"

	"github.com/Moonyongjung/eth-test-tool/util"
)

var nonceInstance *NonceStruct
var nonceOnce sync.Once

//-- Manage account nonce
type NonceStruct struct {
	Nonce uint64
}

func NonceMng() *NonceStruct {
	nonceOnce.Do( func() {
		nonceInstance = &NonceStruct{}
	})
	return nonceInstance
}

func (n *NonceStruct) NewNonce() {
	_, accountNonce, _, _ := util.GetConfigRpc().Get()
	n.Nonce = accountNonce
}

func (n *NonceStruct) NowNonce() uint64{
	return n.Nonce
}

func (n *NonceStruct) AddNonce() {
	nonce := n.Nonce
	nonce = nonce + 1
	n.Nonce = nonce
}