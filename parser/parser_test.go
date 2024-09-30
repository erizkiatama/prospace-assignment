package parser

import (
	"github.com/erizkiatama/prospace-assignment/constant"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ParsedInput
	}{
		{
			name:  "Roman numeral assignment",
			input: "xyz is I",
			expected: ParsedInput{
				InputType:  Assignment,
				ItemType:   Roman,
				FirstToken: []string{"xyz", "is", "I"},
			},
		},
		{
			name:  "Credits assignment",
			input: "xyz xyz Silver is 34 credits",
			expected: ParsedInput{
				InputType:     Assignment,
				ItemType:      Credits,
				FirstToken:    []string{"xyz", "xyz"},
				FirstCurrency: "Silver",
				Credits:       34,
			},
		},
		{
			name:  "Credits assignment error parse comparison",
			input: " is xyz abc def credits",
			expected: ParsedInput{
				InputType: Invalid,
				Error:     constant.ErrInvalidFormat,
			},
		},
		{
			name:  "Credits assignment error not a number",
			input: "xyz xyz is NaN credits",
			expected: ParsedInput{
				InputType: Invalid,
				Error:     constant.ErrInvalidCredit,
			},
		},
		{
			name:  "Roman numeral calculation",
			input: "how much is jkl rst xyz xyz ?",
			expected: ParsedInput{
				InputType:  Calculation,
				ItemType:   Roman,
				FirstToken: []string{"jkl", "rst", "xyz", "xyz"},
			},
		},
		{
			name:  "Credits calculation",
			input: "how many credits is xyz abc Silver ?",
			expected: ParsedInput{
				InputType:     Calculation,
				ItemType:      Credits,
				FirstToken:    []string{"xyz", "abc"},
				FirstCurrency: "Silver",
			},
		},
		{
			name:  "Roman numeral comparison (smaller)",
			input: "is jkl smaller than xyz ?",
			expected: ParsedInput{
				InputType:   Comparison,
				ItemType:    Roman,
				FirstToken:  []string{"jkl"},
				SecondToken: []string{"xyz"},
			},
		},
		{
			name:  "Roman numeral comparison (larger)",
			input: "is xyz abc larger than jkl rst ?",
			expected: ParsedInput{
				InputType:   Comparison,
				ItemType:    Roman,
				FirstToken:  []string{"xyz", "abc"},
				SecondToken: []string{"jkl", "rst"},
			},
		},
		{
			name:  "Roman numeral comparison (error)",
			input: "is xyz abc invalid jkl rst ?",
			expected: ParsedInput{
				InputType: Invalid,
				Error:     constant.ErrInvalidParse,
			},
		},
		{
			name:  "Credits comparison (less)",
			input: "does xyz abc Silver has less credits than xyz abc Gold ?",
			expected: ParsedInput{
				InputType:      Comparison,
				ItemType:       Credits,
				FirstToken:     []string{"xyz", "abc"},
				SecondToken:    []string{"xyz", "abc"},
				FirstCurrency:  "Silver",
				SecondCurrency: "Gold",
			},
		},
		{
			name:  "Credits comparison (more)",
			input: "does xyz abc Gold has more credits than xyz abc Silver ?",
			expected: ParsedInput{
				InputType:      Comparison,
				ItemType:       Credits,
				FirstToken:     []string{"xyz", "abc"},
				SecondToken:    []string{"xyz", "abc"},
				FirstCurrency:  "Gold",
				SecondCurrency: "Silver",
			},
		},
		{
			name:  "Credits comparison (error)",
			input: "does xyz abc Gold invalid xyz abc Silver ?",
			expected: ParsedInput{
				InputType: Invalid,
				Error:     constant.ErrInvalidParse,
			},
		},
		{
			name:  "Invalid input",
			input: "how much wood could a woodchuck chuck if a woodchuck could chuck wood ?",
			expected: ParsedInput{
				InputType: Invalid,
				Error:     constant.ErrInvalidParse,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Parse(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Parse() = %v, want %v", result, tt.expected)
			}
		})
	}
}
