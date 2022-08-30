package characterString

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	// init pattern
	pat := make([][]int, 128)
	for i:=0; i<128; i++ {
		pat[i] = make([]int, len(needle))
	}

	pat[needle[0]][0] = 1
	X := 0
	for i:=1; i<len(needle); i++ {
		for j:=0; j<128; j++ {
			pat[j][i] = pat[j][X]
		}
		pat[needle[i]][i] = i+1
		X = pat[needle[i]][X]
	}

	j:=0
	for i:=0; i<len(haystack); i++ {
		j = pat[haystack[i]][j]
		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}


