package leetcode

import (
	"math"
	"strconv"
	"testing"
)

var bottom = math.Pow(-2, 31)
var top = math.Pow(2, 31) - 1

func TestReverseInt(t *testing.T) {
	result := reverse(1534236469)
	if result != 0 {
		t.Fail()
	}

	result = reverse(357)
	if result != 753 {
		t.Fail()
	}

	result = reverse(-137)
	if result != -731 {
		t.Fail()
	}
}

func reverse(x int) int {
	if (float64(x) >= bottom) && (float64(x) <= top) {
		multiplier := 1
		if x < 0 {
			multiplier = -1
			x = x * multiplier
		}
		return validateIntBounds(strconv.Atoi(ReverseString(strconv.Itoa(x)))) * multiplier
	}
	return 0
}

func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func validateIntBounds(res int, err error) int {
	if err != nil {
		return 0
	}
	if (float64(res) >= bottom) && (float64(res) <= top) {
		return res
	} else {
		return 0
	}
}
