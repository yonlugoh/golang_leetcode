package _497__Check_If_Array_Pairs_Are_Divisible_by_k

import "testing"

func TestCanArrange(t *testing.T) {

	tests := [][]int{
		[]int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9},
		[]int{1, 2, 3, 4, 5, 6},
		[]int{1, 2, 3, 4, 5, 6},
		[]int{-10, 10},
		[]int{-1, 1, -2, 2, -3, 3, -4, 4},
		[]int{1, 2, 3, 4, 12, 13, 22, 23, 11, 21},
	}
	targets := []int{
		5,
		7,
		10,
		2,
		3,
		5,
	}

	results := []bool{
		true,
		true,
		false,
		true,
		true,
		false,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := canArrange(tests[i], targets[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
