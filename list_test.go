package helpers

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

func TestList(t *testing.T) {
	t.Run("should append a value to the end of the list", func(t *testing.T) {
		list := List[int]{1, 2, 3}
		list.Append(4)
		expected := List[int]{1, 2, 3, 4}
		assert.Equal(t, expected, list)
	})

	t.Run("should return the value at the given index", func(t *testing.T) {
		list := List[int]{1, 2, 3}
		item := list.At(-2)
		expectedItem := 2
		assert.Equal(t, expectedItem, item)
	})

	t.Run("should add the elements from the given list", func(t *testing.T) {
		list := List[int]{1, 2, 3}
		list.Extend([]int{4, 5})
		expected := List[int]{1, 2, 3, 4, 5}
		assert.Equal(t, expected, list)
	})

	t.Run("should insert an element at the given index", func(t *testing.T) {
		list := List[int]{1, 2, 3}

		list.Insert(0, 0)
		expected := List[int]{0, 1, 2, 3}
		assert.Equal(t, expected, list)

		list.Insert(99, 2)
		expected = List[int]{0, 1, 99, 2, 3}
		assert.Equal(t, expected, list)

		list.Insert(100, len(list))
		expected = List[int]{0, 1, 99, 2, 3, 100}
		assert.Equal(t, expected, list)
	})

	t.Run("should remove an item from the end of the list and return the item", func(t *testing.T) {
		list := List[int]{1, 2, 3}
		popped := list.Pop()
		expectedItem := 3
		assert.Equal(t, expectedItem, popped, "Failed to return value")
		expected := List[int]{1, 2}
		assert.Equal(t, expected, list, "Failed to modify list")
	})

	t.Run("should remove an item at the given index from the list and return the item", func(t *testing.T) {
		list := List[int]{1, 2, 3}
		removed := list.Remove(1)
		expectedItem := 2
		assert.Equal(t, expectedItem, removed, "Failed to return value")
		expected := List[int]{1, 3}
		assert.Equal(t, expected, list, "Failed to modify list")
	})

	t.Run("should reverse the list", func(t *testing.T) {
		list := List[int]{1, 2, 3}
		list.Reverse()
		expected := List[int]{3, 2, 1}
		assert.Equal(t, expected, list)
	})

	t.Run("should remove the first item from the list and return it", func(t *testing.T) {
		list := List[int]{1, 2, 3}
		shifted := list.Shift()
		expectedItem := 1
		assert.Equal(t, expectedItem, shifted, "Failed to return value")
		expected := List[int]{2, 3}
		assert.Equal(t, expected, list, "Failed to modify list")
	})

	t.Run("should add the given value to the beginning of the list", func(t *testing.T) {
		list := List[int]{1, 2, 3}
		list.UnShift(0)
		expected := List[int]{0, 1, 2, 3}
		assert.Equal(t, expected, list)
	})
}

func TestIndex(t *testing.T) {
	t.Run("should return value", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}
		index := Index(list, 3)
		expectedIndex := 2
		assert.Equal(t, expectedIndex, index)
	})

	t.Run("should return -1 (value is not present)", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}
		index := Index(list, 6)
		expectedIndex := -1
		assert.Equal(t, expectedIndex, index)
	})
}

func TestFind(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	t.Run("should find value", func(t *testing.T) {
		value, found := Find(list, func(num int) bool {
			return num%2 == 0
		})
		assert.True(t, found, "failed to find value")
		assert.Equal(t, 2, value, "returned incorrect value")
	})

	t.Run("should not find value", func(t *testing.T) {
		value, found := Find(list, func(num int) bool {
			return num > 10
		})
		assert.False(t, found, "incorrectly found a value")
		assert.Equal(t, 0, value, "returned non-zero value for non-existing condition")
	})
}

func TestFilter(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	filtered := Filter(list, func(num int) bool {
		return num%2 == 0
	})
	expected := []int{2, 4}
	assert.Equal(t, expected, filtered)
}

func TestMap(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	mapped := Map(list, func(num int) int {
		return num * num
	})
	expected := []int{1, 4, 9, 16, 25}
	assert.Equal(t, expected, mapped, "failed to map numbers to their squares")

	mappedString := Map[int, string](list, func(num int) string {
		return "Number: " + strconv.Itoa(num)
	})
	expectedStrings := []string{"Number: 1", "Number: 2", "Number: 3", "Number: 4", "Number: 5"}
	assert.Equal(t, expectedStrings, mappedString, "failed to map numbers to strings")
}

func TestReduce(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}

	sum := Reduce(list, func(acc, num int) int {
		return acc + num
	}, 0)
	expectedSum := 15
	assert.Equal(t, expectedSum, sum, "failed to calculate the sum")

	product := Reduce(list, func(acc, num int) int {
		return acc * num
	}, 1)
	expectedProduct := 120
	assert.Equal(t, expectedProduct, product, "failed to calculate the product")

	m := Reduce(list, func(acc map[int]int, num int) map[int]int {
		acc[num] = num * num
		return acc
	}, make(map[int]int))
	expectedMap := map[int]int{1: 1, 2: 4, 3: 9, 4: 16, 5: 25}
	assert.True(t, maps.Equal(m, expectedMap), "failed to make square map")
}

func TestEvery(t *testing.T) {
	// Test case 1: Every element satisfies the condition
	list1 := List[int]{2, 4, 6, 8, 10}
	every1 := list1.Every(func(num int) bool {
		return num%2 == 0
	})
	assert.True(t, every1, "failed to evaluate correctly for all elements satisfying the condition")

	// Test case 2: Not every element satisfies the condition
	list2 := List[int]{2, 4, 6, 7, 10}
	every2 := list2.Every(func(num int) bool {
		return num%2 == 0
	})
	assert.False(t, every2, "Everyfailed to evaluate correctly for not all elements satisfying the condition")

	// Test case 3: Empty list
	var emptyList List[int]
	every3 := emptyList.Every(func(num int) bool {
		return num%2 == 0
	})
	assert.True(t, every3, "failed to evaluate correctly for an empty list")
}

func TestSome(t *testing.T) {
	// Test case 1: Some elements satisfy the condition
	list1 := List[int]{1, 2, 3, 4, 5}
	some1 := list1.Some(func(num int) bool {
		return num%2 == 0
	})
	assert.True(t, some1, "failed to evaluate correctly for some elements satisfying the condition")

	// Test case 2: No element satisfies the condition
	list2 := List[int]{1, 3, 5, 7, 9}
	some2 := list2.Some(func(num int) bool {
		return num%2 == 0
	})
	assert.False(t, some2, "failed to evaluate correctly for no elements satisfying the condition")

	// Test case 3: Empty list
	var emptyList List[int]
	some3 := emptyList.Some(func(num int) bool {
		return num%2 == 0
	})
	assert.False(t, some3, "failed to evaluate correctly for an empty list")
}

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
	result, err := Zip(slice1, slice2, slice3)
	assert.Nil(t, err, "Zip() returned an unexpected error")
	assert.Equal(t, expectedResult, result, "Zip() returned an incorrect result")

	// Test case 2: Slices with different lengths
	slice4 := []int{10, 11}
	_, err = Zip(slice1, slice2, slice3, slice4)
	expectedError := errors.New("Zip() received slices of different lengths")
	assert.EqualError(t, err, expectedError.Error(), "Zip() did not return the expected error")

	// Test case 3: No input slices
	_, err = Zip[any]()
	expectedError = errors.New("Zip() expected at least 1 argument but got 0")
	assert.EqualError(t, err, expectedError.Error(), "Zip() did not return the expected error")
}
