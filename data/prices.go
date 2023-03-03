package data

type Prices struct {
	Price            Price             `json:"price"`
	WasPrice         string            `json:"wasPrice"`
	ComparisonPrices []ComparisonPrice `json:"comparisonPrices"`
}
