package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func tokenize(expression string) ([]Token, error) {
	operators := []string{"(", ")", "+", "-", "*", "/"}

	parsedExpression := strings.Split(expression, "")
	textTokens := make([]string, 0, len(parsedExpression))

	tempStorage := ""

	for index, elem := range parsedExpression {

		if containsString(operators, elem) {

			if tempStorage != "" {
				textTokens = append(textTokens, tempStorage)
			}
			tempStorage = ""

			textTokens = append(textTokens, elem)

			continue
		}

		if index == len(parsedExpression)-1 {

			if tempStorage != "" {
				textTokens = append(textTokens, tempStorage+elem)
			} else {

				textTokens = append(textTokens, elem)
			}

			continue
		}

		tempStorage += elem

	}

	parsedTokens := make([]Token, 0, len(textTokens))
	for _, token := range textTokens {

		isOperator := containsString(operators, token)

		if isOperator {
			parsedTokens = append(parsedTokens, Token{
				tokenType:      "operator",
				textValue:      token,
				hasNumberValue: false,
				numberValue:    0,
			})

			continue
		}

		numberValue, err := strconv.ParseFloat(token, 64)

		if err != nil {
			return nil, fmt.Errorf("failed to parse token %s: %s", token, err.Error())

		}

		parsedTokens = append(parsedTokens, Token{
			tokenType:      "number",
			textValue:      token,
			hasNumberValue: true,
			numberValue:    numberValue,
		})

	}

	return parsedTokens, nil
}

func evaluateExpression(expression []Token) (Token, error) {
	err := openExpressionBrackets(&expression)

	if err != nil {
		return newEmptyToken(), err
	}

	expression = solveUnaryOperators(expression)

	err = scanForMathOperators(&expression, func(expression []Token) bool {
		return containsTokensValue(expression, "*") || containsTokensValue(expression, "/")
	})

	if err != nil {
		return newEmptyToken(), err
	}

	err = scanForMathOperators(&expression, func(expression []Token) bool {
		return containsTokensValue(expression, "+") || containsTokensValue(expression, "-")
	})

	if err != nil {
		return newEmptyToken(), err
	}

	if len(expression) == 1 {
		return expression[0], nil
	}

	return newEmptyToken(), errors.New("something is wrong with expression")
}

func solveUnaryOperators(expression []Token) []Token {
	for i := len(expression) - 1; i >= 0; i-- {
		token := expression[i]

		if token.textValue != "+" && token.textValue != "-" {
			continue
		}

		// check if there is a number after the operator, but not before to handle (+90) or (-5) correctly
		if i >= len(expression)-1 || expression[i+1].tokenType != "number" {
			continue
		}

		// we shall skip the situations like (1 + 2) for now
		if i != 0 && expression[i-1].tokenType == "number" {
			continue
		}

		if token.textValue == "-" {
			expression[i+1].numberValue = -expression[i+1].numberValue
		}

		expression[i] = newEmptyToken()

	}

	return cleanExpression(expression)
}

func openExpressionBrackets(expression *[]Token) error {
	stackLen := 0
	entryIndex := 0
	isOpeningBracketFound := false

	for index, token := range *expression {
		if token.tokenType != "operator" {
			continue
		}

		if token.textValue == ")" && !isOpeningBracketFound {
			return errors.New("bracket is not closed")
		}

		if token.textValue == "(" && !isOpeningBracketFound {
			entryIndex = index
			isOpeningBracketFound = true
			stackLen += 1
			continue
		}

		if token.textValue == "(" {
			stackLen += 1
			continue
		}

		if token.textValue == ")" {
			stackLen -= 1
			if stackLen < 0 {
				return errors.New("too many closing brackets")
			}

			if stackLen == 0 {
				expressionSlice := make([]Token, index-entryIndex-1)
				copy(expressionSlice, (*expression)[entryIndex+1:index])

				expressionCompute, err := evaluateExpression(expressionSlice)

				if err != nil {
					return err
				}

				eraseFromSlice((*expression)[entryIndex : index+1])
				(*expression)[index] = expressionCompute

				isOpeningBracketFound = false
			}
		}
	}

	*expression = cleanExpression(*expression)

	return nil

}

func scanForMathOperators(expression *[]Token, operatorsCheckFunc func(expression []Token) bool) error {
	for operatorsCheckFunc(*expression) {

		for index, token := range *expression {
			if !operatorsCheckFunc([]Token{token}) {
				continue
			}

			operator := (*expression)[index]

			if (index-1) < 0 || (index+1) >= len(*expression) {
				return fmt.Errorf("operator %v is not in the correct position", (*expression)[index].textValue)
			}

			operand1, operand2 := (*expression)[index-1], (*expression)[index+1]

			if operand1.tokenType != "number" || operand2.tokenType != "number" || operator.tokenType != "operator" {
				return fmt.Errorf("something is wrong with expression %s %s %s", operand1.textValue, operator.textValue, operand2.textValue)
			}

			result, err := conductArithmeticOperation(operand1.numberValue, operand2.numberValue, operator.textValue)
			if err != nil {
				return err
			}

			(*expression)[index-1] = newEmptyToken()
			(*expression)[index+1] = newEmptyToken()
			(*expression)[index] = result

			break
		}

		*expression = cleanExpression(*expression)
	}

	return nil
}

func Calc(expression string) (float64, error) {

	formattedExpression := formatExpression(expression)

	tokens, err := tokenize(formattedExpression)

	if err != nil {
		return 0, err
	}

	result, err := evaluateExpression(tokens)

	if err != nil {
		return 0, err
	}

	return result.numberValue, nil
}

func conductArithmeticOperation(val1, val2 float64, operator string) (Token, error) {
	operation, ok := operations[operator]

	if !ok {
		return newEmptyToken(), errors.New("math operator does not exist")
	}

	res, err := operation(val1, val2)

	if err != nil {
		return newEmptyToken(), err
	}

	return Token{
		tokenType:      "number",
		hasNumberValue: true,
		textValue:      strconv.FormatFloat(res, 'f', -1, 64),
		numberValue:    res,
	}, nil

}
