package parser

import (
	"errors"
	"fmt"
)

var parseErr = errors.New("parse error")

func parsePlaceholder(in string) (*Node, error) {
	return nil, fmt.Errorf("%w: unimplemented", parseErr)
}
