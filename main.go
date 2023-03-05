package main

import(
	"os"
	"log"
	//"net/smtp"
	//"github.com/joho/godotenv"
	"github.com/NobleUche/code-generator"
	"github.com/gorilla/mux"
//	"github.com/mazen160/go-random"
	"net/http"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"fmt"
)


// structs

type Product struct{
	gorm.Model
	Name string `json:"name"`
	Price uint `json:"price"`
	Category string `json:"category"`
	VendorId uint `json:vendorid`
	Vendoremail string `json:"vendoremail"`
	//Image string `json:"image"`
}

type Vendor struct{
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Walletaddress string `json:"walletaddress"`
	VendorId string `json:"vendorid"`
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
	Server:= http.Server{
		Addr:":9000",
		Handler: router,
	}
	log.Fatal(Server.ListenAndServe())
}


func ConnectDB(){
	// Database connection
	//godotenv.Load()
	DSN:=os.Getenv("Postgresql-url")

	Db,err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil{
		panic("failed to connect")
	}
	fmt.Println("Successfully connected")
	Db.AutoMigrate(&Product{}, Vendor{})
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


// Vendor functions

func Vendorsignup(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var vendors Vendor
	//var vendorsid string
	json.NewDecoder(r.Body).Decode(&vendors)
	// Vendor-Id generation
	Id:=func (n string){
		data, err := generate.GetId(vendors.Email)
		if err != nil{
			panic(err)
		}
		vendors.VendorId=data
	}
	// Email-verification generation and sending
	/*Code:= func (n string){
		dta,err=generate.GetCode(vendors.Name)
		if err !=nil{
			panic(err)
		}
		Email_Code:=dta
		fmt.Println(Email_Code)  
	}*/
	
	Id(vendors.Email)
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

// Email function

/*func EmailSetup(n ,j string){
	from:=os.Getenv("Email")
	password:=os.Getenv("Password")
	//toEmail:=// user email here
	host:="smtp.gmail.com"
	port:="587"
	address:= host+":"+port
	subject:= "Shoppleverse Verification Code"
	//body:= // verification code here
	message:=[]byte(subject+body)
	//email authentication
	auth:=smtp.PlainAuth("",from,password, host)
	//send email
	err:=smtp.SendEmail(address, auth, from, to, message)
	if err !=nil{
		panic(err)
	}

}*/





































