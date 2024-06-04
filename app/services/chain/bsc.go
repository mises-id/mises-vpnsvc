package chain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mises-id/mises-vpnsvc/app/models"
	"github.com/mises-id/mises-vpnsvc/app/models/enum"
	"github.com/mises-id/mises-vpnsvc/app/services/order"
	"github.com/mises-id/mises-vpnsvc/config/env"
	"github.com/mises-id/mises-vpnsvc/lib/utils"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

const (
	BscTestnetGetTransactionsEndPoint = "https://api-testnet.bscscan.com/api?module=account&action=txlist&address=%s&startblock=%d&endblock=99999999&page=1&offset=%d&sort=asc&apikey=%s"
	BscGetTransactionsEndPoint        = "https://api.bscscan.com/api?module=account&action=txlist&address=%s&startblock=%d&endblock=99999999&page=1&offset=%d&sort=asc&apikey=%s"
)

type BscTransaction struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxReceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
	MethodId          string `json:"methodId"`
	FunctionName      string `json:"functionName"`
}

type BscTransactionsResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Result  []*BscTransaction `json:"result"`
}

type Bsc struct {
	GetTransactionsEndPoint string
	ApiKey                  string
	ContractAddress         string
	TokenAddress            string
	ChainID                 uint64
}

func NewBsc(test bool) (*Bsc, error) {
	if env.Envs.BscApiKey == "" {
		return nil, errors.New("env error: BscApiKey")
	}
	obj := new(Bsc)
	if test {
		if env.Envs.BscTestContractAddress == "" || env.Envs.BscTestUsdtAddress == "" {
			return nil, errors.New("env error: address error")
		}
		obj.GetTransactionsEndPoint = BscTestnetGetTransactionsEndPoint
		obj.ContractAddress = strings.ToLower(env.Envs.BscTestContractAddress)
		obj.TokenAddress = strings.ToLower(env.Envs.BscTestUsdtAddress)
		obj.ChainID = enum.ChainIDBscTest
	} else {
		if env.Envs.BscContractAddress == "" || env.Envs.BscUsdtAddress == "" {
			return nil, errors.New("env error: address error")
		}
		obj.GetTransactionsEndPoint = BscGetTransactionsEndPoint
		obj.ContractAddress = strings.ToLower(env.Envs.BscContractAddress)
		obj.TokenAddress = strings.ToLower(env.Envs.BscUsdtAddress)
		obj.ChainID = enum.ChainIDBsc
	}
	obj.ApiKey = env.Envs.BscApiKey
	return obj, nil
}

func (bt *Bsc) VerifyOrders(ctx context.Context, startBlock int64) error {
	// startBlock
	if startBlock <= 0 {
		lastBlock, err := models.GetLastBlockNumberFromChain(ctx, enum.ChainIDBscTest)
		if err != nil {
			return err
		}
		startBlock = lastBlock
	}

	// get transactions from chai
	rs, err := bt.GetTransactions(startBlock, 100)
	if err != nil {
		return err
	}
	if len(rs) == 0 {
		logrus.Info("no new transactions")
		return nil
	}

	// filter
	ups, err := bt.FilterTransactions(rs)
	if err != nil {
		return err
	}
	if len(ups) == 0 {
		logrus.Info("no successful transactions")
		return nil
	}

	// update db
	for _, v := range ups {
		startBlock = v.BlockNumber
		// log error
		if err := order.UpdateOrderAndAccount(ctx, v); err != nil {
			logrus.Error("UpdateOrderAndAccount failed:", v, err)
			continue
		}
	}

	// update last block number
	vc := new(models.VpnChain)
	vc.ChainID = bt.ChainID
	vc.LastBlockNumber = startBlock
	if err := vc.UpsertBlockNumber(ctx); err != nil {
		logrus.Error("UpsertBlockNumber failed:", err)
	}

	return nil
}

func (bt *Bsc) FilterTransactions(rs []*BscTransaction) ([]*order.TransactionDataForOrderUpdate, error) {
	if len(rs) == 0 {
		return nil, nil
	}
	ret := make([]*order.TransactionDataForOrderUpdate, 0, len(rs))
	for _, v := range rs {
		if v.From == "" || v.Hash == "" {
			continue
		}
		if v.IsError != "0" {
			continue
		}
		if v.TxReceiptStatus != "1" {
			continue
		}
		if strings.ToLower(v.To) != bt.ContractAddress {
			continue
		}
		if v.FunctionName != receiverFunction {
			continue
		}
		if v.Input == "" {
			continue
		}
		if v.BlockNumber == "" {
			continue
		}
		blockNumber, err := strconv.ParseInt(v.BlockNumber, 10, 64)
		if err != nil {
			logrus.Errorf("txn:%s, block number error: %s", v.Hash, v.BlockNumber)
			continue
		}
		dec, err := DecodeTransactionInput(v.Input)
		if err != nil {
			logrus.Errorf("txn:%s, decode hex string %s error: %v", v.Hash, v.Input, err)
			continue
		}
		if strings.ToLower(dec.TokenAddress) != bt.TokenAddress {
			logrus.Errorf("txn:%s, token address error: %s", v.Hash, dec.TokenAddress)
			continue
		}
		if dec.Amount.Cmp(priceInUSDTWithExp) != 0 {
			logrus.Errorf("txn:%s, amount error: %v", v.Hash, dec.Amount)
			continue
		}
		if dec.UniqueKey == "" {
			logrus.Errorf("txn:%s, empty unique key", v.Hash)
			continue
		}
		ret = append(ret, &order.TransactionDataForOrderUpdate{
			OrderId:     dec.UniqueKey,
			TxnHash:     v.Hash,
			BlockNumber: blockNumber,
		})
	}
	return ret, nil
}

func (bt *Bsc) GetTransactions(startBlock int64, limit uint64) ([]*BscTransaction, error) {
	url := fmt.Sprintf(bt.GetTransactionsEndPoint, bt.ContractAddress, startBlock, limit, bt.ApiKey)
	ret, err := utils.HttpGet(url)
	if err != nil {
		return nil, err
	}
	if len(ret) == 0 {
		return nil, errors.New("no data")
	}
	var resp BscTransactionsResponse
	if err := json.Unmarshal(ret, &resp); err != nil {
		return nil, err
	}
	if resp.Status != "1" {
		return nil, errors.New("resp status error")
	}
	return resp.Result, nil
}
