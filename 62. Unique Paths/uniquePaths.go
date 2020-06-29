package _2__Unique_Paths

func uniquePaths(m int, n int) int {
	paths := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		paths[i] = make([]int, m+1)

		for j := 0; j <= m; j++ {
			if i == 0 || j == 0 {
				paths[i][j] = 0
			} else if i == 1 && j == 1 {
				paths[1][1] = 1
			} else {
				paths[i][j] = paths[i-1][j] + paths[i][j-1]
			}
		}
	}

	return paths[n][m]

}
