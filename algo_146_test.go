package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test146WithCap2(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)
	// 2 会被删除
	cache.Put(3, 3)

	assert.Equal(t, -1, cache.Get(2))
	assert.Equal(t, 1, cache.Get(1))
}

func Test146WithCap1(t *testing.T) {
	cache := Constructor(1)
	cache.Put(1, 1)
	cache.Put(2, 2)

	assert.Equal(t, -1, cache.Get(1))
	assert.Equal(t, 2, cache.Get(2))
}

func Test146WithWrong(t *testing.T) {
	// leetcode 报错的 case
	cache := Constructor(2)
	cache.Put(1, 0)
	cache.Put(2, 2)
	cache.Get(1)
	cache.Put(3, 3)
	cache.Get(2)  // 上一步应该是 3 1
	cache.Put(4, 4)
	result := cache.Get(1)

	assert.Equal(t, -1, result)
}
