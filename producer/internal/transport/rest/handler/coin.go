package handler

import (
	"bina/internal/core"
	"github.com/gin-gonic/gin"
	"net/http"
)
type getCoinPriceRequest struct {
	Coin string `json:"coin" binding:"required"`
}

type getCoinPriceResponse struct {
	Prices []core.SymbolPrice `json:"prices"`
}
func (h *Handler) getPrice(c *gin.Context){
	var request getCoinPriceRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	coinPrices, err := h.service.Coin.GetCoinPrice(request.Coin)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error getting coin prices")
		return
	}

	c.JSON(http.StatusOK, getCoinPriceResponse{
		Prices: coinPrices,
	})
}

