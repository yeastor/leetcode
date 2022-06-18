package main

import "testing"

func TestPoli(t *testing.T) {

	ln := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				//Next: &ListNode{
				//	Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				//	},
			},
		},
	}

	res := isPalindrome(ln)

	if !res {
		t.Errorf("no res")
	}

	ln = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	res = isPalindrome(ln)

	if !res {
		t.Errorf("no res")
	}

	ln = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	res = isPalindrome(ln)

	if !res {
		t.Errorf("no res")
	}
}
