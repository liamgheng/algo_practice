package algo_practice

func maxSubArray(nums []int) int {
	prev := nums[0]
	result := prev
	for _, num := range nums[1:] {
		prev = max(num, num + prev)
		result = max(result, prev)
	}
	return result
}

func max(a, b int) int  {
	if a > b {
		return a
	}
	return b
}
