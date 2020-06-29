package _52__Maximum_Product_Subarray

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := [][]int{
		[]int{2, 3, -2, 4},
		[]int{-2, 0, -1},
		[]int{2, 3, -2, 4, 5, -6, 4, 2, 0, 4, 6, 7, 2, -3, 0, 4, 3, -6, 4, -5, 2, 0},
	}

	results := []int{
		6,
		0,
		11520,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := maxProduct(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
