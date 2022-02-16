package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test206WithNil(t *testing.T) {
	result := reverseList(nil)
	var expect *ListNode
	assert.Equal(t, expect, result)
}

func Test206WithOne(t *testing.T) {
	data := &ListNode{1, nil}
	result := reverseList(data)
	assert.Equal(t, data, result)
}

func Test206WithTwo(t *testing.T) {
	data := &ListNode{1, &ListNode{2, nil}}
	expect := &ListNode{2, &ListNode{1, nil}}
	result := reverseList(data)
	assert.Equal(t, expect, result)
}
