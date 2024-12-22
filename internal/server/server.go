package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ashtotakoe/calculator-web-service/pkg/calculator"
)

const (
	internalServerError    = "Internal server error"
	invalidExpressionError = "Expression is not valid"
)

var isDetailedValidationResponse = false

func handleExpression(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			writeToResponse(w, &ErrorBody{internalServerError}, http.StatusInternalServerError)
		}
	}()

	req := Request{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)

	if err != nil {
		log.Println("body decoder: ", err)
		writeToResponse(w, &ErrorBody{invalidExpressionError}, http.StatusUnprocessableEntity)

		return
	}

	result, err := calculator.Calc(req.Expression)

	if err != nil {
		log.Println("calculator: ", err)

		if isDetailedValidationResponse {
			writeToResponse(w, &ErrorBody{err.Error()}, http.StatusUnprocessableEntity)
			return
		}

		writeToResponse(w, &ErrorBody{invalidExpressionError}, http.StatusUnprocessableEntity)
		return
	}

	writeToResponse(w, &ResultBody{result.TextValue}, http.StatusOK)

}

func NewServer(dV bool) *http.ServeMux {
	isDetailedValidationResponse = dV

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/calculate", handleExpression)

	return mux
}
