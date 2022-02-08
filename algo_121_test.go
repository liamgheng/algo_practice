package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test121(t *testing.T) {
	result := maxProfit([]int{7,1,5,3,6,4})
	assert.Equal(t, 5, result)

	result = maxProfit([]int{7,6})
	assert.Equal(t, 0, result)
}
