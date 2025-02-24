package learn

import (
	"github.com/gaillard/go-neural"
)

type Deltas [][]float64

type Sample struct {
	In    []float64
	Ideal []float64
}

func Learn(n *neural.Network, in, ideal []float64, speed float64) {
	Backpropagation(n, in, ideal, speed)
}

// Note the calc returned is before the training update.
func Backpropagation(n *neural.Network, in, ideal []float64, speed float64) float64 {
	out := n.Calculate(in)

	deltas := make([][]float64, len(n.Layers))

	last := len(n.Layers) - 1
	l := n.Layers[last]
	deltas[last] = make([]float64, len(l.Neurons))
	for i, n := range l.Neurons {
		deltas[last][i] = n.Out * (1 - n.Out) * (ideal[i] - n.Out)
	}

	for i := last - 1; i >= 0; i-- {
		l := n.Layers[i]
		deltas[i] = make([]float64, len(l.Neurons))
		di := deltas[i+1]
		for j, n := range l.Neurons {

			nn := n.Out * (1 - n.Out)

			var sum float64 = 0
			for k, s := range n.OutSynapses {
				sum += s.Weight * di[k]
			}

			deltas[i][j] = nn * sum
		}
	}

	for i, l := range n.Layers {
		di := deltas[i]
		for j, n := range l.Neurons {
			dj := di[j] * speed
			for _, s := range n.InSynapses {
				s.Weight += dj * s.In
			}
		}
	}

	return leastSquares(out, ideal)
}
