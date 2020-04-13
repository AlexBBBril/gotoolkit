package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntAsString(t *testing.T) {
	t.Parallel()

	t.Run("Create random number", func(t *testing.T) {
		randomInt := Int(123, 456)

		assert.True(t, randomInt > 0)
	})
}

func TestIntAsString1(t *testing.T) {
	t.Parallel()

	t.Run("Create random number in string format", func(t *testing.T) {
		randomString := IntAsString(123, 456)

		assert.NotNil(t, randomString)
		assert.True(t, "" != randomString)
	})
}

func TestString(t *testing.T) {
	t.Parallel()

	t.Run("Create random string", func(t *testing.T) {
		randomString := String(12)

		assert.NotNil(t, randomString)
		assert.True(t, "" != randomString)
		assert.Len(t, randomString, 12)
	})
}
