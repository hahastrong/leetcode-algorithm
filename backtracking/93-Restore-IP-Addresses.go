package backtracking

import "strings"

func restoreIpAddresses(s string) []string {
	var res []string
	backtrackingIp(s, 0, 4, []string{}, &res)
	return res
}

func backtrackingIp(s string, idx, k int, path []string, res *[]string) {
	if k == 0 && idx == len(s) {
		*res = append(*res, strings.Join(path, "."))
		return
	}

	if k == 0 {
		return
	}

	for i:=idx; i<len(s); i++ {
		if isValidIp(s, idx, i) {
			path = append(path, s[idx:i+1])
			backtrackingIp(s, i+1, k-1, path, res)
			path = path[:len(path)-1]
		}
	}
}

func isValidIp(s string, start, end int) bool {
	if s[start] == '0' {
		if start == end {
			return true
		}
		return false
	}

	sum := 0
	for start <= end {
		sum *= 10
		sum += int(s[start] - '0')
		start++
	}

	if sum > 255 {
		return false
	}
	return true
}
