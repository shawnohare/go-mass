package rel

import (
	"sort"
)

func Sort(c List) Slice {
	s := MakeSlice(c)
	sort.Sort(s)
	return s
}
