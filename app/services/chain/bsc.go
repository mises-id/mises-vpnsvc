package chain

type Bsc struct {}

func NewBsc() *Bsc {
	return &Bsc{}
}

func (*Bsc) VerifyOrders (startTime int64) error {
	return nil
}
