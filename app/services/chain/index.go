package chain

import (
	"context"
	"errors"
	"github.com/mises-id/mises-vpnsvc/app/models/enum"
	pb "github.com/mises-id/mises-vpnsvc/proto"
)

type Chain interface {
	VerifyOrders(ctx context.Context, startBlock int64) error
}

func NewChain(chainID uint64) (Chain, error) {
	switch chainID {
	case enum.ChainIDBscTest:
		return NewBscTestNet(), nil
	case enum.ChainIDBsc:
		return NewBsc(), nil
	default:
		return nil, errors.New("chain error")
	}
}

func VerifyOrderFromChain(ctx context.Context, in *pb.VerifyOrderFromChainRequest) error {
	c, err := NewChain(in.Chain)
	if err != nil {
		return err
	}
	return c.VerifyOrders(ctx, in.StartBlock)
}
