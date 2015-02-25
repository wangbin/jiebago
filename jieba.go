package jiebago

import (
	"fmt"
	"github.com/wangbin/jiebago/finalseg"
	"math"
	"regexp"
	"sort"
)

var (
	Dictionary     = "dict.txt"
	UserWordTagTab = make(map[string]string)
)

type Route struct {
	Freq  float64
	Index int
}

func (route Route) String() string {
	return fmt.Sprintf("(%f, %d)", route.Freq, route.Index)
}

type Routes []*Route

func (routes Routes) Len() int {
	return len(routes)
}

func (routes Routes) Less(i, j int) bool {
	routei := routes[i]
	routej := routes[j]
	if routei.Freq < routej.Freq {
		return true
	} else if routei.Freq == routej.Freq {
		return routei.Index < routej.Index
	}
	return false
}

func (routes Routes) Swap(i, j int) {
	routes[i], routes[j] = routes[j], routes[i]
}

func RegexpSplit(r *regexp.Regexp, sentence string) []string {
	result := make([]string, 0)
	locs := r.FindAllStringIndex(sentence, -1)
	lastLoc := 0
	if len(locs) == 0 {
		return []string{sentence}
	}
	for _, loc := range locs {
		if loc[0] == lastLoc {
			result = append(result, sentence[loc[0]:loc[1]])
		} else {
			result = append(result, sentence[lastLoc:loc[0]])
			result = append(result, sentence[loc[0]:loc[1]])
		}
		lastLoc = loc[1]
	}
	if lastLoc < len(sentence) {
		result = append(result, sentence[lastLoc:])
	}

	return result
}

