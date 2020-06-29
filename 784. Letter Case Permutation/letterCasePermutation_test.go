package _84__Letter_Case_Permutation

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []string{
		"a1b2",
		"3zflop",
		"12345",
	}

	results := [][]string{
		{"a1b2", "a1B2", "A1b2", "A1B2"},
		{"3zflop", "3zfloP", "3zflOp", "3zflOP", "3zfLop", "3zfLoP", "3zfLOp", "3zfLOP", "3zFlop", "3zFloP", "3zFlOp", "3zFlOP",
			"3zFLop", "3zFLoP", "3zFLOp", "3zFLOP", "3Zflop", "3ZfloP", "3ZflOp", "3ZflOP", "3ZfLop", "3ZfLoP", "3ZfLOp", "3ZfLOP",
			"3ZFlop", "3ZFloP", "3ZFlOp", "3ZFlOP", "3ZFLop", "3ZFLoP", "3ZFLOp", "3ZFLOP"},
		{"12345"},
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := letterCasePermutation(tests[i]); !sameStringSlice(ret, results[i]) {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}

func sameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}
