package godino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCounter(t *testing.T) {
	c := NewCounter([]string{"foo", "bar", "baz", "foo", "foo", "baz"})
	expected := counterElements[string]{{"foo", 3}, {"bar", 1}, {"baz", 2}}
	assert.Equal(t, expected, c.Elements())
}

func TestCounterAdd(t *testing.T) {
	c := NewCounter([]string{"foo", "bar", "baz", "foo", "foo", "baz"})
	c.Add("bar")
	assert.Equal(t, 2, c.counts["bar"])
}

func TestGet(t *testing.T) {
	c := NewCounter([]string{"foo", "bar", "baz", "foo", "foo", "baz"})
	assert.Equal(t, 3, c.Get("foo"))
}

func TestMostCommon(t *testing.T) {
	c := NewCounter([]string{"foo", "bar", "baz", "foo", "foo", "baz"})

	t.Run("get all elements sorted by most common", func(t *testing.T) {
		mostCommon := c.MostCommon(-1)
		expected := counterElements[string]{
			{"foo", 3},
			{"baz", 2},
			{"bar", 1},
		}
		assert.Equal(t, expected, mostCommon)
	})

	t.Run("get n most common elements", func(t *testing.T) {
		mostCommon := c.MostCommon(2)
		expected := counterElements[string]{
			{"foo", 3},
			{"baz", 2},
		}
		assert.Equal(t, expected, mostCommon)
	})

	t.Run("ties are broken by element first inserted", func(t *testing.T) {
		c := NewCounter([]string{"foo", "bar", "baz", "baz"})
		mostCommon := c.MostCommon(1)
		expected := counterElements[string]{{"baz", 2}}
		assert.Equal(t, expected, mostCommon)

		c.Add("foo")
		mostCommon2 := c.MostCommon(1)
		expected2 := counterElements[string]{{"foo", 2}}
		assert.Equal(t, expected2, mostCommon2)
	})
}

func TestSubtract(t *testing.T) {
	c := NewCounter([]string{"foo", "bar", "baz", "foo", "foo", "baz"})
	assert.Equal(t, 3, c.Get("foo"))

	c.Subtract("foo")
	assert.Equal(t, 2, c.Get("foo"))
}

func TestTotal(t *testing.T) {
	c := NewCounter([]string{"foo", "bar", "baz", "foo", "foo", "baz"})
	assert.Equal(t, 6, c.Total())
}

func TestUpdate(t *testing.T) {
	c := NewCounter([]int{1, 1, 1, 2, 3, 3})
	arr1 := []int{2, 3, 3, 4, 5}
	arr2 := []int{4, 5, 5, 6}
	c.Update(arr1, arr2)
	expected := counterElements[int]{
		{1, 3}, {2, 2}, {3, 4}, {4, 2}, {5, 3}, {6, 1},
	}
	assert.Equal(t, expected, c.Elements())
}
