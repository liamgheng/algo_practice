package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test3(t *testing.T) {
	result := lengthOfLongestSubstring("abcabcbb")
	assert.Equal(t, 3, result)
}

func Test3Case2(t *testing.T) {
	result := lengthOfLongestSubstring("pwwkew")
	assert.Equal(t, 3, result)
}

func Test3Case4(t *testing.T) {
	result := lengthOfLongestSubstring(" ")
	assert.Equal(t, 1, result)
}