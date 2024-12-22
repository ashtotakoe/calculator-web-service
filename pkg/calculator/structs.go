package calculator

type Token struct {
	tokenType      string
	textValue      string
	hasNumberValue bool
	numberValue    float64
}

type Result struct {
	NumberValue float64
	TextValue   string
}

func newEmptyToken() Token {
	return Token{
		tokenType: emptyTokenType,
	}
}
