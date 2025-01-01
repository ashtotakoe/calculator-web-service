package calc_server

import (
	"encoding/json"
	"net/http"
)

type CalcRequest struct {
	Expression string `json:"expression"`
}

func extractCalcRequest(r *http.Request) (*CalcRequest, error) {
	req := &CalcRequest{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(req)

	if err != nil {
		return nil, err
	}

	return req, nil

}
