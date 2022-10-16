package app

import (
	"github.com/bayudava29/go-clean-arc/config"
	"github.com/bayudava29/go-clean-arc/delivery"
	"github.com/bayudava29/go-clean-arc/repository"
	"github.com/bayudava29/go-clean-arc/usecase"
	"github.com/gin-gonic/gin"
)

func InitRouter(hbaseConn []config.HbaseConnection) *gin.Engine {
	// Init Layer
	repository := repository.InitRepository(hbaseConn)
	usecase := usecase.InitUsecase(repository)
	delivery := delivery.InitDelivery(usecase)

	router := gin.Default()
	svc := router.Group("/")
	svc.GET("/get", delivery.GetData)
	return router
}
