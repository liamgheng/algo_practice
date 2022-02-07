package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test15(t *testing.T) {
	//result := threeSum([]int{-1,0,1,2,-1,-4})
	//assert.ElementsMatch(t, result, [][]int{
	//	{-1, 0, 1},
	//	{-1, -1, 2},
	//})

	result := threeSum([]int{1, -1, -1, 0})
	assert.ElementsMatch(t, result, [][]int{
		{-1, 0, 1},
	})
}
