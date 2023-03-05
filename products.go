package main

import(
	"encoding/json"
	"fmt"
	"net/http"
	"gorm.io/gorm"
	"github.com/gorilla/mux"
)



// struct

type Product struct{
	gorm.Model
	Name string `json:"name"`
	Price uint `json:"price"`
	Category string `json:"category"`
	VendorId uint `json:vendorid`
	Vendoremail string `json:"vendoremail"`
	//Image string `json:"image"`
}




// Product functions

func GetAllProducts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var allproducts []Product
	Db.Find(&allproducts)
	json.NewEncoder(w).Encode(allproducts)
}

func Createproduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var products Product
	json.NewDecoder(r.Body).Decode(&products)
	Db.Create(&products)
	fmt.Println("Product created")
	json.NewEncoder(w).Encode(products)
}
func GetProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r)
	var products Product
	if r.Method=="GET"{
		//Set that if the if the product id is not equal to the db id return null
		Db.First(&products, params["id"])
		json.NewEncoder(w).Encode(products)

	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	var products Product
	if r.Method=="PUT"{
		Db.First(&products, params["id"])
		json.NewDecoder(r.Body).Decode(&products)
		Db.Save(&products)
		json.NewEncoder(w).Encode(products)
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	var products Product
	Db.Delete(&products,params["id"])
	json.NewEncoder(w).Encode("Product deleted successfully")
}

