package _6__Merge_Intervals

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	tests := [][][]int{
		[][]int{{1, 4}, {4, 5}},
		[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		[][]int{{8, 10}, {15, 16}, {17, 19}, {1, 3}, {2, 6}},
	}

	results := [][][]int{
		[][]int{{1, 5}},
		[][]int{{1, 6}, {8, 10}, {15, 18}},
		[][]int{{1, 6}, {8, 10}, {15, 16}, {17, 19}},
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := mergeIntervals(tests[i]); !reflect.DeepEqual(ret, results[i]) {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
