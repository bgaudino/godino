package helpers

import (
	"fmt"
	"strings"
)

func ExampleAll() {
	x := 9
	isValid := All(x > 0, x < 10, x%2 != 0)
	fmt.Println(isValid)
	// Output:
	// true
}

func ExampleAny() {
	fruits := "apple banana grape"
	isValid := Any(strings.HasPrefix(fruits, "grape"), strings.Contains(fruits, "banana"), strings.HasSuffix(fruits, "apple"))
	fmt.Println(isValid)
	// Output:
	// true
}

func ExampleMax() {
	max, err := Max(1, 23, 100, -5, 10)
	fmt.Println(max, err)

	max2, err2 := Max[int]()
	fmt.Println(max2, err2)
	// Output:
	// 100 <nil>
	// 0 Max() expected at least 1 argument, got 0
}

func ExampleMin() {
	min, err := Min(1, 23, 100, -5, 10)
	fmt.Println(min, err)

	min2, err2 := Min[int]()
	fmt.Println(min2, err2)
	// Output:
	// -5 <nil>
	// 0 Min() expected at least 1 argument, got 0
}

func ExampleProd() {
	product := Prod(1, 23, 100, -5, 10)
	fmt.Println(product)
	// Output -11500
}

func ExampleSum() {
	sum := Sum(1, 23, 100, -5, 10)
	fmt.Println(sum)
	// Output: 129
}
