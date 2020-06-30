package _21__Best_Time_to_Buy_and_Sell_Stock

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := [][]int{
		[]int{7, 1, 5, 3, 6, 4},
		[]int{},
		[]int{8, 7, 6, 4, 2, 44, 5, 3, 52, 5, 25, 234, 432, 33, 34, 2, 52, 2, 2, 3, 4, 5, 5, 1, 5, 3},
	}

	results := []int{
		5,
		0,
		430,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := maxProfit(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
