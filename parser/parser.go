package parser

import (
	"errors"
	"fmt"
)

var ParseErr = errors.New("parse error")

func ParsePlaceholder(in string) (*Node, error) {
	return nil, fmt.Errorf("%w: unimplemented", ParseErr)
}
