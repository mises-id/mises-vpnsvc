package services

import (
	"errors"
	"github.com/mises-id/mises-vpnsvc/app/models/enum"
	"github.com/mises-id/mises-vpnsvc/app/services/chain"
	"github.com/mises-id/mises-vpnsvc/config/env"
	pb "github.com/mises-id/mises-vpnsvc/proto"
)

func GetVpnConfig() (*pb.GetVpnConfigResult, error) {
	if env.Envs.PriceInUSDT <= 0 {
		return nil, errors.New("config error: price")
	}
	ret := new(pb.GetVpnConfigResult)
	ret.PriceInUsdt = env.Envs.PriceInUSDT
	ret.StrContractABI = chain.StrContractABI
	ret.PurchaseConfigOnChain = []*pb.VpnChain{
		{
			ChainID: enum.ChainIDBscTest,
			TokenAddress: env.Envs.BscTestUsdtAddress,
			ContractAddress: env.Envs.BscTestContractAddress,
		},
	}
	return ret, nil
}
