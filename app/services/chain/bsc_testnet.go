package chain

type BscTestNet struct {}

func NewBscTestNet() *BscTestNet {
	return &BscTestNet{}
}

func (*BscTestNet) VerifyOrders (startBlock int64) error {
	if startBlock <= 0 {

	}
	return nil
}
