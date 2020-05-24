package common

import (
	"encoding/json"
	"log"
	"net/http"
)

// WriteError send a response with a status code.
func WriteError(w http.ResponseWriter, code int, e string) {
	w.WriteHeader(code)
	values := map[string]string{
		"error": e,
	}
	valuesBytes, _ := json.Marshal(values)
	_, err := w.Write(valuesBytes)
	if err != nil {
		log.Println(err.Error())
	}
}
