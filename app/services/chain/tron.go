package chain

type Tron struct {}

func NewTron() *Tron {
	return &Tron{}
}

func (*Tron) VerifyOrders (startBlock int64) error {
	return nil
}
