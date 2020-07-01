package _094__Car_Pooling

import "testing"

func TestCarPooling(t *testing.T) {
	tests := [][][]int{
		[][]int{{2, 1, 5}, {3, 3, 7}},
		[][]int{{2, 1, 5}, {3, 3, 7}},
		[][]int{{2, 1, 5}, {3, 5, 7}},
		[][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}},
	}

	targets := []int{
		4,
		5,
		3,
		11,
	}

	results := []bool{
		false,
		true,
		true,
		true,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := carPooling(tests[i], targets[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
