package algo_practice

func lengthOfLongestSubstring(s string) int {
	result := 0
	for left := 0; left < len(s); left++ {
		kv := map[uint8]bool{}
		right := left
		for ; right < len(s); right ++ {
			ch := s[right]
			if _, ok := kv[ch]; ok {
				break
			} else {
				kv[ch] = true
			}
		}
		if right - left > result {
			result = right - left
		}
	}
	return result
}
