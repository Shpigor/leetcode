package leetcode

import (
	"log"
	"testing"
)

type MedianCase struct {
	name           string
	first          []int
	second         []int
	expectedResult float32
}

func TestTwoListMedian(t *testing.T) {
	cases := []MedianCase{
		{
			first:          []int{0, 1, 2, 3},
			second:         []int{5, 6, 7, 8},
			expectedResult: 4,
		},
		{
			first:          []int{2, 3},
			second:         []int{1},
			expectedResult: 2,
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			median := getMedian(testCase.first, testCase.second)
			if median != testCase.expectedResult {
				log.Printf("Received: %f, Expected: %f", median, testCase.expectedResult)
				t.Fail()
			}
		})
	}
}

func getMedian(first, second []int) float32 {
	isOdd := false
	if (len(first)+len(second))%2 != 0 {
		isOdd = true
	}
	if first[len(first)-1] < second[0] {
		return getMedianFromTwoSortedArray(first, second, isOdd)
	} else if first[0] > second[len(second)-1] {
		return getMedianFromTwoSortedArray(second, first, isOdd)
	}
	if len(first) < len(second) {
		tmp := first
		first = second
		second = tmp
	}

	return 0
}

func getMedianFromTwoSortedArray(first, second []int, isOdd bool) float32 {
	fIdx := ((len(first) + len(second)) / 2) - 1 + ((len(first) + len(second)) % 2)
	median := getByIdxFromTwoArray(fIdx, first, second)
	if !isOdd {
		return (median + getByIdxFromTwoArray(fIdx+1, first, second)) / 2
	}
	return median
}

func getByIdxFromTwoArray(idx int, first, second []int) float32 {
	median := 0
	if idx < len(first) {
		median = first[idx]
	} else {
		median = second[idx-len(first)]
	}
	return float32(median)
}
