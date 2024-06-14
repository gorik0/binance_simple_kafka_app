package service

import (
	"bina/internal/core"
	ser "bina/internal/service/webapi"
)

type CoinService struct {

	WebApi *ser.BinanceWEBapi

}

func (c *CoinService) 	GetCoinPrice(symbol string) ([]core.SymbolPrice,error) {
	return c.WebApi.GetPrices(symbol)
}


func NewCoinService(webApi *ser.BinanceWEBapi) *CoinService {
	return &CoinService{WebApi: webApi}
}