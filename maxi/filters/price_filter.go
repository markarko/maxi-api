package filters

import (
	"github.com/markarko/maxi-api/data"
)

func GetProductsWithBiggerPrice(products data.Products, price float64) data.Products {
	var filteredProducts data.Products
	for _, product := range products {
		if float64(product.Prices.Price.Value) > price {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}

func GetProductsWithSmallerPrice(products data.Products, price float64) data.Products {
	var filteredProducts data.Products
	for _, product := range products {
		if float64(product.Prices.Price.Value) < price {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts
}
