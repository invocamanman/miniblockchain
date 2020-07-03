package main

import (
	"encoding/json"
	"fmt"
)

//"encoding/hex"

//account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

func main() {
	db := openDb()
	defer db.Close()

	addressA := generateRandomAddress()
	addressB := generateRandomAddress()
	const amount = uint64(1000)
	transaction := newTx(addressA, addressB, amount)

	fmt.Println("test ", transaction)

	fmt.Println("transactionByte ", transaction)
	transactionByte := transaction.toBytes()
	fmt.Println("transactionByte ", transactionByte)

	transactionByte2, errm := json.Marshal(&transaction)
	fmt.Println("marshal ", transactionByte2, errm)

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
	fmt.Println("transactionFromByte ", transaction2)
}
