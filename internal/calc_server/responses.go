package calc_server

import (
	"encoding/json"
	"net/http"
)

type ResultBody struct {
	Result string `json:"result"`
}

type ErrorBody struct {
	Error string `json:"error"`
}

func writeToResponse(w http.ResponseWriter, responseBody interface{}, status int) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	err := encoder.Encode(responseBody)

	if err != nil {
		return err
	}

	return nil
}
