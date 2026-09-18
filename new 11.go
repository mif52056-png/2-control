type Product struct {
	Name     string
	Category string
	Price    float64
}

func filterProducts(products []Product, maxPrice float64, category string) []Product {
	var result []Product
	for _, p := range products {
		if p.Price < maxPrice && p.Category == category {
			result = append(result, p)
		}
	}
	return result
}