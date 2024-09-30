package constant

import "errors"

var (
	ErrInvalidFormat = errors.New("requested number is in invalid format")
	ErrInvalidParse  = errors.New("i have no idea what are you talking about")
	ErrInvalidCredit = errors.New("credits is not a number")
)
