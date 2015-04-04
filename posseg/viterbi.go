package posseg

import (
	"fmt"
	"sort"
)

type probState struct {
	prob  float64
	state string
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

func viterbi(obs []rune) []string {
	obsLength := len(obs)
	V := make([]map[string]float64, obsLength)
	V[0] = make(map[string]float64)
	mem_path := make([]map[string]string, obsLength)
	mem_path[0] = make(map[string]string)
	ys := charStateTab.get(obs[0]) // default is all_states
	for _, y := range ys {
		V[0][y] = probEmit[y].get(obs[0]) + probStart[y]
		mem_path[0][y] = ""
	}
	for t := 1; t < obsLength; t++ {
		prev_states := make([]string, 0)
		for x := range mem_path[t-1] {
			if len(probTrans[x]) > 0 {
				prev_states = append(prev_states, x)
			}
		}
		//use Go's map to implement Python's Set()
		prev_states_expect_next := make(map[string]int)
		for _, x := range prev_states {
			for y := range probTrans[x] {
				prev_states_expect_next[y] = 1
			}
		}
		tmp_obs_states := charStateTab.get(obs[t])

		obs_states := make([]string, 0)
		for index := range tmp_obs_states {
			if _, ok := prev_states_expect_next[tmp_obs_states[index]]; ok {
				obs_states = append(obs_states, tmp_obs_states[index])
			}
		}
		if len(obs_states) == 0 {
			for key := range prev_states_expect_next {
				obs_states = append(obs_states, key)
			}
		}
		if len(obs_states) == 0 {
			obs_states = probTransKeys
		}
		mem_path[t] = make(map[string]string) // TODO: value needed or not?
		V[t] = make(map[string]float64)
		for _, y := range obs_states {
			var max, ps probState
			for i, y0 := range prev_states {
				ps = probState{
					prob:  V[t-1][y0] + probTrans[y0].Get(y) + probEmit[y].get(obs[t]),
					state: y0}
				if i == 0 || ps.prob > max.prob || (ps.prob == max.prob && ps.state > max.state) {
					max = ps
				}
			}
			V[t][y] = max.prob
			mem_path[t][y] = max.state
		}
	}
	last := make(probStates, 0)
	length := len(mem_path)
	vlength := len(V)
	for y := range mem_path[length-1] {
		ps := probState{prob: V[vlength-1][y], state: y}
		last = append(last, ps)
	}
	sort.Sort(sort.Reverse(last))
	state := last[0].state
	route := make([]string, len(obs))

	for i := obsLength - 1; i >= 0; i-- {
		route[i] = state
		state = mem_path[i][state]
	}
	return route
}
