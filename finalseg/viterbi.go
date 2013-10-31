package finalseg

import (
	"fmt"
	"sort"
)

const MIN_FLOAT = -3.14e100

var PrevStatus = make(map[byte][]byte)

func init() {
	PrevStatus['B'] = []byte{'E', 'S'}
	PrevStatus['M'] = []byte{'M', 'B'}
	PrevStatus['S'] = []byte{'S', 'E'}
	PrevStatus['E'] = []byte{'B', 'M'}
}

type Viterbi struct {
	prob  float64
	state byte
}

func (v Viterbi) String() string {
	return fmt.Sprintf("(%f, %s)", v.prob, v.state)
}

type Viterbis []*Viterbi

func (vs Viterbis) Len() int {
	return len(vs)
}

func (vs Viterbis) Less(i, j int) bool {
	if vs[i].prob == vs[j].prob {
		return vs[i].state < vs[j].state
	}
	return vs[i].prob < vs[j].prob
}

func (vs Viterbis) Swap(i, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}

func viterbi(obs []rune, states []byte) (float64, []byte) {
	path := make(map[byte][]byte)
	V := make([]map[byte]float64, len(obs))
	V[0] = make(map[byte]float64)
	for _, y := range states {
		if val, ok := ProbEmit[y][obs[0]]; ok {
			V[0][y] = val + ProbStart[y]
		} else {
			V[0][y] = MIN_FLOAT + ProbStart[y]
		}
		path[y] = []byte{y}
	}

	for t := 1; t < len(obs); t++ {
		newPath := make(map[byte][]byte)
		V[t] = make(map[byte]float64)
		for _, y := range states {
			vs0 := make(Viterbis, 0)
			var em_p float64
			if val, ok := ProbEmit[y][obs[t]]; ok {
				em_p = val
			} else {
				em_p = MIN_FLOAT
			}
			for _, y0 := range PrevStatus[y] {
				var transP float64
				if tp, ok := ProbTrans[y0][y]; ok {
					transP = tp
				} else {
					transP = MIN_FLOAT
				}
				prob0 := V[t-1][y0] + transP + em_p
				vs0 = append(vs0, &Viterbi{prob: prob0, state: y0})
			}
			sort.Sort(sort.Reverse(vs0))
			V[t][y] = vs0[0].prob
			pp := make([]byte, len(path[vs0[0].state]))
			copy(pp, path[vs0[0].state])
			newPath[y] = append(pp, y)
		}
		path = newPath
	}
	vs := make(Viterbis, 0)
	for _, y := range []byte{'E', 'S'} {
		vs = append(vs, &Viterbi{V[len(obs)-1][y], y})
	}
	sort.Sort(sort.Reverse(vs))
	v := vs[0]
	return v.prob, path[v.state]
}
