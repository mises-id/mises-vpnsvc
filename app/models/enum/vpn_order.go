package enum

type (
	VpnOrderStatus int
	ChainID        uint64
	VpnOrderStatusMap map[VpnOrderStatus]string
)

const (
	// VpnOrderStatus
	VpnOrderInit    VpnOrderStatus = 0
	VpnOrderSuccess VpnOrderStatus = 1
	VpnOrderFail    VpnOrderStatus = 2
	VpnOrderPending VpnOrderStatus = 3
	VpnOrderRetry   VpnOrderStatus = 4

	// ChainID
	ChainETH     ChainID = 1
	ChainBSC     ChainID = 56
	ChainBSCTest ChainID = 97
)

var (
	VpnOrderStatusText = VpnOrderStatusMap{
		VpnOrderInit: "unpaid",
		VpnOrderSuccess: "paid",
		VpnOrderFail: "failed",
		VpnOrderPending: "pending",
		VpnOrderRetry: "pending",
	}
)
