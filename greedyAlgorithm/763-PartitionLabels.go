package greedyAlgorithm

func partitionLabels(s string) []int {
	dict := make(map[byte]int, 0)
	for i:=0; i<len(s); i++ {
		dict[s[i]] = i
	}

	var res []int

	start := 0
	right := 0
	for i:= 0; i<len(s); i++ {
		right = Max(right, dict[s[i]])
		if right == i {
			res = append(res, i-start+1)
			start = i+1
		}
	}

	return res
}
