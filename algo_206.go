package algo_practice

// 时间复杂度 N
// 空间复杂度 1
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		// MARK: 把这两步写反了，应该先修改 prev，再修改 current
		prev = current
		current = next
	}
	return prev
}
