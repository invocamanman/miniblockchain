package main

import (
	"net/http"
	"testing"

	"github.com/dghubble/sling"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

const baseURL = "http://localhost:8080"

type IssueService struct {
	sling *sling.Sling
}

func NewIssueService(httpClient *http.Client) *IssueService {
	return &IssueService{
		sling: sling.New().Client(httpClient).Base(baseURL),
	}
}

// type Params struct {
// 	Coins uint64 `json:"coins,omitempty"  binding:"required" url:"coins,omitempty"`
// 	Addr  string `json:"addr,omitempty"  binding:"required" url:"addr,omitempty"`
// }

func TestMint(t *testing.T) {

	base := sling.New().Base(baseURL)
	path := "mint"

	paramsResponse := new(Params)
	var err string

	body := &Params{
		Coins: 1000,
		Addr:  "0xa0571569334517d77e1a1B03Cb9D345312eC8275",
	}
	base.New().Post(path).BodyJSON(body).Receive(paramsResponse, err) // can be usefull req, error := ..

	paramsResponse2 := new(Params)
	var err2 string
	base.New().Post(path).BodyJSON(body).Receive(paramsResponse2, err2)

	assert.Equal(t, paramsResponse.Coins, uint64(1000))
	assert.Equal(t, paramsResponse2.Coins, uint64(2000))
}

func TestGetBalance(t *testing.T) {

	base := sling.New().Base(baseURL)
	path := "balance/0xa0571569334517d77e1a1B03Cb9D345312eC8275"

	type balanceParam struct {
		Addr  string `json:"addr,omitempty"  binding:"required" url:"addr,omitempty"`
		Coins uint64 `json:"coins,omitempty"  binding:"required" url:"coins,omitempty"`
	}
	params := &balanceParam{Addr: "0xa0571569334517d77e1a1B03Cb9D345312eC8275"}

	paramsResponse := new(balanceParam)
	var errS string
	base.New().Get(path).QueryStruct(params).Receive(paramsResponse, errS)

	//aternative:
	// client := http.Client{}
	// req, err :=base.New().Get(path).QueryStruct(params).Request()
	// response, errC := client.Do(req)
	// body, _ := ioutil.ReadAll(response.Body) //bytes

	assert.Equal(t, uint64(2000), paramsResponse.Coins)
}

func TestHandlePostTx(t *testing.T) {

	privateA := generatePrivateKey()
	pubA := publicFromPrivateKey(privateA)
	addressA := addressFromPrivateKey(privateA)
	addressB := common.HexToAddress("0x9FE6Db844980ac50dc995DCA767963f1317882bF")
	const amount = uint64(1000)
	Tx := newTx(addressA, addressB, amount)
	Tx.signTx(privateA)
	TxSend := sendTx{Transaction: Tx, PubKey: pubA}

	base := sling.New().Base(baseURL)
	path := "tx"

	req, _ := base.New().Post(path).BodyJSON(&TxSend).Request() // can be usefull req, error := ..
	client := http.Client{}
	response, _ := client.Do(req)

	assert.Equal(t, 200, response.StatusCode)
}
