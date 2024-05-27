package chain

import (
	"context"
	"errors"
	pb "github.com/mises-id/mises-vpnsvc/proto"
)

type Chain interface {
	VerifyOrders(startTime int64) error
}

func NewChain(chainName string) (Chain, error) {
	switch chainName {
	case "bsc":
		return NewBsc(), nil
	case "bsc_testnet":
		return NewBscTestNet(), nil
	case "tron":
		return NewTron(), nil
	case "tron_testnet":
		return NewTronTestNet(), nil
	default:
		return nil, errors.New("chain error")
	}
}

func VerifyOrderFromChain(ctx context.Context, in *pb.VerifyOrderFromChainRequest) error {
	c, err := NewChain(in.Chain)
	if err != nil {
		return err
	}
	return c.VerifyOrders(in.StartTime)
}
