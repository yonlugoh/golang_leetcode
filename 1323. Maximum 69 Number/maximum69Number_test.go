package _323__Maximum_69_Number

import "testing"

func TestMaximum69Number(t *testing.T) {
	tests := []int{
		9669,
		9999,
		9996,
		699969,
	}

	results := []int{
		9969,
		9999,
		9999,
		999969,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := maximum69Number(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
