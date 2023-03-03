package data

type Price struct {
	Value      float32 `json:"value"`
	Unit       string  `json:"unit"`
	Quantity   int     `json:"quantity"`
	ReasonCode int     `json:"reasonCode"`
	Type       string  `json:"type"`
}
