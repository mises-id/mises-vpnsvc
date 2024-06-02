package enum

import (
	"time"
)

type (
	StatusType uint8
)

const (
	// StatusType
	StatusOpen   StatusType = 1
	StatusClose  StatusType = 2
	StatusDelete StatusType = 3

	// Chain ID
	ChainIDETH     uint64 = 1
	ChainIDBscTest uint64 = 97
	ChainIDBsc     uint64 = 56
)

// config
type PlanItem struct {
	TokenAmount string
	TokenName   string
	TimeRange   time.Duration
}

type Plan map[uint64]PlanItem

type Chain map[uint64]string

var (
	Plans = Plan{
		1: {
			TokenAmount: "3.00",
			TokenName:   "USDT",
			TimeRange:   30 * 24 * time.Hour,
		},
	}

	Chains = Chain{
		ChainIDBscTest: "bsc test net",
		ChainIDBsc:     "bsc",
	}
)
