package algo_practice

// 10:14 10:41
func maxProfit(prices []int) int {
	// 先转为每日收益变化
	changes := make([]int, len(prices)-1)
	for index := 1; index < len(prices); index++ {
		changes = append(changes, prices[index]-prices[index-1])
	}

	// 求连续最大
	result := 0
	prev := 0
	for index := 0; index < len(changes); index++ {
		prev = max(changes[index], changes[index] + prev)
		result = max(result, prev)
	}
	return result
}


