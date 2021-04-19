package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		// ToDo: Special handling here? Throw 500?
		panic(fmt.Sprintf("failed to create response body, %v", err))
	}
}

func ResponseEmpty(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}
