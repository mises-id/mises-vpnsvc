package chain

type TronTestNet struct {}

func NewTronTestNet() *TronTestNet {
	return &TronTestNet{}
}

func (*TronTestNet) VerifyOrders (startBlock int64) error {
	return nil
}
