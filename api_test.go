package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
	addressA := common.HexToAddress("0xa0571569334517d77e1a1B03Cb9D345312eC8275")
	addressB := common.HexToAddress("0x9FE6Db844980ac50dc995DCA767963f1317882bF")
	const amount = uint64(1000)
	Tx := newTx(addressA, addressB, amount)
	assert.Equal(t, addressA, Tx.From)
	assert.Equal(t, addressB, Tx.To)
	assert.Equal(t, amount, Tx.Amount)
	assert.Equal(t, "0x0093bec823ce02d25563d72695d53ce2f31a569902717d95f228b34e09bf28a7", common.Hash(Tx.Hash).Hex())
	assert.Equal(t, "a0571569334517d77e1a1b03cb9d345312ec82759fe6db844980ac50dc995dca767963f1317882bfe8030000000000000093bec823ce02d25563d72695d53ce2f31a569902717d95f228b34e09bf28a7", Tx.hex())
}
