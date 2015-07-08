package partition

type Masses []float64

func (m Masses) Len() int {
	return len(m)
}

func (m Masses) Mass(i int) float64 {
	return m[i]
}

func TestEqualWidthBins(t *testing.T) {

}
