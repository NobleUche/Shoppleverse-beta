package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"io"
	"net/http"
)

// struct

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Category    string `json:"category"`
	VendorId    uint   `json:vendorid`
	Vendoremail string `json:"vendoremail"`
	Image       []byte `json:"image"`
}

/** Error coming from image upload and serach query functions. Look into it**/

// Product functions

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var allproducts []Product
	Db.Find(&allproducts)
	json.NewEncoder(w).Encode(allproducts)
}

func Createproduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products Product
	// file upload retrieval
	r.MultipartForm(10 << 20) // max file size 10mb
	file, handler, err := r.FormFile("file")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()
	// read file contents
	filecontents, err := io.ReadAll(file)
	if err != nil {
		panic(err)
		return
	}
	// pass every data into struct
	json.NewDecoder(r.Body).Decode(&products)
	Db.Create(&products)
	fmt.Println("Product created")
	json.NewEncoder(w).Encode(products)
}
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var products Product
	if r.Method == "GET" {
		Db.First(&products, params["id"])
		json.NewEncoder(w).Encode(products)
	} else {
		json.NewEncoder(w).Encode("ITEM NOT AVAILABLE")
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var products Product
	if r.Method == "PUT" {
		Db.First(&products, params["id"])
		json.NewDecoder(r.Body).Decode(&products)
		Db.Save(&products)
		json.NewEncoder(w).Encode(products)
		json.NewEncoder(w).Encode("ITEM UPDATED")
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var products Product
	Db.Delete(&products, params["id"])
	json.NewEncoder(w).Encode("PRODUCT DELETED SUCCESSFULLY")
}


func ProductSearch(w http.ResponseWriter, r *http.Request){
	nameQuery:=r.URL.Query()
	var products []Product
	Db.Where("name LIKE ? OR category LIKE ?", "%"+nameQuery+"%", "%"+nameQuery+"%").Find(&products)
	json.NewEncoder(w).Encode(products)
}