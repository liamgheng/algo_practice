package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test146(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	result := cache.Get(1)
	assert.Equal(t, 1, result, "error")
	cache.Put(3, 3)
	result = cache.Get(2)
	assert.Equal(t, -1, result, "error")
}
