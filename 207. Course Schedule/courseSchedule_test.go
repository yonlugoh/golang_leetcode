package _07__Course_Schedule

import "testing"

func TestCourseSchedule1(t *testing.T) {
	nums := []int{
		2,
		2,
		5,
	}
	tests := [][][]int{
		[][]int{{1, 0}},
		[][]int{{1, 0}, {0, 1}},
		[][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {3, 4}},
	}

	results := []bool{
		true,
		false,
		true,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := courseSchedule1(nums[i], tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
