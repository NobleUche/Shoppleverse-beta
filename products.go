package main

import (
	"os"
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
	Price       string   `json:"price"`
	Category    string `json:"category"`
	VendorId    string   `json:vendorid`
	Image    	string `json:"image"`
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
	var product Product

	// parse the multipart form
	err:=r.ParseMultipartForm(32 << 20) // max file size 32mb
	if err != nil{
		http.Error(w, "Error parsing multipart form", http.StatusBadRequest)
		return
	}

	//Get the form values
	name:=r.FormValue("name")
	price:=r.FormValue("price")
	category:=r.FormValue("category")
	vendorid:=r.FormValue("vendorid")

	//Get the file from the request
	file, handler, err := r.FormFile("image")
	if err != nil{
		http.Error(w, "Error getting file from request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	//Save file to disk
	imagePath:="./images"+handler.Filename
	f,err:=os.OpenFile(imagePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil{
		http.Error(w, "Error saving file to disk", http.StatusInternalServerError)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil{
		http.Error(w,"Error saving file to disk", http.StatusInternalServerError)
		return
	}

	// pass every data into struct
	product=Product{
		Name: name,
		Price:price,
		Image:imagePath,
		VendorId:vendorid,
		Category:category,
	}

	// pass into database
	result :=Db.Create(&product)
	if result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Product created")
	json.NewEncoder(w).Encode(product)
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