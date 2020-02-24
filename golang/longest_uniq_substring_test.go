package leetcode

import (
	"testing"
)

type SubStringTestCase struct {
	received string
	expected int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []SubStringTestCase{
		{"42", 2},
		{"      -42", 4},
		{"4193 with words", 9},
		{"words and 987", 8},
		{"dvdf", 3},
		{"abcabcbb", 3},
		{"abcdabd", 4},
	}
	for _, test := range cases {
		runTestSubstring(t, test)
	}
}

func runTestSubstring(t *testing.T, testCase SubStringTestCase) {
	if result := lengthOfLongestSubstring(testCase.received); result != testCase.expected {
		t.Logf("Received: %d -- Expected: %d", result, testCase.expected)
		t.Fail()
	}
}

func lengthOfLongestSubstring(line string) int {
	result := 0
	currentResult := 0
	uniqMap := make(map[uint8]int)
	for i := 0; i < len(line); i++ {
		char := line[i]
		if val, ok := uniqMap[char]; ok {
			if currentResult > result {
				result = currentResult
			}
			uniqMap = make(map[uint8]int)
			i = val
			currentResult = 0
		} else {
			uniqMap[char] = i
			currentResult++
		}
	}
	if currentResult > result {
		return currentResult
	}
	return result
}
