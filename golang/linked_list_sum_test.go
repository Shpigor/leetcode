package leetcode

import (
	"strconv"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestNode(t *testing.T) {
	expected := "[7,0,8]"
	left := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	right := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	result := addTwoNumbers(left, right).ToString()
	t.Logf("expected: %s   result: %s", expected, result)
	if result != expected {
		t.Fail()
	}
}

func (l *ListNode) ToString() string {
	message := "["
	level := l
	for hasValue(level) {
		message = message + strconv.Itoa(level.Val)
		if hasNextValue(level) {
			message = message + ","
		}
		level = level.Next
	}
	return message + "]"
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	nextNode := head
	nextLeft := l1
	nextRight := l2
	for hasValue(nextLeft) || hasValue(nextRight) {
		nextNode = sum(nextLeft, nextRight, nextNode)
		if nextNode.Next == nil && (hasNextValue(nextLeft) || hasNextValue(nextRight)) {
			nextNode.Next = &ListNode{}
		}
		if hasValue(nextLeft) {
			nextLeft = nextLeft.Next
		}
		if hasValue(nextRight) {
			nextRight = nextRight.Next
		}
		nextNode = nextNode.Next
	}
	return head
}

func sum(left, right, nextNode *ListNode) *ListNode {
	var newVal int
	if nextNode != nil {
		newVal = sumInt(left, right) + nextNode.Val
	} else {
		newVal = sumInt(left, right)
		nextNode = &ListNode{}
	}
	if newVal >= 10 {
		nextNode.Val = newVal - 10
		nextNode.Next = &ListNode{
			Val: 1,
		}
	} else {
		nextNode.Val = newVal
	}
	return nextNode
}

func sumInt(left, right *ListNode) int {
	result := 0
	if left == nil {
		result = right.Val
	} else if right == nil {
		result = left.Val
	} else {
		result = left.Val + right.Val
	}
	return result
}

func hasValue(node *ListNode) bool {
	result := false
	if node != nil {
		result = true
	}
	return result
}

func hasNextValue(node *ListNode) bool {
	result := false
	if node != nil && node.Next != nil {
		result = true
	}
	return result
}
