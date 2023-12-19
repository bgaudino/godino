package godino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssert(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			assert.True(t, ok, "Assert() failed to panic with an error")
			assert.Error(t, err)
		} else {
			t.Errorf("Assert() did not panic as expected")
		}
	}()

	Assert(false)
	AssertEqual(1, 2)
	AssertTrue(false)
	AssertFalse(true)
	AssertGreaterThan(1, 2)
	AssertGreaterThanOrEqual(1, 2)
	AssertLessThan(2, 1)
	AssertLessThanOrEqual(2, 1)
}
