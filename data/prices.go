package data

type Prices struct {
	Price            Price             `json:"price"`
	WasPrice         WasPrice          `json:"wasPrice"`
	ComparisonPrices []ComparisonPrice `json:"comparisonPrices"`
}
