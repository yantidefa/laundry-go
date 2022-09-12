package handlers

import (
	config "Laundry/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Paket struct {
	Id          int    `form:"id" json:"id"`
	Nama_paket  string `form:"nama_paket" json:"nama_paket"`
	Jenis_paket string `form:"jenis_paket" json:"jenis_paket"`
	Id_outlet   int    `form:"id_outlet" json:"id_outlet"`
	Harga       string `form:"harga" json:"harga"`
}

func CreatePaket(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var paket Paket
	json.Unmarshal(payloads, &paket)

	db, _ := config.DB()
	db.Create(&paket)

	res := Result{Code: 200, Data: paket, Message: "Success create paket"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetPakets(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: pakets")

	db, _ := config.DB()
	pakets := []Paket{}
	db.Find(&pakets)

	res := Result{Code: 200, Data: pakets, Message: "Success get paket"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetPaket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paketID := vars["id"]

	var paket Paket

	db, _ := config.DB()
	db.First(&paket, paketID)

	res := Result{Code: 200, Data: paket, Message: "Success get paket"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdatePaket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paketID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var paketUpdate Paket
	json.Unmarshal(payloads, &paketUpdate)

	var paket Paket

	db, _ := config.DB()
	db.First(&paket, paketID)
	db.Model(&paket).Updates(paketUpdate)

	res := Result{Code: 200, Data: paket, Message: "Success update paket"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeletePaket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paketID := vars["id"]

	var paket Paket

	db, _ := config.DB()
	db.First(&paket, paketID)
	db.Delete(&paket)

	res := Result{Code: 200, Message: "Success delete paket"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
