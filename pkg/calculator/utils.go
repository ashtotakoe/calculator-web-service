package calculator

import (
	"strings"
)

func cleanExpression(expression *[]Token) []Token {
	res := make([]Token, 0, len(*expression))

	for _, token := range *expression {

		if token.tokenType != emptyTokenType {
			res = append(res, token)
		}
	}

	return res
}

func containsString(slice []string, item string) bool {

	for _, value := range slice {

		if value == item {

			return true

		}

	}

	return false
}

func containsTokensValue(tokens *[]Token, value string) bool {
	for _, token := range *tokens {
		if token.textValue == value {
			return true
		}
	}
	return false
}

func formatExpression(expression string) string {

	expression = strings.ReplaceAll(expression, " ", "")
	expression = strings.ReplaceAll(expression, ",", ".")

	return expression
}

func eraseFromSlice(slice []Token) []Token {
	for index := range slice {
		slice[index] = newEmptyToken()
	}

	return slice
}
