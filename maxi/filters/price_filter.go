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

func GetProductsFilterByPriceAsc(products data.Products) data.Products {
	return GetProductsFilterByPrice(products, true)
}

func GetProductsFilterByPriceDesc(products data.Products) data.Products {
	return GetProductsFilterByPrice(products, false)
}

func GetProductsFilterByPrice(products data.Products, asc bool) data.Products {
	var filteredProducts data.Products
	for _, product := range products {
		price := float32(product.Prices.Price.Value)
		if len(filteredProducts) == 0 {
			filteredProducts = append(filteredProducts, product)
		} else {
			for i, filteredProduct := range filteredProducts {
				filteredPrice := float32(filteredProduct.Prices.Price.Value)
				if asc && price < filteredPrice ||
					!asc && price > filteredPrice {
					filteredProducts = append(filteredProducts, &data.Product{})
					copy(filteredProducts[i+1:], filteredProducts[i:])
					filteredProducts[i] = product
					break
				}
			}
		}
	}
	return filteredProducts
}

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
