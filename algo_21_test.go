package algo_practice

import "testing"

func Test21(t *testing.T) {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	result := mergeTwoLists(l1, l2)
	printLinkList(result)
}
