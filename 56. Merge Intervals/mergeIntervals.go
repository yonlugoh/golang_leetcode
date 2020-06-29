package _6__Merge_Intervals

import "sort"

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	// sort by the first element in interval
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{}
	// current keeps track of last interval to be appended to the result
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		// if interval is incorporated in current
		if intervals[i][0] <= current[1] {
			// if interval has a higher 2nd element
			if intervals[i][1] > current[1] {
				current[1] = intervals[i][1]
			}
		} else {
			// add current interval to result
			result = append(result, current)
			// set the interval as current
			current = intervals[i]
		}
	}
	// add remaining interval to result
	result = append(result, current)
	return result
}
