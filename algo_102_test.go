package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test102(t *testing.T) {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{
		20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil},
	}}
	result := levelOrder(root)
	assert.Equal(t, [][]int{
		{3},
		{9, 20},
		{15, 7},
	}, result)
}
