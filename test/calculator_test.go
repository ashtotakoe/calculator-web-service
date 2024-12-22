package test

import (
	"testing"

	"github.com/ashtotakoe/calculator-web-service/pkg/calculator"
)

func TestCalculator(t *testing.T) {

	for _, testCase := range validTestCases {
		t.Run("valid expressions", func(t *testing.T) {
			result, err := calculator.Calc(testCase.expression)

			if err != nil {
				t.Errorf("program with input %s should return %f, but returned an error: %s",
					testCase.expression, testCase.expected, err.Error())

				return
			}

			if result.NumberValue != testCase.expected {
				t.Errorf("program with input %s should return %f, but returned with %f",
					testCase.expression, testCase.expected, result.NumberValue)

				return
			}

		})

	}

	for _, testCase := range failTestCases {
		t.Run("invalid expressions", func(t *testing.T) {

			result, err := calculator.Calc(testCase.expression)

			if err == nil {
				t.Errorf("program with input %s should return an error, but returned the result %s",
					testCase.expression, result.TextValue)
			}

		})
	}
}
