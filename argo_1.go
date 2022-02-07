package algo_practice

func twoSum(nums []int, target int) []int {
	kv := make(map[int]bool)
	for _, num := range nums {
		kv[num] = true
	}
	for firstIndex, num := range nums {
		newTarget := target - num
		if _, ok := kv[newTarget]; ok {
			for secondIndex, item := range nums[firstIndex+1:] {
				if item == newTarget {
					// NOTE 注意结果不是 secondIndex
					return []int{firstIndex, secondIndex + firstIndex + 1}
				}
			}
		}
	}
	return []int{-1, -1}
}
