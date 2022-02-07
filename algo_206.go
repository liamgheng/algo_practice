package algo_practice

func reverseList(head *ListNode) *ListNode {
	// 用来在遍历时存储前一个节点
	var prev *ListNode
	current := head

	for current != nil {
		// 把 next 先存储下来，因为之后会改动 current.next
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
