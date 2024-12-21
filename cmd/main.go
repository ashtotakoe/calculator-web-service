package main

import (
	"calculator/pkg/calculator"
	"fmt"
)

func main() {

	// expression := os.Args[1]
	res, err := calculator.Calc("((1+4) * (1+2) +10) *4")

	// res, err := Calc(expression)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(int(res))
	}

}
