package main

import (
	"bytes"
	"github.com/erizkiatama/prospace-assignment/constant"
	"reflect"
	"testing"
)

// MockDatabase implements the Database interface for testing
type MockDatabase struct {
	isError bool
}

func (m *MockDatabase) AddUnitToRomanMapping(unit, roman string)                     {}
func (m *MockDatabase) AddCurrencyToCreditsMapping(currency string, credits float32) {}
func (m *MockDatabase) GetRomanFromUnit(unit string) (string, error) {
	if m.isError {
		return "", constant.ErrInvalidFormat
	}
	return "I", nil
}
func (m *MockDatabase) GetCreditsFromCurrency(currency string) (float32, error) {
	if m.isError {
		return 0, constant.ErrInvalidFormat
	}
	return 1.0, nil
}

// MockCalculator implements the Calculator interface for testing
type MockCalculator struct {
	isError bool
}

func (m *MockCalculator) ConvertUnitsToInt(units []string) (int, error) {
	if m.isError {
		return 0, constant.ErrInvalidFormat
	}
	return 1, nil
}
func (m *MockCalculator) CalculateCreditsCurrency(unit float32, currency string) (float32, error) {
	if m.isError {
		return 0, constant.ErrInvalidFormat
	}
	return 1.0, nil
}
func (m *MockCalculator) CompareTwoUnits(first, second []string) (string, error) {
	if m.isError {
		return "", constant.ErrInvalidFormat
	}
	return "smaller than", nil
}
func (m *MockCalculator) CompareTwoCurrency(first, second []string, firstCurrency, secondCurrency string) (string, error) {
	if m.isError {
		return "", constant.ErrInvalidFormat
	}
	return "has less credits than", nil
}

func TestRunIntergalacticConverter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
		hasError bool
	}{
		{
			name:     "Roman numeral assignment",
			input:    "glob is I\n",
			expected: []string{},
		},
		{
			name:     "Credits assignment",
			input:    "glob glob Silver is 34 credits\n",
			expected: []string{},
		},
		{
			name:     "Roman numeral calculation",
			input:    "how much is pish tegj glob glob ?\n",
			expected: []string{"pish tegj glob glob is 1"},
		},
		{
			name:     "Credits calculation",
			input:    "how many credits is glob prok Silver ?\n",
			expected: []string{"glob prok silver is 1.00 Credits"},
		},
		{
			name:     "Roman numeral comparison",
			input:    "is pish smaller than glob ?\n",
			expected: []string{"pish is smaller than glob"},
		},
		{
			name:     "Credits comparison",
			input:    "does glob prok Silver has less credits than glob prok Gold ?\n",
			expected: []string{"glob prok silver has less credits than glob prok gold"},
		},
		{
			name:     "Invalid input",
			input:    "how much wood could a woodchuck chuck if a woodchuck could chuck wood ?\n",
			expected: []string{constant.ErrInvalidParse.Error()},
		},
		{
			name:     "empty input",
			input:    "\n",
			expected: []string{},
		},
		{
			name:     "error on assignment",
			input:    "xyz xyz is 34 credits\n",
			expected: []string{constant.ErrInvalidFormat.Error()},
			hasError: true,
		},
		{
			name:     "error on roman numeral calculation",
			input:    "how much is pish tegj glob glob ?\n",
			expected: []string{constant.ErrInvalidFormat.Error()},
			hasError: true,
		},
		{
			name:     "error on credits calculation",
			input:    "how many credits is glob prok Silver ?\n",
			expected: []string{constant.ErrInvalidFormat.Error()},
			hasError: true,
		},
		{
			name:     "error on roman numeral comparison",
			input:    "is pish smaller than glob ?\n",
			expected: []string{constant.ErrInvalidFormat.Error()},
			hasError: true,
		},
		{
			name:     "error on redits comparison",
			input:    "does glob prok Silver has less credits than glob prok Gold ?\n",
			expected: []string{constant.ErrInvalidFormat.Error()},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare input
			input := bytes.NewBufferString(tt.input)

			// Run the function
			db := &MockDatabase{
				isError: tt.hasError,
			}
			calc := &MockCalculator{
				isError: tt.hasError,
			}
			got := runIntergalacticConverter(db, calc, input)

			// Check the output
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("runIntergalacticConverter() = %v, want %v", got, tt.expected)
			}
		})
	}
}
