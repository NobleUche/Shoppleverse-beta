package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

//  Functions

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	ConnectDB()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/products", GetAllProducts).Methods("GET")
	router.HandleFunc("/api/v1/createproducts", Createproduct).Methods("POST")
	router.HandleFunc("/api/v1/product/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/api/v1/updateproduct/{id}", UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/v1/deleteproduct/{id}", DeleteProduct).Methods("DELETE")
	router.HandleFunc("/api/v1/vendorsignup", Vendorsignup).Methods("POST")
	router.HandleFunc("/api/v1/vendor/{id}", GetVendor).Methods("GET")
	router.HandleFunc("/api/v1/vendors", GetAllVendors).Methods("GET")
	Server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}
	log.Fatal(Server.ListenAndServe())
}
