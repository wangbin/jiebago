package posseg

import (
	"fmt"
	"sort"
)

type stateTag struct {
	State byte
	Tag   string
}

func (st stateTag) String() string {
	return fmt.Sprintf("(%q, %s)", st.State, st.Tag)
}

func emptyStateTag() stateTag {
	return stateTag{' ', ""}
}

type ProbState struct {
	Prob float64
	ST   stateTag
}

func (ps ProbState) String() string {
	return fmt.Sprintf("(%v: %f)", ps.ST, ps.Prob)
}

type ProbStates []ProbState

func (pss ProbStates) Len() int {
	return len(pss)
}

func (pss ProbStates) Less(i, j int) bool {
	if pss[i].Prob == pss[j].Prob {
		if pss[i].ST.State == pss[j].ST.State {
			return pss[i].ST.Tag < pss[j].ST.Tag
		}
		return pss[i].ST.State < pss[j].ST.State
	}
	return pss[i].Prob < pss[j].Prob
}

func (pss ProbStates) Swap(i, j int) {
	pss[i], pss[j] = pss[j], pss[i]
}

func viterbi(obs []rune) (float64, []stateTag) {
	obsLength := len(obs)
	V := make([]map[stateTag]float64, obsLength)
	V[0] = make(map[stateTag]float64)
	mem_path := make([]map[stateTag]stateTag, obsLength)
	mem_path[0] = make(map[stateTag]stateTag)
	// all_states := ProbTransKeys
	ys := charStateTab.get(obs[0]) // default is all_states
	for _, y := range ys {
		V[0][y] = ProbEmit[y].Get(obs[0]) + ProbStart[y]
		mem_path[0][y] = emptyStateTag()
	}
	for t := 1; t < obsLength; t++ {
		prev_states := make([]stateTag, 0)
		for x, _ := range mem_path[t-1] {
			if len(ProbTrans[x]) > 0 {
				prev_states = append(prev_states, x)
			}
		}
		//use Go's map to implement Python's Set()
		prev_states_expect_next := make(map[stateTag]stateTag)
		for _, x := range prev_states {
			for y, _ := range ProbTrans[x] {
				prev_states_expect_next[y] = y
			}
		}
		tmp_obs_states := charStateTab.get(obs[t])

		obs_states := make([]stateTag, 0)
		for index, _ := range tmp_obs_states {
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
			obs_states = ProbTransKeys
		}
		mem_path[t] = make(map[stateTag]stateTag) // TODO: value needed or not?
		V[t] = make(map[stateTag]float64)
		for _, y := range obs_states {
			pss := make(ProbStates, 0)
			for _, y0 := range prev_states {
				ps := ProbState{
					Prob: V[t-1][y0] + ProbTrans[y0].Get(y) + ProbEmit[y].Get(obs[t]),
					ST:   y0}
				pss = append(pss, ps)
			}
			sort.Sort(sort.Reverse(pss))
			V[t][y] = pss[0].Prob
			mem_path[t][y] = pss[0].ST
		}
	}
	last := make(ProbStates, 0)
	length := len(mem_path)
	vlength := len(V)
	for y, _ := range mem_path[length-1] {
		ps := ProbState{Prob: V[vlength-1][y], ST: y}
		last = append(last, ps)
	}
	sort.Sort(sort.Reverse(last))
	prob := last[0].Prob
	state := last[0].ST
	route := make([]stateTag, len(obs))
	i := obsLength - 1
	for {
		if i < 0 {
			break
		}
		route[i] = state
		state = mem_path[i][state]
		i -= 1
	}
	return prob, route
}
