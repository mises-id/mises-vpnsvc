package chain
import (
	"context"
)
type Bsc struct {}

func NewBsc() *Bsc {
	return &Bsc{}
}

func (b *Bsc) VerifyOrders (ctx context.Context, startBlock int64) error {
	return nil
}
