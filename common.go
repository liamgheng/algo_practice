package algo_practice

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func printLinkList(listNode *ListNode)  {
	current := listNode
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}