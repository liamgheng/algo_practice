package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test160(t *testing.T) {
	headA := &ListNode{1, &ListNode{2, nil}}
	headB := &ListNode{1, &ListNode{2, nil}}
	result := getIntersectionNode(headA, headB)
	assert.Nil(t, result)
}

func Test160_case2(t *testing.T) {
	mergeNode := &ListNode{1, nil}
	headA := &ListNode{1, mergeNode}
	headB := &ListNode{1, mergeNode}
	result := getIntersectionNode(headA, headB)
	assert.Equal(t, mergeNode, result)
}

func Test160_case3(t *testing.T) {
	mergeNode := &ListNode{2, &ListNode{4, nil}}
	headA := &ListNode{1, &ListNode{9, &ListNode{1, mergeNode}}}
	headB := &ListNode{3, mergeNode}
	result := getIntersectionNode(headA, headB)
	assert.Equal(t, mergeNode, result)
}
