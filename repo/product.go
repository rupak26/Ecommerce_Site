package repo

import (
	 "github.com/jmoiron/sqlx"
)

type Product struct {
	Id   int           `json:"id"`
	ProductName string `json:"productname"`
	Url  string        `json:"url"`
	Quantity int       `json:"quantity"`
}

type ProductRepo interface{
	Create(p Product) (*Product , error) 
	Get(productID int) (*Product , error)
	List() (*[]Product , error)
} 

type productRepo struct {
    db *sqlx.DB
}


func NewProductRepo(db *sqlx.DB) ProductRepo {
     repo := &productRepo{
		db : db ,
	 }
	 return repo
}

func (r *productRepo) Create(p Product) (*Product , error) {
	query := `
		INSERT INTO products (
			product_name,
			url,
			quantity
		) VALUES (
			:product_name,
			:url,
			:quantity
		)
		RETURNING id;
	`

	rows, err := r.db.NamedQuery(query, p)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&p.Id)
	}

	return &p, nil
}

func (r productRepo) Get(productId int) (*Product , error) {
	var prod Product
    query := `SELECT FROM products WHERE id=$1`
	err := r.db.Get(&prod , query , productId) 
    if err != nil {
		return nil , err
	}
	return &prod , nil
} 

func (r *productRepo) List() (*[]Product , error) {
	var prod []Product
	query := `SELECT * FROM products`
	err := r.db.Select(&prod, query)
	if err != nil {
		return nil, err
	}
	return &prod, nil
} 

	

