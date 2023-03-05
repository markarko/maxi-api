package data

import (
	"encoding/json"
	"io"
)

type ProductsContainer struct {
	query              json.RawMessage `json:"-"`
	RequestId          json.RawMessage `json:"-"`
	CorrectedQuery     json.RawMessage `json:"-"`
	Results            Products        `json:"results"`
	SponsoredCarousels json.RawMessage `json:"-"`
	Pagination         json.RawMessage `json:"-"`
	Sorts              json.RawMessage `json:"-"`
	FilterGroups       json.RawMessage `json:"-"`
	Breadcrumbs        json.RawMessage `json:"-"`
	CategoryName       json.RawMessage `json:"-"`
	ModelVersion       json.RawMessage `json:"-"`
	SearchVariation    json.RawMessage `json:"-"`
	Dym                json.RawMessage `json:"-"`
	OfferTypes         json.RawMessage `json:"-"`
}

func (pc *ProductsContainer) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(pc)
}

func (pc *ProductsContainer) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(pc)
}
