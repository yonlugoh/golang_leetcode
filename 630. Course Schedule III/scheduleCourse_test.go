package _30__Course_Schedule_III

import "testing"

func TestCourseSchedule3(t *testing.T) {
	tests := [][][]int{
		[][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}},
		[][]int{{200, 1300}, {600, 950}, {2000, 3200}, {100, 200}, {500, 600}, {200, 1950}},
		[][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}, {100, 300}, {1000, 1300}, {200, 350}},
	}

	results := []int{
		3,
		5,
		4,
	}

	caseNum := len(tests)
	for i := 0; i < caseNum; i++ {
		if ret := scheduleCourse3(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
