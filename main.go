package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ashtotakoe/calculator-web-service/internal/calc_server"
)

func main() {

	port := os.Args[1]

	runWithDetailedErrors := false

	if len(os.Args) > 2 {
		runWithDetailedErrors = os.Args[2] == "detailed-validation"
	}

	s := calc_server.NewServer(
		calc_server.ServerConf{
			DetailedErrors: runWithDetailedErrors,
		})

	log.Printf("Server is running on port %s \n detailed validation = %t\n", port, runWithDetailedErrors)

	err := http.ListenAndServe(":"+port, s)

	if err != nil {
		log.Println(err)
	}
}
