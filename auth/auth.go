package auth

import (
	"fmt"
	"net/http"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Find All User")
}

func Search(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Search User")
}
