package utils

import "testing"

func TestSliceInt(t *testing.T) {
	testCases := [][][]int{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{1, 2, 4}},
	}
	testResults := []bool{
		true,
		false,
	}
	for idx := range testCases {
		testCase := testCases[idx]
		actual := IsEqual(testCase[0], testCase[1], false)
		expect := testResults[idx]
		if actual != expect {
			t.Logf("Expect: %v, get: %v", expect, actual)
			t.Fail()
		}
	}
}
