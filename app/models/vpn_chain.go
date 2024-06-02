package models

import (
	"context"
	"github.com/mises-id/mises-vpnsvc/lib/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
			"last_block_number": m.LastBlockNumber,
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
