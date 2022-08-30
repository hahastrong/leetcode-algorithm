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