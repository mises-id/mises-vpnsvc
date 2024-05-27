package chain

type Tron struct {}

func NewTron() *Tron {
	return &Tron{}
}

func (*Tron) VerifyOrders (startTime int64) error {
	return nil
}
