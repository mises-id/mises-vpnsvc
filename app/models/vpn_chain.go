package models

import (
	"context"
	"errors"
	"github.com/mises-id/mises-vpnsvc/lib/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type (
	VpnChain struct {
		ID              primitive.ObjectID `bson:"_id,omitempty"`
		ChainID         uint64             `bson:"chain_id"`
		LastBlockNumber int64              `bson:"last_block_number"`
		CreatedAt       time.Time          `bson:"created_at"`
		UpdatedAt       time.Time          `bson:"updated_at"`
	}
)

func (m *VpnChain) UpsertBlockNumber(ctx context.Context) error {
	t := time.Now()
	filter := bson.D{
		{"chain_id", m.ChainID},
	}
	update := bson.M{
		"$setOnInsert": bson.M{
			"chain_id":   m.ChainID,
			"created_at": t,
		},
		"$set": bson.M{
			"updated_at":        t,
			"last_block_number": m.LastBlockNumber,
		},
	}
	_, err := db.DB().Collection("vpnchain").UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

func GetLastBlockNumberFromChain(ctx context.Context, chainID uint64) (int64, error) {
	res := &VpnChain{}
	result := db.DB().Collection("vpnchain").FindOne(ctx, &bson.M{
		"chain_id": chainID,
	})
	err := result.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		return 0, nil
	} else {
		return 0, err
	}
	if err := result.Decode(res); err != nil {
		return 0, err
	}
	return res.LastBlockNumber, nil
}
