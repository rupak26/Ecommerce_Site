package productdatabase



type Product struct {
	Id   int    `json:"id"`
	ProductName string `json:"productname"`
	Url  string    `json:"url"`
	Quantity int `json:"quantity"`
}

var ProductList[] Product



func StoreProduct(product Product) {
	ProductList = append(ProductList, product)
}

func GetProductList() []Product {
	return ProductList
}
func GetProduct(id int) *Product {
	for _,prod := range(ProductList) {
		if prod.Id == id {
			return &prod
		}
	}
	return nil
}
