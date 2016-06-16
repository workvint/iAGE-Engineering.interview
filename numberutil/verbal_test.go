package numberutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = map[int]string{
	0:       "ноль",
	1000000: "один миллион",
	1:       "один",
	3:       "три",
	13:      "тринадцать",
	30:      "тридцать",
	45:      "сорок пять",
	100:     "сто",
	109:     "сто девять",
	183:     "сто восемьдесят три",
	1000:    "одна тысяча",
	1001:    "одна тысяча один",
	1069:    "одна тысяча шестьдесят девять",
	1234:    "одна тысяча двести тридцать четыре",
	2000:    "две тысячи",
	2001:    "две тысячи один",
	2069:    "две тысячи шестьдесят девять",
	2234:    "две тысячи двести тридцать четыре",
	5000:    "пять тысяч",
	5001:    "пять тысяч один",
	5069:    "пять тысяч шестьдесят девять",
	5234:    "пять тысяч двести тридцать четыре",
	21000:   "двадцать одна тысяча",
	21001:   "двадцать одна тысяча один",
	21069:   "двадцать одна тысяча шестьдесят девять",
	21234:   "двадцать одна тысяча двести тридцать четыре",
	52000:   "пятьдесят две тысячи",
	52001:   "пятьдесят две тысячи один",
	52069:   "пятьдесят две тысячи шестьдесят девять",
	52234:   "пятьдесят две тысячи двести тридцать четыре",
	99000:   "девяносто девять тысяч",
	99001:   "девяносто девять тысяч один",
	99069:   "девяносто девять тысяч шестьдесят девять",
	99234:   "девяносто девять тысяч двести тридцать четыре",
	100000:  "сто тысяч",
	100001:  "сто тысяч один",
	205000:  "двести пять тысяч",
	205001:  "двести пять тысяч один",
	205069:  "двести пять тысяч шестьдесят девять",
	205234:  "двести пять тысяч двести тридцать четыре",
	421000:  "четыреста двадцать одна тысяча",
	421001:  "четыреста двадцать одна тысяча один",
	421069:  "четыреста двадцать одна тысяча шестьдесят девять",
	421234:  "четыреста двадцать одна тысяча двести тридцать четыре",
	652000:  "шетьсот пятьдесят две тысячи",
	652001:  "шетьсот пятьдесят две тысячи один",
	652069:  "шетьсот пятьдесят две тысячи шестьдесят девять",
	652234:  "шетьсот пятьдесят две тысячи двести тридцать четыре",
	999000:  "девятьсот девяносто девять тысяч",
	999001:  "девятьсот девяносто девять тысяч один",
	999069:  "девятьсот девяносто девять тысяч шестьдесят девять",
	999234:  "девятьсот девяносто девять тысяч двести тридцать четыре",
	999999:  "девятьсот девяносто девять тысяч девятьсот девяносто девять",
}

func TestVerbal(t *testing.T) {
	assert := assert.New(t)

	_, err := Verbal(-1)
	if assert.Error(err) {
		assert.Equal(err.Error(), "Min number is 0")
	}

	_, err = Verbal(1000001)
	if assert.Error(err) {
		assert.Equal(err.Error(), "Max number is 1000000")
	}

	for number, verbal := range tests {
		if result, err := Verbal(number); result != verbal {
			if err != nil {
				t.Error(err)
			}

			t.Error(
				"for", number,
				"expected", verbal,
				"got", result,
			)
		}
	}
}
