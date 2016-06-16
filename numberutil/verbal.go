// Package numberutil contains utility functions for working with numbers.
package numberutil

import (
	"errors"
	"strconv"
	"strings"
)

const (
	minNumber = 0
	maxNumber = 1000000
	base      = 10
)

// Returns the verbal representation of the given number.
// Supported languages: Russian.
// Min number is 0
// Max number is 1000000
func Verbal(number int) (string, error) {
	switch {
	case number < minNumber:
		return "", errors.New("Min number is " + strconv.Itoa(minNumber))
	case number > maxNumber:
		return "", errors.New("Max number is " + strconv.Itoa(maxNumber))
	}

	result := []string{}
	if number != 0 {
		for i, j := 6, maxNumber; j >= base; i, j = i-1, j/base {
			if n := number / j; n > 0 {
				switch i {
				case 6:
					result = append(result, order7(n))
				case 5:
					result = append(result, order6(n, number))
				case 4:
					result = append(result, order5(n, &number))
				case 3:
					result = append(result, order4(n))
				case 2:
					result = append(result, order3(n))
				case 1:
					result = append(result, order2(n, &number))
				}
			}
			number %= j
		}
		result = append(result, order1(number))
	} else {
		result = append(result, "ноль")
	}

	s := strings.TrimSpace(strings.Join(result, " "))

	return s, nil
}

// Millions
func order7(n int) string {
	var s string

	s = order1(n) + " миллион"

	return s
}

// Hundreds of thousands
func order6(n int, num int) string {
	var s string

	s = order3(n)
	if (num/1000)%100 == 0 {
		s += " тысяч"
	}

	return s
}

// Tens of thousands
func order5(n int, num *int) string {
	var s string

	s = order2(n, num)
	if *num/1000 == 0 {
		s += " тысяч"
	}

	return s
}

// Thousands
func order4(n int) string {
	var s string

	switch n {
	case 1:
		s = "одна тысяча"
	case 2:
		s = "две тысячи"
	case 3:
		s = "три тысячи"
	case 4:
		s = "четыре тысячи"
	default:
		s = order1(n) + " тысяч"
	}

	return s
}

// Hundreds
func order3(n int) string {
	var s string

	switch n {
	case 1:
		s = "сто"
	case 2:
		s = "двести"
	case 3:
		s = "триста"
	case 4:
		s = "четыреста"
	case 5:
		s = "пятьсот"
	case 6:
		s = "шестьсот"
	case 7:
		s = "семьсот"
	case 8:
		s = "восемьсот"
	case 9:
		s = "девятьсот"
	}

	return s
}

// Tens
func order2(n int, num *int) string {
	var s string

	switch n {
	case 1:
		switch *num {
		case 10:
			s = "десять"
		case 11:
			s = "одиннадцать"
		case 12:
			s = "двенадцать"
		case 13:
			s = "тринадцать"
		case 14:
			s = "четырнадцать"
		case 15:
			s = "пятнадцать"
		case 16:
			s = "шестнадцать"
		case 17:
			s = "семнадцать"
		case 18:
			s = "восемнадцать"
		case 19:
			s = "девятнадцать"
		}

		if *num%base < 10 {
			*num = 0
		}
	case 2:
		s = "двадцать"
	case 3:
		s = "тридцать"
	case 4:
		s = "сорок"
	case 5:
		s = "пятьдесят"
	case 6:
		s = "шестьдесят"
	case 7:
		s = "семьдесят"
	case 8:
		s = "восемьдесят"
	case 9:
		s = "девяносто"
	}

	return s
}

// Units
func order1(n int) string {
	var s string

	switch n {
	case 1:
		s = "один"
	case 2:
		s = "два"
	case 3:
		s = "три"
	case 4:
		s = "четыре"
	case 5:
		s = "пять"
	case 6:
		s = "шесть"
	case 7:
		s = "семь"
	case 8:
		s = "восемь"
	case 9:
		s = "девять"
	}

	return s
}
