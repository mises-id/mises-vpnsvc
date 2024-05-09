package services

import (
	"context"
	"errors"
	"github.com/mises-id/mises-vpnsvc/app/models"
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
		}else {
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
