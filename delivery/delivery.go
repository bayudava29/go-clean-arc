package delivery

import (
	"github.com/bayudava29/go-clean-arc/usecase"
	"github.com/gin-gonic/gin"
)

type Delivery interface {
	GetData(c *gin.Context)
}

type delivery struct {
	Usecase usecase.Usecase
}

func InitDelivery(usecase usecase.Usecase) Delivery {
	return &delivery{
		Usecase: usecase,
	}
}

func (delivery *delivery) GetData(c *gin.Context) {
	// Instruction for GetData Delivery
}
