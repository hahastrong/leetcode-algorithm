package backtracking

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	backtrackingWithDup(nums, 0, []int{}, &res)
	return res
}

func backtrackingWithDup(nums []int, idx int, path []int, res *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)

	if idx == len(nums) {
		return
	}

	flag := make([]bool, len(nums))
	for i:=idx; i<len(nums); i++ {
		flag[i] = true
		if i>0 && nums[i-1] == nums[i] && flag[i-1] {
			continue
		}
		path = append(path, nums[i])
		backtrackingWithDup(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}
