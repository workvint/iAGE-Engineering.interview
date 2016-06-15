package numberutil

import "testing"

var tests = map[int]string{
        0: "ноль",
        1000000: "один миллион",
}

func TestVerbal(t *testing.T) {
        for number, verbal := range tests {
                result := Verbal(number)
                if result != verbal {
                        t.Error(
                                "for", number,
                                "expected", verbal,
                                "got", result,
                        )
                }
        }
}