func GetDAG(sentence string) map[int][]int {
	dag := make(map[int][]int)
	runes := []rune(sentence)
	n := len(runes)
	i := 0
	var frag string
	for k := 0; k < n; k++ {
		tmpList := make([]int, 0)
		i = k
		frag = string(runes[k])
		for {
			if !T.Nodes.Contains(frag) {
				break
			}
			if _, ok := T.Freq[frag]; ok {
				tmpList = append(tmpList, i)
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
		dag[k] = tmpList
	}
	return dag
}

func Calc(sentence string, dag map[int][]int) map[int]*Route {
	runes := []rune(sentence)
	number := len(runes)
	routes := make(map[int]*Route)
	routes[number] = &Route{Freq: 0.0, Index: 0}
	logTotal := math.Log(T.Total)
	for idx := number - 1; idx >= 0; idx-- {
		candidates := make(Routes, 0)
		for _, i := range dag[idx] {
			word := string(runes[idx : i+1])
			var route *Route
			if _, ok := T.Freq[word]; ok {
				route = &Route{Freq: math.Log(T.Freq[word]) - logTotal + routes[i+1].Freq, Index: i}
			} else {
				route = &Route{Freq: math.Log(1.0) - logTotal + routes[i+1].Freq, Index: i}
			}
			candidates = append(candidates, route)
		}
		sort.Sort(sort.Reverse(candidates))
		routes[idx] = candidates[0]
	}
	return routes
}

type cutAction func(sentence string) []string

func cut_DAG(sentence string) []string {
	dag := GetDAG(sentence)
	routes := Calc(sentence, dag)
	x := 0
	var y int
	runes := []rune(sentence)
	length := len(runes)
	result := make([]string, 0)
	buf := make([]rune, 0)
	for {
		if x >= length {
			break
		}
		y = routes[x].Index + 1
		l_word := runes[x:y]
		if y-x == 1 {
			buf = append(buf, l_word...)
		} else {
			if len(buf) > 0 {
				if len(buf) == 1 {
					result = append(result, string(buf))
					buf = make([]rune, 0)
				} else {
					bufString := string(buf)
					if _, ok := T.Freq[bufString]; !ok {
						recognized := finalseg.Cut(bufString)
						for _, t := range recognized {
							result = append(result, t)
						}
					} else {
						for _, elem := range buf {
							result = append(result, string(elem)) // TODO: I don't get this?
						}
					}
					buf = make([]rune, 0)
				}
			}
			result = append(result, string(l_word))
		}
		x = y
	}

	if len(buf) > 0 {
		if len(buf) == 1 {
			result = append(result, string(buf))
		} else {
			bufString := string(buf)
			if _, ok := T.Freq[bufString]; !ok {
				recognized := finalseg.Cut(bufString)
				for _, t := range recognized {
					result = append(result, t)
				}
			} else {
				for _, elem := range buf {
					result = append(result, string(elem)) // TODO: I don't get this?
				}
			}
		}
	}
	return result
}

func cut_DAG_NO_HMM(sentence string) []string {
	result := make([]string, 0)
	re_eng := regexp.MustCompile(`[[:alnum:]]`)
	dag := GetDAG(sentence)
	routes := Calc(sentence, dag)
	x := 0
	var y int
	runes := []rune(sentence)
	length := len(runes)
	buf := make([]rune, 0)
	for {
		if x >= length {
			break
		}
		y = routes[x].Index + 1
		l_word := runes[x:y]
		if re_eng.MatchString(string(l_word)) && len(l_word) == 1 {
			buf = append(buf, l_word...)
			x = y
		} else {
			if len(buf) > 0 {
				result = append(result, string(buf))
				buf = make([]rune, 0)
			}
			result = append(result, string(l_word))
			x = y
		}
	}
	if len(buf) > 0 {
		result = append(result, string(buf))
		buf = make([]rune, 0)
	}
	return result
}

func cut_All(sentence string) []string {
	result := make([]string, 0)
	runes := []rune(sentence)
	dag := GetDAG(sentence)
	old_j := -1
	ks := make([]int, 0)
	for k := range dag {
		ks = append(ks, k)
	}
	sort.Ints(ks)
	for k := range ks {
		l := dag[k]
		if len(l) == 1 && k > old_j {
			result = append(result, string(runes[k:l[0]+1]))
			old_j = l[0]
		} else {
			for _, j := range l {
				if j > k {
					result = append(result, string(runes[k:j+1]))
					old_j = j
				}
			}
		}
	}
	return result
}

func Cut(sentence string, cut_all bool, HMM bool) []string {
	result := make([]string, 0)
	var re_han, re_skip *regexp.Regexp
	if cut_all {
		re_han = regexp.MustCompile(`\p{Han}+`)
		re_skip = regexp.MustCompile(`[^[:alnum:]+#\n]`)
	} else {
		re_han = regexp.MustCompile(`([\p{Han}+[:alnum:]+#&\._]+)`)
		re_skip = regexp.MustCompile(`(\r\n|\s)`)
	}
	blocks := RegexpSplit(re_han, sentence)
	var cut_block cutAction
	if HMM {
		cut_block = cut_DAG
	} else {
		cut_block = cut_DAG_NO_HMM
	}
	if cut_all {
		cut_block = cut_All
	}
	for _, blk := range blocks {
		if len(blk) == 0 {
			continue
		}
		if re_han.MatchString(blk) {
			for _, word := range cut_block(blk) {
				result = append(result, word)
			}
		} else {
			type skipSplitFunc func(sentence string) []string
			var ssf skipSplitFunc
			if cut_all {
				ssf = func(sentence string) []string {
					return re_skip.Split(sentence, -1)
				}
			} else {
				ssf = func(sentence string) []string {
					return RegexpSplit(re_skip, sentence)
				}
			}

			for _, x := range ssf(blk) {
				if re_skip.MatchString(x) {
					result = append(result, x)
				} else if !cut_all {
					for _, xx := range x {
						result = append(result, string(xx))
					}
				} else {
					result = append(result, x)
				}
			}
		}
	}
	return result
}

func CutForSearch(sentence string, hmm bool) []string {
	result := make([]string, 0)
	words := Cut(sentence, false, hmm)
	for _, word := range words {
		runes := []rune(word)
		for _, increment := range []int{2, 3} {
			if len(runes) > increment {
				var gram2 string
				for i := 0; i < len(runes)-increment+1; i++ {
					gram2 = string(runes[i : i+increment])
					if _, ok := T.Freq[gram2]; ok {
						result = append(result, gram2)
					}
				}
			}
		}
		result = append(result, word)
	}
	return result
}
