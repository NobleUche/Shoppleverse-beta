package main

import (
	"encoding/json"
	"github.com/NobleUche/code-generator"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"net/smtp"
	"os"
)

// structs

type Vendor struct {
	gorm.Model
	Name          string `json:"name"`
	Email         string `json:"email"`
	Walletaddress string `json:"walletaddress"`
	VendorId      string `json:"vendorid"`
}

// Vendor functions

func Vendorsignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var vendors Vendor
	//var vendorsid string
	json.NewDecoder(r.Body).Decode(&vendors)
	// Vendor-Id generation
	Id := func(email string) {
		data, err := generate.GetId(email)
		if err != nil {
			panic(err)
		}
		vendors.VendorId = data
	}
	// Email-verification code generation and sending
	Code := func(name string) {
		data, err := generate.GetCode(name)
		if err != nil {
			panic(err)
		}
		Email_Code := data
		EmailSetup(vendors.Email, Email_Code)
	}

	Id(vendors.Email)
	Code(vendors.Name)
	Db.Create(&vendors)
	json.NewEncoder(w).Encode(vendors)
}

func GetVendor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var vendors Vendor
	if r.Method == "GET" {
		Db.First(&vendors, params["id"])
		json.NewEncoder(w).Encode(vendors)
	}
}

func GetAllVendors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var allvendors []Vendor
	Db.Find(&allvendors)
	json.NewEncoder(w).Encode(allvendors)
}

// Email function

func EmailSetup(receiver_email, verification_code string) {
	from := os.Getenv("Email") // string
	password := os.Getenv("Password")
	to := []string{receiver_email} // user email here said it is []string though
	host := "smtp.gmail.com"       //string
	port := "587"
	address := host + ":" + port
	subject := "Shoppleverse Verification Code"
	body := verification_code // verification code here
	message := []byte(subject + body)
	//email authentication
	auth := smtp.PlainAuth("", from, password, host)
	//send email
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}

}
