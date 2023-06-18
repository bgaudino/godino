package helpers

import "fmt"

func ExampleNewCounter() {
	fruits := []string{"apple", "banana", "banana", "banana", "orange", "orange"}
	counter := NewCounter(fruits)
	fmt.Println(counter.Elements())
	// Output:
	// [{apple 1} {banana 3} {orange 2}]
}

func ExampleCounter_Add() {
	fruits := []string{"apple", "banana", "banana", "banana", "orange", "orange"}
	counter := NewCounter(fruits)
	fmt.Println(counter.Get("apple"))

	// Adding an element already in the counter
	counter.Add("apple")
	fmt.Println(counter.Get("apple"))

	// Adding an element not yet in the counter
	counter.Add("grape")
	fmt.Println(counter.Get("grape"))
	// Output:
	// 1
	// 2
	// 1
}

func ExampleCounter_Elements() {
	fruits := []string{"apple", "banana", "banana", "banana", "orange", "orange"}
	counter := NewCounter(fruits)
	fmt.Println(counter.Elements())
	// Output:
	// [{apple 1} {banana 3} {orange 2}]
}

func ExampleCounter_Get() {
	fruits := []string{"apple", "banana", "banana", "banana", "orange", "orange"}
	counter := NewCounter(fruits)
	fmt.Println(counter.Get("banana"))
	fmt.Println(counter.Get("grape"))
	// Output:
	// 3
	// 0
}

func ExampleCounter_MostCommon() {
	fruits := []string{"apple", "banana", "banana", "banana", "orange", "orange"}
	counter := NewCounter(fruits)
	top2 := counter.MostCommon(2)
	fmt.Println(top2)

	counter.Add("orange") // Ties are broken by the order elements were added
	sortedElements := counter.MostCommon(-1)
	fmt.Println(sortedElements)

	// Output:
	// [{banana 3} {orange 2}]
	// [{banana 3} {orange 3} {apple 1}]
}

func ExampleCounter_Subtract() {
	fruits := []string{"apple", "banana", "banana", "banana", "orange", "orange"}
	counter := NewCounter(fruits)

	// Subtracting an element in the counter
	counter.Subtract("banana")
	fmt.Println(counter.Get("banana"))

	// Subtracting an element not in the counter
	counter.Subtract("grape")
	fmt.Println(counter.Get("grape"))

	// Output:
	// 2
	// -1
}

func ExampleCounter_Total() {
	fruits := []string{"apple", "banana", "banana", "banana", "orange", "orange"}
	counter := NewCounter(fruits)
	fmt.Println(counter.Total())
	// Output: 6
}

func ExampleCounter_Update() {
	fruits := []string{"apple", "banana", "banana", "banana", "orange", "orange"}
	counter := NewCounter(fruits)

	moreFruits := []string{"apple", "banana", "orange", "grape", "grape"}
	evenMoreFruits := []string{"apple", "banana", "pineapple"}
	counter.Update(moreFruits, evenMoreFruits)

	fmt.Println(counter.Elements())
	// Output:
	// [{apple 3} {banana 5} {orange 3} {grape 2} {pineapple 1}]
}
