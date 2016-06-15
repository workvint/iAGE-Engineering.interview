// Package numberutil contains utility functions for working with numbers.
package numberutil

import (
	"errors"
)

const (
	minNumber = 0
	maxNumber = 1000000
)

// Returns the verbal representation of the given number.
// Supported languages: Russian.
// Min number is 0
// Max number is 1000000
func Verbal(number int) (string, error) {
	switch {
	case number < minNumber:
		return "", errors.New("Min number is 0")
	case number > maxNumber:
		return "", errors.New("Max number is 1000000")
	}

	return "", nil
}
