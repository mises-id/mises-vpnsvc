package chain

import (
	"fmt"
	"testing"
)

func TestBscTestNet_GetTransactions(t *testing.T) {
	obj := new(BscTestNet)
	ts, err := obj.GetTransactions(0, 10)
	fmt.Println(err)
	fmt.Println(ts)
}
