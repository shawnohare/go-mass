package mass

type Pair struct {
	Object interface{}
	Mass   float64
}

// A List is a concrete implementation of a Collection interface.
type List struct {
	Pairs []*Pair
}

// Implement a mass.Collection and sort interface.
func (c *List) Get(i int) interface{} { return c.Pairs[i].Object }
func (c *List) Mass(i int) float64    { return c.Pairs[i].Mass }
func (c *List) Len() int              { return len(c.Pairs) }
func (c *List) Less(i, j int) bool    { return c.Mass(i) < c.Mass(j) }
func (c *List) Swap(i, j int)         { c.Pairs[i], c.Pairs[j] = c.Pairs[j], c.Pairs[i] }

func NewList(col Collection) *List {
	// Create a new List struct that represents the input Collection.
	c := &List{
		Pairs: make([]*Pair, col.Len()),
	}

	// Populate the collection struct with the input Collection's data.
	for i := range c.pairs {
		c.Pairs[i] = &Pair{
			Object: col.Get(i),
			Mass:   col.Mass(i),
		}
	}

	return c
}
