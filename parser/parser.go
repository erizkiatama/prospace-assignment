package parser

import (
	"github.com/erizkiatama/prospace-assignment/constant"
	"strconv"
	"strings"
)

type InputType int

const (
	Assignment InputType = iota
	Calculation
	Comparison
	Invalid
)

type ItemType int

const (
	Roman ItemType = iota
	Credits
)

type ParsedInput struct {
	InputType      InputType
	ItemType       ItemType
	FirstToken     []string
	SecondToken    []string
	FirstCurrency  string
	SecondCurrency string
	Credits        int
	Error          error
}

func Parse(line string) ParsedInput {
	// Clean up the question mark & trailing whitespace
	line, _ = strings.CutSuffix(line, "?")
	line, _ = strings.CutSuffix(line, " ?")
	line = strings.TrimSpace(line)

	tokens := strings.Split(line, " ")
	switch {
	case len(tokens) == 3:
		return ParsedInput{
			InputType:  Assignment,
			ItemType:   Roman,
			FirstToken: tokens,
		}
	case len(tokens) > 4 && tokens[len(tokens)-1] == "credits":
		first, second, err := parseComparison(line, " is ")
		if err != nil {
			return ParsedInput{InputType: Invalid, Error: err}
		}

		credits, err := strconv.Atoi(second[len(second)-2])
		if err != nil {
			return ParsedInput{InputType: Invalid, Error: constant.ErrInvalidCredit}
		}
		return ParsedInput{
			InputType:     Assignment,
			ItemType:      Credits,
			FirstToken:    first[:len(first)-1],
			FirstCurrency: first[len(first)-1],
			Credits:       credits,
			Error:         nil,
		}
	case strings.HasPrefix(line, "how much is"):
		return ParsedInput{
			InputType:  Calculation,
			ItemType:   Roman,
			FirstToken: tokens[3:],
		}
	case strings.HasPrefix(line, "how many credits is"):
		return ParsedInput{
			InputType:     Calculation,
			ItemType:      Credits,
			FirstToken:    tokens[4 : len(tokens)-1],
			FirstCurrency: tokens[len(tokens)-1],
		}
	case strings.HasPrefix(line, "is"):
		separator := ""
		if strings.Contains(line, "smaller than") {
			separator = "smaller than"
		} else if strings.Contains(line, "larger than") {
			separator = "larger than"
		}

		line, _ = strings.CutPrefix(line, "is")
		first, second, err := parseComparison(line, separator)
		if err != nil {
			return ParsedInput{InputType: Invalid, Error: err}
		}

		return ParsedInput{
			InputType:   Comparison,
			ItemType:    Roman,
			FirstToken:  first,
			SecondToken: second,
		}
	case strings.HasPrefix(line, "does"):
		separator := ""
		if strings.Contains(line, "less") {
			separator = "has less credits than"
		} else if strings.Contains(line, "more") {
			separator = "has more credits than"
		}

		line, _ = strings.CutPrefix(line, "does")
		first, second, err := parseComparison(line, separator)
		if err != nil {
			return ParsedInput{InputType: Invalid, Error: err}
		}

		lastIdxFirst, lastIdxSecond := len(first)-1, len(second)-1
		return ParsedInput{
			InputType:      Comparison,
			ItemType:       Credits,
			FirstToken:     first[:lastIdxFirst],
			SecondToken:    second[:lastIdxSecond],
			FirstCurrency:  first[lastIdxFirst],
			SecondCurrency: second[lastIdxSecond],
		}
	default:
		return ParsedInput{
			InputType: Invalid,
			Error:     constant.ErrInvalidParse,
		}
	}
}

func parseComparison(line, sep string) ([]string, []string, error) {
	if sep == "" {
		return nil, nil, constant.ErrInvalidParse
	}

	before, after, _ := strings.Cut(line, sep)

	before = strings.TrimSpace(before)
	after = strings.TrimSpace(after)

	if len(before) == 0 || len(after) == 0 {
		return nil, nil, constant.ErrInvalidFormat
	}

	return strings.Split(before, " "), strings.Split(after, " "), nil
}
