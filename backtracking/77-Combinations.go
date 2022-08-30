package backtracking

func combine(n int, k int) [][]int {
	var res [][]int
	backtrackingCombine(n, k, 1, []int{}, &res)
	return res
}

func backtrackingCombine(n, k, idx int, path []int, res *[][]int) {
	if 0 == k {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i:=idx; i<=n-k+1; i++ {
		path = append(path, i)
		backtrackingCombine(n, k-1, i+1, path[:], res)
		path = path[:len(path)-1]
	}
}
