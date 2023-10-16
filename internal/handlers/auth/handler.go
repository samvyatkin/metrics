package auth

import (
	"net/http"
)

func Handle(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusUnauthorized)
	res.Write([]byte("Куда лезешь, собака сутулая?"))
}
