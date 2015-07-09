package mass

// An Interface reprsenting a real-valued statistical model.
type Model interface {
	Eval(interface{}) float64
}

// Map a real-valued statistical model over a collection of input data.
func Map(m Model, inputs []interface{}) Slice {
	outputs := make(Slice, len(inputs))
	for i, x := range inputs {
		outputs.Set(i, x, m.Eval(x))
	}
	return outputs
}
