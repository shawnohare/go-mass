package mass

// An Model interface a real-valued statistical model.
type StatModel interface {
	Eval(interface{}) float64
}

// Map a real-valued statistical model over a collection of input data.
func MapPairs(m StatModel, inputs []interface{}) Slice {
	outputs := make(Slice, len(inputs))
	for i, x := range inputs {
		outputs.Set(i, x, m.Eval(x))
	}
	return outputs
}

func Map(m StatModel, inputs []interface{}) []float64 {
	outputs := make([]float64, len(inputs))
	for i, x := range inputs {
		outputs[i] = m.Eval(x)
	}
	return outputs
}
