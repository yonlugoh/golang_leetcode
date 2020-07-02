package _60__Subarray_Sum_Equals_K

import (
	"testing"
)

func TestSubarraySum(t *testing.T) {
	tests := [][]int{
		[]int{1, 1, 1},
		[]int{1, 3, 4, 5, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 1, 1, 1, 1, 1},
		[]int{1, 0, -1, 1, 0, -1, 0, 1, -1, 0, 0, 2, 3, 0, -2},
	}

	targets := []int{
		2,
		5,
		0,
	}

	results := []int{
		2,
		6,
		32,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := subarraySum(tests[i], targets[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
