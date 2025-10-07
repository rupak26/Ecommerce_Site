package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/products"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface{
	 products.ProductRepo // embeding
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

func (r *productRepo) Create(p domain.Product) (*domain.Product , error) {
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

func (r productRepo) GetById(pId int) (*domain.Product , error) {
	var prod domain.Product
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

func (r *productRepo) List(page , limit int64) (*[]domain.Product , error) {
	var prod []domain.Product
	offset := (page - 1) * limit
    query := `SELECT * FROM products LIMIT $1 OFFSET $2`
	err := r.db.Select(&prod, query , limit , offset)
	if err != nil {
		return nil, err
	}
	return &prod, nil
} 

func (r *productRepo) Count() (int64, error) {
    query := `SELECT COUNT(*) FROM products`
    var count int
    err := r.db.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, err 
    }
    return int64(count), nil
}

func (r *productRepo) Update(prod domain.Product) (*domain.Product , error) {
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

func (r *productRepo) PatchProduct(id int, req domain.UpdateProductReq) (*domain.Product, error) {
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

	var prod domain.Product
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