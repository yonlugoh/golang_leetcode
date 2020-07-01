package _62__Maximum_Width_Ramp

import "testing"

func TestMaxWidthRamp(t *testing.T) {

	tests := [][]int{
		[]int{6, 0, 8, 2, 1, 5},
		[]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1},
		[]int{6, 0, 8, 2, 1, 5, 4, 2, 1, 0},
		[]int{6, 9, 8, 2, 1, 5, 4, 2, 1, 0},
	}

	results := []int{
		4,
		7,
		8,
		4,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := maxWidthRamp(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
