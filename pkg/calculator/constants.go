package calculator

import "errors"

const (
	emptyTokenType     = "empty"
	operationTokenType = "operator"
	numberTokenType    = "number"
)

var operations = map[string]func(float64, float64) (float64, error){
	"+": func(a, b float64) (float64, error) { return a + b, nil },
	"-": func(a, b float64) (float64, error) { return a - b, nil },
	"*": func(a, b float64) (float64, error) { return a * b, nil },

	"/": func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, errors.New("cannot devide by 0")
		}
		return a / b, nil
	},
}
