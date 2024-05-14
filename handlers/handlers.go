package handlers

import (
	"context"
	"errors"
	"github.com/mises-id/mises-vpnsvc/app/models"
	"github.com/mises-id/mises-vpnsvc/app/models/enum"
	"github.com/mises-id/mises-vpnsvc/app/services"
	pb "github.com/mises-id/mises-vpnsvc/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.VpnsvcServer {
	return vpnsvcService{}
}

type vpnsvcService struct{}

func (s vpnsvcService) UpdateOrder(ctx context.Context, in *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	// todo: 校验order数据

	// todo: 从链上取数据并校验金额

	if err := services.UpdateOrderAndAccount(ctx, in); err != nil {
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

	// todo: 校验24小时订单数，超过限制不予下单

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
	order, err := models.FindVpnOrderByID(ctx, in.EthAddress, id)
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
		resp.Data.Orders = vo
	}
	return &resp, nil
}

func (s vpnsvcService) GetServerList(ctx context.Context, in *pb.GetServerListRequest) (*pb.GetServerListResponse, error) {
	serverList, err := services.GetServerList(ctx, in)
	if err != nil {
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
	link, err := services.GetServerLink(ctx, in)
	if err != nil {
		return nil, err
	}
	var resp pb.GetServerLinkResponse
	resp.Code = 0
	resp.Data = link
	return &resp, nil
}
