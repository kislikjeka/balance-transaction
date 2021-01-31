package handler

import (
	"TransactionTest/internal/service/balance"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	balanceService balance.UseCase
}

func NewHandler(bs balance.UseCase) *Handler {
	return &Handler{bs}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	b := router.Group("/balances")
	{
		b.GET("/", h.listBalances)
		b.POST("/", h.createBalance)
		b.GET("/:id", h.getBalanceById)
		b.PUT("/:id", h.updateBalance)
		b.DELETE("/:id", h.deleteBalance)
		b.POST("/transaction", h.makeTransaction)
	}

	return router
}
