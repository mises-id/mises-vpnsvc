package models

import (
	"context"
	"time"

	"github.com/mises-id/mises-vpnsvc/lib/db"
	"github.com/mises-id/mises-vpnsvc/lib/db/odm"
	"github.com/mises-id/mises-vpnsvc/lib/pagination"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type (
	IAdminParams interface {
		BuildAdminSearch(chain *odm.DB) *odm.DB
	}
	IAdminPageParams interface {
		BuildAdminSearch(chain *odm.DB) *odm.DB
		GetPageParams() *pagination.TraditionalParams
	}
	IAdminQuickPageParams interface {
		BuildAdminSearch(chain *odm.DB) *odm.DB
		GetQuickPageParams() *pagination.PageQuickParams
	}
)

func EnsureIndex() {

	opts := options.CreateIndexes().SetMaxTime(30 * time.Second)
	trueBool := true

	_, err := db.DB().Collection("vpnorder").Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{
				Key: "chain_id", Value: bsonx.Int32(1),
			}, {
				Key: "txn_hash", Value: bsonx.Int32(1)},
			},
			Options: &options.IndexOptions{
				Unique: &trueBool,
			},
		},
		{
			Keys: bson.M{"misesid": 1},
		},
	}, opts)

	if err != nil {
		logrus.Debug(err)
	}

	_, err = db.DB().Collection("vpnaccount").Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{
				Key: "misesid", Value: bsonx.Int32(1),
			}},
			Options: &options.IndexOptions{
				Unique: &trueBool,
			},
		},
		{
			Keys: bson.M{"end_at": 1},
		},
	}, opts)

	if err != nil {
		logrus.Debug(err)
	}

	_, err = db.DB().Collection("vpnchain").Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{
				Key: "chain_id", Value: bsonx.Int32(1),
			}},
			Options: &options.IndexOptions{
				Unique: &trueBool,
			},
		},
	}, opts)

	if err != nil {
		logrus.Debug(err)
	}

}

func WithTransaction(ctx context.Context, callback func(ctx mongo.SessionContext) (interface{}, error)) (err error) {
	session := mongo.SessionFromContext(ctx)
	if session == nil {
		_session, err := db.Client().StartSession()
		if err != nil {
			return err
		}
		session = _session
	}
	defer session.EndSession(ctx)
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}
	return nil
}
