package main

import (
	"testing"
)

type TestCase struct {
	Vals int
	Want string
}

func getTestCases() []*TestCase {

	return []*TestCase{
		{
			Vals: 1,
			Want: "1",
		},
		{
			Vals: 2,
			Want: "11",
		},
		{
			Vals: 3,
			Want: "21",
		},
		{
			Vals: 4,
			Want: "1211",
		},
	}
}

func TestSay(t *testing.T) {
	testCases := getTestCases()

	for _, tc := range testCases {
		res := countAndSay(tc.Vals)
		if res != tc.Want {
			t.Errorf("error res = %v, w = %v\n", res, tc.Want)
		}

	}
}
