package godino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeque(t *testing.T) {
	deque := NewDeque[string]()

	// Test PushLeft and PushRight
	deque.PushLeft("Apple")
	deque.PushRight("Banana")
	deque.PushLeft("Cherry")
	deque.PushRight("Durian")
	assert.Equal(t, []string{"Cherry", "Apple", "Banana", "Durian"}, deque.Elements())

	// Test PopLeft and PopRight
	item := deque.PopLeft()
	assert.Equal(t, "Cherry", item)
	item = deque.PopRight()
	assert.Equal(t, "Durian", item)

	// Test Len
	assert.Equal(t, 2, deque.Len())

	// Test ExtendLeft
	deque.ExtendLeft([]string{"Grape", "Kiwi"})
	assert.Equal(t, []string{"Kiwi", "Grape", "Apple", "Banana"}, deque.Elements())

	// Test ExtendRight
	deque.ExtendRight([]string{"Lemon", "Mango"})
	assert.Equal(t, []string{"Kiwi", "Grape", "Apple", "Banana", "Lemon", "Mango"}, deque.Elements())

	// Test Copy
	dequeCopy := deque.Copy()
	assert.Equal(t, deque.Elements(), dequeCopy.Elements())

	// Test Clear
	deque.Clear()
	assert.Equal(t, []string{}, deque.Elements())
	assert.Equal(t, 0, deque.Len())
}

func TestRotate(t *testing.T) {
	// Create a new deque
	deque := NewDeque[string]()

	// Push elements
	deque.PushRight("Apple")
	deque.PushRight("Banana")
	deque.PushRight("Cherry")
	deque.PushRight("Durian")

	// Rotate with positive value
	deque.Rotate(2)
	assert.Equal(t, []string{"Cherry", "Durian", "Apple", "Banana"}, deque.Elements())

	// Rotate with negative value
	deque.Rotate(-3)
	assert.Equal(t, []string{"Durian", "Apple", "Banana", "Cherry"}, deque.Elements())
}

func TestReverseDeque(t *testing.T) {
	// Create a new deque
	deque := NewDeque[string]()

	// Push elements
	deque.PushRight("Apple")
	deque.PushRight("Banana")
	deque.PushRight("Cherry")
	deque.PushRight("Durian")

	// Reverse the deque
	deque.Reverse()
	assert.Equal(t, []string{"Durian", "Cherry", "Banana", "Apple"}, deque.Elements())

	// Reverse an empty deque
	emptyDeque := NewDeque[string]()
	emptyDeque.Reverse()
	assert.Equal(t, []string{}, emptyDeque.Elements())
}

func TestDequeIndex(t *testing.T) {
	// Create a new deque
	deque := NewDeque[string]()

	// Push elements
	deque.PushRight("Apple")
	deque.PushRight("Banana")
	deque.PushRight("Cherry")
	deque.PushRight("Durian")

	// Test indexing
	assert.Equal(t, []string{"Banana"}, deque.Index(1))
	assert.Equal(t, []string{"Banana", "Cherry"}, deque.Index(1, 3))
	assert.Equal(t, []string{"Apple", "Banana", "Cherry", "Durian"}, deque.Index(0, 4))
}

func TestResize(t *testing.T) {
	deque := NewDeque[int](4)
	assert.Equal(t, 4, deque.Capacity())
	for i := 0; i < 5; i++ {
		deque.PushRight(i)
	}
	assert.Equal(t, 8, deque.Capacity())
	for i := 0; i < 3; i++ {
		deque.PopLeft()
	}
	assert.Equal(t, 4, deque.Capacity())
}
