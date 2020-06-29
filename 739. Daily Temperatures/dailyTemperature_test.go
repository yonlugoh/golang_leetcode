package _39__Daily_Temperatures

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := [][]int{
		[]int{73, 74, 75, 71, 69, 72, 76, 73},
		[]int{73, 72, 71, 70, 69, 68, 67},
		[]int{66, 66, 77, 77, 55, 50, 60, 61, 62, 60, 59},
	}
	results := [][]int{
		[]int{1, 1, 4, 2, 1, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0},
		[]int{2, 1, 0, 0, 2, 1, 1, 1, 0, 0, 0},
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := dailyTemperatures(tests[i]); !reflect.DeepEqual(ret, results[i]) {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
