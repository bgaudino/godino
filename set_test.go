package godino

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

func TestNewSet(t *testing.T) {
	t.Run("creating a new set of ints", func(t *testing.T) {
		set := NewSet(1, 2, 3)
		expected := map[int]struct{}{1: {}, 2: {}, 3: {}}
		assert.Equal(t, maps.Equal(set, expected), true)
	})
	t.Run("creating a new set of strings", func(t *testing.T) {
		set := NewSet("foo", "bar", "baz")
		expected := map[string]struct{}{"foo": {}, "bar": {}, "baz": {}}
		assert.Equal(t, maps.Equal(set, expected), true)
	})
}

func TestAdd(t *testing.T) {
	t.Run("adding a single value", func(t *testing.T) {
		set := NewSet[int]()
		val := 1
		set.Add(val)
		assert.Equal(t, set.Has(val), true)
	})
	t.Run("adding many values", func(t *testing.T) {
		values := []string{"foo", "bar", "baz"}
		set := NewSet[string](values...)
		assert.ElementsMatch(t, set.Members(), values)
	})
}

func TestCopy(t *testing.T) {
	t.Run("copying a set", func(t *testing.T) {
		set1 := NewSet(1, 2, 3)
		set2 := set1.Copy()
		assert.ElementsMatch(t, set1.Members(), set2.Members())
	})
}

func TestClear(t *testing.T) {
	set := NewSet(1, 2, 3)
	set.Clear()
	assert.Equal(t, len(set), 0)
}

func TestDifference(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	set3 := NewSet(-1, 0, 1)
	diff := set1.Difference(set2, set3)
	assert.ElementsMatch(t, diff.Members(), []int{2})
}

func TestDifferenceUpdate(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	set3 := NewSet(-1, 0, 1)
	set1.DifferenceUpdate(set2, set3)
	assert.ElementsMatch(t, set1.Members(), []int{2})
}

func TestDiscard(t *testing.T) {
	set := NewSet(1, 2, 3)
	value := 1
	set.Discard(value)
	assert.False(t, set.Has(value))
}

func TestHas(t *testing.T) {
	set := NewSet(1, 2, 3)
	t.Run("has value", func(t *testing.T) {
		assert.Equal(t, set.Has(1), true)
	})
	t.Run("does not have value", func(t *testing.T) {
		assert.Equal(t, set.Has(4), false)
	})
}

func TestIntersection(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	set3 := NewSet(3, 4, 5)
	intersection := set1.Intersection(set2, set3)
	assert.ElementsMatch(t, intersection.Members(), []int{3})
}

func TestIntersectionUpdate(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	set3 := NewSet(3, 4, 5)
	set1.IntersectionUpdate(set2, set3)
	assert.ElementsMatch(t, set1.Members(), []int{3})
}

func TestIsDisjoint(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	t.Run("is disjoint", func(t *testing.T) {
		set2 := NewSet(4, 5, 6)
		assert.Equal(t, set1.IsDisjoint(set2), true)
	})
	t.Run("is not disjoint", func(t *testing.T) {
		set2 := NewSet(3, 4, 5)
		assert.Equal(t, set1.IsDisjoint(set2), false)
	})
}

func TestIsSubset(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	t.Run("is subset", func(t *testing.T) {
		set2 := NewSet(1, 2, 3, 4)
		assert.Equal(t, set1.IsSubset(set2), true)
	})
	t.Run("is not subset", func(t *testing.T) {
		set2 := NewSet(2, 3, 4, 5)
		assert.Equal(t, set1.IsSubset(set2), false)
	})
}

func TestIsSuperset(t *testing.T) {
	set1 := NewSet(1, 2, 3, 4)
	t.Run("is superset", func(t *testing.T) {
		set2 := NewSet(1, 2, 3)
		assert.Equal(t, set1.IsSuperset(set2), true)
	})
	t.Run("is not superset", func(t *testing.T) {
		set2 := NewSet(3, 4, 5)
		assert.Equal(t, set1.IsSuperset(set2), false)
	})
}

func TestMembers(t *testing.T) {
	set := NewSet(1, 2, 3)
	assert.ElementsMatch(t, set.Members(), []int{1, 2, 3})
}

func TestPop(t *testing.T) {
	t.Run("pop from an empty set", func(t *testing.T) {
		set := NewSet[int]()
		popped, err := set.Pop()
		assert.Error(t, err)
		assert.Zero(t, popped)
		assert.Zero(t, len(set))
	})
	t.Run("pop from a set with one element", func(t *testing.T) {
		set := NewSet("foo")
		popped, err := set.Pop()
		assert.Nil(t, err)
		assert.Equal(t, "foo", popped)
		assert.Zero(t, len(set))
	})
	t.Run("pop from a set with multiple elements", func(t *testing.T) {
		set := NewSet(1, 2, 3)
		popped, err := set.Pop()
		assert.Nil(t, err)
		assert.NotContains(t, set, popped)
		assert.Equal(t, len(set), 2)
	})
}

func TestRemove(t *testing.T) {
	t.Run("removing an item that exists in the set", func(t *testing.T) {
		set := NewSet(1, 2, 3)
		value := 1
		removed := set.Remove(value)
		assert.False(t, set.Has(value))
		assert.True(t, removed)
	})
	t.Run("removing an item that does not exist in the set", func(t *testing.T) {
		set := NewSet(1, 2, 3)
		removed := set.Remove(4)
		assert.False(t, removed)
	})
}

func TestSymmetricDifference(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	diff := set1.SymmetricDifference(set2)
	assert.ElementsMatch(t, diff.Members(), []int{1, 4})
}

func TestSymmetricDifferenceUpdate(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	set1.SymmetricDifferenceUpdate(set2)
	assert.ElementsMatch(t, set1.Members(), []int{1, 4})
}

func TestUnion(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	union := set1.Union(set2)
	assert.ElementsMatch(t, union.Members(), []int{1, 2, 3, 4, 5})
}

func TestUnionUpdate(t *testing.T) {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	set1.Update(set2)
	assert.ElementsMatch(t, set1.Members(), []int{1, 2, 3, 4, 5})
}
