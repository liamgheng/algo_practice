package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test3(t *testing.T) {
	result := lengthOfLongestSubstring("abcabcbb")
	assert.Equal(t, result, 3)
}
