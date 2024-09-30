package calculator

import (
	"errors"
	"testing"
)

// Mock database for testing
type mockDB struct {
	unitToRoman       map[string]string
	currencyToCredits map[string]float32
}

func (m *mockDB) AddUnitToRomanMapping(unit, roman string) {
	m.unitToRoman[unit] = roman
}

func (m *mockDB) GetRomanFromUnit(unit string) (string, error) {
	if roman, ok := m.unitToRoman[unit]; ok {
		return roman, nil
	}
	return "", errors.New("unit not found")
}

func (m *mockDB) AddCurrencyToCreditsMapping(currency string, credits float32) {
	m.currencyToCredits[currency] = credits
}

func (m *mockDB) GetCreditsFromCurrency(currency string) (float32, error) {
	if credits, ok := m.currencyToCredits[currency]; ok {
		return credits, nil
	}
	return 0, errors.New("currency not found")
}

func newMockDatabase() *mockDB {
	return &mockDB{
		unitToRoman:       make(map[string]string),
		currencyToCredits: make(map[string]float32),
	}
}

func TestConvertUnitsToInt(t *testing.T) {
	mockDB := newMockDatabase()
	mockDB.AddUnitToRomanMapping("xyz", "I")
	mockDB.AddUnitToRomanMapping("abc", "V")
	mockDB.AddUnitToRomanMapping("def", "X")
	mockDB.AddUnitToRomanMapping("jkl", "L")
	mockDB.AddUnitToRomanMapping("rst", "M")

	calc := NewCalculator(mockDB)

	tests := []struct {
		input    []string
		expected int
		hasError bool
	}{
		{[]string{"xyz", "abc"}, 4, false},
		{[]string{"def", "jkl", "xyz", "xyz"}, 42, false},
		{[]string{"xyz", "xyz", "xyz", "xyz", "xyz"}, 0, true},
		{[]string{"jkl", "rst"}, 0, true},
		{[]string{"xyz", "jkl"}, 0, true},
		{[]string{"unknown"}, 0, true},
	}

	for _, test := range tests {
		result, err := calc.ConvertUnitsToInt(test.input)
		if test.hasError && err == nil {
			t.Errorf("Expected error for input %v, got none", test.input)
		}
		if !test.hasError && err != nil {
			t.Errorf("Unexpected error for input %v: %v", test.input, err)
		}
		if result != test.expected {
			t.Errorf("For input %v, expected %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestCalculateCreditsCurrency(t *testing.T) {
	mockDB := newMockDatabase()
	mockDB.AddCurrencyToCreditsMapping("gold", 14450.0)
	mockDB.AddCurrencyToCreditsMapping("silver", 17.0)

	calc := NewCalculator(mockDB)

	tests := []struct {
		unitResult float32
		currency   string
		expected   float32
		hasError   bool
	}{
		{2, "gold", 28900.0, false},
		{3, "silver", 51.0, false},
		{1, "unknown", 0, true},
	}

	for _, test := range tests {
		result, err := calc.CalculateCreditsCurrency(test.unitResult, test.currency)
		if test.hasError && err == nil {
			t.Errorf("Expected error for input %v %s, got none", test.unitResult, test.currency)
		}
		if !test.hasError && err != nil {
			t.Errorf("Unexpected error for input %v %s: %v", test.unitResult, test.currency, err)
		}
		if result != test.expected {
			t.Errorf("For input %v %s, expected %f, got %f", test.unitResult, test.currency, test.expected, result)
		}
	}
}

func TestCompareTwoUnits(t *testing.T) {
	mockDB := newMockDatabase()
	mockDB.AddUnitToRomanMapping("xyz", "I")
	mockDB.AddUnitToRomanMapping("abc", "V")
	mockDB.AddUnitToRomanMapping("def", "X")

	calc := NewCalculator(mockDB)

	tests := []struct {
		first    []string
		second   []string
		expected string
		hasError bool
	}{
		{[]string{"xyz", "abc"}, []string{"def"}, "smaller than", false},
		{[]string{"def"}, []string{"xyz", "abc"}, "larger than", false},
		{[]string{"xyz", "abc"}, []string{"xyz", "abc"}, "equal to", false},
		{[]string{"unknown"}, []string{"xyz"}, "", true},
		{[]string{"xyz"}, []string{"unknown"}, "", true},
	}

	for _, test := range tests {
		result, err := calc.CompareTwoUnits(test.first, test.second)
		if test.hasError && err == nil {
			t.Errorf("Expected error for input %v and %v, got none", test.first, test.second)
		}
		if !test.hasError && err != nil {
			t.Errorf("Unexpected error for input %v and %v: %v", test.first, test.second, err)
		}
		if result != test.expected {
			t.Errorf("For input %v and %v, expected %s, got %s", test.first, test.second, test.expected, result)
		}
	}
}

func TestCompareTwoCurrency(t *testing.T) {
	mockDB := newMockDatabase()
	mockDB.AddUnitToRomanMapping("xyz", "I")
	mockDB.AddUnitToRomanMapping("abc", "V")
	mockDB.AddCurrencyToCreditsMapping("gold", 14450.0)
	mockDB.AddCurrencyToCreditsMapping("silver", 17.0)

	calc := NewCalculator(mockDB)

	tests := []struct {
		firstUnits     []string
		secondUnits    []string
		firstCurrency  string
		secondCurrency string
		expected       string
		hasError       bool
	}{
		{[]string{"xyz", "abc"}, []string{"xyz"}, "gold", "silver", "has more credits than", false},
		{[]string{"xyz"}, []string{"xyz", "abc"}, "silver", "gold", "has less credits than", false},
		{[]string{"xyz"}, []string{"xyz"}, "gold", "gold", "has equal credits with", false},
		{[]string{"unknown"}, []string{"xyz"}, "gold", "silver", "", true},
		{[]string{"xyz"}, []string{"xyz"}, "unknown", "gold", "", true},
		{[]string{"xyz"}, []string{"xyz"}, "gold", "unknown", "", true},
	}

	for _, test := range tests {
		result, err := calc.CompareTwoCurrency(test.firstUnits, test.secondUnits, test.firstCurrency, test.secondCurrency)
		if test.hasError && err == nil {
			t.Errorf("Expected error for input %v, %v, %s, %s, got none", test.firstUnits, test.secondUnits, test.firstCurrency, test.secondCurrency)
		}
		if !test.hasError && err != nil {
			t.Errorf("Unexpected error for input %v, %v, %s, %s: %v", test.firstUnits, test.secondUnits, test.firstCurrency, test.secondCurrency, err)
		}
		if result != test.expected {
			t.Errorf("For input %v, %v, %s, %s, expected %s, got %s", test.firstUnits, test.secondUnits, test.firstCurrency, test.secondCurrency, test.expected, result)
		}
	}
}
