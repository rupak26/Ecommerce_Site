package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/infra/db"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product_handler"
	"ecommerce/rest/handlers/user_handler"
	"ecommerce/rest/middleware"
	"ecommerce/users"
	"ecommerce/products"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()
    dbCon , err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
    
    err = db.MigrateDB(dbCon , "./migrations")
    if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)
   
    //domains
    usrSvc := users.NewService(userRepo)
    prodSvc := products.NewService(productRepo)

	middleware := middleware.NewMiddleware(cnf)
	
    productHandler := product_handler.NewHandler(middleware , prodSvc)
	userHandler := user_handler.NewHandler(usrSvc)

    server := rest.NewServer(cnf , productHandler , userHandler)
	server.Start()
}