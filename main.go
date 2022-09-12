package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"Laundry/auth"
	"Laundry/handlers"
	"Laundry/middlewares"
)

var err error

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// func (user) TableName() string {
//     return "user"
// }

func main() {

	if err != nil {
		log.Println("koneksi gagal", err)
	} else {
		log.Println("koneksi berhasil")
	}

	handleRequests()
}

func handleRequests() {
	log.Println("Menjalankan Server http://172.0.0.1:9999")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", handlers.HomePage)
	//Handle User
	myRouter.HandleFunc("/api/user", handlers.CreateUser).Methods("POST")
	myRouter.HandleFunc("/api/user", handlers.GetUsers).Methods("GET")
	myRouter.HandleFunc("/api/user/{id}", handlers.GetUser).Methods("GET")
	myRouter.HandleFunc("/api/user/{id}", handlers.UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/api/user/{id}", handlers.DeleteUser).Methods("DELETE")
	//Handle User

	//Handle Outlet
	myRouter.HandleFunc("/api/outlet", handlers.CreateOutlet).Methods("POST")
	myRouter.HandleFunc("/api/outlet", handlers.GetOutlets).Methods("GET")
	myRouter.HandleFunc("/api/outlet/{id}", handlers.GetOutlet).Methods("GET")
	myRouter.HandleFunc("/api/outlet/{id}", handlers.UpdateOutlet).Methods("PUT")
	myRouter.HandleFunc("/api/outlet/{id}", handlers.DeleteOutlet).Methods("DELETE")
	//Handle Outlet

	//Handle Member
	myRouter.HandleFunc("/api/member", handlers.CreateMember).Methods("POST")
	myRouter.HandleFunc("/api/member", handlers.GetMembers).Methods("GET")
	myRouter.HandleFunc("/api/member/{id}", handlers.GetMember).Methods("GET")
	myRouter.HandleFunc("/api/member/{id}", handlers.UpdateMember).Methods("PUT")
	myRouter.HandleFunc("/api/member/{id}", handlers.DeleteMember).Methods("DELETE")
	//Handle Member

	//Handle Paket
	myRouter.HandleFunc("/api/paket", handlers.CreatePaket).Methods("POST")
	myRouter.HandleFunc("/api/paket", handlers.GetPakets).Methods("GET")
	myRouter.HandleFunc("/api/paket/{id}", handlers.GetPaket).Methods("GET")
	myRouter.HandleFunc("/api/paket/{id}", handlers.UpdatePaket).Methods("PUT")
	myRouter.HandleFunc("/api/paket/{id}", handlers.DeletePaket).Methods("DELETE")
	//Handle Paket

	//Handle Transaksi
	myRouter.HandleFunc("/api/transaksi", handlers.CreateTransaksi).Methods("POST")
	myRouter.HandleFunc("/api/transaksi", handlers.GetTransaksis).Methods("GET")
	myRouter.HandleFunc("/api/transaksi/{id}", handlers.GetTransaksi).Methods("GET")
	myRouter.HandleFunc("/api/transaksi/{id}", handlers.UpdateTransaksi).Methods("PUT")
	myRouter.HandleFunc("/api/transaksi/{id}", handlers.DeleteTransaksi).Methods("DELETE")
	//Handle Transaksi

	//Handle detailTransaksi
	myRouter.HandleFunc("/api/detailtransaksi", handlers.CreateDetailtransaksi).Methods("POST")
	myRouter.HandleFunc("/api/detailtransaksi", handlers.GetDetailtransaksis).Methods("GET")
	myRouter.HandleFunc("/api/detailtransaksi/{id}", handlers.GetDetailtransaksi).Methods("GET")
	myRouter.HandleFunc("/api/detailtransaksi/{id}", handlers.UpdateDetailtransaksi).Methods("PUT")
	myRouter.HandleFunc("/api/detailtransaksi/{id}", handlers.DeleteDetailtransaksi).Methods("DELETE")
	//Handle detailTransaksi

	myRouter.HandleFunc("/api/login", middlewares.BasicAuth(auth.FindAll)).Methods("GET")
	// myRouter.HandleFunc("/api/search", middlewares.BasicAuth(auth.Search)).Methods("GET")

	log.Fatal(http.ListenAndServe(":9999", myRouter))
}
