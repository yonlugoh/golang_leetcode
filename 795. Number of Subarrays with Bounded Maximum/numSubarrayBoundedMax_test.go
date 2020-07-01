package _95__Number_of_Subarrays_with_Bounded_Maximum

import "testing"

func TestNumSubarrayBoundedMax(t *testing.T) {

	tests := [][]int{
		[]int{2, 1, 4, 3},
		[]int{1, 2, 3, 4, 5, 6},
		[]int{6, 5, 3, 2, 1, 0},
		[]int{1},
	}
	left := []int{
		2,
		2,
		3,
		4,
	}
	right := []int{
		3,
		3,
		4,
		4,
	}
	results := []int{
		3,
		5,
		4,
		0,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := numSubarrayBoundedMax(tests[i], left[i], right[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
