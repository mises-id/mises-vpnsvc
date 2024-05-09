package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func EthHash(msg string) (hash common.Hash) {

	data := []byte(msg)
	hash = crypto.Keccak256Hash(data)

	return
}
