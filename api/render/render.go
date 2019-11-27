package render

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, response interface{}, code int) {
	json, err := json.Marshal(response)
	if err != nil {
		ResponseError(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	if response != nil {
		w.Write(json)
	}
}

func ResponseError(w http.ResponseWriter, err error, code int) {
	json, err := json.Marshal(ErrorResponse{code, err.Error()})
	if err != nil {
		ResponseError(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	w.Write(json)
}
