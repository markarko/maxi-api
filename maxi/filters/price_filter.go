package filters

import (
	"github.com/markarko/maxi-api/data"
)

type PriceFilter int8

const (
	BIGGER PriceFilter = iota
	SMALLER
	EQUAL
)

func GetProductsWithBiggerPrice(products data.Products, price float32) data.Products {
	return GetProductsAccordingToPrice(products, price, BIGGER)
}

func GetProductsWithSmallerPrice(products data.Products, price float32) data.Products {
	return GetProductsAccordingToPrice(products, price, SMALLER)
}

func GetProductsWithEqualPrice(products data.Products, price float32) data.Products {
	return GetProductsAccordingToPrice(products, price, EQUAL)
}

func GetProductsAccordingToPrice(products data.Products, price float32, filter PriceFilter) data.Products {
	var filteredProducts data.Products
	for _, product := range products {

		priceToCompare := float32(product.Prices.Price.Value)

		if filter == BIGGER && priceToCompare > price ||
			filter == SMALLER && priceToCompare < price ||
			filter == EQUAL && priceToCompare == price {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}
