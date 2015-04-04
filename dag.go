package jiebago

import (
	"math"
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
	var i int
	for k := 0; k < n; k++ {
		d[k] = make([]int, 0)
		i = k
		frag = string(runes[k])
		for {
			if freq, ok := s.Freq(frag); !ok {
				break
			} else {
				if freq > 0.0 {
					d[k] = append(d[k], i)
				}
			}
			i += 1
			if i >= n {
				break
			}
			frag = string(runes[k : i+1])
		}
		if len(d[k]) == 0 {
			d[k] = append(d[k], k)
		}
	}
	return d
}

func Routes(s Segmenter, runes []rune, d dag) map[int]route {
	n := len(runes)
	rs := make(map[int]route)
	rs[n] = route{Freq: 0.0, Index: 0}
	logTotal := s.LogTotal()
	var r route
	for idx := n - 1; idx >= 0; idx-- {
		for _, i := range d[idx] {
			word := string(runes[idx : i+1])
			if freq, ok := s.Freq(word); ok {
				r = route{Freq: math.Log(freq) - logTotal + rs[i+1].Freq, Index: i}
			} else {
				r = route{Freq: math.Log(1.0) - logTotal + rs[i+1].Freq, Index: i}
			}

			if v, ok := rs[idx]; !ok || v.Freq < r.Freq || (v.Freq == r.Freq && v.Index < r.Index) {
				rs[idx] = r
			}
		}
	}
	return rs
}
