package main

import (
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"

	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type tx struct {
	From      common.Address
	To        common.Address
	Amount    uint64
	Hash      [32]byte
	Signature [64]byte
	Nonce     uint64
	Timestamp int64
}

func (tx *tx) hex() string {
	b, _ := tx.toBytes()
	s := hex.EncodeToString(b)
	return s
}

func (tx *tx) signTx(prv *ecdsa.PrivateKey) (err error) {
	signatureS, err := crypto.Sign(tx.Hash[:], prv)
	var signatureA [64]byte
	copy(signatureA[:], signatureS[:])
	tx.Signature = signatureA
	return err
}

func (tx *tx) verifyTx(pubkey []byte) bool {
	verify := crypto.VerifySignature(pubkey, tx.Hash[:], tx.Signature[:])
	return verify
}

// total bytes: 20 + 20 + 8 + 32 = 80 bytes
func (tx *tx) fromBytes(b []byte) error {

	tx.From = common.BytesToAddress(b[:20])
	tx.To = common.BytesToAddress(b[20:40])
	tx.Amount = binary.LittleEndian.Uint64(b[40:48])
	var hash [32]byte
	copy(hash[:], b[48:80])
	tx.Hash = hash
	return nil
}

func (tx *tx) toBytes() ([]byte, error) {

	var b [80]byte
	copy(b[:20], tx.From.Bytes())
	copy(b[20:40], tx.To.Bytes())

	amountBin := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountBin, tx.Amount)

	copy(b[40:48], amountBin[:])
	copy(b[48:80], tx.Hash[:])

	return b[:], nil
}

func newTx(from common.Address, to common.Address, amount uint64) tx {
	amountBin := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountBin, amount)

	var i = uint64(0)
	nonceBin := make([]byte, 8)
	var hashKek [32]byte
	for {
		binary.LittleEndian.PutUint64(nonceBin, i)
		hashKek = crypto.Keccak256Hash(from.Bytes(), to.Bytes(), amountBin, nonceBin)
		if hashKek[0] == byte(0) {
			break
		} else {
			i++
		}
	}

	time := time.Now().Unix()
	transaction := tx{From: from, To: to, Amount: amount, Hash: hashKek, Timestamp: time, Nonce: i}
	return transaction
}
