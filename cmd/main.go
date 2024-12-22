package main

import (
	"calculator/pkg/calculator"
	"fmt"
	"os"
)

func main() {
	expression := os.Args[1]
	res, err := calculator.Calc(expression)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
