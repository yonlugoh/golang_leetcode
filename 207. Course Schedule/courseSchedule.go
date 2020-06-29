package _07__Course_Schedule

func courseSchedule1(numCourses int, prerequisites [][]int) bool {
	// visited is an array to keep track of whether we have visited each course
	visited := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		visited[i] = 0
	}
	// store prerequisites in a hashmap
	// maps prerequisite course x (int) -> courses that have x as prerequisite (slice of ints)
	pre := map[int][]int{}
	for _, req := range prerequisites {
		x := req[0]
		y := req[1]
		val, found := pre[x]
		if found {
			pre[x] = append(val, y)
		} else {
			pre[x] = []int{y}
		}
	}
	// perform DFS on course 9 to numCourses-1
	for i := 0; i < numCourses; i++ {
		if !dfs(visited, pre, i) {
			return false
		}
	}
	return true
}

func dfs(visited []int, pre map[int][]int, x int) bool {
	// our prerequisite course has been taken
	if visited[x] == 1 {
		return true
		// not possible to complete a course as prerequisite not complete
	} else if visited[x] == -1 {
		return false
	}
	// mark current course as not possible to avoid cycle
	visited[x] = -1

	//check if any postrequisite, when running dfs, gives us false.
	// any postrequisite returns false, then this will return false

	for _, i := range pre[x] {
		if !dfs(visited, pre, i) {
			return false
		}
	}
	// mark visited
	visited[x] = 1
	return true

}
