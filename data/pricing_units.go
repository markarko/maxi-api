package data

type PricingUnits struct {
	Type             string  `json:"type"`
	Unit             string  `json:"unit"`
	Interval         float32 `json:"interval"`
	MinOrderQuantity float32 `json:"minOrderQuantity"`
	MaxOrderQuantity float32 `json:"maxOrderQuantity"`
	Weighted         bool    `json:"weighted"`
}
