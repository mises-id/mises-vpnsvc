package order

import (
	"context"
	"fmt"
	"github.com/mises-id/mises-vpnsvc/app/models"
	"github.com/mises-id/mises-vpnsvc/app/models/enum"
	"github.com/mises-id/mises-vpnsvc/app/services/vpn"
	pb "github.com/mises-id/mises-vpnsvc/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TransactionDataForOrderUpdate struct {
	OrderId     string
	TxnHash     string
	BlockNumber int64
}

func UpdateOrderOnPay(ctx context.Context, in *TransactionDataForOrderUpdate) error {
	id, err := primitive.ObjectIDFromHex(in.OrderId)
	if err != nil {
		return err
	}
	order := &models.VpnOrder{
		ID:      id,
		TxnHash: in.TxnHash,
		Status: enum.VpnOrderSuccess,
		BlockNumber: in.BlockNumber,
		OrderAt: time.Now(),
	}
	err = order.UpdateOrderOnPayById(ctx)
	if err != nil {
		return err
	}
	return nil
}

func UpdateOrderOnPayForTest(ctx context.Context, in *pb.UpdateOrderRequest) error {
	id, err := primitive.ObjectIDFromHex(in.OrderId)
	if err != nil {
		return err
	}
	order := &models.VpnOrder{
		ID:      id,
		TxnHash: in.TxnHash,
		Status: enum.VpnOrderSuccess,
		MisesID: in.EthAddress,
		OrderAt: time.Now(),
	}
	err = order.UpdateOrderOnPayByIdForTest(ctx)
	if err != nil {
		return err
	}
	return nil
}

func UpdateOrderOnPending(ctx context.Context, in *pb.UpdateOrderRequest) error {
	id, err := primitive.ObjectIDFromHex(in.OrderId)
	if err != nil {
		return err
	}
	order := &models.VpnOrder{
		ID:      id,
		MisesID: in.EthAddress,
		TxnHash: in.TxnHash,
		Status: enum.VpnOrderPending,
	}
	err = order.UpdateOrderOnPendingById(ctx)
	if err != nil {
		return err
	}
	return nil
}

func UpdateOrderAndAccount(ctx context.Context, in *TransactionDataForOrderUpdate) error {
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// 1. update order
		if err := UpdateOrderOnPay(sessCtx, in); err != nil {
			return nil, err
		}

		// 2. get order info
		id, _ := primitive.ObjectIDFromHex(in.OrderId)
		order, err := models.FindVpnOrderByID(sessCtx, id)
		if err != nil {
			return nil, err
		}

		// 3. update vpn account
		if err := vpn.ModifyVpnAccount(sessCtx, order); err != nil {
			return nil, fmt.Errorf("modify vpn account: %w", err)
		}

		return nil, nil
	}

	return models.WithTransaction(ctx, callback)
}

func UpdateOrderAndAccountForTest(ctx context.Context, in *pb.UpdateOrderRequest) error {
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// 1. update order
		if err := UpdateOrderOnPayForTest(sessCtx, in); err != nil {
			return nil, err
		}

		// 2. get order info
		id, _ := primitive.ObjectIDFromHex(in.OrderId)
		order, err := models.FindUserVpnOrderByID(sessCtx, in.EthAddress, id)
		if err != nil {
			return nil, err
		}

		// 3. update vpn account
		if err := vpn.ModifyVpnAccount(sessCtx, order); err != nil {
			return nil, fmt.Errorf("modify vpn account: %w", err)
		}

		return nil, nil
	}

	return models.WithTransaction(ctx, callback)
}
