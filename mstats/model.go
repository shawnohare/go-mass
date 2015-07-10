package mstats

import "github.com/shawnohare/go-mass"

// An Model interface a real-valued statistical model.
type Model interface {
	Eval(interface{}) float64
}

// Map a real-valued statistical model over a collection of input data.
func MapPairs(m Model, inputs []interface{}) mass.Slice {
	outputs := make(mass.Slice, len(inputs))
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
// 	outputs := make(mass.Slice, len(inputs))
// 	for i, x := range inputs {
// 		outputs[i] = m.Apply(x)
// 	}
// 	return outputs
// }
