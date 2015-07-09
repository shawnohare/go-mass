package mass

// A Slice is a concrete implementation of a Collection interface
// that still affords standard slice mechanics.
type Slice []*Pair

// Implement a mass.Collection and sort interface.
func (s Slice) Get(i int) interface{}                  { return s[i].Element }
func (s Slice) Set(i int, e interface{}, mass float64) { s[i] = &Pair{e, mass} }
func (s Slice) Mass(i int) float64                     { return s[i].Mass }
func (s Slice) Len() int                               { return len(s) }
func (s Slice) Less(i, j int) bool                     { return s.Mass(i) < s.Mass(j) }
func (s Slice) Swap(i, j int)                          { s[i], s[j] = s[j], s[i] }

func MakeSlice(col Collection) Slice {
	// Create a new Slice struct that represents the input Collection.
	s := make(Slice, col.Len())

	// Populate the collection struct with the input Collection's data.
	for i := range s {
		s[i] = &Pair{
			Element: col.Get(i),
			Mass:    col.Mass(i),
		}
	}

	return s
}
