package _30__Surrounded_Regions

func solveBoard(board [][]byte) {
	height := len(board)
	if height <= 2 {
		return
	}
	width := len(board[0])
	// Run dfs on the boundary rows of index 0 and height-1
	for row := 0; row < height; row++ {
		if board[row][0] == 'O' {
			dfs(&board, height, width, row, 0)
		}
		if board[row][width-1] == 'O' {
			dfs(&board, height, width, row, width-1)
		}
	}
	// Run dfs on the boundary cols of index 0 and width-1
	for col := 0; col < width; col++ {
		if board[0][col] == 'O' {
			dfs(&board, height, width, 0, col)
		}
		if board[height-1][col] == 'O' {
			dfs(&board, height, width, height-1, col)
		}
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			// All 'O's are surrounded, changed to 'X'
			if board[row][col] == 'O' {
				board[row][col] = 'X'
			}
			// All previously 'O', now marked as '*', changed back to 'O'
			if board[row][col] == '*' {
				board[row][col] = 'O'
			}

		}

	}

}

func dfs(board *[][]byte, height int, width int, row int, col int) {

	if row < 0 || col < 0 || row >= height || col >= width || (*board)[row][col] != 'O' {
		return

	}
	if (*board)[row][col] == 'O' {
		(*board)[row][col] = '*'
	}
	dfs(board, height, width, row-1, col)
	dfs(board, height, width, row+1, col)
	dfs(board, height, width, row, col+1)
	dfs(board, height, width, row, col-1)

}
