package algo_practice

var result [][]int

// 8:47 典型回溯法
func permute(nums []int) [][]int {
	backtrace(0, []int{}, nums)

	return result
}

func backtrace(nextIndex int, tmpResult []int, nums []int)  {
	if nextIndex == len(nums) {
		// 没有位置需要填充，说明产生了新的组合
		tmp := make([]int, nextIndex)
		// MARK 注意这里需要 copy 下，否则后面会改掉
		copy(tmp, tmpResult)
		result = append(result, tmp)
		return
	}
	// 说明 nextIndex 之前（不包括 nextIndex）已经填充了，所以接下来挨个填充剩下的
	for index := nextIndex; index < len(nums); index++ {
		nums[index], nums[nextIndex] = nums[nextIndex], nums[index]
		tmpResult = append(tmpResult, nums[index])
		backtrace(nextIndex + 1, tmpResult, nums)
		nums[nextIndex], nums[index] = nums[index], nums[nextIndex]
	}
}