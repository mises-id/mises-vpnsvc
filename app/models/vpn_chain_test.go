package models

import (
	"context"
	"fmt"
	"github.com/mises-id/mises-vpnsvc/app/models/enum"
	"github.com/mises-id/mises-vpnsvc/lib/db"
	"testing"
	"time"
)

func TestVpnChain_UpsertBlockNumber(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	fmt.Println("setup mongo...")
	db.SetupMongo(ctx)
	EnsureIndex()

	vc := new(VpnChain)
	vc.ChainID = enum.ChainIDBscTest
	vc.LastBlockNumber = 0
	err := vc.UpsertBlockNumber(ctx)
	fmt.Println(err)
}
