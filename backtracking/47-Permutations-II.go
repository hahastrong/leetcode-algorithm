package backtracking

func permuteUnique(nums []int) [][]int {
	var res [][]int
	flag := make([]bool, len(nums))
	backtrackingPermuteUnique(nums, flag, []int{}, &res)
	return res
}

func backtrackingPermuteUnique(nums []int, flag []bool, path []int, res *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	used := make([]bool, 21)
	for i:=0; i<len(nums); i++ {
		if flag[i] {
			continue
		}
		if used[nums[i]+10] {
			continue
		}

		used[nums[i]+10] = true
		flag[i] = true
		path = append(path, nums[i])
		backtrackingPermuteUnique(nums, flag, path, res)
		path = path[:len(path)-1]
		flag[i] = false
	}
}
