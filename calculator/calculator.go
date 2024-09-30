package calculator

import (
	"github.com/erizkiatama/prospace-assignment/constant"
	"github.com/erizkiatama/prospace-assignment/database"
	"strings"
)

type Calculator interface {
	ConvertUnitsToInt([]string) (int, error)
	CompareTwoUnits([]string, []string) (string, error)
	CalculateCreditsCurrency(unitResult float32, currency string) (float32, error)
	CompareTwoCurrency([]string, []string, string, string) (string, error)
}

type calculator struct {
	db database.Database
}

var (
	RomanValues = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	validSubtractions = map[byte][]byte{
		'I': {'V', 'X'},
		'X': {'L', 'C'},
		'C': {'D', 'M'},
	}
)

func NewCalculator(db database.Database) Calculator {
	return &calculator{db: db}
}

func (c *calculator) convertUnitToRoman(units []string) (string, error) {
	romanNumeral := ""
	for _, unit := range units {
		roman, err := c.db.GetRomanFromUnit(strings.ToLower(unit))
		if err != nil {
			return "", err
		}
		romanNumeral += roman
	}

	return romanNumeral, nil
}

func (c *calculator) convertRomanToInt(roman string) (int, error) {
	total := 0
	repeatCount := 0
	prevValue := 0
	prevSymbol := byte(0)

	for i := len(roman) - 1; i >= 0; i-- {
		currSymbol := roman[i]
		currValue := RomanValues[currSymbol]

		// Same roman symbol can only be repeated 3 times
		if currSymbol == prevSymbol {
			repeatCount++
			if repeatCount > 3 || (currSymbol == 'D' || currSymbol == 'L' || currSymbol == 'V') {
				return 0, constant.ErrInvalidFormat
			}
		} else {
			repeatCount = 1
		}

		// A roman symbol can only be subtracted by specific roman symbols
		if currValue >= prevValue {
			total += currValue
		} else {
			if !isValidSubtraction(currSymbol, prevSymbol) {
				return 0, constant.ErrInvalidFormat
			}
			total -= currValue
		}

		prevValue = currValue
		prevSymbol = currSymbol
	}

	return total, nil
}

func isValidSubtraction(smaller, larger byte) bool {
	allowed, exists := validSubtractions[smaller]
	if !exists {
		return false
	}

	for _, symbol := range allowed {
		if symbol == larger {
			return true
		}
	}

	return false
}

func (c *calculator) ConvertUnitsToInt(tokens []string) (int, error) {
	romanNumeral, err := c.convertUnitToRoman(tokens)
	if err != nil {
		return 0, err
	}

	return c.convertRomanToInt(romanNumeral)
}

func (c *calculator) CalculateCreditsCurrency(unitResult float32, currency string) (float32, error) {
	credits, err := c.db.GetCreditsFromCurrency(strings.ToLower(currency))
	if err != nil {
		return 0, err
	}

	return unitResult * credits, nil
}

func (c *calculator) getUnitResults(first, second []string) (int, int, error) {
	firstResult, err := c.ConvertUnitsToInt(first)
	if err != nil {
		return 0, 0, err
	}

	secondResult, err := c.ConvertUnitsToInt(second)
	if err != nil {
		return 0, 0, err
	}

	return firstResult, secondResult, err
}

func (c *calculator) CompareTwoUnits(firstUnits, secondUnits []string) (string, error) {
	firstResult, secondResult, err := c.getUnitResults(firstUnits, secondUnits)
	if err != nil {
		return "", err
	}

	if firstResult > secondResult {
		return "larger than", nil
	} else if firstResult < secondResult {
		return "smaller than", nil
	}
	return "equal to", nil
}

func (c *calculator) CompareTwoCurrency(firstUnits, secondUnits []string, firstCurrency, secondCurrency string) (string, error) {
	firstUnitResult, secondUnitResult, err := c.getUnitResults(firstUnits, secondUnits)
	if err != nil {
		return "", err
	}

	firstResult, err := c.CalculateCreditsCurrency(float32(firstUnitResult), firstCurrency)
	if err != nil {
		return "", err
	}

	secondResult, err := c.CalculateCreditsCurrency(float32(secondUnitResult), secondCurrency)
	if err != nil {
		return "", err
	}

	if firstResult > secondResult {
		return "has more credits than", nil
	} else if firstResult < secondResult {
		return "has less credits than", nil
	}
	return "has equal credits with", nil
}
