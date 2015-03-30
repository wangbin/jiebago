package jiebago

import (
	"math"
	"sort"
)

type route struct {
	Freq  float64
	Index int
}

type routes []route

func (rs routes) Len() int {
	return len(rs)
}

func (rs routes) Less(i, j int) bool {
	if rs[i].Freq < rs[j].Freq {
		return true
	}
	if rs[i].Freq == rs[j].Freq {
		return rs[i].Index < rs[j].Index
	}
	return false
}

func (rs routes) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

type dag map[int][]int

func DAG(s Segmenter, runes []rune) dag {
	d := make(dag)
	n := len(runes)
	var frag string
	for k := 0; k < n; k++ {
		tmpList := make([]int, 0)
		i := k
		frag = string(runes[k])
		for {
			if freq, ok := s.Freq(frag); !ok {
				break
			} else {
				if freq > 0.0 {
					tmpList = append(tmpList, i)
				}
			}
			i += 1
			if i >= n {
				break
			}
			frag = string(runes[k : i+1])
		}
		if len(tmpList) == 0 {
			tmpList = append(tmpList, k)
		}
		d[k] = tmpList
	}
	return d
}

func Routes(s Segmenter, runes []rune, d dag) map[int]route {
	n := len(runes)
	rs := make(map[int]route)
	rs[n] = route{Freq: 0.0, Index: 0}
	logTotal := math.Log(s.Total())
	for idx := n - 1; idx >= 0; idx-- {
		candidates := make(routes, len(d[idx]))
		for index, i := range d[idx] {
			word := string(runes[idx : i+1])
			if freq, ok := s.Freq(word); ok {
				candidates[index] = route{Freq: math.Log(freq) - logTotal + rs[i+1].Freq, Index: i}
			} else {
				candidates[index] = route{Freq: math.Log(1.0) - logTotal + rs[i+1].Freq, Index: i}
			}
		}
		sort.Sort(sort.Reverse(candidates))
		rs[idx] = candidates[0]
	}
	return rs
}
