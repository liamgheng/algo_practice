package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test704(t *testing.T) {
	result := search([]int{-1,0,3,5,9,12}, 9)
	assert.Equal(t, 4, result)
}

func Test704Case2(t *testing.T) {
	result := search([]int{-1,0,3,5,9,12}, 2)
	assert.Equal(t, -1, result)
}
