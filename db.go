package main

import "github.com/syndtr/goleveldb/leveldb"

func openDb() *leveldb.DB {
	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		panic(err)
	}
	return db
}
