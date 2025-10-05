package repo

import (
	"database/sql"
	"fmt"
	"strings"
	
	"github.com/jmoiron/sqlx"
)

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


type ProductRepo interface{
	Create(p Product) (*Product , error) 
	GetById(pID int) (*Product , error)
	List() (*[]Product , error)
	Update(prod Product) (*Product , error)
	PatchProduct(id int, req UpdateProductReq) (*Product, error)
	Delete(id int) error
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
			productname,
			url,
			quantity
		) VALUES (
			:productname,
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
		if err := rows.Scan(&p.Id); err != nil {
			return nil, err
		}
	}
	return &p, nil
}

func (r productRepo) GetById(pId int) (*Product , error) {
	var prod Product
    query := `SELECT id, productname, url, quantity FROM products WHERE id = $1`
	err := r.db.Get(&prod , query , pId) 
    if err != nil {
		if err == sql.ErrNoRows {
			return nil , nil
		}
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

func (r *productRepo) Update(prod Product) (*Product , error) {
	query :=`   UPDATE products
	            SET   
			    productname =  :productname,
				url = :url ,
				quantity = :quantity
			    WHERE id = :id
		        RETURNING id, productname, url, quantity;	
			`
    rows, err := r.db.NamedQuery(query, prod)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.StructScan(&prod)
		if err != nil {
			return nil, err
		}
		return &prod, nil
	}
	return nil, nil	
}


func (r *productRepo) PatchProduct(id int, req UpdateProductReq) (*Product, error) {
	query := "UPDATE products SET "
	params := map[string]interface{}{"id": id}
	setClauses := []string{}

	if req.ProductName != nil {
		setClauses = append(setClauses, "productname = :productname")
		params["productname"] = *req.ProductName
	}
	if req.Url != nil {
		setClauses = append(setClauses, "url = :url")
		params["url"] = *req.Url
	}
	if req.Quantity != nil {
		setClauses = append(setClauses, "quantity = :quantity")
		params["quantity"] = *req.Quantity
	}
	
	if len(setClauses) == 0 {
		return nil, fmt.Errorf("no fields provided to update")
	}

	query += strings.Join(setClauses, ", ") + " WHERE id = :id RETURNING id, productname , url , quantity;"
    
	rows, err := r.db.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prod Product
	if rows.Next() {
		if err := rows.StructScan(&prod); err != nil {
			return nil, err
		}
		return &prod, nil
	}
	return nil, fmt.Errorf("user with id %d not found", id)
}


func (r *productRepo) Delete(id int) error {
	query := `DELETE FROM products WHERE id=$1` 
	
	_,err := r.db.Exec(query , id) 
	if err != nil {
		return  err
	}
	return nil
}