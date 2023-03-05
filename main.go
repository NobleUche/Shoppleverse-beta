package main

import(
	"log"
	"crypto/md5"
	//"error"
	"github.com/gorilla/mux"
	"github.com/mazen160/go-random"
	"net/http"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"fmt"
)

// Database Connection 

const DSN = "postgresql://postgres:peculiar@localhost:5432/test"

// structs

type Product struct{
	gorm.Model
	Name string `json:"name"`
	Price uint `json:"price"`
	Category string `json:"category"`
	VendorId uint `json:vendorid`
	Vendoremail string `json:"vendoremail"`
	//Vendors Vendor `gorm:"association_foreignkey:Seller"`
	//Image string `json:"image"`
}

type Vendor struct{
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Walletaddress string `json:"walletaddress"`
	VendorId int `json:"vendorid"`
}




// Variables

var Db *gorm.DB
var err error


//  Functions

func main(){
	ConnectDB()
	router:= mux.NewRouter()
	router.HandleFunc("/api/v1/products",GetAllProducts).Methods("GET")
	router.HandleFunc("/api/v1/createproducts",Createproduct).Methods("POST")
	router.HandleFunc("/api/v1/product/{id}",GetProduct).Methods("GET")
	router.HandleFunc("/api/v1/updateproduct/{id}",UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/v1/deleteproduct/{id}",DeleteProduct).Methods("DELETE")
	router.HandleFunc("/api/v1/vendorsignup",Vendorsignup).Methods("POST")
	router.HandleFunc("/api/v1/vendor/{id}",GetVendor).Methods("GET")
	router.HandleFunc("/api/v1/vendors",GetAllVendors).Methods("GET")
	router.HandleFunc("/api/v1/search",SearchProducts).Methods("POST")
	Server:= http.Server{
		Addr:":9000",
		Handler: router,
	}
	log.Fatal(Server.ListenAndServe())
}


func ConnectDB(){
	// Database connection
	Db,err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil{
		panic("failed to connect")
	}
	fmt.Println("Successfully connected")
	Db.AutoMigrate(&Product{}, Vendor{})
 }

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

	}/*else if r.Method=="PUT"{
		// format is get request then post request
		Db.First(&products,params["productid"])
		json.NewDecoder(r.Body).Decode(&products)
		Db.Save(&products)
		json.NewEncoder(w).Encode(products)

		}else if r.Method == "DELETE"{
			Db.Delete(&products,params["productid"])
			//json.NewDecoder(r.Body).Decode(&products)
			//Db.Save(&products)
			json.NewEncoder(w).Encode("The User is deleted Successfully ")


	}*/

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

func Vendorsignup(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var vendors Vendor
	var data int
	var err error
	//var vendorsid string
	json.NewDecoder(r.Body).Decode(&vendors)
	V:=vendors.Name
	Id:=func (n string){
		data, err = random.GetInt(len(n))
		if err != nil{
			log.Fatal(err)
		}
		//gdata:=data+
		vendors.VendorId=md5.Sum(data)
	}
	// Use cryptography and seed on the rand result 
	Id(V)
	Db.Create(&vendors)
	json.NewEncoder(w).Encode(vendors)
}

func GetVendor(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r)
	var vendors Vendor
	if r.Method=="GET"{
		Db.First(&vendors,params["id"])
		json.NewEncoder(w).Encode(vendors)
	}
}

func GetAllVendors(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var allvendors []Vendor
	Db.Find(&allvendors)
	json.NewEncoder(w).Encode(allvendors)
}

// Search functions 

/* The search functions is not working as it is not returning the right values*/

func SearchProducts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var products Product  
	var search string
	//json.NewDecoder(r.Body).Decode(&search)
	//Db.Where("name = ?",search_query).Find(&products)
	Db.Find(&products, search)
	json.NewEncoder(w).Encode(products)
}

/*func SearchVendors(w http.ResponseWriter, r *http.Request){
s	w.Header().Set("Content-Type", "application/json")	
} */


/* Setup email code verification and vendor Id code generation*/





































