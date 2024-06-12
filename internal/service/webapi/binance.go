package ser

type BinanceWebApiCFG struct {
	APIKey string
	APISecret string

}

type BinanceWEBapi struct {
	client *binance.client
	 cfg *BinanceWebApiCFG
}

func NewBinanceWebApi(cfg *BinanceWebApiCFG) *BinanceWEBapi {
return &BinanceWEBapi{
	cfg: cfg,
	client: nil,
}
}

