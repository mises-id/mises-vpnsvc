package chain

type TronTestNet struct {}

func NewTronTestNet() *TronTestNet {
	return &TronTestNet{}
}

func (*TronTestNet) VerifyOrders (startTime int64) error {
	return nil
}
