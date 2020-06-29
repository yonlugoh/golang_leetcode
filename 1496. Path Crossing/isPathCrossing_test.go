package _496__Path_Crossing

import (
	"testing"
)

func TestPathCrossing(t *testing.T) {
	tests := []string{
		"NES",
		"NESWW",
		"NNNNWWEEEWWESSSS",
	}
	results := []bool{
		false,
		true,
		true,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := isPathCrossing(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
