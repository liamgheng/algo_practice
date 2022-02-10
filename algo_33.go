package algo_practice

// start 9:50 10:30
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (right - left) / 2 + left
		if nums[middle] == target {
			return middle
		}

		if nums[left] <= nums[middle] {
			// 左边有序
			if nums[left] <= target && target < nums[middle] {
				// 说明在左边区间
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else {
			// 右边有序
			if nums[middle] < target && target <= nums[right] {
				// 说明在右区间
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}
	return -1
}
