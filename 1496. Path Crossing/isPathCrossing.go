package _496__Path_Crossing

func isPathCrossing(path string) bool {
	type Tuple struct {
		X int
		Y int
	}

	visited := make(map[Tuple]int)
	//starting positions
	x := 0
	y := 0
	// add initial position to visited set
	visited[Tuple{X: x, Y: y}] = 0
	for _, ch := range path {
		if string(ch) == "N" {
			y++
		} else if string(ch) == "S" {
			y--
		} else if string(ch) == "W" {
			x--
		} else if string(ch) == "E" {
			x++
		}
		// found returns true if current position exists in set
		_, found := visited[Tuple{X: x, Y: y}]
		if found {
			return true
		}
		// add current position to visited set
		visited[Tuple{X: x, Y: y}] = 0

	}
	// no overlap after iterating through path
	return false

}
