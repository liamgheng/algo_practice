package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test33(t *testing.T) {
	//result := search([]int{4,5,6,7,0,1,2}, 6)
	//assert.Equal(t, 2, result)

	result := search([]int{1}, 0)
	assert.Equal(t, -1, result)
}

func Test33Case2(t *testing.T) {
	result := search([]int{4,5,6,7,0,1,2}, 0)
	assert.Equal(t, 4, result)
}
