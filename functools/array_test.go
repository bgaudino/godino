package functools

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	arr := []int{1, 2, 3}
	Append(&arr, 4)
	assert.Equal(t, []int{1, 2, 3, 4}, arr)
}

func TestClearArray(t *testing.T) {
	arr := []int{1, 2, 3}
	Clear(&arr)
	assert.Empty(t, arr)
}

func TestCopyArray(t *testing.T) {
	arr := []int{1, 2, 3}
	assert.Equal(t, Copy(arr), arr)
}

func TestContains(t *testing.T) {
	arr := []int{1, 2, 3}
	assert.True(t, Contains(arr, 3))
	assert.False(t, Contains(arr, 0))
}

func TestExtend(t *testing.T) {
	arr := []int{1, 2, 3}
	Extend(&arr, []int{4, 5, 6})
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, arr)
}

func TestIndex(t *testing.T) {
	t.Run("should return value", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		index := Index(arr, 3)
		expectedIndex := 2
		assert.Equal(t, expectedIndex, index)
	})

	t.Run("should return -1 (value is not present)", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		index := Index(arr, 6)
		expectedIndex := -1
		assert.Equal(t, expectedIndex, index)
	})
}

func TestInsert(t *testing.T) {
	arr := []int{1, 2, 3}

	Insert(&arr, 0, 0)
	expected := []int{0, 1, 2, 3}
	assert.Equal(t, expected, arr)

	Insert(&arr, 99, 2)
	expected = []int{0, 1, 99, 2, 3}
	assert.Equal(t, expected, arr)

	Insert(&arr, 100, len(arr))
	expected = []int{0, 1, 99, 2, 3, 100}
	assert.Equal(t, expected, arr)
}

func TestPopArray(t *testing.T) {
	arr := []int{1, 2, 3}
	popped := Pop(&arr)
	expectedItem := 3
	assert.Equal(t, expectedItem, popped, "Failed to return value")
	expected := []int{1, 2}
	assert.Equal(t, expected, arr, "Failed to modify list")
}

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	Reverse(&arr)
	expected := []int{4, 3, 2, 1}
	assert.Equal(t, expected, arr)
}

func TestShift(t *testing.T) {
	arr := []int{1, 2, 3}
	shifted := Shift(&arr)
	expectedItem := 1
	assert.Equal(t, expectedItem, shifted, "Failed to return value")
	expected := []int{2, 3}
	assert.Equal(t, expected, arr, "Failed to modify list")
}

func TestValueAt(t *testing.T) {
	t.Run("should return the value at the given index", func(t *testing.T) {
		arr := []int{1, 2, 3}
		item := ValueAt(arr, 2)
		assert.Equal(t, 3, item)
	})

	t.Run("should support negative index", func(t *testing.T) {
		arr := []int{1, 2, 3}
		item := ValueAt(arr, -2)
		assert.Equal(t, 2, item)
	})
}

func TestZip(t *testing.T) {
	// Test case 1: Valid input slices
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice3 := []int{7, 8, 9}
	expectedResult := [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
	result, err := Zip[[]int, int](slice1, slice2, slice3)
	assert.Nil(t, err, "Zip() returned an unexpected error")
	assert.Equal(t, expectedResult, result, "Zip() returned an incorrect result")

	// Test case 2: Slices with different lengths
	slice4 := []int{10, 11}
	_, err = Zip[[]int, int](slice1, slice2, slice3, slice4)
	expectedError := errors.New("Zip() received slices of different lengths")
	assert.EqualError(t, err, expectedError.Error(), "Zip() did not return the expected error")

	// Test case 3: No input slices
	_, err = Zip[[]int, int]()
	expectedError = errors.New("Zip() expected at least 1 argument but got 0")
	assert.EqualError(t, err, expectedError.Error(), "Zip() did not return the expected error")
}
