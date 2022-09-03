package backtracking

func solveNQueens(n int) [][]string {
	path := make([][]byte, n)
	for i:=0; i<n; i++ {
		path[i] = make([]byte, n)
		for j := 0; j<n; j++ {
			path[i][j] = '.'
		}
	}
	var res [][]string
	backtrackingSolveQueens(0, &path, &res)
	return res
}

func backtrackingSolveQueens(idx int, path *[][]byte, res *[][]string) {
	if idx == len(*path) {
		var tmp []string
		for i:=0; i<len(*path); i++ {
			tmp = append(tmp, string((*path)[i]))
		}
		*res = append(*res, tmp)
		return
	}

	for i:=0; i<len(*path); i++ {
		if isQueenValid(idx, i, path) {
			(*path)[idx][i] = 'Q'
			backtrackingSolveQueens(idx+1, path, res)
			(*path)[idx][i] = '.'
		}
	}
}

func isQueenValid(row, column int, path *[][]byte) bool {
	for i:=0; i<row; i++ {
		if (*path)[i][column] == 'Q' {
			return false
		}
	}

	for i, j :=row, column; i>=0 && j>=0; i-- {
		if (*path)[i][j] == 'Q' {
			return false
		}
		j--
	}

	for i, j := row, column; i>=0 && j<len(*path); j++ {
		if (*path)[i][j] == 'Q' {
			return false
		}
		i--
	}
	return true
}
