package main

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type tx struct {
	From   common.Address
	To     common.Address
	Amount uint64
	Hash   [32]byte
}

func (tx *tx) toBytes() []byte {
	b, err := json.Marshal(tx)
	if err != nil {
		fmt.Println("error:", err)
	}
	return []byte(b)
}

func (tx *tx) fromBytes(bytes []byte) {
	err := json.Unmarshal(bytes, &tx)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func (tx *tx) hex() string {
	s := hex.EncodeToString(tx.toBytes())
	return s
}

func newTx(from common.Address, to common.Address, amount uint64) tx {
	amountBin := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountBin, amount)
	hashKek := crypto.Keccak256Hash(from.Bytes(), to.Bytes(), amountBin)
	transaction := tx{From: from, To: to, Amount: amount, Hash: hashKek}
	return transaction
}
