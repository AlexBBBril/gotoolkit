package env

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func TestRequired(t *testing.T) {
	assert := assert.New(t)

	t.Parallel()

	t.Run("Panic on receiving a nonexistent variable", func(t *testing.T) {
		_ = os.Unsetenv("TEST")

		defer func() {
			if r := recover(); nil != r {
				assert.Equal("Missing environment variable named: TEST", fmt.Sprintf("%v", r))
			} else {
				t.Fatal("did panic")
			}
		}()

		Required("TEST")
	})

	t.Run("Success", func(t *testing.T) {
		test := "test"

		_ = os.Setenv("TEST", "test")

		rTest := Required("TEST")

		assert.Equal(test, rTest)

		_ = os.Unsetenv("TEST")
	})
}

func TestRequiredAsInt(t *testing.T) {
	assert := assert.New(t)

	t.Parallel()

	t.Run("Panic on receiving a nonexistent variable", func(t *testing.T) {
		_ = os.Unsetenv("TEST")

		defer func() {
			if r := recover(); nil != r {
				assert.Equal("Missing environment variable named: TEST", fmt.Sprintf("%v", r))
			} else {
				t.Fatal("did panic")
			}
		}()

		RequiredAsInt("TEST")
	})

	t.Run("Success", func(t *testing.T) {
		test := 123

		_ = os.Setenv("TEST", strconv.Itoa(test))

		rTest := RequiredAsInt("TEST")

		assert.Equal(test, rTest)

		_ = os.Unsetenv("TEST")
	})

	t.Run("Translation error string to int", func(t *testing.T) {
		defer func() {
			if r := recover(); nil != r {
				assert.Equal("strconv.Atoi: parsing \"test\": invalid syntax", fmt.Sprintf("%v", r))
			} else {
				t.Fatal("did panic")
			}
		}()

		test := "test"

		_ = os.Setenv("TEST", test)

		rTest := RequiredAsInt("TEST")

		assert.Equal(test, rTest)

		_ = os.Unsetenv("TEST")
	})
}
