package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product_handler"
	"ecommerce/rest/handlers/user_handler"
	"ecommerce/rest/middleware"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()
    dbCon , err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
    


	middleware := middleware.NewMiddleware(cnf)
	productRepo := repo.NewProductRepo(dbCon)
    productHandler := product_handler.NewHandler(middleware , productRepo)
	userRepo := repo.NewUserRepo(dbCon)
	userHandler := user_handler.NewHandler(userRepo)

    server := rest.NewServer(cnf , productHandler , userHandler)
	server.Start()
}