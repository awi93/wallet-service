package httputil

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type HandleFunc func(r *http.Request) Response

type Response struct {
	StatusCode int
	Body       interface{}
}

func buildJsonResponse(response interface{}) []byte {
	res, _ := json.Marshal(response)
	var dest bytes.Buffer
	json.Compact(&dest, res)
	return dest.Bytes()
}

func HandleRequest(w http.ResponseWriter, r *http.Request, Func HandleFunc) {
	resp := Func(r)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if resp.StatusCode == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(resp.StatusCode)
	}
	res := buildJsonResponse(resp.Body)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write(res)
	if err != nil {
		log.Printf("Fail to Write Response")
	}
}
