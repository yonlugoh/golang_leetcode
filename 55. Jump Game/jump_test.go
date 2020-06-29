package Jump_Game

import "testing"

func TestJump(t *testing.T) {
	tests := [][]int{
		[]int{2, 3, 1, 1, 4},
		[]int{3, 2, 1, 0, 4},
		[]int{2, 3, 1, 0, 3, 3, 1, 3, 0, 1, 4, 2, 0, 3, 0, 2, 0},
	}

	results := []bool{
		true,
		false,
		true,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := canJump(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
