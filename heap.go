package rel

// Heap pointers implement the standard heap interface.
type Heap []*Pair

func (h *Heap) Set(i int, e interface{}, val float64) { (*h)[i] = NewPair(e, val) }

func (h *Heap) SetPair(i int, p *Pair) { (*h)[i] = p }
func (h *Heap) Get(i int) interface{}  { return (*h)[i].Element }
func (h *Heap) Val(i int) float64      { return (*h)[i].Value }
func (h *Heap) Len() int               { return len(*h) }
func (h *Heap) Less(i, j int) bool     { return h.Val(i) < h.Val(j) }
func (h *Heap) Swap(i, j int)          { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *Heap) Push(x interface{})     { *h = append(*h, x.(*Pair)) }

func (h *Heap) Pop() interface{} {
	old := *h
	n := old.Len()
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// NewHeap creates a pointer to a Heap struct, which satisfies
// the heap, sort, and List interfaces.
func NewHeap() *Heap {
	return new(Heap)
}
