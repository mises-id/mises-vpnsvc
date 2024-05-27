package chain

type BscTestNet struct {}

func NewBscTestNet() *BscTestNet {
	return &BscTestNet{}
}

func (*BscTestNet) VerifyOrders (startTime int64) error {
	return nil
}
