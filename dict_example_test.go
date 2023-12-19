package godino

import (
	"fmt"
	"sort"
)

func ExampleDict_Clear() {
	d := Dict[int, string]{1: "foo"}
	d.Clear()
	fmt.Println(d.Items())
	// Output: []
}

func ExampleDict_Copy() {
	d := Dict[int, string]{1: "foo"}
	d2 := d.Copy()
	fmt.Println(d2.Items())
	// Output: [{1 foo}]
}

func ExampleDict_Get() {
	d := Dict[int, string]{1: "foo"}

	fmt.Println(d.Get(1))
	fmt.Println(d.Get(2))
	fmt.Println(d.Get(2, "bar"))
	// Output:
	// foo
	//
	// bar
}

func ExampleDict_Has() {
	d := Dict[int, string]{1: "foo"}

	fmt.Println(d.Has(1))
	fmt.Println(d.Has(2))
	// Output:
	// true
	// false
}

func ExampleDict_Items() {
	d := Dict[int, string]{
		1: "foo",
		2: "bar",
		3: "baz",
	}
	items := d.Items()
	sort.Slice(items, func(i, j int) bool { return items[i].Key < items[j].Key })
	fmt.Println(items)
	// Output:
	// [{1 foo} {2 bar} {3 baz}]
}

func ExampleDict_Keys() {
	d := Dict[int, string]{
		1: "foo",
		2: "bar",
		3: "baz",
	}

	keys := d.Keys()
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	fmt.Println(keys)
	// Output:
	// [1 2 3]
}

func ExampleDict_Pop() {
	d := Dict[int, string]{
		1: "foo",
		2: "bar",
		3: "baz",
	}
	popped, ok := d.Pop(1)
	fmt.Println(popped, ok)
	fmt.Println(d.Has(1))
	fmt.Println(len(d))

	popped2, ok2 := d.Pop(4, "fallback")
	fmt.Println(popped2, ok2)

	popped3, ok3 := d.Pop(5)
	fmt.Println(popped3, ok3)
	// Output:
	// foo true
	// false
	// 2
	// fallback false
	//  false
}

func ExampleDict_SetDefault() {
	d := Dict[int, string]{
		1: "foo",
		2: "bar",
	}
	d2 := Dict[int, string]{
		2: "baz",
		3: "bar",
	}
	d.Update(d2)

	items := d.Items()
	sort.Slice(items, func(i, j int) bool {
		return items[i].Key < items[j].Key
	})
	fmt.Println(items)
	// Output:
	// [{1 foo} {2 baz} {3 bar}]
}

func ExampleDict_Values() {
	d := Dict[int, string]{
		1: "foo",
		2: "bar",
		3: "baz",
	}
	values := d.Values()
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	fmt.Println(values)
	// Output:
	// [bar baz foo]
}
