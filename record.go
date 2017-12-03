package cryptoping

type Record struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	Rank         int8    `json:"rank,string"`
	PriceUSD     float32 `json:"price_usd,string"`
	PriceBTC     float32 `json:"price_btc,string"`
	VolumeUSD    float32 `json:"34h_volume_usd,string"`
	MarketCapUsd float32 `json:"market_cap_usd,string"`
	Supply       float32 `json:"available_supply,string"`
	Change1h     float32 `json:"percent_change_1h,string"`
	Change24h    float32 `json:"percent_change_24h,string"`
	Change7d     float32 `json:"percent_change_7d,string"`
	LastUpdated  string  `json:"last_updated"`
}
