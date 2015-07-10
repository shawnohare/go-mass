package rel

// A Slice is a concrete implementation of a List interface
// that still affords standard slice mechanics.
type Slice []*Pair

// Implement a List and sort interface.
func (s Slice) Get(i int) interface{}                 { return s[i].Element }
func (s Slice) Set(i int, e interface{}, val float64) { s[i] = &Pair{e, val} }
func (s Slice) Val(i int) float64                     { return s[i].Value }
func (s Slice) Len() int                              { return len(s) }
func (s Slice) Less(i, j int) bool                    { return s.Val(i) < s.Val(j) }
func (s Slice) Swap(i, j int)                         { s[i], s[j] = s[j], s[i] }

func MakeSlice(l List) Slice {
	// Create a new Slice struct that represents the input List.
	s := make(Slice, l.Len())

	// Populate the list struct with the input List's data.
	for i := range s {
		s[i] = &Pair{
			Element: l.Get(i),
			Value:   l.Val(i),
		}
	}

	return s
}
