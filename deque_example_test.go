package helpers

import "fmt"

func ExampleNewDeque() {
	deque := NewDeque[int](8)
	deque.ExtendRight([]int{1, 2, 3})
	fmt.Println(deque)
	// Output:
	// [1 2 3]
}

func ExampleDeque_Capacity() {
	deque := NewDeque[int](8) // Provide minimum capacity
	fmt.Println(deque.Capacity())

	for i := 0; i < 10; i++ {
		deque.PushRight(i)
	}
	fmt.Println(deque.Capacity()) // Automatically grows as needed

	deque.Clear()
	fmt.Println(deque.Capacity()) // Automatically shrinks as needed
	// Output:
	// 8
	// 16
	// 8
}

func ExampleDeque_Clear() {
	deque := NewDeque[int]()
	deque.ExtendRight([]int{1, 2, 3})
	deque.Clear()
	fmt.Println(deque.Len())
	// Output: 0
}

func ExampleDeque_Copy() {
	deque := NewDeque[int]()
	deque.ExtendRight([]int{1, 2, 3})
	deque2 := deque.Copy()
	fmt.Println(deque, deque2)
	// Output:
	// [1 2 3] [1 2 3]
}

func ExampleDeque_Elements() {
	deque := NewDeque[int]()
	deque.ExtendRight([]int{1, 2, 3})
	elements := deque.Elements()
	fmt.Println(elements)
	// Output:
	// [1 2 3]
}

func ExampleDeque_ExtendLeft() {
	deque := NewDeque[int]()
	deque.ExtendLeft([]int{1, 2, 3})
	fmt.Println(deque)
	// Output:
	// [3 2 1]
}

func ExampleDeque_ExtendRight() {
	deque := NewDeque[int]()
	deque.ExtendRight([]int{1, 2, 3})
	fmt.Println(deque)
	// Output:
	// [1 2 3]
}

func ExampleDeque_Index() {
	deque := NewDeque[int]()
	deque.ExtendRight([]int{1, 2, 3, 4, 5})

	// Sindle index
	fmt.Println(deque.Index(2))

	// Two indices
	fmt.Println(deque.Index(1, 4))

	// Output:
	// [3]
	// [2 3 4]
}

func ExampleDeque_PeekLeft() {
	deque := NewDeque[int]()
	deque.ExtendRight([]int{1, 2, 3})
	head := deque.PeekLeft()
	fmt.Println(head)
	// Output: 1
}

func ExampleDeque_PopLeft() {
	deque := NewDeque[int]()
	deque.ExtendRight([]int{1, 2, 3})
	popped := deque.PopLeft()
	fmt.Println(popped)
	fmt.Println(deque)
	// Output:
	// 1
	// [2 3]
}

func ExampleDeque_PeekRight() {
	deque := NewDeque[int]()
	deque.ExtendRight([]int{1, 2, 3})
	tail := deque.PeekRight()
	fmt.Println(tail)
	// Output: 3
}

func ExampleDeque_PopRight() {
	deque := NewDeque[int]()
	deque.ExtendRight([]int{1, 2, 3})
	popped := deque.PopRight()
	fmt.Println(popped)
	fmt.Println(deque)
	// Output:
	// 3
	// [1 2]
}

func ExampleDeque_PushLeft() {
	deque := NewDeque[int]()
	deque.ExtendRight([]int{1, 2, 3})
	deque.PushLeft(4)
	fmt.Println(deque)
	// Output:
	// [4 1 2 3]
}

func ExampleDeque_PushRight() {
	deque := NewDeque[int]()
	deque.ExtendRight([]int{1, 2, 3})
	deque.PushRight(4)
	fmt.Println(deque)
	// Output:
	// [1 2 3 4]
}

func ExampleDeque_Reverse() {
	deque := NewDeque[int]()
	deque.ExtendRight([]int{1, 2, 3})
	fmt.Println(deque)
	deque.Reverse()
	fmt.Println(deque)
	// Output:
	// [1 2 3]
	// [3 2 1]
}

func ExampleDeque_Rotate() {
	deque := NewDeque[int]()
	deque.ExtendRight([]int{1, 2, 3})
	fmt.Println(deque)

	deque.Rotate(2)
	fmt.Println(deque)

	deque.Rotate(-2)
	fmt.Println(deque)
	// Output:
	// [1 2 3]
	// [3 1 2]
	// [1 2 3]
}
