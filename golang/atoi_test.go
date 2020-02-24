package leetcode

import (
	"math"
	"testing"
)

type TestCase struct {
	received string
	expected int
}

func TestAtoi(t *testing.T) {
	cases := []TestCase{
		{"42", 42},
		{"      -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"-0", 0},
		{"0", 0},
		{"--1", 0},
		{"01", 1},
		{"+1", 1},
		{"9223372036854775808", 2147483647},
		{"-2147483649", -2147483648},
		{"-9223372036854775809", -2147483648},
		{"18446744073709551617", 2147483647},
		{"-18446744073709551619", -2147483648},
	}
	for _, test := range cases {
		runTest(t, test)
	}
}

func runTest(t *testing.T, testCase TestCase) {
	if result := stringToInt(testCase.received); result != testCase.expected {
		t.Logf("Received: %d -- Expected: %d", result, testCase.expected)
		t.Fail()
	}
}

func stringToInt(number string) int {
	if len(number) > 0 {
		return parseInt(number)
	}
	return 0
}

func parseInt(number string) int {
	value := uint64(0)
	positionNumber := 0
	minus := 1
Loop:
	for _, char := range number {
		switch char {
		case ' ':
			if positionNumber > 0 {
				break Loop
			}
			continue
		case '-':
			if positionNumber > 0 {
				break Loop
			}
			minus = -1
			positionNumber++
		case '+':
			if positionNumber > 0 {
				break Loop
			}
			minus = 1
			positionNumber++
		case '0':
			value = value * 10
		case '1':
			value = value*10 + 1
		case '2':
			value = value*10 + 2
		case '3':
			value = value*10 + 3
		case '4':
			value = value*10 + 4
		case '5':
			value = value*10 + 5
		case '6':
			value = value*10 + 6
		case '7':
			value = value*10 + 7
		case '8':
			value = value*10 + 8
		case '9':
			value = value*10 + 9
		default:
			break Loop
		}
		positionNumber++
		if value >= math.MaxInt32 && minus == 1 {
			return math.MaxInt32
		} else if value >= 2147483648 && minus == -1 {
			return math.MinInt32
		}
	}
	return int(value) * minus
}
