package main

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var start *ListNode

func isPalindrome(head *ListNode) bool {
	start = head
	return rec(head)
}

func rec(curr *ListNode) bool {
	if curr != nil {
		if rec(curr.Next) == false {
			return false
		}
		//fmt.Println(start.Val, curr.Val)

		if start.Val != curr.Val {
			return false
		}
		//fmt.Println(start.Val, curr.Val)
		start = start.Next
	}
	return true
}

func isPalindrom2e(head *ListNode) bool {
	var positions [10]int
	var odd [10]bool

	var i = 1
	for head.Next != nil {
		positions[head.Val] += i
		odd[head.Val] = !odd[head.Val]
		i++
		head = head.Next
	}
	positions[head.Val] += i
	odd[head.Val] = !odd[head.Val]

	if i == 1 {
		return true
	}
	allowNotOdd := false
	if i%2 != 0 {
		allowNotOdd = true
	}

	hasOdd := false
	hasBad := false
	halfVal := int(math.Ceil(float64(i) / 2))
	for _, posSum := range positions {
		if posSum == 0 {
			continue
		}
		res := posSum % (i + 1)
		if odd[head.Val] && !allowNotOdd {
			return false
		} else if odd[head.Val] && allowNotOdd || odd[head.Val] && hasBad {
			hasBad = true
		}
		if allowNotOdd && !hasOdd && res == halfVal {
			hasOdd = true
		} else if res != 0 {
			return false
		}
	}

	return true
}
