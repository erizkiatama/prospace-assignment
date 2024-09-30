package main

import (
	"fmt"
	"github.com/erizkiatama/prospace-assignment/calculator"
	"github.com/erizkiatama/prospace-assignment/database"
	"os"
)

func main() {
	db := database.NewDatabase()
	calc := calculator.NewCalculator(db)

	responses := runIntergalacticConverter(db, calc, os.Stdin)
	for _, response := range responses {
		fmt.Println(response)
	}
}
