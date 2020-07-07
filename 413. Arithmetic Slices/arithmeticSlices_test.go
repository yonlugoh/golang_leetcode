package _13__Arithmetic_Slices

import "testing"

func TestArithmeticSlices(t *testing.T) {
	tests := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 1, 2, 5, 7},
		[]int{1, 3, 5, 7, 9, 13, 0, 1, 1},
	}

	results := []int{
		6,
		0,
		6,
	}
	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := arithmeticSlices(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
