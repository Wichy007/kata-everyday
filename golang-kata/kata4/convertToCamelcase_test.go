package kata4_test

import (
	"golang-kata/kata4"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToCamelcase(t *testing.T) {
	t.Run("convert word to capitalize except first word", func(t *testing.T) {
		result := kata4.ToCamelCase("the")
		assert.Equal(t, "the", result)
	})
}
