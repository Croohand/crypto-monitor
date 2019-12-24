package helpers

import (
	"encoding/json"
	"net/http"
)

// JSON format
func WriteHttp(w http.ResponseWriter, src interface{}) {
	data, err := json.Marshal(src)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
