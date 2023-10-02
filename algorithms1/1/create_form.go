package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrOddParameterNumber = errors.New("odd number of parameters passed")

type ErrNotANumber struct {
	Parameter interface{}
}

func (err ErrNotANumber) Error() string {
	return fmt.Sprintf("Parameter %q is not a number", err.Parameter)
}

func CreateForm(params ...string) (map[string]float64, error) {
	if len(params)%2 != 0 {
		return nil, ErrOddParameterNumber
	}

	zipped := make(map[string]float64, len(params)/2)
	for i := 0; i < len(params); i += 2 {
		num, err := strconv.ParseFloat(params[i+1], 64)
		if err != nil {
			return nil, ErrNotANumber{Parameter: params[i+1]}
		}

		zipped[params[i]] = num
	}

	return zipped, nil
}
