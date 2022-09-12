package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	config "Laundry/config"
)

// var db *gorm.DB

// const (
// 	db = gorm.Open("mysql", "root:0303@/laundry?charset=utf8&parseTime=True")
// )

type Detailtransaksi struct {
	Id           int    `form:"id" json:"id"`
	Id_transaksi int    `form:"id_transaksi" json:"id_transaksi"`
	Id_paket     int    `form:"id_paket" json:"id_paket"`
	Qty          int    `form:"qty" json:"qty"`
	Keterangan   string `form:"keterangan" json:"keterangan"`
}

func CreateDetailtransaksi(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var detailtransaksi Detailtransaksi
	json.Unmarshal(payloads, &detailtransaksi)

	db, _ := config.DB()
	db.Create(&detailtransaksi)

	res := Result{Code: 200, Data: detailtransaksi, Message: "Success create detailtransaksi"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetDetailtransaksis(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get detailtransaksis")

	db, _ := config.DB()
	detailtransaksis := []Detailtransaksi{}
	db.Find(&detailtransaksis)

	res := Result{Code: 200, Data: detailtransaksis, Message: "Success get detailtransaksis"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetDetailtransaksi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	detailtransaksiID := vars["id"]

	var detailtransaksi Detailtransaksi

	db, _ := config.DB()
	db.First(&detailtransaksi, detailtransaksiID)

	res := Result{Code: 200, Data: detailtransaksi, Message: "Success get detailtransaksi"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateDetailtransaksi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	detailtransaksiID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var detailtransaksiUpdate Detailtransaksi
	json.Unmarshal(payloads, &detailtransaksiUpdate)

	var detailtransaksi Detailtransaksi
	db, _ := config.DB()
	db.First(&detailtransaksi, detailtransaksiID)
	db.Model(&detailtransaksi).Updates(detailtransaksiUpdate)

	res := Result{Code: 200, Data: detailtransaksi, Message: "Success update detailtransaksi"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteDetailtransaksi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	detailtransaksiID := vars["id"]

	var detailtransaksi Detailtransaksi

	db, _ := config.DB()
	db.First(&detailtransaksi, detailtransaksiID)
	db.Delete(&detailtransaksi)

	res := Result{Code: 200, Message: "Success delete detailtransaksi"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
