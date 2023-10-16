package collector

import (
	"errors"
	"metrics/internal/database"
	"net/http"
	"strconv"
	"strings"
)

var (
	ErrMetricMethod = errors.New("wrong metric methid")
	ErrMetricType   = errors.New("wrong metric type")
	ErrMetricName   = errors.New("wrong metric name")
	ErrMetricValue  = errors.New("wrong metric value")
)

var (
	ErrNotFound = errors.New("value not found")
)

const (
	method  = "update"
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
	mType, mName, mValue, err := parse(p)

	if err != nil {
		switch {
		case errors.Is(err, ErrMetricMethod):
			http.NotFound(res, req)
		case errors.Is(err, ErrMetricType):
			writeStatus(res, http.StatusBadRequest)
		case errors.Is(err, ErrMetricName):
			http.NotFound(res, req)
		case errors.Is(err, ErrMetricValue):
			writeStatus(res, http.StatusBadRequest)
		}

		return
	}

	switch {
	case *mType == gauge:
		handleGauge(res, *mName, *mValue)
	case *mType == counter:
		handleCounter(res, *mName, *mValue)
	default:
		writeStatus(res, http.StatusBadRequest)
	}
}

func parse(path []string) (*string, *string, *string, error) {
	mMethod, err := find(0, path)
	if err != nil || *mMethod != method {
		return nil, nil, nil, ErrMetricMethod
	}

	mType, err := find(1, path)
	if err != nil {
		return nil, nil, nil, ErrMetricType
	}

	mName, err := find(2, path)
	if err != nil || *mName == "" {
		return nil, nil, nil, ErrMetricName
	}

	mValue, err := find(3, path)
	if err != nil || *mValue == "" {
		return nil, nil, nil, ErrMetricName
	}

	return mType, mName, mValue, nil
}

func handleGauge(res http.ResponseWriter, name, value string) {
	_, err := strconv.ParseFloat(value, 64)

	if err != nil {
		writeStatus(res, http.StatusBadRequest)
		return
	}

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

func find(index int, path []string) (*string, error) {
	for i, v := range path {
		if i == index {
			return &v, nil
		}
	}

	return nil, ErrNotFound
}
