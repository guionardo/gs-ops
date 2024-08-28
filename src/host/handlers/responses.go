package handlers

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, payload interface{}, statusCode int) (err error) {
	w.WriteHeader(statusCode)
	if statusCode != http.StatusNoContent {
		var body []byte
		if body, err = json.Marshal(payload); err == nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			_, err = w.Write(body)
		}
	}
	return
}
