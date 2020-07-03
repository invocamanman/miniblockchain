package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestgenerateRandomAddress(t *testing.T) {
	var addressA interface{} = generateRandomAddress()
	assert.Equal(t, true, addressA.(common.Address))
}
