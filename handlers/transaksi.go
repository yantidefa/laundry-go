package handlers

import (
	config "Laundry/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Transaksi struct {
	Id                int       `form:"id" json:"id"`
	Id_outlet         int       `form:"id_outlet" json:"id_outlet"`
	Id_member         int       `form:"id_member" json:"id_member"`
	Id_user           int       `form:"id_user" json:"id_user"`
	Kode_invoice      string    `form:"kode_invoice" json:"kode_invoice"`
	Tanggal           time.Time `form:"tanggal" json:"tanggal"`
	Tgl_bayar         time.Time `form:"tgl_bayar" json:"tgl_bayar"`
	Batas_waktu       time.Time `form:"batas_waktu" json:"batas_waktu"`
	Biaya_tambahan    int       `form:"biaya_tambahan" json:"biaya_tambahan"`
	Diskon            int       `form:"diskon" json:"diskon"`
	Pajak             int       `form:"pajak" json:"pajak"`
	Status_paket      string    `form:"status_paket" json:"status_paket"`
	Status_pembayaran string    `form:"status_pembayaran" json:"status_pembayaran"`
}

type TransaksiCreateStruct struct {
	Id_outlet      int    `form:"id_outlet" json:"id_outlet"`
	Id_member      int    `form:"id_member" json:"id_member"`
	Id_user        int    `form:"id_user" json:"id_user"`
	Kode_invoice   string `form:"kode_invoice" json:"kode_invoice"`
	Biaya_tambahan int    `form:"biaya_tambahan" json:"biaya_tambahan"`
}

func (t *Transaksi) BeforeCreate(tx *gorm.DB) (err error) {
	t.Tgl_bayar = time.Now()
	t.Tanggal = time.Now()
	t.Batas_waktu = time.Now().AddDate(0, 0, 3)
	t.Status_paket = "baru"
	t.Status_pembayaran = "belum dibayar"

	// err = errors.New("uwawjos")

	return
}

func CreateTransaksi(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var body TransaksiCreateStruct
	err := json.Unmarshal(payloads, &body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	transaksi := Transaksi{
		Id_outlet:      body.Id_outlet,
		Id_member:      body.Id_member,
		Id_user:        body.Id_user,
		Kode_invoice:   body.Kode_invoice,
		Biaya_tambahan: body.Biaya_tambahan,
	}

	fmt.Println(string(payloads))
	fmt.Println(transaksi)

	db, _ := config.DB()
	if err := db.Create(&transaksi).Error; err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := Result{Code: 200, Data: transaksi, Message: "Success create transaksi"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create detail transaksi

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func GetTransaksis(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: transaksis")

	transaksis := []Transaksi{}
	db, _ := config.DB()
	db.Find(&transaksis)

	res := Result{Code: 200, Data: transaksis, Message: "Success get transaksi"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetTransaksi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transaksiID := vars["id"]

	var transaksi Transaksi

	db, _ := config.DB()
	db.First(&transaksi, transaksiID)

	res := Result{Code: 200, Data: transaksi, Message: "Success get transaksi"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateTransaksi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transaksiID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var transaksiUpdate Transaksi
	json.Unmarshal(payloads, &transaksiUpdate)

	var transaksi Transaksi

	db, _ := config.DB()
	db.First(&transaksi, transaksiID)
	db.Model(&transaksi).Updates(transaksiUpdate)

	res := Result{Code: 200, Data: transaksi, Message: "Success update transaksi"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteTransaksi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transaksiID := vars["id"]

	var transaksi Transaksi

	db, _ := config.DB()
	db.First(&transaksi, transaksiID)
	db.Delete(&transaksi)

	res := Result{Code: 200, Message: "Success delete transaksi"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
