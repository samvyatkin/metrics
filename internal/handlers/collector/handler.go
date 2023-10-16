package collector

import (
	"metrics/internal/database"
	"net/http"
	"strconv"
	"strings"
)

const (
	gauge   = "gauge"
	counter = "counter"
)

type CollectorHandler struct {
	DB *database.MemStorage
}

func (h *CollectorHandler) Handle(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		writeStatus(res, http.StatusMethodNotAllowed)
		return
	}

	p := strings.Split(req.URL.Path, "/")[1:]
	n := len(p)

	switch {
	case n == 4 && p[1] == gauge && p[2] != "" && p[3] != "":
		handleGauge(res)
	case n == 4 && p[1] == counter && p[2] != "" && p[3] != "":
		handleCounter(res, p[2], p[3])
	default:
		http.NotFound(res, req)
	}
}

func handleGauge(res http.ResponseWriter) {
	writeStatus(res, http.StatusOK)
}

func handleCounter(res http.ResponseWriter, name, value string) {
	_, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		writeStatus(res, http.StatusBadRequest)
		return
	}

	writeStatus(res, http.StatusOK)
}

func writeStatus(res http.ResponseWriter, status int) {
	res.WriteHeader(status)
	res.Write([]byte(http.StatusText(status)))
}
