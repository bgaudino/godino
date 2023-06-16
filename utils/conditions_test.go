package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	// Test case 1: All conditions are true
	result1 := All(true, true, true, true)
	assert.True(t, result1, "failed to evaluate correctly for all true conditions")

	// Test case 2: Some conditions are false
	result2 := All(true, false, true, true)
	assert.False(t, result2, "failed to evaluate correctly for some false conditions")

	// Test case 3: All conditions are false
	result3 := All(false, false, false)
	assert.False(t, result3, "failed to evaluate correctly for all false conditions")

	// Test case 4: No conditions provided
	result4 := All()
	assert.True(t, result4, "failed to evaluate correctly for no conditions provided")
}

func TestAny(t *testing.T) {
	// Test case 1: Some conditions are true
	result1 := Any(true, false, false, true)
	assert.True(t, result1, "failed to evaluate correctly for some true conditions")

	// Test case 2: All conditions are false
	result2 := Any(false, false, false)
	assert.False(t, result2, "failed to evaluate correctly for all false conditions")

	// Test case 3: No conditions provided
	result3 := Any()
	assert.False(t, result3, "failed to evaluate correctly for no conditions provided")
}
