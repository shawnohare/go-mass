package rel

// An Model interface a real-valued statistical model.
type Model interface {
	Eval(interface{}) float64
}

// Map a real-valued statistical model over a list of input data.
func MapPairs(m Model, inputs []interface{}) Slice {
	outputs := make(Slice, len(inputs))
	for i, x := range inputs {
		outputs.Set(i, x, m.Eval(x))
	}
	return outputs
}

func MapEval(m Model, inputs []interface{}) []float64 {
	outputs := make([]float64, len(inputs))
	for i, x := range inputs {
		outputs[i] = m.Eval(x)
	}
	return outputs
}

//
// func MapApply(m Model, inputs []interface{}) []float64 {
// 	outputs := make(Slice, len(inputs))
// 	for i, x := range inputs {
// 		outputs[i] = m.Apply(x)
// 	}
// 	return outputs
// }
