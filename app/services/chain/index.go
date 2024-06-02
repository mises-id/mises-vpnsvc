package chain

import (
	"context"
	"errors"
	pb "github.com/mises-id/mises-vpnsvc/proto"
)

type Chain interface {
	VerifyOrders(startBlock int64) error
}

func NewChain(chainName string) (Chain, error) {
	switch chainName {
	case "bsc_testnet":
		return NewBscTestNet(), nil
	case "bsc":
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
	return c.VerifyOrders(in.StartBlock)
}
