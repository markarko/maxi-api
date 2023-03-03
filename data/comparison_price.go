package data

type ComparisonPrice struct {
	Value      float32 `json:"value"`
	Unit       string  `json:"unit"`
	Quantity   int     `json:"quantity"`
	ReasonCode string  `json:"reasonCode"`
	Type       string  `json:"type"`
}
