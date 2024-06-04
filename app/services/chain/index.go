package chain

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/mises-id/mises-vpnsvc/app/models/enum"
	"github.com/mises-id/mises-vpnsvc/config/env"
	pb "github.com/mises-id/mises-vpnsvc/proto"
	"github.com/sirupsen/logrus"
	"math/big"
	"strings"
)

const (
	StrContractABI = `[
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "initialOwner",
        "type": "address"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "owner",
        "type": "address"
      }
    ],
    "name": "OwnableInvalidOwner",
    "type": "error"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "account",
        "type": "address"
      }
    ],
    "name": "OwnableUnauthorizedAccount",
    "type": "error"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "previousOwner",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "newOwner",
        "type": "address"
      }
    ],
    "name": "OwnershipTransferred",
    "type": "event"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "tokenAddress",
        "type": "address"
      }
    ],
    "name": "getTokenBalance",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "owner",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "tokenAddress",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      },
      {
        "internalType": "string",
        "name": "uniqueKey",
        "type": "string"
      }
    ],
    "name": "receiveTokens",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "renounceOwnership",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "newOwner",
        "type": "address"
      }
    ],
    "name": "transferOwnership",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "tokenAddress",
        "type": "address"
      }
    ],
    "name": "withdrawToken",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
]`
	receiverFunction = "receiveTokens(address _tokenAddress,uint256 _amount,string _orderHash)"
)

var (
	contractABI        *abi.ABI
	priceInUSDTWithExp *big.Int
)

type TransactionInput struct {
	Amount       *big.Int
	TokenAddress string
	UniqueKey    string
}

type Chain interface {
	VerifyOrders(ctx context.Context, startBlock int64) error
}

func init() {
	_contractABI, err := abi.JSON(strings.NewReader(StrContractABI))
	if err != nil {
		logrus.Error("contractABI error")
		return
	}
	contractABI = &_contractABI

	if env.Envs.PriceInUSDT <= 0 {
		logrus.Error("env PriceInUSDT error")
		return
	}
	powerOfTen := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	bigPriceInUSDT := new(big.Int).SetInt64(env.Envs.PriceInUSDT)
	priceInUSDTWithExp = new(big.Int).Mul(bigPriceInUSDT, powerOfTen)
}

func NewChain(chainID uint64) (Chain, error) {
	switch chainID {
	case enum.ChainIDBscTest:
		return NewBsc(true)
	case enum.ChainIDBsc:
		return NewBsc(false)
	default:
		return nil, errors.New("chain error")
	}
}

func VerifyOrderFromChain(ctx context.Context, in *pb.VerifyOrderFromChainRequest) error {
	c, err := NewChain(in.Chain)
	if err != nil {
		return err
	}
	return c.VerifyOrders(ctx, in.StartBlock)
}

func DecodeTransactionInput(input string) (*TransactionInput, error) {
	if contractABI == nil {
		return nil, errors.New("contractABI is nil")
	}
	if input[:2] == "0x" {
		input = input[2:]
	}
	decodedBytes, err := hex.DecodeString(input)
	if err != nil {
		return nil, err
	}
	if len(decodedBytes) == 0 {
		return nil, errors.New("no input data")
	}
	methodSigData := decodedBytes[:4]
	method, err := contractABI.MethodById(methodSigData)
	if err != nil {
		return nil, err
	}
	inputsSigData := decodedBytes[4:]
	inputsMap := make(map[string]interface{})
	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		return nil, err
	}
	rawAmount, ok := inputsMap["amount"]
	if !ok {
		return nil, errors.New("param error: amount")
	}
	amount, ok := rawAmount.(*big.Int)
	if !ok {
		return nil, errors.New("param type error: amount")
	}
	rawTokenAddress, ok := inputsMap["tokenAddress"]
	if !ok {
		return nil, errors.New("param error: tokenAddress")
	}
	rawUniqueKey, ok := inputsMap["uniqueKey"]
	if !ok {
		return nil, errors.New("param error: uniqueKey")
	}
	return &TransactionInput{
		Amount: amount,
		TokenAddress: fmt.Sprintf("%s", rawTokenAddress),
		UniqueKey: fmt.Sprintf("%s", rawUniqueKey),
	}, nil
}
