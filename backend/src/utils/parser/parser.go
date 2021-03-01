package parser

import (
	"errors"
	"strconv"
)

func ParseID(IDParam string) (uint64, error) {
	ID, err := strconv.ParseUint(IDParam, 10, 64)
	if err != nil {
		return 0, errors.New("invalid id param")
	}

	return ID, nil
}
