package algo_practice

func findKthLargest(nums []int, k int) int {
	middleNum := len(nums)/2
	middle := nums[middleNum]
	var left []int
	var right []int
	for _, num := range append(nums[:middleNum], nums[middleNum+1:]...) {
		if num < middle {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	if len(right) + 1 == k {
		return middle
	} else if len(right) + 1 > k {
		// 在右侧
		return findKthLargest(right, k)
	} else {
		// 在左侧
		return findKthLargest(left, k - len(right) - 1)
	}
}
