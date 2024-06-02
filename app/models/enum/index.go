package enum

import (
	"time"
)

type (
	StatusType uint8
)

const (
	//StatusType
	StatusOpen   StatusType = 1
	StatusClose  StatusType = 2
	StatusDelete StatusType = 3
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
		97: "bsc test net", // bsc testnet
		56: "bsc",          // bsc
	}
)
