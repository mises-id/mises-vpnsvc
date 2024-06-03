package chain

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mises-id/mises-vpnsvc/app/models"
	"github.com/mises-id/mises-vpnsvc/app/models/enum"
	"github.com/mises-id/mises-vpnsvc/config/env"
	"github.com/mises-id/mises-vpnsvc/lib/utils"
	"io"
	"net/http"
)

const BscTestnetGetTransactionsEndPoint = "https://api-testnet.bscscan.com/api?module=account&action=txlist&address=%s&startblock=%d&endblock=99999999&page=1&offset=%d&sort=asc&apikey=%s"

type BscTestnetTransaction struct {
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

type BscTestNet struct{}

func NewBscTestNet() *BscTestNet {
	return &BscTestNet{}
}

func (bt *BscTestNet) VerifyOrders(ctx context.Context, startBlock int64) error {
	if env.Envs.BscApiKey == "" || env.Envs.BscTestContractAddress == "" || env.Envs.BscTestUsdtAddress == "" {
		return errors.New("config error")
	}
	if startBlock <= 0 {
		_, err := models.GetLastBlockNumberFromChain(ctx, enum.ChainIDBscTest)
		if err != nil {
			return err
		}
	}

	return nil
}

func (bt *BscTestNet) GetTransactions(startBlock int64, limit uint64) ([]*BscTestnetTransaction, error) {
	url := fmt.Sprintf(BscTestnetGetTransactionsEndPoint, env.Envs.BscTestContractAddress, startBlock, limit, env.Envs.BscApiKey)
	ret, err := utils.HttpGet(url)
	if err != nil {
		return nil, err
	}
	if len(ret) == 0 {
		return nil, errors.New("no data")
	}
	ts := make([]*BscTestnetTransaction, 0)
	if err := json.Unmarshal(ret, &ts); err != nil {
		return nil, err
	}
	return ts, nil
}
