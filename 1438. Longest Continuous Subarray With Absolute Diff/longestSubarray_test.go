package _438__Longest_Continuous_Subarray_With_Absolute_Diff

import "testing"

func TestLongestSubarray(t *testing.T) {
	tests := [][]int{
		[]int{8, 2, 4, 7},
		[]int{10, 1, 2, 4, 7, 2},
		[]int{4, 2, 2, 2, 4, 4, 2, 2},
		[]int{10, 1, 2, 4, 7, 2},
	}

	limits := []int{
		4,
		5,
		0,
		5,
	}

	results := []int{
		2,
		4,
		3,
		4,
	}
	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := longestSubarray(tests[i], limits[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
