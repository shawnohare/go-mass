package mass

import (
	"sort"
)

func Sort(c Collection) Slice {
	s := MakeSlice(c)
	sort.Sort(s)
	return s
}
