package calculator

import "errors"

type TokenType = string

const (
	emptyTokenType     TokenType = "empty"
	operationTokenType TokenType = "operator"
	numberTokenType    TokenType = "number"
)

var operations = map[string]func(float64, float64) (float64, error){
	"+": func(a, b float64) (float64, error) { return a + b, nil },
	"-": func(a, b float64) (float64, error) { return a - b, nil },
	"*": func(a, b float64) (float64, error) { return a * b, nil },

	"/": func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, errors.New("cannot divide by 0")
		}
		return a / b, nil
	},
}
