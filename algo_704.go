package algo_practice

// MARK
func searchBinary(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	for low <= high {
		// 这段没必要，因为之后会进入外面的条件
		//if low == high && nums[low] != target {
		//	return -1
		//}

		// 注意这里加 low
		middle := (high - low) / 2 + low
		if target == nums[middle] {
			return middle
		} else if target < nums[middle] {
			// 在左侧
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -1
}
