package api

import (
	"encoding/json"
	"net/http"
)

func Data(w http.ResponseWriter, statusCode int, contentType string, data []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	w.Write(data)
}

func BindJson(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func JsonOK(w http.ResponseWriter, payload any) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

	return nil
}

func JsonErr(w http.ResponseWriter, statusCode int, cause error) error {
	jsonData, err := json.Marshal(cause.Error())
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonData)

	return nil
}
