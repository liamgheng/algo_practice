package algo_practice

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	currentList1 := list1
	currentList2 := list2
	result := &ListNode{-1, nil}  // 哨兵节点
	current := result

	for currentList1 != nil || currentList2 != nil {
		if currentList1 == nil {
			current.Next = currentList2
			break
		}
		if currentList2 == nil {
			current.Next = currentList1
			break
		}
		if currentList1.Val < currentList2.Val {
			current.Next = currentList1
			currentList1 = currentList1.Next
		} else {
			current.Next = currentList2
			currentList2 = currentList2.Next
		}
		current = current.Next
	}

	return result.Next
}
