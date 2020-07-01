package _9__Word_Search

import "testing"

func TestWordSearch(t *testing.T) {
	tests := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}

	targets := []string{
		"ABCCED",
		"SEE",
		"ABCB",
		"ABCESCFSADEZ",
		"ABCESCFSADEE",
	}
	results := []bool{
		true,
		true,
		false,
		false,
		true,
	}

	caseNum := len(targets)
	for i := 0; i < caseNum; i++ {
		if ret := wordSearch(tests, targets[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
