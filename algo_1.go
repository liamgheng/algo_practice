package algo_practice

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for index, num := range nums {
		// 按照顺序建立 hash 表，如果之后有数据匹配到，则返回结果
		if targetIndex, ok := hashTable[target-num]; ok {
			return []int{targetIndex, index}
		}
		hashTable[num] = index
	}
	return nil
}
