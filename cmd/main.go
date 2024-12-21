package main

import (
	"calculator/pkg/calculator"
	"fmt"
)

func main() {
	// expression := os.Args[1]
	// res, err := Calc(expression)

	// res, err := calculator.Calc("((1+4) * (1+2) +10) *4")
	res, err := calculator.Calc("-90 + 90")

	if err != nil {
		panic(err)
	}
	fmt.Println(int(res))

}
