package helpers

import (
	"fmt"
)

type Deque[T any] struct {
	store           []T
	head            int
	tail            int
	length          int
	initialCapacity int
}

func NewDeque[T any](capacity ...int) *Deque[T] {
	cap := 1
	if len(capacity) >= 1 {
		if capacity[0] > 1 {
			cap = capacity[0]
		}
	}
	return &Deque[T]{
		store:           make([]T, cap),
		initialCapacity: cap,
	}
}

func (d Deque[T]) Capacity() int {
	return len(d.store)
}

func (d *(Deque[T])) Clear() {
	d.store = make([]T, d.initialCapacity)
	d.head, d.tail, d.length = 0, 0, 0
}

func (d Deque[T]) Copy() *Deque[T] {
	c := NewDeque[T](len(d.store))
	c.ExtendRight(d.Elements())
	return c
}

func (d Deque[T]) decrement(n int) int {
	return (n - 1 + len(d.store)) % len(d.store)
}

func (d *(Deque[T])) growIfFull() {
	capacity := len(d.store)
	if d.length == capacity {
		d.resize(capacity * 2)
	}
}

func (d Deque[T]) Elements() []T {
	elements := make([]T, d.length)
	for i := 0; i < d.length; i++ {
		elements[i] = d.store[(d.head+i)%len(d.store)]
	}
	return elements
}

func (d *(Deque[T])) growIfExtendWouldMakeFull(size int) {
	capacity := len(d.store)
	for capacity < d.length+size {
		capacity *= 2
	}
	if capacity > len(d.store) {
		d.resize(capacity)
	}
}

func (d *(Deque[T])) ExtendLeft(arr []T) {
	d.growIfExtendWouldMakeFull(len(arr))
	for _, v := range arr {
		d.PushLeft(v)
	}
}

func (d *(Deque[T])) ExtendRight(arr []T) {
	d.growIfExtendWouldMakeFull(len(arr))
	for _, v := range arr {
		d.PushRight(v)
	}
}

func (d Deque[T]) increment(n int) int {
	return (n + 1) % len(d.store)
}

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

func (d *(Deque[T])) Len() int {
	return d.length
}

func (d *(Deque[T])) PeekLeft() T {
	return d.store[d.head%len(d.store)]
}

func (d *(Deque[T])) PopLeft() T {
	value := d.store[d.head%len(d.store)]
	var dummy T
	d.store[d.head%len(d.store)] = dummy
	d.head = d.increment(d.head)
	d.length--
	d.shrinkIfSparse()
	return value
}

func (d *(Deque[T])) PeekRight() T {
	return d.store[d.decrement(d.tail)]
}

func (d *(Deque[T])) PopRight() T {
	newTail := d.decrement(d.tail)
	value := d.store[newTail]
	var dummy T
	d.store[newTail] = dummy
	d.tail = newTail
	d.length--
	d.shrinkIfSparse()
	return value
}

func (d *(Deque[T])) PushRight(value T) {
	d.growIfFull()
	d.store[d.tail%len(d.store)] = value
	d.tail = d.increment(d.tail)
	d.length++
}

func (d *(Deque[T])) PushLeft(value T) {
	d.growIfFull()
	d.head = d.decrement(d.head)
	d.store[d.head] = value
	d.length++
}

func (d *(Deque[T])) resize(size int) {
	newStore := make([]T, size)
	for i := 0; i <= d.length; i++ {
		newStore[i] = d.store[(d.head+i)%len(d.store)]
	}
	d.store = newStore
	d.head = 0
	d.tail = d.length
}

func (d *(Deque[T])) shrinkIfSparse() {
	capacity := len(d.store)
	if d.length <= capacity/4 {
		d.resize(capacity / 2)
	}
}

func (d *(Deque[T])) Reverse() {
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

func (d *(Deque[T])) Rotate(n int) {
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
