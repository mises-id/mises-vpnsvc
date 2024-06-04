package models

import (
	"context"
	"fmt"
	"github.com/mises-id/mises-vpnsvc/lib/db"
	"testing"
	"time"
)

func TestGetLastBlockNumberFromChain(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	fmt.Println("setup mongo...")
	db.SetupMongo(ctx)
	EnsureIndex()

	no, err := GetLastBlockNumberFromChain(ctx, 97)

	fmt.Println(err)
	fmt.Println(no)
}
