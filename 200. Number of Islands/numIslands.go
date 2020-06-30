package _00__Number_of_Islands

func numIslands(grid [][]byte) int {

	height := len(grid)
	if height == 0 {
		return 0
	}
	width := len(grid[0])

	res := 0

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {

			if grid[row][col] == '1' {
				res++
				dfs(grid, height, width, row, col)

			}

		}

	}

	return res

}

func dfs(grid [][]byte, height int, width int, row int, col int) {

	if row < 0 || col < 0 || row >= height || col >= width || grid[row][col] != '1' {
		return
	}
	grid[row][col] = '2'

	dfs(grid, height, width, row+1, col)
	dfs(grid, height, width, row-1, col)
	dfs(grid, height, width, row, col+1)
	dfs(grid, height, width, row, col-1)

}
