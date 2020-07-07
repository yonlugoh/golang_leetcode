package _56__132_Pattern

import "testing"

func TestFind132pattern(t *testing.T) {
	tests := [][]int{
		[]int{1, 2, 3, 4},
		[]int{3, 1, 4, 2},
		[]int{-1, 3, 2, 0},
	}

	results := []bool{
		false,
		true,
		true,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := find132pattern(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
