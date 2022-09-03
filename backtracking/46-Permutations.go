package backtracking

func permute(nums []int) [][]int {
	var res [][]int
	flag := make([]bool, 21)
	backtrackingPermute(nums, 0, flag, []int{}, &res)
	return res
}

func backtrackingPermute(nums []int, idx int, flag []bool, path []int, res *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i:= 0; i < len(nums); i++ {
		if flag[nums[i]+10] {
			continue
		}
		flag[nums[i]+10] = true
		path = append(path, nums[i])
		backtrackingPermute(nums, 0, flag, path, res)
		path = path[:len(path)-1]
		flag[nums[i]+10] = false
	}
}
