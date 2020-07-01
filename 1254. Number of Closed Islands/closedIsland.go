package _254__Number_of_Closed_Islands

type Tuple struct {
	Row int
	Col int
}

var (
	visited map[Tuple]int
)

func closedIsland(grid [][]int) int {

	height := len(grid)
	if height <= 2 {
		return 0
	}
	width := len(grid[0])
	res := 0
	visited = make(map[Tuple]int)
	for row := 1; row < height-1; row++ {
		for col := 1; col < width-1; col++ {
			_, ok := visited[Tuple{row, col}]
			if grid[row][col] == 0 && !ok {

				is_closed := dfs(grid, height, width, row, col)
				if is_closed {
					res++
				}

			}

		}

	}

	return res
}

func dfs(grid [][]int, height int, width int, row int, col int) bool {
	if row < 0 || col < 0 || row >= height || col >= width {
		return false
	}
	if grid[row][col] == 1 {
		return true
	}
	_, ok := visited[Tuple{row, col}]

	if ok {
		return true
	}
	visited[Tuple{Row: row, Col: col}] = 1

	return dfs(grid, height, width, row-1, col) && dfs(grid, height, width, row+1, col) && dfs(grid, height, width, row, col+1) && dfs(grid, height, width, row, col-1)

}
