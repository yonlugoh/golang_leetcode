package _5__Jump_Game_II

import "testing"

func TestJump2(t *testing.T) {
	tests := [][]int{
		[]int{2, 3, 1, 1, 4},
		[]int{2, 3, 0, 3, 1, 1, 1, 4, 0, 0, 2, 0},
		[]int{2, 3, 1, 0, 3, 3, 1, 3, 0, 1, 4, 2, 0, 3, 0, 2, 0},
	}

	results := []int{
		2,
		5,
		6,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := jump2(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
