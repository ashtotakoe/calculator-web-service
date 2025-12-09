package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/ashtotakoe/calculator-web-service/internal/calc_server"
)

const expressionHandlerURL = "/api/v1/calculate"

func TestServer(t *testing.T) {
	s := httptest.NewServer(calc_server.NewServer(calc_server.ServerConfig{
		DetailedErrors: false,
	}))
	defer s.Close()

	url := s.URL

	for _, testCase := range validTestCases {

		t.Run(formatExpression(testCase.expression), func(t *testing.T) {
			response := sendRequest(url+expressionHandlerURL, testCase.expression)

			if response.StatusCode != http.StatusOK {
				t.Errorf("server responded with invalid response code: %v", response.StatusCode)
				return
			}

			parsedBody := &calc_server.ResultBody{}

			decoder := json.NewDecoder(response.Body)
			err := decoder.Decode(parsedBody)

			if err != nil {
				t.Errorf("failed to parse the request body:  %s", err.Error())
				return
			}

			result, err := strconv.ParseFloat(parsedBody.Result, 64)

			if err != nil {
				t.Errorf("failed to parse the request body:  %s", err.Error())
				return
			}

			if result != testCase.expected {
				t.Errorf("%f (expected) is not equal to %f (received)", testCase.expected, result)
			}
		})
	}

	for _, testCase := range failTestCases {

		t.Run(formatExpression(testCase.expression), func(t *testing.T) {
			response := sendRequest(url+expressionHandlerURL, testCase.expression)

			if response.StatusCode != http.StatusUnprocessableEntity {
				t.Errorf("server responded with invalid response code: %v, \n expected: %v", response.StatusCode, http.StatusUnprocessableEntity)
				return
			}

			parsedBody := &calc_server.ResultBody{}

			decoder := json.NewDecoder(response.Body)
			err := decoder.Decode(parsedBody)

			if err != nil {
				t.Errorf("failed to parse the request body:  %s", err.Error())
				return
			}

		})
	}
}

func sendRequest(url, expression string) *http.Response {
	reqBody := &calc_server.CalcRequest{
		Expression: expression,
	}
	reqBodyBytes, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	response, _ := http.DefaultClient.Do(req)

	return response
}
