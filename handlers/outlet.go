package handlers

import (
	config "Laundry/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Outlet struct {
	Id            int    `form:"id" json:"id"`
	Nama_outlet   string `form:"nama_outlet" json:"nama_outlet"`
	Alamat_outlet string `form:"alamat_outlet" json:"alamat_outlet"`
	Telp_outlet   string `form:"telp_outlet" json:"telp_outlet"`
}

func CreateOutlet(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var outlet Outlet
	json.Unmarshal(payloads, &outlet)

	db, _ := config.DB()
	db.Create(&outlet)

	res := Result{Code: 201, Data: outlet, Message: "Success create outlet"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	// fmt.Fprintf(w, "Welcome!")
}

func GetOutlets(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get outlets")

	db, _ := config.DB()
	outlets := []Outlet{}
	db.Find(&outlets)

	res := Result{Code: 200, Data: outlets, Message: "Success get outlets"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetOutlet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	outletID := vars["id"]

	var outlet Outlet

	db, _ := config.DB()
	db.First(&outlet, outletID)

	res := Result{Code: 200, Data: outlet, Message: "Success get user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateOutlet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	outletID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var outletUpdate Outlet
	json.Unmarshal(payloads, &outletUpdate)

	var outlet Outlet

	db, _ := config.DB()
	db.First(&outlet, outletID)
	db.Model(&outlet).Updates(outletUpdate)

	res := Result{Code: 200, Data: outlet, Message: "Success update user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteOutlet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	outletID := vars["id"]

	var outlet Outlet

	db, _ := config.DB()
	db.First(&outlet, outletID)
	db.Delete(&outlet)

	res := Result{Code: 200, Message: "Success delete outlet"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
