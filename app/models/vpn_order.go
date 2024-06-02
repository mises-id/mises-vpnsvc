package models

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"github.com/google/uuid"
	"github.com/mises-id/mises-vpnsvc/app/models/enum"
	"github.com/mises-id/mises-vpnsvc/lib/db"
	"github.com/mises-id/mises-vpnsvc/lib/pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	NativeTokenAddress = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
)

type (
	Transaction struct {
		Timestamp        string `json:"timestamp"`
		BlockHash        string `json:"blockHash" bson:"blockHash"`
		BlockNumber      string `json:"blockNumber" bson:"blockNumber"`
		From             string `json:"from" bson:"from"`
		Gas              string `json:"gas" bson:"gas"`
		GasPrice         string `json:"gasPrice" bson:"gasPrice"`
		GasUsed          string `json:"gasUsed" bson:"gasUsed"`
		Hash             string `json:"hash" bson:"hash"`
		TxreceiptStatus  string `json:"txreceipt_status" bson:"txreceipt_status"`
		Input            string `json:"input" bson:"input"`
		Nonce            string `json:"nonce" bson:"nonce"`
		To               string `json:"to" bson:"to"`
		Type             string `json:"type" bson:"type"`
		TransactionIndex string `json:"transactionIndex" bson:"transactionIndex"`
		Value            string `json:"value" bson:"value"`
	}

	TxDecodedLogParams struct {
		Name  string `json:"name"`
		Value string `json:"value"`
		Type  string `json:"type"`
	}

	TxDecodedLog struct {
		Address      string        `json:"address"`
		DecodedEvent *DecodedEvent `json:"decoded_event"`
	}

	DecodedEvent struct {
		Label  string                `json:"label"`
		Params []*TxDecodedLogParams `json:"params"`
	}

	TransactionDecodedReceipt struct {
		Status         string          `json:"receipt_status"`
		BlockTimestamp string          `json:"block_timestamp"`
		Logs           []*TxDecodedLog `json:"logs"`
		Message        string          `json:"message"`
	}

	BlockData struct {
		Number       string         `json:"number"`
		Hash         string         `json:"hash"`
		Nonce        string         `json:"nonce"`
		Timestamp    string         `json:"timestamp"`
		Transactions []*Transaction `json:"transactions"`
	}

	VpnOrder struct {
		ID              primitive.ObjectID  `bson:"_id,omitempty"`
		MisesID         string              `bson:"misesid"`
		TokenAmount    string              `bson:"token_account"`
		TokenName       string              `bson:"token_name"`
		TimeRange       time.Duration       `bson:"time_range"`
		ChainID         uint64              `bson:"chain_id"`
		TxnHash         string              `bson:"txn_hash"`
		ContractAddress string              `bson:"contract_address"`
		FromAddress     string              `bson:"from_address"`
		Status          enum.VpnOrderStatus `bson:"status"`
		Transaction     *Transaction        `bson:"tx,omitempty"`
		OrderAt         time.Time          `bson:"order_at,omitempty"`
		BlockAt         *time.Time          `bson:"block_at,omitempty"`
		UpdatedAt       time.Time           `bson:"updated_at"`
		CreatedAt       time.Time           `bson:"created_at"`
	}
)

func (u *VpnOrder) BeforeCreate(ctx context.Context) error {
	if u.MisesID == "" {
		return errors.New("user error")
	}
	if u.ChainID == 0 {
		return errors.New("chain error")
	}
	if u.TxnHash == "" {
		u.TxnHash = "tmp:" + uuid.New().String()
	}
	u.CreatedAt = time.Now()
	return u.BeforeUpdate(ctx, u.CreatedAt)
}

