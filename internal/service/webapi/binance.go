package ser

import (
	"bina/internal/core"
	"context"
	binance "github.com/adshao/go-binance/v2"
	"strings"
)

type BinanceWebApiCFG struct {
	APIKey    string
	APISecret string
}

type BinanceWEBapi struct {
	client *binance.Client
	cfg    *BinanceWebApiCFG
}

func NewBinanceWebApi(cfg *BinanceWebApiCFG) *BinanceWEBapi {
	return &BinanceWEBapi{
		cfg:    cfg,
		client: nil,
	}
}


func (b *BinanceWEBapi) GetPrices(coin string) ([]core.SymbolPrice, error) {
	prices, err := b.client.NewListPricesService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	coinPrices := make([]core.SymbolPrice, 0)
	for _, price := range prices {
		if strings.HasPrefix(price.Symbol, coin) {
			coinPrices = append(coinPrices, core.SymbolPrice{
				Symbol: price.Symbol,
				Price:  price.Price,
			})
		}
	}
	return coinPrices, nil
}
