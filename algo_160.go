package algo_practice

// 10:47:10:59
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	headAHash := map[*ListNode]struct{}{}
	headBHash := map[*ListNode]struct{}{}
	headACurrent := headA
	headBCurrent := headB
	for headACurrent != nil || headBCurrent != nil {
		if headACurrent == nil {
			// 只处理 B
			if _, ok := headAHash[headBCurrent]; ok {
				return headBCurrent
			}
			headBCurrent = headBCurrent.Next
			continue
		}

		if headBCurrent == nil {
			//  只处理 A
			if _, ok := headBHash[headACurrent]; ok {
				return headACurrent
			}
			headACurrent = headACurrent.Next
			continue
		}

		// 两者都没结束，也没相交
		headAHash[headACurrent] = struct{}{}
		headBHash[headBCurrent] = struct{}{}
		if _, ok := headAHash[headBCurrent]; ok {
			return headBCurrent
		}
		if _, ok := headBHash[headACurrent]; ok {
			return headACurrent
		}
		headACurrent = headACurrent.Next
		headBCurrent = headBCurrent.Next
	}
	return nil
}