package algo_practice

func isValid(s string) bool {
	var stack []int
	bracketMap := map[int]int{
		']': '[',
		')': '(',
		'}': '{',
	}
	for _, ch := range s {
		item := int(ch)
		if len(stack) == 0 {
			stack = append(stack, item)
			continue
		}

		stackTop := stack[len(stack)-1]
		if matched, ok := bracketMap[item]; ok && matched == stackTop {
			stack = stack[:len(stack)-1]  // pop
		} else {
			stack = append(stack, item)
		}
	}
	return len(stack) == 0
}
