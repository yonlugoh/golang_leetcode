package _18__Coin_Change_2

import "testing"

func TestMaxProduct(t *testing.T) {

	amounts := []int{
		5,
		3,
		2040,
	}
	tests := [][]int{
		[]int{1, 2, 5},
		[]int{2},
		[]int{2, 7, 13, 29},
	}

	results := []int{
		4,
		0,
		278245,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := change2(amounts[i], tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
