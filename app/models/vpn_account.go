package models

import (
	"context"
	"errors"
	"github.com/mises-id/mises-vpnsvc/lib/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type (
	VpnAccountStatus int
)

const (
	AccountUnavailable VpnAccountStatus = 0
	AccountAvailable   VpnAccountStatus = 1
)

type VpnAccount struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	MisesID     string             `bson:"misesid"`
	LastOrderId primitive.ObjectID `bson:"last_order_id"`
	Status      VpnAccountStatus   `bson:"status"`
	StartAt     time.Time          `bson:"start_at"`
	EndAt       time.Time          `bson:"end_at"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func (m *VpnAccount) Upsert(ctx context.Context) error {
	t := time.Now()
	filter := bson.D{
		{"misesid", m.MisesID},
	}
	update := bson.M{
		"$setOnInsert": bson.M{
			"misesid":       m.MisesID,
			"status":        AccountAvailable,
			"created_at":    t,
			"updated_at":    t,
		},
		"$set": bson.M{
			"last_order_id": m.LastOrderId,
			"updated_at":    t,
		},
	}
	result, err := db.DB().Collection("vpnaccount").UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	if result.UpsertedCount == 0 {
		return errors.New("vpn account upsert error")
	}
	return nil
}

func (m *VpnAccount) UpdateByLastOrderId(ctx context.Context) error {
	t := time.Now()
	filter := bson.M{
		"misesid": m.MisesID,
		"last_order_id": m.LastOrderId,
	}
	update := bson.M{
		"$set": bson.M{
			"start_at": m.StartAt,
			"end_at":   m.EndAt,
			"update_at": t,
		},
	}
	result, err := db.DB().Collection("vpnaccount").UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("vpn account modify error")
	}
	return nil
}

func FindVpnAccountByLastOrderId(ctx context.Context, misesId string, orderId primitive.ObjectID) (*VpnAccount, error) {
	res := &VpnAccount{}
	result := db.DB().Collection("vpnaccount").FindOne(ctx, &bson.M{
		"misesid":       misesId,
		"last_order_id": orderId,
	})
	if err := result.Err(); err != nil {
		return nil, err
	}
	if err := result.Decode(res); err != nil {
		return nil, err
	}
	return res, nil
}

func FindVpnAccountByMisesId(ctx context.Context, misesId string) (*VpnAccount, error) {
	res := &VpnAccount{}
	result := db.DB().Collection("vpnaccount").FindOne(ctx, &bson.M{
		"misesid":       misesId,
	})
	if err := result.Err(); err != nil {
		return nil, err
	}
	if err := result.Decode(res); err != nil {
		return nil, err
	}
	return res, nil
}
