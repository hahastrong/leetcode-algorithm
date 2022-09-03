package backtracking

func findSubsequences(nums []int) [][]int {
	var res [][]int
	backtrackingSubsequences(nums, 0, []int{}, &res)
	return res
}

func backtrackingSubsequences(nums []int, idx int, path []int, res *[][]int) {
	if len(path) > 1 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}

	if idx == len(nums) {
		return
	}

	flag := make([]bool, 201)
	for i:=idx; i<len(nums); i++ {

		if i>0 && flag[nums[i]+100] {
			continue
		}
		flag[nums[i]+100] = true
		if len(path) > 0 && nums[i] < path[len(path)-1] {
			continue
		}

		path = append(path, nums[i])
		backtrackingSubsequences(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}
