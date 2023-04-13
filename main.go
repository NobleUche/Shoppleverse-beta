package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

/** Look into the handlers middleware **/

//  Functions

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	ConnectDB()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/products", GetAllProducts).Methods("GET")
	router.HandleFunc("/api/v1/vendorlogin", VendorLogin).Methods("GET")
	router.HandleFunc("/api/v1/ordernow", OrderCreation).Methods("POST")
	router.HandleFunc("/api/v1/createproducts", Createproduct).Methods("POST")
	router.HandleFunc("/api/v1/product/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/api/v1/updateproduct/{id}", UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/v1/deleteproduct/{id}", DeleteProduct).Methods("DELETE")
	router.HandleFunc("/api/v1/vendorsignup", Vendorsignup).Methods("POST")
	router.HandleFunc("/api/v1/vendor/{id}", GetVendor).Methods("GET")
	router.HandleFunc("/api/v1/vendors", GetAllVendors).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(router)))
}
