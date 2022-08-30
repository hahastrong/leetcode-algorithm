package dynamicProgramming

import (
	"fmt"
	"testing"
)

func TestBase(t *testing.T) {
	fib(4)
}

func TestUni(t *testing.T) {
	obstacleGird := make([][]int, 0)
	obstacleGird = append(obstacleGird, []int{0, 0, 0})
	obstacleGird = append(obstacleGird, []int{0, 1, 0})
	obstacleGird = append(obstacleGird, []int{0, 0, 0})

	res := uniquePathsWithObstacles(obstacleGird)
	fmt.Println(res)
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, m)
	if obstacleGrid[0][0] == 1 || obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 {
		return 0
	}
	for i:=0; i<m; i++ {
		if obstacleGrid[0][i] != 1 {
			dp[i] = 1
		}
	}

	for i:=1;i<n; i++ {
		for j:=0;j<m; j++ {

			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[m-1]
}

func TestIntegerBreak(t *testing.T) {

	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(10))

}

func TestNumTrees(t *testing.T) {
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(3))
}

func TestCanPartition(t *testing.T) {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}

func TestLastStonesWeightII(t *testing.T) {
	fmt.Println(lastStoneWeightII([]int{2,7,4,1,8,1}))
	fmt.Println(lastStoneWeightII([]int{31,26,33,21,40}))
}

func TestFindTargetSumWays(t *testing.T) {
	fmt.Println(findTargetSumWays([]int{1,1,1,1,1}, 3))
	fmt.Println(findTargetSumWays([]int{1}, 1))
}


func TestFindMaxForm(t *testing.T) {
	fmt.Println(findMaxForm([]string{"10","0001","111001","1","0"}, 5, 3))
	fmt.Println(findMaxForm([]string{"10","0","1"}, 1, 1))
}

func TestChange(t *testing.T) {
	fmt.Println(change(5, []int{1, 2, 5}))
	fmt.Println(change(3, []int{2}))
}

func TestCombinationSum4(t *testing.T) {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
	fmt.Println(combinationSum4([]int{9}, 3))
}

func TestCoinChange(t *testing.T) {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{2}, 0))
}

func TestNumSquares(t *testing.T) {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
}

func TestWordBreak(t *testing.T) {
	fmt.Println(wordBreak("leetcode", []string{"leet","code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple","pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats","dog","sand","and","cat"}))
}

func TestRob(t *testing.T) {
	fmt.Println(rob([]int{1, 2, 3, 4}))
	fmt.Println(rob([]int{2,7,9,3,1}))
}

func TestRob2(t *testing.T) {
	fmt.Println(rob2([]int{2,3,2}))
	fmt.Println(rob2([]int{1,2,3,1}))
	fmt.Println(rob2([]int{1,2,3}))
	fmt.Println(rob2([]int{1}))
}

func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}

func TestMaxProfitII(t *testing.T) {
	fmt.Println(maxProfitII([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfitII([]int{1,2,3,4,5}))
	fmt.Println(maxProfitII([]int{7,6,4,3,1}))
}

func TestMaxProfitIII(t *testing.T) {
	fmt.Println(maxProfitIII([]int{3,3,5,0,0,3,1,4}))
	fmt.Println(maxProfitIII([]int{1,2,3,4,5}))
	fmt.Println(maxProfitIII([]int{7,6,4,3,1}))
}

func TestMaxProfitIV(t *testing.T) {
	fmt.Println(maxProfitIV(2, []int{2,4,1}))
	fmt.Println(maxProfitIV(2, []int{3,2,6,5,0,3}))
	//
	fmt.Println(maxProfitIV(2, []int{}))
}

func TestMaxProfitV(t *testing.T) {
	fmt.Println(maxProfitV([]int{1,2,3,0,2}))
	fmt.Println(maxProfitV([]int{1}))
}

func TestMaxProfitVI(t *testing.T) {
	fmt.Println(maxProfitVI([]int{1,3,2,8,4,9}, 2))
	fmt.Println(maxProfitVI([]int{1}, 2))
}

func TestLengthOfLIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{1,3,6,7,9,4,10,5,6}))
	fmt.Println(lengthOfLIS([]int{10,9,2,5,3,7,101,18}))
	fmt.Println(lengthOfLIS([]int{0,1,0,3,2,3}))
}

func TestFindLengthOfLCIS(t *testing.T) {
	fmt.Println(findLengthOfLCIS([]int{1,3,5,4,7}))
	fmt.Println(findLengthOfLCIS([]int{2,2,2,2,2}))
}

func TestFindLength(t *testing.T) {
	fmt.Println(findLength([]int{1,2,3,2,1}, []int{3,2,1,4,7}))
	fmt.Println(findLength([]int{0,0,0,0,0}, []int{0,0,0,0,0}))
	fmt.Println(findLength([]int{0,1,1,1,1}, []int{1,0,1,0,1}))
}

func TestLongestCommonSubsequence(t *testing.T) {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("abc", "abc"))
	fmt.Println(longestCommonSubsequence("abc", "def"))
	fmt.Println(longestCommonSubsequence("", "def"))
}

func TestMaxUncrossedLines(t *testing.T) {
	fmt.Println(maxUncrossedLines([]int{1,4,2}, []int{1,2,4}))
	fmt.Println(maxUncrossedLines([]int{2,5,1,2,5}, []int{10,5,2,1,5,2}))
	fmt.Println(maxUncrossedLines([]int{1,3,7,1,7,5}, []int{1,9,2,5,1}))
}

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5,4,-1,7,8}))
	fmt.Println(maxSubArray([]int{-5,-4,-3,-2,-1}))
}

func TestIsSubsequence(t *testing.T) {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
	fmt.Println(isSubsequence("", "ahbgdc"))
}

func TestNumDistinct(t *testing.T) {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println(numDistinct("babgbag", "bag"))
	fmt.Println(numDistinct("babgbag", ""))
}

func TestMinDistance(t *testing.T) {
	fmt.Println(minDistance("sea", "eat"))
	fmt.Println(minDistance("leetcode", "etco"))
	fmt.Println(minDistance("babgbag", ""))
}

func TestMinDistance2(t *testing.T) {
	fmt.Println(minDistance2("horse", "ros"))
	fmt.Println(minDistance2("intention", "execution"))
	fmt.Println(minDistance2("babgbag", ""))
}

func TestCountSubstrings(t *testing.T) {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
	fmt.Println(countSubstrings(""))
}

func TestLongestPalindromeSubseq(t *testing.T) {
	fmt.Println(longestPalindromeSubseq("bbbab"))
	fmt.Println(longestPalindromeSubseq("cbbd"))
}

