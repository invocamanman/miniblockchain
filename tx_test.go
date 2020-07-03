package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestNewTx(t *testing.T) {
	addressA := common.HexToAddress("0xa0571569334517d77e1a1B03Cb9D345312eC8275")
	addressB := common.HexToAddress("0x9FE6Db844980ac50dc995DCA767963f1317882bF")
	const amount = uint64(1000)
	Tx := newTx(addressA, addressB, amount)
	assert.Equal(t, addressA, Tx.From)
	assert.Equal(t, addressB, Tx.To)
	assert.Equal(t, amount, Tx.Amount)
	assert.Equal(t, "0x7a2f1fdf8ca41d2b711bb910edc277efaf757dfde6256a7d87149b01df9e84e2", common.Hash(Tx.Hash).Hex())
	assert.Equal(t, "7b2246726f6d223a22307861303537313536393333343531376437376531613162303363623964333435333132656338323735222c22546f223a22307839666536646238343439383061633530646339393564636137363739363366313331373838326266222c22416d6f756e74223a313030302c2248617368223a5b3132322c34372c33312c3232332c3134302c3136342c32392c34332c3131332c32372c3138352c31362c3233372c3139342c3131392c3233392c3137352c3131372c3132352c3235332c3233302c33372c3130362c3132352c3133352c32302c3135352c312c3232332c3135382c3133322c3232365d7d", Tx.hex())
}

func TestMarshalTx(t *testing.T) {
	addressA := common.HexToAddress("0xa0571569334517d77e1a1B03Cb9D345312eC8275")
	addressB := common.HexToAddress("0x9FE6Db844980ac50dc995DCA767963f1317882bF")
	const amount = uint64(1000)
	Tx := newTx(addressA, addressB, amount)

	b, err := json.Marshal(&Tx)

	fmt.Println("marshal:", string(b))

	Tx2 := tx{}
	err2 := json.Unmarshal(b, &Tx2)

	assert.Equal(t, nil, string(b))

	assert.Equal(t, nil, err)
	assert.Equal(t, nil, err2)
	//assert.Equal(t, "0x7a2f1fdf8ca41d2b711bb910edc277efaf757dfde6256a7d87149b01df9e84e2", string(b))
	assert.Equal(t, Tx, Tx2)
}
