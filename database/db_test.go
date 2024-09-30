package database

import (
	"testing"
)

func TestNewDatabase(t *testing.T) {
	db := NewDatabase()
	if db == nil {
		t.Error("NewDatabase() returned nil")
	}
}

func TestAddUnitToRomanMapping(t *testing.T) {
	db := NewDatabase()
	db.AddUnitToRomanMapping("xyz", "I")

	roman, err := db.GetRomanFromUnit("xyz")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if roman != "I" {
		t.Errorf("Expected 'I', got '%s'", roman)
	}
}

func TestGetRomanFromUnit(t *testing.T) {
	db := NewDatabase()
	db.AddUnitToRomanMapping("abc", "V")

	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"abc", "V", false},
		{"ABC", "V", false},
		{"unknown", "", true},
	}

	for _, test := range tests {
		result, err := db.GetRomanFromUnit(test.input)
		if test.hasError && err == nil {
			t.Errorf("Expected error for input '%s', got none", test.input)
		}
		if !test.hasError && err != nil {
			t.Errorf("Unexpected error for input '%s': %v", test.input, err)
		}
		if result != test.expected {
			t.Errorf("For input '%s', expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestAddCurrencyToCreditsMapping(t *testing.T) {
	db := NewDatabase()
	db.AddCurrencyToCreditsMapping("Gold", 14450.0)

	credits, err := db.GetCreditsFromCurrency("Gold")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if credits != 14450.0 {
		t.Errorf("Expected 14450.0, got %f", credits)
	}
}

func TestGetCreditsFromCurrency(t *testing.T) {
	db := NewDatabase()
	db.AddCurrencyToCreditsMapping("Silver", 17.0)

	tests := []struct {
		input    string
		expected float32
		hasError bool
	}{
		{"Silver", 17.0, false},
		{"SILVER", 17.0, false},
		{"unknown", 0, true},
	}

	for _, test := range tests {
		result, err := db.GetCreditsFromCurrency(test.input)
		if test.hasError && err == nil {
			t.Errorf("Expected error for input '%s', got none", test.input)
		}
		if !test.hasError && err != nil {
			t.Errorf("Unexpected error for input '%s': %v", test.input, err)
		}
		if result != test.expected {
			t.Errorf("For input '%s', expected %f, got %f", test.input, test.expected, result)
		}
	}
}
