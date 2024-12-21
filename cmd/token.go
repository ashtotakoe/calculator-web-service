package main

type Token struct {
	tokenType      string
	textValue      string
	hasNumberValue bool
	numberValue    float64
}

func newEmptyToken() Token {
	return Token{
		tokenType: emptyTokenType,
	}
}
