package helpers

import (
	"fmt"
	"sort"
)

func ExampleNewSet() {
	set := NewSet(1, 2, 3, 3) // Values are unique
	fmt.Println(len(set))
	// Output: 3
}
func ExampleSet_Add() {
	set := NewSet(1, 2, 3)
	set.Add(4)
	fmt.Println(set.Has(4))

	set.Add(5, 6, 7)
	expected := NewSet(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(expected.Equals(set))
	// Output:
	// true
	// true
}

func ExampleSet_Clear() {
	set := NewSet(1, 2, 3)
	set.Clear()
	fmt.Println(set.Members())
	// Output: []
}

func ExampleSet_Copy() {
	set1 := NewSet(1, 2, 3)
	set2 := set1.Copy()

	fmt.Println(set1.Equals(set2))
	// Output: true
}

func ExampleSet_Difference() {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	difference := set1.Difference(set2)

	expected := NewSet(1)
	fmt.Println(difference.Equals(expected))
	// Output: true
}

func ExampleSet_DifferenceUpdate() {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	set1.DifferenceUpdate(set2)

	expected := NewSet(1)
	fmt.Println(expected.Equals(set1))
	// Output: true
}

func ExampleSet_Discard() {
	set := NewSet(1, 2, 3)
	set.Discard(1)
	fmt.Println(set.Has(1))
	// Output: false
}

func ExampleSet_Equals() {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	fmt.Println(set1.Equals(set2))

	set1.Add(4)
	set2.Add(1)
	fmt.Println(set1.Equals(set2))
	// Output:
	// false
	// true
}

func ExampleSet_Has() {
	set := NewSet(1, 2, 3)
	fmt.Println(set.Has(1))
	fmt.Println(set.Has(4))
	// Output:
	// true
	// false
}

func ExampleSet_Intersection() {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	intersection := set1.Intersection(set2)

	expected := NewSet(2, 3)
	fmt.Println(intersection.Equals(expected))
	// Output: true
}

func ExampleSet_IntersectionUpdate() {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	set1.IntersectionUpdate(set2)

	expected := NewSet(2, 3)
	fmt.Println(set1.Equals(expected))
	// Output: true
}

func ExampleSet_IsDisjoint() {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(4, 5, 6)
	fmt.Println(set1.IsDisjoint(set2))

	set1.Add(4)
	fmt.Println(set1.IsDisjoint(set2))
	// Output:
	// true
	// false
}

func ExampleSet_IsSubset() {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(1, 2, 3, 4)
	fmt.Println(set1.IsSubset(set2))

	set1.Add(5)
	fmt.Println(set1.IsSubset(set2))
	// Output:
	// true
	// false
}

func ExampleSet_IsSuperset() {
	set1 := NewSet(1, 2, 3, 4)
	set2 := NewSet(2, 3, 4)
	fmt.Println(set1.IsSuperset(set2))

	set2.Add(5)
	fmt.Println(set1.IsSuperset(set2))
	// Output:
	// true
	// false
}

func ExampleSet_Members() {
	set := NewSet(1, 2, 3)
	members := set.Members()

	sortable := sort.IntSlice(members)
	sortable.Sort()
	fmt.Println(sortable)
	// Output: [1 2 3]
}

func ExampleSet_Pop() {
	set := NewSet(1)
	val, err := set.Pop()
	fmt.Println(val, err, len(set))

	val2, err2 := set.Pop()
	fmt.Println(val2, err2, len(set))
	// Output:
	// 1 <nil> 0
	// 0 cannot pop item from an empty set 0
}

func ExampleSet_Remove() {
	set := NewSet(1, 2, 3)
	ok := set.Remove(3)
	fmt.Println(ok)
	fmt.Println(set.Has(3))

	ok2 := set.Remove(4)
	fmt.Println(ok2)
	// Output:
	// true
	// false
	// false
}

func ExampleSet_SymmetricDifference() {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	symmetricDifference := set1.SymmetricDifference(set2)

	expected := NewSet(1, 4)
	fmt.Println(expected.Equals(symmetricDifference))
	// Output: true
}

func ExampleSet_SymmetricDifferenceUpdate() {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(2, 3, 4)
	set1.SymmetricDifferenceUpdate(set2)

	expected := NewSet(1, 4)
	fmt.Println(expected.Equals(set1))
	// Output: true
}

func ExampleSet_Union() {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	union := set1.Union(set2)
	expected := NewSet(1, 2, 3, 4, 5)
	fmt.Println(expected.Equals(union))
}

func ExampleSet_Update() {
	set1 := NewSet(1, 2, 3)
	set2 := NewSet(3, 4, 5)
	set1.Update(set2)
	expected := NewSet(1, 2, 3, 4, 5)
	fmt.Println(expected.Equals(set1))
}
