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

// total bytes: 20 + 20 + 8 + 32 = 80 bytes
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

// total bytes: 20 + 20 + 8 + 32 = 80 bytes

func (tx *tx) UnmarshalJSON(b []byte) error {

	tx.From = common.BytesToAddress(b[:20])
	tx.To = common.BytesToAddress(b[20:40])
	tx.Amount = binary.LittleEndian.Uint64(b[40:48])
	var hash [32]byte
	copy(hash[:], b[48:80])
	tx.Hash = hash
	return nil
}

func (tx *tx) MarshalJSON() ([]byte, error) {

	var b [80]byte
	copy(b[:20], tx.From.Bytes())
	copy(b[20:40], tx.To.Bytes())

	amountBin := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountBin, tx.Amount)

	copy(b[40:48], amountBin[:])
	copy(b[48:80], tx.Hash[:])
	fmt.Println("b \n", b)

	// b, err := json.Marshal(b)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	return b[:], nil
}

func newTx(from common.Address, to common.Address, amount uint64) tx {
	amountBin := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountBin, amount)
	hashKek := crypto.Keccak256Hash(from.Bytes(), to.Bytes(), amountBin)
	transaction := tx{From: from, To: to, Amount: amount, Hash: hashKek}
	return transaction
}
