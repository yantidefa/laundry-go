package middlewares

import (
	config "Laundry/config"
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
)

func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		username, password, ok := request.BasicAuth()
		fmt.Println("username : ", username)
		fmt.Println("password : ", password)
		fmt.Println("ok : ", ok)

		if !ok || !CheckUsernameAndPassword(username, password) {
			response.Header().Set("wwww-Authenticte", `Basic realm="Account Invalid"`)
			response.WriteHeader(401)
			response.Write([]byte("Unauthorised\n"))
			return
		}
		handler(response, request)
	}
}

type User struct {
	Id       int    `form:"id" json:"id"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func CheckUsernameAndPassword(username, password string) bool {
	db, _ := config.DB()
	users := User{}
	db.First(&users).Where("username = ?", username)

	if users.Id == 0 {
		return false
	}
	if users.Password != password {
		return false
	}

	return true
}
