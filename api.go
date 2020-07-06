package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
)

// type APIServer struct {
// 	Router *gin.Engine

// 	B *leveldb.DB

// 	M map[string]int
// }

var router *gin.Engine

var db *leveldb.DB

var m map[string]uint64

func server(database *leveldb.DB) {
	m = make(map[string]uint64)
	db = database
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
	addr := c.Param("addr")
	c.String(http.StatusOK, "hi %s", addr)
}

func handlePostMint(c *gin.Context) {

	// message := c.PostForm("message")
	// nick := c.DefaultPostForm("nick", "anonymous")

	var json tx
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if json.User != "manu" || json.Password != "123" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 	return
	// }

	//c.BindJSON(&u)
	// c.JSON(http.StatusOK, gin.H{
	// 	"user": u.Username,
	// 	"pass": u.Password,
	// })
	// }
}

func handlePostTx(c *gin.Context) {

	// message := c.PostForm("message")
	// nick := c.DefaultPostForm("nick", "anonymous")

	var json tx
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if json.User != "manu" || json.Password != "123" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 	return
	// }

	//c.BindJSON(&u)
	// c.JSON(http.StatusOK, gin.H{
	// 	"user": u.Username,
	// 	"pass": u.Password,
	// })
	// }
}

// func handleOptions(c *gin.Context) {
// 	c.Header("Allow", "POST, GET, OPTIONS")
// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.Header("Access-Control-Allow-Headers", "origin, content-type, accept")
// 	c.Header("Content-Type", "application/json")
// 	c.Status(http.StatusOK)
// }
