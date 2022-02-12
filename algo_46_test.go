package algo_practice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test46(t *testing.T) {
	result := permute([]int{1,2,3})
	assert.Equal(t, 6, len(result))
	fmt.Println(result)
}
