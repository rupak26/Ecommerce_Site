package domain 


type Product struct {
	Id          int    `db:"id" json:"id"`
	ProductName string `db:"productname" json:"productname"`
	Url         string `db:"url" json:"url"`
	Quantity    int    `db:"quantity" json:"quantity"`
}

type UpdateProductReq struct {
	ProductName *string
	Url *string
	Quantity *int
}
