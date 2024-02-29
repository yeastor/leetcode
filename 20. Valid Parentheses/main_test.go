package main

import (
	"testing"
)

type TestCase struct {
	Vals string
	Want bool
}

func getTestCases() []*TestCase {

	return []*TestCase{
		{
			Vals: "()",
			Want: true,
		},
		{
			Vals: "()[]{}",
			Want: true,
		},
		{
			Vals: "()[]{{{}}}}",
			Want: false,
		},
		{
			Vals: "()[[[[]]]]{{}}[]",
			Want: true,
		},
		{
			Vals: "()[[[[[]]]]{{}}[]",
			Want: false,
		},
	}
}

func TestSay(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		res := isValid(tc.Vals)
		if res != tc.Want {
			t.Errorf("error res = %v, w = %v\n", res, tc.Want)
		}

	}
}
