package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test215(t *testing.T) {
	result := findKthLargest([]int{3,2,1,5,6,4}, 2)
	assert.Equal(t, 5, result)
}
