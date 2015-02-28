package finalseg

import (
	"fmt"
	"sort"
)

const MinFloat = -3.14e100

var (
	prevStatus = make(map[byte][]byte)
	probStart  = make(map[byte]float64)
)

func init() {
	prevStatus['B'] = []byte{'E', 'S'}
	prevStatus['M'] = []byte{'M', 'B'}
	prevStatus['S'] = []byte{'S', 'E'}
	prevStatus['E'] = []byte{'B', 'M'}
	probStart['B'] = -0.26268660809250016
	probStart['E'] = -3.14e+100
	probStart['M'] = -3.14e+100
	probStart['S'] = -1.4652633398537678
}

type Viterbi struct {
	prob  float64
	state byte
}

func (v Viterbi) String() string {
	return fmt.Sprintf("(%f, %x)", v.prob, v.state)
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
		if val, ok := probEmit[y][obs[0]]; ok {
			V[0][y] = val + probStart[y]
		} else {
			V[0][y] = MinFloat + probStart[y]
		}
		path[y] = []byte{y}
	}

	for t := 1; t < len(obs); t++ {
		newPath := make(map[byte][]byte)
		V[t] = make(map[byte]float64)
		for _, y := range states {
			vs0 := make(Viterbis, 0)
			var em_p float64
			if val, ok := probEmit[y][obs[t]]; ok {
				em_p = val
			} else {
				em_p = MinFloat
			}
			for _, y0 := range prevStatus[y] {
				var transP float64
				if tp, ok := probTrans[y0][y]; ok {
					transP = tp
				} else {
					transP = MinFloat
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
