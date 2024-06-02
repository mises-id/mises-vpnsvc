package models

import (
	"context"
	"fmt"
	"github.com/mises-id/mises-vpnsvc/app/models/enum"
	"github.com/mises-id/mises-vpnsvc/lib/db"
	"testing"
	"time"
)

func TestCountUserVpnOrdersInTimeRange(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	fmt.Println("setup mongo...")
	db.SetupMongo(ctx)
	EnsureIndex()
	cnt, err := CountUserVpnOrdersInTimeRange(ctx, "0x3836f698d4e7d7249ccc3291d9ccd608ee718988", 8 * time.Hour, enum.VpnOrderPending)
	fmt.Println(err)
	fmt.Println(cnt)
}
