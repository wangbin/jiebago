package posseg

import (
	"fmt"
	"sort"
)

type probState struct {
	prob  float64
	state uint16
}

func (ps probState) String() string {
	return fmt.Sprintf("(%v: %f)", ps.state, ps.prob)
}

type probStates []probState

func (pss probStates) Len() int {
	return len(pss)
}

func (pss probStates) Less(i, j int) bool {
	if pss[i].prob == pss[j].prob {
		return pss[i].state < pss[j].state
	}
	return pss[i].prob < pss[j].prob
}

func (pss probStates) Swap(i, j int) {
	pss[i], pss[j] = pss[j], pss[i]
}

func viterbi(obs []rune) []tag {
	obsLength := len(obs)
	V := make([]map[uint16]float64, obsLength)
	V[0] = make(map[uint16]float64)
	memPath := make([]map[uint16]uint16, obsLength)
	memPath[0] = make(map[uint16]uint16)
	ys := charStateTab.get(obs[0]) // default is all_states
	for _, y := range ys {
		V[0][y] = probEmit[y].get(obs[0]) + probStart[y]
		memPath[0][y] = 0
	}
	for t := 1; t < obsLength; t++ {
		var prevStates []uint16
		for x := range memPath[t-1] {
			if len(probTrans[x]) > 0 {
				prevStates = append(prevStates, x)
			}
		}
		//use Go's map to implement Python's Set()
		prevStatesExpectNext := make(map[uint16]int)
		for _, x := range prevStates {
			for y := range probTrans[x] {
				prevStatesExpectNext[y] = 1
			}
		}
		tmpObsStates := charStateTab.get(obs[t])

		var obsStates []uint16
		for index := range tmpObsStates {
			if _, ok := prevStatesExpectNext[tmpObsStates[index]]; ok {
				obsStates = append(obsStates, tmpObsStates[index])
			}
		}
		if len(obsStates) == 0 {
			for key := range prevStatesExpectNext {
				obsStates = append(obsStates, key)
			}
		}
		if len(obsStates) == 0 {
			obsStates = probTransKeys
		}
		memPath[t] = make(map[uint16]uint16)
		V[t] = make(map[uint16]float64)
		for _, y := range obsStates {
			var max, ps probState
			for i, y0 := range prevStates {
				ps = probState{
					prob:  V[t-1][y0] + probTrans[y0].Get(y) + probEmit[y].get(obs[t]),
					state: y0}
				if i == 0 || ps.prob > max.prob || (ps.prob == max.prob && ps.state > max.state) {
					max = ps
				}
			}
			V[t][y] = max.prob
			memPath[t][y] = max.state
		}
	}
	last := make(probStates, 0)
	length := len(memPath)
	vlength := len(V)
	for y := range memPath[length-1] {
		ps := probState{prob: V[vlength-1][y], state: y}
		last = append(last, ps)
	}
	sort.Sort(sort.Reverse(last))
	state := last[0].state
	route := make([]tag, len(obs))

	for i := obsLength - 1; i >= 0; i-- {
		route[i] = tag(state)
		state = memPath[i][state]
	}
	return route
}
