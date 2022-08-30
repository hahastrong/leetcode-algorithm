package backtracking

func partition(s string) [][]string {
	var res [][]string
	backtrackingPartition(s, 0, []string{}, &res)
	return res
}

func backtrackingPartition(s string, idx int, path []string, res *[][]string) {
	if idx >= len(s) {
		temp := make([]string, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i:=idx; i<len(s); i++ {
		if isValid(s, idx, i) {
			path = append(path, s[idx:i+1])
			backtrackingPartition(s, i+1, path, res)
			path = path[:len(path)-1]
		}
	}
}

func isValid(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		// 下标要自增啊
		start++
		end--
	}
	return true
}