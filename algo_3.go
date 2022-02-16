package algo_practice

// 第一次 n * n 复杂度太高了
// 第二次改用滑动窗口算法
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	preLeft := 0
	left := 0
	kv := make(map[uint8]int)  // char: index

	for right := 0; right < len(s); right++ {
		ch := s[right]
		if index, ok := kv[ch]; ok {
			// 字符串之前出现过，所以 left 移动到出现位置的后一位
			left = index + 1

			// index 之前的 kv 需要删除，因为重新开始计算了
			for i := preLeft; i < left; i ++ {
				delete(kv, s[i])
			}

			// 本元素重新加入 kv，并且计算本轮之后的 max
			kv[ch] = right

			preLeft = left
		} else {
			// MARK: 第一次忘记添加
			kv[ch] = right
		}
		if right - left + 1 > maxLength {
			maxLength = right - left + 1
		}
	}

	return maxLength
}
