package handlers

import (
	config "Laundry/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Member struct {
	Id            int    `form:"id" json:"id"`
	Nama_member   string `form:"nama_member" json:"nama_member"`
	Alamat_member string `form:"alamat_member" json:"alamat_member"`
	Jenis_kelamin string `form:"jenis_kelamin" json:"jenis_kelamin"`
	Telp_member   string `form:"telp_member" json:"telp_member"`
}

func CreateMember(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var member Member
	json.Unmarshal(payloads, &member)

	db, _ := config.DB()
	db.Create(&member)

	res := Result{Code: 200, Data: member, Message: "Success create member"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get members")

	db, _ := config.DB()
	members := []Member{}
	db.Find(&members)

	res := Result{Code: 200, Data: members, Message: "Success get members"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := vars["id"]

	var member Member

	db, _ := config.DB()
	db.First(&member, memberID)

	res := Result{Code: 200, Data: member, Message: "Success get member"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var memberUpdate Member
	json.Unmarshal(payloads, &memberUpdate)

	var member Member

	db, _ := config.DB()
	db.First(&member, memberID)
	db.Model(&member).Updates(memberUpdate)

	res := Result{Code: 200, Data: member, Message: "Success update member"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := vars["id"]

	var member Member

	db, _ := config.DB()
	db.First(&member, memberID)
	db.Delete(&member)

	res := Result{Code: 200, Message: "Success delete member"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