func (u *VpnOrder) BeforeUpdate(ctx context.Context, ts time.Time) error {
	u.UpdatedAt = ts
	//u.FromAddress = strings.ToLower(u.FromAddress)
	//u.ContractAddress = strings.ToLower(u.ContractAddress)
	//if u.Transaction != nil {
	//	u.Transaction.Hash = strings.ToLower(u.Transaction.Hash)
	//}
	//if u.Transaction != nil {
	//	u.Transaction.From = strings.ToLower(u.Transaction.From)
	//	u.Transaction.To = strings.ToLower(u.Transaction.To)
	//}

	return nil
}

func CreateVpnOrder(ctx context.Context, data *VpnOrder) (*VpnOrder, error) {

	if err := data.BeforeCreate(ctx); err != nil {
		return nil, err
	}
	res, err := db.DB().Collection("vpnorder").InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	data.ID = res.InsertedID.(primitive.ObjectID)
	return data, nil
}

func ListVpnOrder(ctx context.Context, params IAdminParams) ([]*VpnOrder, error) {
	res := make([]*VpnOrder, 0)
	chain := params.BuildAdminSearch(db.ODM(ctx))
	err := chain.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func PageVpnOrder(ctx context.Context, params IAdminPageParams) ([]*VpnOrder, pagination.Pagination, error) {
	out := make([]*VpnOrder, 0)
	chain := params.BuildAdminSearch(db.ODM(ctx))
	pageParams := params.GetPageParams()
	paginator := pagination.NewTraditionalPaginator(pageParams.PageNum, pageParams.PageSize, chain)
	page, err := paginator.Paginate(&out)
	if err != nil {
		return nil, nil, err
	}
	return out, page, nil
}

func FindVpnOrderByID(ctx context.Context, misesId string, orderId primitive.ObjectID) (*VpnOrder, error) {
	res := &VpnOrder{}
	result := db.DB().Collection("vpnorder").FindOne(ctx, &bson.M{
		"misesid": misesId,
		"_id":     orderId,
	})
	if err := result.Err(); err != nil {
		return nil, err
	}
	if err := result.Decode(res); err != nil {
		return nil, err
	}
	return res, nil
}

func FindVpnOrdersByMisesID(ctx context.Context, misesId string) ([]*VpnOrder, error) {
	opts := options.Find().SetSort(bson.D{{"created_at", -1}})
	cursor, err := db.DB().Collection("vpnorder").Find(ctx, &bson.M{
		"misesid": misesId,
	}, opts)
	if err != nil {
		return nil, err
	}
	var results []*VpnOrder
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func CountUserVpnOrdersInTimeRange(ctx context.Context, misesId string, timeRange time.Duration, status enum.VpnOrderStatus) (int64, error) {
	cnt, err := db.DB().Collection("vpnorder").CountDocuments(ctx, &bson.M{
		"misesid": misesId,
		"created_at": bson.M{
			"$gt": time.Now().Add(-1 * timeRange),
		},
		"status": enum.VpnOrderPending,
	})
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func (u *VpnOrder) UpdateOrderOnPayById(ctx context.Context) error {
	update := bson.M{}
	update["txn_hash"] = u.TxnHash
	update["status"] = u.Status
	update["order_at"] = u.OrderAt
	currentStatus := []enum.VpnOrderStatus{enum.VpnOrderInit, enum.VpnOrderPending}
	ret, err := db.DB().Collection("vpnorder").UpdateOne(ctx, &bson.M{"_id": u.ID, "misesid": u.MisesID, "status": bson.M{"$in": currentStatus}}, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return err
	}
	if ret.ModifiedCount == 0 {
		return errors.New("update order error")
	}
	return nil
}

func (u *VpnOrder) UpdateOrderOnPendingById(ctx context.Context) error {
	update := bson.M{}
	update["txn_hash"] = u.TxnHash
	update["status"] = u.Status
	ret, err := db.DB().Collection("vpnorder").UpdateOne(ctx, &bson.M{"_id": u.ID, "misesid": u.MisesID, "status": enum.VpnOrderInit}, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return err
	}
	if ret.ModifiedCount == 0 {
		return errors.New("update order error")
	}
	return nil
}
