package algo_practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test20(t *testing.T) {
	result := isValid("()")
	assert.Equal(t, true, result)

	result = isValid("([)]")
	assert.Equal(t, false, result)
}
