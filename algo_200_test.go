package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	result := numIslands(grid)
	assert.Equal(t, 1, result)
}
