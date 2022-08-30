package dynamicProgramming

func countSubstrings(s string) int {
	res := 0
	for i:=0; i<len(s); i++ {
		j:=i
		for k:=i; k>=0 && j<len(s) && s[j] == s[k]; k-- {
			res++
			j++
		}
		j=i+1
		for k:=i; k>=0 && j<len(s) && s[j] == s[k]; k-- {
			res++
			j++
		}

	}
	return res
}