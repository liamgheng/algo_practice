package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test415(t *testing.T) {
	result := addStrings("11", "123")
	assert.Equal(t, "134", result)
}

func Test425Case2(t *testing.T) {
	result := addStrings("1", "9")
	assert.Equal(t, "10", result)
}