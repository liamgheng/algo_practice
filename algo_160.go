package algo_practice

// 10:47:10:59
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	headAHash := map[*ListNode]struct{}{}
	headACurrent := headA

	for headACurrent != nil {
		headAHash[headACurrent] = struct{}{}
		headACurrent = headACurrent.Next
	}

	headBCurrent := headB
	for headBCurrent != nil {
		if _, ok := headAHash[headBCurrent]; ok {
			return headBCurrent
		}
		headBCurrent = headBCurrent.Next
	}

	return nil
}