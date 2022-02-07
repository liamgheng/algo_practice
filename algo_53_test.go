package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test53(t *testing.T) {
	result := maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4})
	assert.Equal(t, 6, result)

	result = maxSubArray([]int{-1})
	assert.Equal(t, -1, result)

	result = maxSubArray([]int{1})
	assert.Equal(t, 1, result)
}
