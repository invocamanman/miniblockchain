package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
)

// type aPIServer struct {
// 	Router *gin.Engine

// 	B *leveldb.DB

// 	M map[string]int
// }

var router *gin.Engine

var db *leveldb.DB

var m map[string]uint64

func server() {
	m = make(map[string]uint64)
	db := openDb()
	defer db.Close()
	router = gin.Default()
	initializeRoutes()
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initializeRoutes() {
	router.POST("/mint", handlePostMint)
	router.POST("/tx", handlePostTx)
	//router.OPTIONS("/api", handleOptions)
	router.GET("/txs/:addr", handleGetTx)
	router.GET("/balance/:addr", handleGetBalance)

}

func handleGetTx(c *gin.Context) {
	addr := c.Param("addr")
	c.String(http.StatusOK, "hi %s", addr)
}

func handleGetBalance(c *gin.Context) {
	//addr2 := c.Param("addr")
	addr := c.Query("addr")

	//fmt.Println("addr", addr, "addr2", addr2)
	coins := m[addr]
	c.JSON(http.StatusOK, gin.H{"coins": coins})
}

func handlePostMint(c *gin.Context) {

	// coins, err := strconv.ParseUint(c.PostForm("coins"), 10, 64)
	// addr := c.PostForm("addr")

	//mt.Println("coins", coins, "addr", addr, "err,", err)

	var json Params
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m[json.Addr] += json.Coins
	c.JSON(http.StatusOK, gin.H{
		"coins": m[json.Addr],
		"addr":  json.Addr,
	})

}

func handlePostTx(c *gin.Context) {

	// message := c.PostForm("message")
	// nick := c.DefaultPostForm("nick", "anonymous")

	var transactionSend sendTx
	if err := c.ShouldBindJSON(&transactionSend); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	verify := transactionSend.Transaction.verifyTx(transactionSend.PubKey)
	if verify == false {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
	addressFrom := transactionSend.Transaction.From.Hex()
	coinsFrom := m[addressFrom]
	amount := transactionSend.Transaction.Amount

	if coinsFrom < amount {
		c.Status(http.StatusPaymentRequired)
		return
	}

	addressTo := transactionSend.Transaction.To.Hex()
	m[addressFrom] -= amount
	m[addressTo] += amount
	transactionByte, _ := transactionSend.Transaction.toBytes()
	err := db.Put(transactionSend.Transaction.Hash[:], transactionByte, nil)
	if err != nil {
		fmt.Println("error? ", err)
	}
	c.Status(http.StatusOK)
	return

}

// func handleOptions(c *gin.Context) {
// 	c.Header("Allow", "POST, GET, OPTIONS")
// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.Header("Access-Control-Allow-Headers", "origin, content-type, accept")
// 	c.Header("Content-Type", "application/json")
// 	c.Status(http.StatusOK)
// }
