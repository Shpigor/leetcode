package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	expected := []int{0, 1}
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	t.Logf("expected: %+v   result: %+v", expected, result)
	if !reflect.DeepEqual(expected, result) {
		t.Fail()
	}
}

func twoSum(nums []int, target int) []int {
	lastIdx := len(nums) - 1
	for idx, number := range nums {
		nextIdx := idx + 1
		if idx != lastIdx {
			for _, nextNumber := range nums[nextIdx:] {
				if number+nextNumber == target {
					return []int{idx, nextIdx}
				}
				nextIdx += 1
			}
		}
	}
	return nil
}
