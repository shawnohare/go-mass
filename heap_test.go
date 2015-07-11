package rel

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	h := new(Heap)
	pair := NewPair(nil, 1)
	// Test Push/Pop methods.
	h.Push(pair)
	assert.Equal(t, 1, h.Len())
	h.Pop()
	assert.Equal(t, 0, h.Len())
	// Test that h satisfies the heap interface.
	heap.Push(h, pair)
	assert.Equal(t, 1, h.Len())
}
