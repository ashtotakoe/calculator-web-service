package calc_server

import (
	"log"
	"net/http"

	"github.com/ashtotakoe/calculator-web-service/pkg/calculator"
)

const (
	internalServerError    = "Internal server error"
	invalidExpressionError = "Expression is not valid"
)

type ServerConf struct {
	DetailedErrors bool
}

var serverConfig ServerConf

func handleExpression(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			writeToResponse(w, &ErrorBody{internalServerError}, http.StatusInternalServerError)
		}
	}()

	req, err := extractCalcRequest(r)

	if err != nil {
		log.Println("body decoder: ", err)
		writeToResponse(w, &ErrorBody{invalidExpressionError}, http.StatusUnprocessableEntity)

		return
	}

	result, err := calculator.Calc(req.Expression)

	if err != nil {
		log.Println("calculator: ", err)

		if serverConfig.DetailedErrors {
			writeToResponse(w, &ErrorBody{err.Error()}, http.StatusUnprocessableEntity)
			return
		}

		writeToResponse(w, &ErrorBody{invalidExpressionError}, http.StatusUnprocessableEntity)
		return
	}

	writeToResponse(w, &ResultBody{result.TextValue}, http.StatusOK)

}

func NewServer(config ServerConf) *http.ServeMux {

	serverConfig = config

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/calculate", handleExpression)

	return mux
}
