package Sum

import (
	"fmt"
	"sort"
	"testing"
)


func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i:=0; i<len(nums)-2; i++ {
		for i > 0 && i < len(nums) && nums[i] == nums[i-1] {
			i++
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i] + nums[left] + nums[right] == 0 {
				tmp := []int{nums[i], nums[left], nums[right]}
				res = append(res, tmp)

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else if nums[i] + nums[left] + nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

func Test3Sum(t *testing.T) {
	nums := []int {1,-1,-1,0}
	res := threeSum(nums)
	fmt.Println(res)

}