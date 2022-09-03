package backtracking

func subsets(nums []int) [][]int {
	var res [][]int
	backtrackingSubsets(nums, 0, []int{}, &res)
	return res
}

func backtrackingSubsets(nums []int, idx int, path []int, res *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)

	if idx == len(nums) {
		return
	}

	for i:=idx; i<len(nums); i++ {
		path = append(path, nums[i])
		backtrackingSubsets(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}
