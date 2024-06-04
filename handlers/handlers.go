package handlers

import (
	"context"
	"errors"
	"github.com/mises-id/mises-vpnsvc/app/models"
	"github.com/mises-id/mises-vpnsvc/app/models/enum"
	"github.com/mises-id/mises-vpnsvc/app/services"
	"github.com/mises-id/mises-vpnsvc/app/services/chain"
	"github.com/mises-id/mises-vpnsvc/app/services/order"
	"github.com/mises-id/mises-vpnsvc/app/services/vpn"
	"github.com/mises-id/mises-vpnsvc/lib/utils"
	pb "github.com/mises-id/mises-vpnsvc/proto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.VpnsvcServer {
	return vpnsvcService{}
}

const LAUNCH_TIME int64 = 1714492800

type vpnsvcService struct{}

func (s vpnsvcService) UpdateOrder(ctx context.Context, in *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {

	if !utils.IsEthAddress(in.EthAddress) || in.OrderId == "" || in.TxnHash == "" {
		return nil, errors.New("param error")
	}

	// todo:only for test
	//if err := order.UpdateOrderAndAccountForTest(ctx, in); err != nil {
	//	return nil, err
	//}

	if err := order.UpdateOrderOnPending(ctx, in); err != nil {
		return nil, err
	}

	var resp pb.UpdateOrderResponse
	resp.Code = 0
	resp.Data = &pb.UpdateOrderResult{
		Status: true,
	}
	return &resp, nil
}

func (s vpnsvcService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// todo: 校验order数据
	if _, ok := enum.Chains[in.ChainId]; !ok {
		return nil, errors.New("chain error")
	}

	p, ok := enum.Plans[in.PlanId]
	if !ok {
		return nil, errors.New("plan error")
	}

	cnt, _ := models.CountUserVpnOrdersInTimeRange(ctx, in.EthAddress, 8*time.Hour, enum.VpnOrderPending)
	if cnt > 10 {
		return nil, errors.New("too much pending orders in 8 hours")
	}

	// data
	order := &models.VpnOrder{
		MisesID:     in.EthAddress,
		ChainID:     in.ChainId,
		TokenName:   p.TokenName,
		TokenAmount: p.TokenAmount,
		TimeRange:   p.TimeRange,
	}

	// create
	order, err := models.CreateVpnOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	var resp pb.CreateOrderResponse
	resp.Code = 0
	resp.Data = &pb.CreateOrderResult{
		OrderId: order.ID.Hex(),
	}
	return &resp, nil
}

func (s vpnsvcService) VpnInfo(ctx context.Context, in *pb.VpnInfoRequest) (*pb.VpnInfoResponse, error) {
	// vpninfo
	va, err := models.FindVpnAccountByMisesId(ctx, in.EthAddress)
	if err != nil {
		va = new(models.VpnAccount)
	}
	t := time.Now()
	var status uint64 = 0
	if va.EndAt.After(t) {
		status = 1
	} else if va.EndAt.After(time.Unix(LAUNCH_TIME, 0)) {
		status = 2
	}

	// orders
	orders, err := models.FindVpnOrdersByMisesID(ctx, in.EthAddress)
	if err != nil {
		return nil, err
	}

	// resp
	var resp pb.VpnInfoResponse
	resp.Code = 0
	resp.Data = &pb.VpnInfoResult{
		Status: status,
		Subscription: &pb.Subscription{
			ExpireTime: va.EndAt.Format(time.DateOnly),
		},
	}
	if len(orders) > 0 {
		vo := make([]*pb.VpnOrder, 0, len(orders))
		for _, v := range orders {
			statusText, _ := enum.VpnOrderStatusText[v.Status]
			chainText, _ := enum.Chains[v.ChainID]
			vo = append(vo, &pb.VpnOrder{
				OrderId:    v.ID.Hex(),
				Status:     statusText,
				Amount:     v.TokenAmount,
				Chain:      chainText,
				Token:      v.TokenName,
				TxnHash:    v.TxnHash,
				CreateTime: v.CreatedAt.Format(time.DateTime),
			})
		}
		resp.Data.Orders = vo
	}

	return &resp, nil
}

func (s vpnsvcService) FetchOrderInfo(ctx context.Context, in *pb.FetchOrderInfoRequest) (*pb.FetchOrderInfoResponse, error) {
	// todo:校验

	id, err := primitive.ObjectIDFromHex(in.OrderId)
	if err != nil {
		return nil, errors.New("order id error")
	}
	order, err := models.FindUserVpnOrderByID(ctx, in.EthAddress, id)
	if err != nil {
		return nil, err
	}
	var resp pb.FetchOrderInfoResponse
	statusText, _ := enum.VpnOrderStatusText[order.Status]
	chainText, _ := enum.Chains[order.ChainID]
	resp.Data = &pb.VpnOrder{
		OrderId:    order.ID.Hex(),
		Status:     statusText,
		Amount:     order.TokenAmount,
		Chain:      chainText,
		Token:      order.TokenName,
		TxnHash:    order.TxnHash,
		CreateTime: order.CreatedAt.Format(time.DateTime),
	}
	return &resp, nil
}

func (s vpnsvcService) FetchOrders(ctx context.Context, in *pb.FetchOrdersRequest) (*pb.FetchOrdersResponse, error) {
	// todo:校验

	orders, err := models.FindVpnOrdersByMisesID(ctx, in.EthAddress)
	if err != nil {
		return nil, err
	}

	var resp pb.FetchOrdersResponse
	resp.Code = 0
	if len(orders) > 0 {
		vo := make([]*pb.VpnOrder, 0, len(orders))
		for _, v := range orders {
			statusText, _ := enum.VpnOrderStatusText[v.Status]
			chainText, _ := enum.Chains[v.ChainID]
			vo = append(vo, &pb.VpnOrder{
				OrderId:    v.ID.Hex(),
				Status:     statusText,
				Amount:     v.TokenAmount,
				Chain:      chainText,
				Token:      v.TokenName,
				TxnHash:    v.TxnHash,
				CreateTime: v.CreatedAt.Format(time.DateTime),
			})
		}
		resp.Data = vo
	}
	return &resp, nil
}

func (s vpnsvcService) GetServerList(ctx context.Context, in *pb.GetServerListRequest) (*pb.GetServerListResponse, error) {
	serverList, err := vpn.GetServerList(ctx, in)
	if err != nil {
		logrus.Error("GetServerList error:", err)
		return nil, err
	}
	var resp pb.GetServerListResponse
	resp.Code = 0
	resp.Data = &pb.GetServerListResult{
		Servers: serverList,
	}
	return &resp, nil
}

func (s vpnsvcService) GetServerLink(ctx context.Context, in *pb.GetServerLinkRequest) (*pb.GetServerLinkResponse, error) {
	link, err := vpn.GetServerLink(ctx, in)
	if err != nil {
		logrus.Error("GetServerLink error:", err)
		return nil, err
	}
	var resp pb.GetServerLinkResponse
	resp.Code = 0
	resp.Data = link
	return &resp, nil
}

func (s vpnsvcService) VerifyOrderFromChain(ctx context.Context, in *pb.VerifyOrderFromChainRequest) (*pb.VerifyOrderFromChainResponse, error) {
	if err := chain.VerifyOrderFromChain(ctx, in); err != nil {
		logrus.Error("VerifyOrderFromChain error:", err)
		return nil, err
	}
	var resp pb.VerifyOrderFromChainResponse
	resp.Code = 0
	return &resp, nil
}

func (s vpnsvcService) CleanExpiredVpnLink(ctx context.Context, in *pb.CleanExpiredVpnLinkRequest) (*pb.CleanExpiredVpnLinkResponse, error) {
	// for production
	//go services.CleanExpiredVpnLink(ctx, in)
	err := vpn.CleanExpiredVpnLink(ctx, in)
	if err != nil {
		return nil, err
	}
	var resp pb.CleanExpiredVpnLinkResponse
	resp.Code = 0
	return &resp, nil
}

func (s vpnsvcService) GetVpnConfig(ctx context.Context, in *pb.GetVpnConfigRequest) (*pb.GetVpnConfigResponse, error) {
	config, err := services.GetVpnConfig()
	if err != nil {
		return nil, err
	}
	var resp pb.GetVpnConfigResponse
	resp.Code = 0
	resp.Data = config
	return &resp, nil
}
