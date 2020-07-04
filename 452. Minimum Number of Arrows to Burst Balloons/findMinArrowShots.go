package _52__Minimum_Number_of_Arrows_to_Burst_Balloons

import "sort"

func findMinArrowShots(points [][]int) int {
	length := len(points)
	if length == 0 {
		return 0
	}
	// sort by 2nd element, ascending
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	currPos := points[0][1]
	arrowCount := 1

	for i := 1; i < length; i++ {
		// if currPos is larger than first element, than it will definitely be able to burst this balloon
		if currPos >= points[i][0] {
			continue
		}
		// set the next position to shoot
		currPos = points[i][1]
		arrowCount++

	}

	return arrowCount

}
