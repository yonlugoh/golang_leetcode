package _52__Minimum_Number_of_Arrows_to_Burst_Balloons

import "testing"

func TestFindMinArrowShots(t *testing.T) {
	tests := [][][]int{
		[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
		[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}, {1, 1}, {100, 200}, {20, 22}, {30, 40}},
		[][]int{{1, 2}, {9, 10}, {8, 9}, {7, 8}, {6, 7}, {5, 6}, {4, 5}, {3, 4}, {2, 3}},
	}

	results := []int{
		2,
		6,
		5,
	}
	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := findMinArrowShots(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
