package ErrorHandler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ErrorMessage struct {
	ErrorMessage string `json:"error"`
	Code         string `json:"code "`
}

func Response(w http.ResponseWriter, status int, message string, code string) {
	w.WriteHeader(status)
	var E ErrorMessage
	E.ErrorMessage = message
	E.Code = code
	json.NewEncoder(w).Encode(E)
}

func Response1(w http.ResponseWriter, status int, message string, name string) {
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]string{
		"Message    ": message,
		"ProductName": name,
	})
}

func Response2(w http.ResponseWriter, status int, message string, name string, total int) {
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]string{
		"Message    ": message,
		"ProductName": name,
		"Total bill ": strconv.Itoa(total),
	})
}
func Response3(w http.ResponseWriter, status int, message string, total int) {
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]string{
		"Message    ": message,
		"Total bill ": strconv.Itoa(total),
	})
}
