package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/erizkiatama/prospace-assignment/calculator"
	"github.com/erizkiatama/prospace-assignment/database"
	"github.com/erizkiatama/prospace-assignment/parser"
)

func runIntergalacticConverter(db database.Database, calc calculator.Calculator, reader io.Reader) []string {
	responses := make([]string, 0)

	scanner := bufio.NewScanner(reader)
	for {
		scanned := scanner.Scan()
		if !scanned {
			break
		}

		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		// Process the line here
		parsed := parser.Parse(strings.ToLower(line))
		switch parsed.InputType {
		case parser.Assignment:
			if parsed.ItemType == parser.Roman {
				db.AddUnitToRomanMapping(parsed.FirstToken[0], parsed.FirstToken[2])
			} else {
				unitResult, err := calc.ConvertUnitsToInt(parsed.FirstToken)
				if err != nil {
					responses = append(responses, err.Error())
					break
				}
				db.AddCurrencyToCreditsMapping(parsed.FirstCurrency, float32(parsed.Credits)/float32(unitResult))
			}
		case parser.Calculation:
			if parsed.ItemType == parser.Roman {
				result, err := calc.ConvertUnitsToInt(parsed.FirstToken)
				if err != nil {
					responses = append(responses, err.Error())
					break
				}
				responses = append(responses, fmt.Sprintf("%s is %d", strings.Join(parsed.FirstToken, " "), result))
			} else {
				unitResult, err := calc.ConvertUnitsToInt(parsed.FirstToken)
				if err != nil {
					responses = append(responses, err.Error())
					break
				}
				result, err := calc.CalculateCreditsCurrency(float32(unitResult), parsed.FirstCurrency)
				if err != nil {
					responses = append(responses, err.Error())
					break
				}
				responses = append(responses, fmt.Sprintf(
					"%s %s is %.2f Credits",
					strings.Join(parsed.FirstToken, " "), parsed.FirstCurrency, result,
				))
			}
		case parser.Comparison:
			if parsed.ItemType == parser.Roman {
				result, err := calc.CompareTwoUnits(parsed.FirstToken, parsed.SecondToken)
				if err != nil {
					responses = append(responses, err.Error())
					break
				}
				responses = append(responses, fmt.Sprintf("%s is %s %s",
					strings.Join(parsed.FirstToken, " "), result, strings.Join(parsed.SecondToken, " ")),
				)
			} else {
				result, err := calc.CompareTwoCurrency(parsed.FirstToken, parsed.SecondToken, parsed.FirstCurrency, parsed.SecondCurrency)
				if err != nil {
					responses = append(responses, err.Error())
					break
				}
				responses = append(responses, fmt.Sprintf("%s %s %s %s %s",
					strings.Join(parsed.FirstToken, " "),
					parsed.FirstCurrency,
					result,
					strings.Join(parsed.SecondToken, " "),
					parsed.SecondCurrency),
				)
			}
		default:
			responses = append(responses, parsed.Error.Error())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return responses
}
