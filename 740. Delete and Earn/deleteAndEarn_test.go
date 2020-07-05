package _40__Delete_and_Earn

import "testing"

func TestDeleteAndEarn(t *testing.T) {
	tests := [][]int{
		[]int{2, 3, 2},
		[]int{1, 3, 4, 5, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 1, 1, 1, 1, 1},
		[]int{1, 2, 3, 1},
		[]int{2, 2, 3, 3, 3, 4},
	}

	results := []int{
		4,
		31,
		5,
		9,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := deleteAndEarn(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
