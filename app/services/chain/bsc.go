package chain

type Bsc struct {}

func NewBsc() *Bsc {
	return &Bsc{}
}

func (*Bsc) VerifyOrders (startBlock int64) error {
	return nil
}
