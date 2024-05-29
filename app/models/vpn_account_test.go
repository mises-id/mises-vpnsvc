package models

import (
	"context"
	"testing"
	"time"
	"github.com/mises-id/mises-vpnsvc/lib/db"
	"fmt"
)

func TestFindVpnAccountByEndTime(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	fmt.Println("setup mongo...")
	db.SetupMongo(ctx)
	EnsureIndex()
	va, err := FindVpnAccountByEndTime(ctx, time.Now(), 100)
	fmt.Println(err)
	fmt.Println(va)
}
