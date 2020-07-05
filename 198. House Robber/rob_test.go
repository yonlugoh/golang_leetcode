package _98__House_Robber

import "testing"

func TestRob(t *testing.T) {
	tests := [][]int{
		[]int{2, 3, 2},
		[]int{1, 3, 4, 5, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 1, 1, 1, 1, 1},
		[]int{1, 2, 3, 1},
		[]int{2, 7, 9, 3, 1},
	}

	results := []int{
		4,
		28,
		4,
		12,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := rob(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
