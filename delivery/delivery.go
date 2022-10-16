package delivery

type Delivery interface{}

type delivery struct{}

func InitDelivery() Delivery {
	return &delivery{}
}
