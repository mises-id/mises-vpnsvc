package services

import (
	"context"
	"errors"
	"github.com/mises-id/mises-vpnsvc/app/models"
	"github.com/mises-id/mises-vpnsvc/app/provider"
	pb "github.com/mises-id/mises-vpnsvc/proto"
	"time"
)

// todo:test get from env

var (
	ServerList = []*pb.GetServerListItem{
		{
			Ip:   "34.230.112.190",
			Name: "Mises Test",
		},
	}
	ServerAddressList = map[string]struct{}{
		"34.230.112.190": {},
	}
)

func ModifyVpnAccount(ctx context.Context, order *models.VpnOrder) error {
	// OrderAt: 支付时间
	// TimeRange: 有效时长

	// StartAt: 当前开始时间
	// EndAt: 当前结束时间

	// 1. set last_order_id
	vpnAccount := &models.VpnAccount{
		MisesID:     order.MisesID,
		LastOrderId: order.ID,
	}
	if err := vpnAccount.Upsert(ctx); err != nil {
		return err
	}

	// 2. 计算新的开始、结束时间
	va, err := models.FindVpnAccountByLastOrderId(ctx, order.MisesID, order.ID)
	if err != nil {
		return err
	}
	if va.StartAt.After(va.EndAt) {
		return errors.New("time range error")
	}

	oldStartAt := va.StartAt
	oldEndAt := va.EndAt

	orderEndAt := order.OrderAt.Add(order.TimeRange)
	if order.OrderAt.Before(va.StartAt) && orderEndAt.After(va.StartAt) {
		va.StartAt = order.OrderAt
		va.EndAt = orderEndAt.Add(va.EndAt.Sub(va.StartAt))
	} else if order.OrderAt.After(va.StartAt) {
		if order.OrderAt.Before(va.EndAt) || order.OrderAt.Equal(va.EndAt) {
			va.EndAt = va.EndAt.Add(order.TimeRange)
		} else {
			va.StartAt = order.OrderAt
			va.EndAt = orderEndAt
		}
	}

	if va.StartAt.Equal(oldStartAt) && va.EndAt.Equal(oldEndAt) {
		// do nothing
		return nil
	}

	// 3. update by last_order_id
	return va.UpdateByLastOrderId(ctx)
}

func CheckVpnAccount(ctx context.Context, misesId string) (*models.VpnAccount, error) {
	va, err := models.FindVpnAccountByMisesId(ctx, misesId)
	if err != nil {
		return nil, err
	}
	if va.Status != models.AccountAvailable {
		return nil, errors.New("account unavailable")
	}
	t := time.Now()
	if !va.EndAt.After(t) {
		return nil, errors.New("subscription expired")
	}
	return va, nil
}

func GetServerList(ctx context.Context, in *pb.GetServerListRequest) ([]*pb.GetServerListItem, error) {
	if _, err := CheckVpnAccount(ctx, in.EthAddress); err != nil {
		return nil, err
	}
	return ServerList, nil
}

func GetServerLink(ctx context.Context, in *pb.GetServerLinkRequest) (string, error) {
	if _, ok := ServerAddressList[in.Server]; !ok {
		return "", errors.New("server error")
	}

	va, err := CheckVpnAccount(ctx, in.EthAddress)
	if err != nil {
		return "", err
	}

	// todo: check本地库的inbound信息

	xui := &provider.MisesXuiClient{}
	return xui.AddInbounds(va.MisesID, va.LastOrderId.Hex(), in.Server, va.EndAt.Unix())

	// todo: 更新本地库的inbound信息
}
