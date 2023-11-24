package kata1_test

import (
	"golang-kata/kata1"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKata1(t *testing.T) {
	t.Run("if input is array 1 number that cannot square root", func(t *testing.T) {
		result := kata1.Kata1([]int{2})
		assert.Equal(t, []int{4}, result)
	})

	t.Run("if input is array 1 number that can square root", func(t *testing.T) {
		result := kata1.Kata1([]int{4})
		assert.Equal(t, []int{2}, result)
	})

	t.Run("if input is array 2 number that both can and cannot square root", func(t *testing.T) {
		result := kata1.Kata1([]int{3, 16})
		assert.Equal(t, []int{9, 4}, result)
	})
}
