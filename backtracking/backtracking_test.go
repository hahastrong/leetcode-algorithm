package backtracking

import (
	"fmt"
	"testing"
)

func Test_combination(t *testing.T) {
	fmt.Println(combine(4, 2))
}

func Test_combinationSum3(t *testing.T) {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(4, 1))
}

func Test_letterCombinations(t *testing.T) {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}

func Test_combinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2,3,6,7}, 7))
	fmt.Println(combinationSum([]int{2,3,5}, 8))
	fmt.Println(combinationSum([]int{2}, 1))
}

func Test_combinationSum2(t *testing.T) {
	fmt.Println(combinationSum2([]int{10,1,2,7,6,1,5}, 8))
	fmt.Println(combinationSum2([]int{2,5,2,1,2}, 5))
	fmt.Println(combinationSum2([]int{}, 1))
}

func Test_partition(t *testing.T) {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))
	fmt.Println(partition("b"))
}

func Test_restoreIpAddresses(t *testing.T) {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("101023"))
}

func Test_subsets(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{0}))
}

func Test_subsetsWithDup(t *testing.T) {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{0}))
}

func Test_findSubsequences(t *testing.T) {
	fmt.Println(findSubsequences([]int{4,6,7,7}))
	fmt.Println(findSubsequences([]int{4,4,3,2,1}))
}

func Test_permute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))
	fmt.Println(permute([]int{1}))
}
func Test_permuteUnique(t *testing.T) {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 2, 3}))
	fmt.Println(permuteUnique([]int{-1,2,-1,2,1,-1,2,1}))
}

func Test_findItinerary(t *testing.T) {
	fmt.Println(findItinerary([][]string {{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}}))
	fmt.Println(findItinerary([][]string{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}}))
	fmt.Println(findItinerary([][]string{{"JFK","KUL"},{"JFK","NRT"},{"NRT","JFK"}}))

}

func Test_solveNQueens(t *testing.T) {
	fmt.Println(solveNQueens(4))
	fmt.Println(solveNQueens(1))
	fmt.Println(solveNQueens(2))
}

func Test_solveSudoku(t *testing.T) {
	board := [][]byte{{'5','3','.','.','7','.','.','.','.'},{'6','.','.','1','9','5','.','.','.'},{'.','9','8','.','.','.','.','6','.'},{'8','.','.','.','6','.','.','.','3'},{'4','.','.','8','.','3','.','.','1'},{'7','.','.','.','2','.','.','.','6'},{'.','6','.','.','.','.','2','8','.'},{'.','.','.','4','1','9','.','.','5'},{'.','.','.','.','8','.','.','7','9'}}
	solveSudoku(board)
	fmt.Println(board)

}