package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	result := twoSum([]int{2,7,11,15}, 9)
	assert.Equal(t, []int{0, 1}, result)

	result = twoSum([]int{3, 2, 4}, 6)
	assert.Equal(t, []int{1, 2}, result)
}
