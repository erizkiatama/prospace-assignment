package database

import (
	"errors"
	"strings"
)

type Database interface {
	AddUnitToRomanMapping(string, string)
	GetRomanFromUnit(string) (string, error)
	AddCurrencyToCreditsMapping(string, float32)
	GetCreditsFromCurrency(string) (float32, error)
}

type database struct {
	unitToRomanValues      map[string]string
	currencyToCreditValues map[string]float32
}

func NewDatabase() Database {
	return &database{
		unitToRomanValues:      make(map[string]string),
		currencyToCreditValues: make(map[string]float32),
	}
}

func (db *database) AddUnitToRomanMapping(unit, roman string) {
	db.unitToRomanValues[strings.ToLower(unit)] = strings.ToUpper(roman)
}

func (db *database) GetRomanFromUnit(unit string) (string, error) {
	if roman, exists := db.unitToRomanValues[strings.ToLower(unit)]; exists {
		return roman, nil
	}

	return "", errors.New(unit + " unit is not defined in the intergalactic database")
}

func (db *database) AddCurrencyToCreditsMapping(currency string, credits float32) {
	db.currencyToCreditValues[strings.ToLower(currency)] = credits
}

func (db *database) GetCreditsFromCurrency(currency string) (float32, error) {
	if credits, exists := db.currencyToCreditValues[strings.ToLower(currency)]; exists {
		return credits, nil
	}

	return 0, errors.New(currency + " currency is not defined in the intergalactic database")
}
