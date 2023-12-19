package godino

import (
	"fmt"
)

// A combination of a stack and queue or "double-ended queue" with fast and memory efficient pushes and pops for both ends.
type Deque[T any] struct {
	store       []T
	head        int
	tail        int
	length      int
	minCapacity int
}

// Returns a deque the specified minimum capacity (defaults to 1)
func NewDeque[T any](capacity ...int) *Deque[T] {
	cap := 1
	if len(capacity) >= 1 {
		if capacity[0] > 1 {
			cap = capacity[0]
		}
	}
	return &Deque[T]{
		store:       make([]T, cap),
		minCapacity: cap,
	}
}

// Returns the current capacity of the deque. This will resize automatically as elements are added or removed
func (d Deque[T]) Capacity() int {
	return len(d.store)
}

// Removes all elements from the deque
func (d *Deque[T]) Clear() {
	d.store = make([]T, d.minCapacity)
	d.head, d.tail, d.length = 0, 0, 0
}

// Returns a copy of the deque
func (d Deque[T]) Copy() *Deque[T] {
	c := NewDeque[T](d.minCapacity)
	c.ExtendRight(d.Elements())
	return c
}

func (d Deque[T]) decrement(n int) int {
	return (n - 1 + len(d.store)) % len(d.store)
}

func (d *Deque[T]) growIfFull() {
	capacity := len(d.store)
	if d.length >= capacity {
		d.resize(capacity * 2)
	}
}

// Returns an array containing all elements from the deque
func (d Deque[T]) Elements() []T {
	elements := make([]T, d.length)
	for i := 0; i < d.length; i++ {
		elements[i] = d.store[(d.head+i)%len(d.store)]
	}
	return elements
}

func (d *Deque[T]) growIfExtendWouldMakeFull(size int) {
	capacity := len(d.store)
	for capacity < d.length+size {
		capacity *= 2
	}
	if capacity > len(d.store) {
		d.resize(capacity)
	}
}

// Adds the given elements to the left side of the deque.
// The series of left pushes results in reversing the order of given elements.
func (d *Deque[T]) ExtendLeft(arr []T) {
	d.growIfExtendWouldMakeFull(len(arr))
	for _, v := range arr {
		d.PushLeft(v)
	}
}

// Adds the given elements to the right side of the deque
func (d *Deque[T]) ExtendRight(arr []T) {
	d.growIfExtendWouldMakeFull(len(arr))
	for _, v := range arr {
		d.PushRight(v)
	}
}

func (d Deque[T]) increment(n int) int {
	return (n + 1) % len(d.store)
}

// Returns a slice of the queue based on the given indices.
// If one index is provided, an slice containing the single item at the index is returned.
// If two indices are provided, a slice containging the elements between the two indicies is returned.
func (d Deque[T]) Index(nums ...int) []T {
	elements := d.Elements()
	if len(nums) == 1 {
		return []T{elements[nums[0]]}
	}
	if len(nums) >= 2 {
		return elements[nums[0]:nums[1]]
	}
	return elements
}

// Returns the number of elements in the deque
func (d *Deque[T]) Len() int {
	return d.length
}

// Returns the first element (leftmost) in the deque.
// Panics if called on an empty deque.
func (d *Deque[T]) PeekLeft() T {
	if d.length == 0 {
		panic("PeekLeft() called on an empty deque")
	}
	return d.store[d.head%len(d.store)]
}

// Removes the first element (leftmost) from the deque and returns it.
// Panics if called on an empty deque.
func (d *Deque[T]) PopLeft() T {
	if d.length == 0 {
		panic("PopLeft() called on an empty deque")
	}
	value := d.store[d.head%len(d.store)]
	var dummy T
	d.store[d.head%len(d.store)] = dummy
	d.head = d.increment(d.head)
	d.length--
	d.shrinkIfSparse()
	return value
}

// Returns the last element (rightmost) of the deque.
// Panics if called on an empty deque.
func (d *Deque[T]) PeekRight() T {
	if d.length == 0 {
		panic("PeekRight() called on an empty deque")
	}
	return d.store[d.decrement(d.tail)]
}

// Removes the last element (rightmost) from the deque and returns it.
// Panics if called on an empty deque.
func (d *Deque[T]) PopRight() T {
	if d.length == 0 {
		panic("PopRight() called on an empty deque")
	}
	newTail := d.decrement(d.tail)
	value := d.store[newTail]
	var dummy T
	d.store[newTail] = dummy
	d.tail = newTail
	d.length--
	d.shrinkIfSparse()
	return value
}

// Adds an element to the end of the deque
func (d *Deque[T]) PushRight(value T) {
	d.growIfFull()
	d.store[d.tail%len(d.store)] = value
	d.tail = d.increment(d.tail)
	d.length++
}

// Adds an element to the beginning of the deque
func (d *Deque[T]) PushLeft(value T) {
	d.growIfFull()
	d.head = d.decrement(d.head)
	d.store[d.head] = value
	d.length++
}

func (d *Deque[T]) resize(size int) {
	newStore := make([]T, size)
	for i := 0; i <= d.length; i++ {
		newStore[i] = d.store[(d.head+i)%len(d.store)]
	}
	d.store = newStore
	d.head = 0
	d.tail = d.length
}

func (d *Deque[T]) shrinkIfSparse() {
	capacity := len(d.store)
	if capacity <= d.minCapacity {
		return
	}
	if d.length <= capacity/4 {
		size := capacity / 2
		if size < d.minCapacity {
			size = d.minCapacity
		}
		d.resize(size)
	}
}

// Reverses the order of the elements in the deque
func (d *Deque[T]) Reverse() {
	if d.length == 0 {
		return
	}
	elements := make([]T, len(d.store))
	for i := 0; i < d.length; i++ {
		elements[i] = d.store[(d.tail-1-i+len(d.store))%len(d.store)]
	}
	d.store = elements
	d.head = 0
	d.tail = d.length
}

func (d Deque[T]) String() string {
	return fmt.Sprintf("%v", d.Elements())
}

// Rotate the deque n elements to the right.
// If n is negative, the deque is rotated to the left.
func (d *Deque[T]) Rotate(n int) {
	if d.length == 0 {
		return
	}
	if n < 0 {
		n = 0 - n
		for i := 0; i < n%d.length; i++ {
			d.PushLeft(d.PopRight())
		}
	} else {
		for i := 0; i < n%d.length; i++ {
			d.PushRight(d.PopLeft())
		}
	}
}
