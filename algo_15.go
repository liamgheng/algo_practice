package algo_practice

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)

	for first := 0; first < len(nums); first ++ {
		// 避免两次结果相同，造成重复
		if first > 0 && nums[first-1] == nums[first] {
			continue
		}

		target := -1 * nums[first]
		// third 和 second 从两侧往中间逼近
		third := len(nums) - 1
		for second := first + 1; second < len(nums); second ++ {
			if second > first +1 && nums[second-1] == nums[second] {
				continue
			}
			// 后两数之和
			for second < third && nums[second] + nums[third] > target {
				third --
			}
			if second == third {
				break
			}
			if nums[second] + nums[third] == target {
				result = append(result, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return result
}
