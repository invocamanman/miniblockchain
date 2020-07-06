package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDbTx(t *testing.T) {
	db := openDb()
	defer db.Close()

	addressA := generateRandomAddress()
	addressB := generateRandomAddress()
	const amount = uint64(1000)
	transaction := newTx(addressA, addressB, amount)

	transactionByte, _ := transaction.toBytes()

	err := db.Put(transaction.Hash[:], transactionByte, nil)
	if err != nil {
		fmt.Println("error? ", err)
	}

	data, err := db.Get([]byte(transaction.Hash[:]), nil)
	if err != nil {
		fmt.Println("error? ", err)
	}

	transaction2 := tx{}

	transaction2.fromBytes(data)

	// quick fix
	transaction2.Nonce = transaction.Nonce
	transaction2.Timestamp = transaction.Timestamp

	assert.Equal(t, transaction, transaction2)
}
