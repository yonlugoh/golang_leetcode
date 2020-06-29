package _28__Longest_Consecutive_Sequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := [][]int{
		[]int{100, 4, 200, 1, 3, 2},
		[]int{100, 4, 200, 1, 5, 1000},
		[]int{8, 7, 6, 4, 2, 44, 5, 3, 52, 5, 25, 234, 432, 33, 34, 2, 52, 2, 2, 3, 4, 5, 5, 1, 5, 3},
	}

	results := []int{
		4,
		2,
		8,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := longestConsecutive(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
