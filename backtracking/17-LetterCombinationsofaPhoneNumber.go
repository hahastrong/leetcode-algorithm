package backtracking

var number = []string {
	"",
	"",
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

func letterCombinations(digits string) []string {
	var res []string
	backtrackingLetter(digits, 0, []byte{}, &res)
	return res
}

func backtrackingLetter(digits string, idx int, path []byte, res *[]string) {
	if len(path) == len(digits) {
		if len(path) > 0 {
			*res = append(*res, string(path))
		}
		return
	}

	numberLetter := number[digits[idx]-'0']
	for j:=0; j<len(numberLetter); j++ {
		path = append(path, numberLetter[j])
		backtrackingLetter(digits, idx+1, path, res)
		path = path[:len(path)-1]
	}

}
