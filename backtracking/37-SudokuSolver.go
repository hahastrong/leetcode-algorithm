package backtracking

func solveSudoku(board [][]byte)  {
	backtrackingSolveSudoku(board)
}

func backtrackingSolveSudoku(board [][]byte) bool {
	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			if (board)[i][j] != '.' {
				continue
			}

			for k:='1'; k<= '9'; k++ {
				if isSudokuValid(board, i, j, byte(k)) {
					(board)[i][j] = byte(k)
					if backtrackingSolveSudoku(board) {
						return true
					}
					(board)[i][j] = '.'
				}
			}
			return false
		}
	}
	return true
}

func isSudokuValid(board [][]byte, row, column int, k byte) bool {
	for i:=0; i<9; i++ {
		if board[row][i] == k || board[i][column] == k {
			return false
		}
	}

	row = (row/3)*3
	column = (column/3)*3
	for i:=row; i<row+3; i++ {
		for j:=column; j<column+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
