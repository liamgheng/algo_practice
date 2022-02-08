package algo_practice

func hasCycle(head *ListNode) bool {
	seen := map[*ListNode]struct{}{}
	current := head
	for current != nil {
		if _, ok := seen[current]; ok {
			return true
		}
		seen[current] = struct{}{}
		current = current.Next
	}
	return false
}
