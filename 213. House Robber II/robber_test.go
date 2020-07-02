package _13__House_Robber_II

import "testing"

func TestRobber(t *testing.T) {
	tests := [][]int{
		[]int{2, 3, 2},
		[]int{1, 3, 4, 5, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 1, 1, 1, 1, 1},
		[]int{1, 2, 3, 1},
	}

	results := []int{
		3,
		28,
		4,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := robber(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
