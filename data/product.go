package data

type Product struct {
	Code              string       `json:"code"`
	Name              string       `json:"name"`
	Description       string       `json:"description"`
	Brand             string       `json:"brand"`
	Link              string       `json:"link"`
	ImageAssets       []ImageAsset `json:"imageAssets"`
	PackageSize       string       `json:"packageSize"`
	Shoppable         bool         `json:"shoppable"`
	Prices            Prices       `json:"prices"`
	PricingUnits      PricingUnits `json:"pricingUnits"`
	Badges            Badges       `json:"badges"`
	StockStatus       string       `json:"stockStatus"`
	ArticleNumber     string       `json:"articleNumber"`
	Uom               string       `json:"uom"`
	Taxes             string       `json:"taxes"`
	Fees              string       `json:"fees"`
	SellerId          string       `json:"sellerId"`
	SellerName        string       `json:"sellerName"`
	IsVariant         bool         `json:"isVariant"`
	OfferType         string       `json:"offerType"`
	HasMultipleOffers bool         `json:"hasMultipleOffers"`
	OfferId           string       `json:"offerId"`
	Aisle             Aisle        `json:"aisle"`
	Sponsored         bool         `json:"sponsored"`
}

type Products []*Product
