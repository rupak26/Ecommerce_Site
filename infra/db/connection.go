package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	//"golang.org/x/text/number"
)

func GetConnectionString() string {
     return "user=postgres password=12345678 host=localhost port=5432 dbname=ecommerce sslmode=disable"
}

func NewConnection() (*sqlx.DB , error) {
    dbSource := GetConnectionString()
	db , err := sqlx.Connect("postgres" , dbSource)
    
	if err != nil {
		fmt.Println(err) 
		return nil , err
	}
	return db , nil
}