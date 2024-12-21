package calculator

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func cleanExpression(expression []Token) []Token {
	res := make([]Token, 0, len(expression))

	for _, token := range expression {

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

func containsTokensValue(tokens []Token, value string) bool {
	for _, token := range tokens {
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

func printArr(arr []Token) {
	fmt.Println()
	fmt.Println()

	for _, token := range arr {
		display := token.textValue

		if token.tokenType == emptyTokenType {
			display = " "
		}

		fmt.Print("  ", display, "  ")
	}
	fmt.Println()
	fmt.Println()
	for i := range arr {
		fmt.Print("  ", i, "  ")
	}
	fmt.Println()
	fmt.Println()

}
