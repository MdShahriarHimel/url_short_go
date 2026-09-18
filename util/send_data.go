package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendData(w http.ResponseWriter, status int, data any) {
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		log.Println("Error Sending Data", err)
	}
}
