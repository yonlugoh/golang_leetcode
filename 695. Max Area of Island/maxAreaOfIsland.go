package _95__Max_Area_of_Island

func maxAreaOfIsland(grid [][]int) int {

	height := len(grid)
	if height == 0 {
		return 0
	}

	width := len(grid[0])

	areas := make([][]int, height)
	maxArea := 0
	for row := 0; row < height; row++ {
		areas[row] = make([]int, width)
		for col := 0; col < width; col++ {
			areas[row][col] = dfs(grid, height, width, row, col)
			if areas[row][col] > maxArea {
				maxArea = areas[row][col]
			}
		}
	}

	return maxArea

}

func dfs(grid [][]int, height int, width int, row int, col int) int {

	if row < 0 || col < 0 || row >= height || col >= width || grid[row][col] != 1 {
		return 0
	}
	// Mark visited to prevent cycle
	grid[row][col] = 0

	top := dfs(grid, height, width, row-1, col)
	bottom := dfs(grid, height, width, row+1, col)
	right := dfs(grid, height, width, row, col+1)
	left := dfs(grid, height, width, row, col-1)

	return 1 + top + bottom + right + left

}
