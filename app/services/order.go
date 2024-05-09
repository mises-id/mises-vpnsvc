package services

import (
	"context"
	"fmt"
	"github.com/mises-id/mises-vpnsvc/app/models"
	"github.com/mises-id/mises-vpnsvc/app/models/enum"
	pb "github.com/mises-id/mises-vpnsvc/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func UpdateOrderOnPay(ctx context.Context, in *pb.UpdateOrderRequest) error {
	id, err := primitive.ObjectIDFromHex(in.OrderId)
	if err != nil {
		return err
	}
	order := &models.VpnOrder{
		ID:      id,
		MisesID: in.EthAddress,
		TxnHash: in.TxnHash,
		Status: enum.VpnOrderSuccess,
		OrderAt: time.Now(), // todo:test tmp
	}
	err = order.UpdateOrderOnPayById(ctx)
	if err != nil {
		return err
	}
	return nil
}

func UpdateOrderAndAccount(ctx context.Context, in *pb.UpdateOrderRequest) error {
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// 1. update order
		if err := UpdateOrderOnPay(ctx, in); err != nil {
			return nil, err
		}

		// 2. get order info
		id, _ := primitive.ObjectIDFromHex(in.OrderId)
		order, err := models.FindVpnOrderByID(ctx, in.EthAddress, id)
		if err != nil {
			return nil, err
		}

		// 3. update vpn account
		if err := ModifyVpnAccount(ctx, order); err != nil {
			return nil, fmt.Errorf("modify vpn account: %w", err)
		}

		return nil, nil
	}

	return models.WithTransaction(ctx, callback)
}
