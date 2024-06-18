package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}, message string) {
	response, err := json.Marshal(Response{
		Status:  status == http.StatusOK,
		Message: message,
		Data:    payload,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
