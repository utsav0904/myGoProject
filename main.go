package main

import (
	"Interface/Dal"
	"Interface/Handler"
	"Interface/Services"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	fmt.Println("Ready....")
	//Items := models.InitialiseItems()
	//Transactions := models.InitialiseTransactions()
	Database := Dal.NewCatalogDataBase()
	ProductRepo := Dal.NewProductRepo(Database)
	ProductService := Services.NewProductServices(ProductRepo)
	ProductController := Handler.Initialise(ProductService)
	r.HandleFunc("/products/all", ProductController.GetAllProduct).Methods("GET")
	r.HandleFunc("/products/{id}", ProductController.GetProductById).Methods("GET")
	r.HandleFunc("/products/create", ProductController.CreateProduct).Methods("POST")
	r.HandleFunc("/products/buy/{id}", ProductController.BuyProduct).Methods("POST")
	r.HandleFunc("/products/purchase", ProductController.BuyProductMany).Methods("POST")
	r.HandleFunc("/transactions/all", ProductController.GetAllTransactions).Methods("GET")
	r.HandleFunc("/transactions/top5", ProductController.GetTop5Products).Methods("GET")
	//log.Fatal(http.ListenAndServe(":8080", r))
	port := os.Getenv("PORT")
	if port == "" {
        	port = "8080"
	}
	log.Fatal(http.ListenAndServe(":port", r))
}
