package coingecko

// ChartItem struct
type ChartItem [2]float64

// CoinsIDMarketChart struct
type CoinsIDMarketChart struct {
	Prices       []ChartItem `json:"prices"`
	MarketCaps   []ChartItem `json:"market_caps"`
	TotalVolumes []ChartItem `json:"total_volumes"`
}
