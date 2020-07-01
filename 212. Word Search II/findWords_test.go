package _12__Word_Search_II

import "testing"

func TestFindWords(t *testing.T) {
	tests := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}

	words := [][]string{
		[]string{
			"oath",
			"pea",
			"eat",
			"rain",
			"make",
			"neat"},
	}
	results := [][]string{
		[]string{
			"eat",
			"oath",
			"neat"},
	}

	caseNum := len(words)
	for i := 0; i < caseNum; i++ {
		if ret := findWords(tests, words[i]); sameStringSlice(ret, results[i]) {
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
