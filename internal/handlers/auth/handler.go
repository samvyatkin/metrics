package auth

import (
	"net/http"
)

func Handle(res http.ResponseWriter, req *http.Request) {
	http.NotFound(res, req)
}
