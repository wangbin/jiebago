package posseg

import (
	"fmt"
	"sort"
)

const MIN_FLOAT = -3.14e100

type StateTag struct {
	State byte
	Tag   string
}

func (st StateTag) String() string {
	return fmt.Sprintf("(%q, %s)", st.State, st.Tag)
}

func emptyStateTag() StateTag {
	return StateTag{' ', ""}
}

type ProbState struct {
	Prob float64
	ST   StateTag
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
		if pss[i].ST.Tag < pss[j].ST.Tag {
			return true
		} else if pss[i].ST.State < pss[j].ST.State {
			return true
		} else {
			return false
		}
	}
	return pss[i].Prob < pss[j].Prob
}

func (pss ProbStates) Swap(i, j int) {
	pss[i], pss[j] = pss[j], pss[i]
}

func Viterbi(obs []rune) (float64, []StateTag) {
	obsLength := len(obs)
	V := make([]map[StateTag]float64, obsLength)
	V[0] = make(map[StateTag]float64)
	mem_path := make([]map[StateTag]StateTag, obsLength)
	mem_path[0] = make(map[StateTag]StateTag)
	// all_states := ProbTransKeys
	ys := CharStateTab.Get(obs[0]) // default is all_states
	for _, y := range ys {
		V[0][y] = ProbEmit[y].Get(obs[0]) + ProbStart[y]
		mem_path[0][y] = emptyStateTag()
	}
	for t := 1; t < obsLength; t++ {
		prev_states := make([]StateTag, 0)
		for x, _ := range mem_path[t-1] {
			if len(ProbTrans[x]) > 0 {
				prev_states = append(prev_states, x)
			}
		}
		//use Go's map to implement Python's Set()
		prev_states_expect_next := make(map[StateTag]StateTag)
		for _, x := range prev_states {
			for y, _ := range ProbTrans[x] {
				prev_states_expect_next[y] = y
			}
		}
		tmp_obs_states := CharStateTab.Get(obs[t])

		obs_states := make([]StateTag, 0)
		for index, _ := range tmp_obs_states {
			if _, ok := prev_states_expect_next[tmp_obs_states[index]]; ok {
				obs_states = append(obs_states, tmp_obs_states[index])
			}
		}
		if len(obs_states) == 0 {
			obs_states = ProbTransKeys
		}
		mem_path[t] = make(map[StateTag]StateTag)
		V[t] = make(map[StateTag]float64)
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
	route := make([]StateTag, len(obs))
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
