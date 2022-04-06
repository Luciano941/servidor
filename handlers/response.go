package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendData(rw http.ResponseWriter, data interface{}, status int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	output, _ := json.Marshal(&data)
	fmt.Fprintln(rw, string(output))
}

func SendError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	fmt.Fprintln(rw, "Resource not found")
}
