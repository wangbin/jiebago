// Golang implemention of jieba (Python Chinese word segmentation module).
package jiebago

import (
	"fmt"
	"github.com/wangbin/jiebago/finalseg"
	"math"
	"regexp"
	"sort"
)

var (
	// Word/Tag Map load from user dictionary
	UserWordTagTab = make(map[string]string)
	reEng          = regexp.MustCompile(`[[:alnum:]]`)
	reHanCutAll    = regexp.MustCompile(`\p{Han}+`)
	reSkipCutAll   = regexp.MustCompile(`[^[:alnum:]+#\n]`)
	reHanDefault   = regexp.MustCompile(`([\p{Han}+[:alnum:]+#&\._]+)`)
	reSkipDefault  = regexp.MustCompile(`(\r\n|\s)`)
)

type route struct {
	Freq  float64
	Index int
}

func (r route) String() string {
	return fmt.Sprintf("(%f, %d)", r.Freq, r.Index)
}

type routes []*route

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

// Build a directed acyclic graph (DAG) for sentence.
func (j *Jieba) DAG(sentence string) map[int][]int {
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
			if freq, ok := j.Freq[frag]; !ok {
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
		dag[k] = tmpList
	}
	return dag
}

func (j *Jieba) Calc(sentence string, dag map[int][]int) map[int]*route {
	runes := []rune(sentence)
	number := len(runes)
	rs := make(map[int]*route)
	rs[number] = &route{Freq: 0.0, Index: 0}
	logTotal := math.Log(j.Total)
	for idx := number - 1; idx >= 0; idx-- {
		candidates := make(routes, 0)
		for _, i := range dag[idx] {
			word := string(runes[idx : i+1])
			var r *route
			if _, ok := j.Freq[word]; ok {
				r = &route{Freq: math.Log(j.Freq[word]) - logTotal + rs[i+1].Freq, Index: i}
			} else {
				r = &route{Freq: math.Log(1.0) - logTotal + rs[i+1].Freq, Index: i}
			}
			candidates = append(candidates, r)
		}
		sort.Sort(sort.Reverse(candidates))
		rs[idx] = candidates[0]
	}
	return rs
}

type cutFunc func(sentence string) chan string

func (j *Jieba) cutDAG(sentence string) chan string {
	result := make(chan string)
	go func() {
		dag := j.DAG(sentence)
		routes := j.Calc(sentence, dag)
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
			if y-x == 1 {
				buf = append(buf, l_word...)
			} else {
				if len(buf) > 0 {
					if len(buf) == 1 {
						result <- string(buf)
						buf = make([]rune, 0)
					} else {
						bufString := string(buf)
						if v, ok := j.Freq[bufString]; !ok || v == 0.0 {
							for x := range finalseg.Cut(bufString) {
								result <- x
							}
						} else {
							for _, elem := range buf {
								result <- string(elem) // TODO: I don't get this?
							}
						}
						buf = make([]rune, 0)
					}
				}
				result <- string(l_word)
			}
			x = y
		}

		if len(buf) > 0 {
			if len(buf) == 1 {
				result <- string(buf)
			} else {
				bufString := string(buf)
				if v, ok := j.Freq[bufString]; !ok || v == 0.0 {
					for t := range finalseg.Cut(bufString) {
						result <- t
					}
				} else {
					for _, elem := range buf {
						result <- string(elem) // TODO: I don't get this?
					}
				}
			}
		}
		close(result)
	}()
	return result
}

func (j *Jieba) cutDAGNoHMM(sentence string) chan string {
	result := make(chan string)

	go func() {
		dag := j.DAG(sentence)
		routes := j.Calc(sentence, dag)
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
			if reEng.MatchString(string(l_word)) && len(l_word) == 1 {
				buf = append(buf, l_word...)
				x = y
			} else {
				if len(buf) > 0 {
					result <- string(buf)
					buf = make([]rune, 0)
				}
				result <- string(l_word)
				x = y
			}
		}
		if len(buf) > 0 {
			result <- string(buf)
			buf = make([]rune, 0)
		}
		close(result)
	}()
	return result
}

func (j *Jieba) cutAll(sentence string) chan string {
	result := make(chan string)

	go func() {
		runes := []rune(sentence)
		dag := j.DAG(sentence)
		old_j := -1
		ks := make([]int, 0)
		for k := range dag {
			ks = append(ks, k)
		}
		sort.Ints(ks)
		for k := range ks {
			l := dag[k]
			if len(l) == 1 && k > old_j {
				result <- string(runes[k : l[0]+1])
				old_j = l[0]
			} else {
				for _, j := range l {
					if j > k {
						result <- string(runes[k : j+1])
						old_j = j
					}
				}
			}
		}
		close(result)
	}()
	return result
}

/*
Cut sentence.

isCutAll controls use full cut mode or accurate mode.

Full Mode gets all the possible words from the sentence. Fast but not accurate.

Accurate Mode attempts to cut the sentence into the most accurate segmentations,
which is suitable for text analysis.

HMM contols whether to use the Hidden Markov Mode.
*/
func (j *Jieba) Cut(sentence string, isCutAll bool, HMM bool) chan string {
	result := make(chan string)
	go func() {
		var reHan, reSkip *regexp.Regexp
		if isCutAll {
			reHan = reHanCutAll
			reSkip = reSkipCutAll
		} else {
			reHan = reHanDefault
			reSkip = reSkipDefault
		}
		var cut cutFunc
		if HMM {
			cut = j.cutDAG
		} else {
			cut = j.cutDAGNoHMM
		}
		if isCutAll {
			cut = j.cutAll
		}
		for blk := range RegexpSplit(reHan, sentence) {
			if len(blk) == 0 {
				continue
			}
			if reHan.MatchString(blk) {
				for x := range cut(blk) {
					result <- x
				}
			} else {
				type skipSplitFunc func(sentence string) chan string
				var ssf skipSplitFunc
				if isCutAll {
					ssf = func(sentence string) chan string {
						ch := make(chan string)
						go func() {
							for _, s := range reSkip.Split(sentence, -1) {
								ch <- s
							}
							close(ch)
						}()
						return ch
					}
				} else {
					ssf = func(sentence string) chan string {
						return RegexpSplit(reSkip, sentence)
					}
				}

				for x := range ssf(blk) {
					if reSkip.MatchString(x) {
						result <- x
					} else if !isCutAll {
						for _, xx := range x {
							result <- string(xx)
						}
					} else {
						result <- x
					}
				}
			}
		}
		close(result)
	}()
	return result
}

// Cut sentence using Search Engine Mode, based on the Accurate Mode, attempts
// to cut long words into several short words, which can raise the recall rate.
// Suitable for search engines.
func (j *Jieba) CutForSearch(sentence string, hmm bool) chan string {
	result := make(chan string)
	go func() {
		for word := range j.Cut(sentence, false, hmm) {
			runes := []rune(word)
			for _, increment := range []int{2, 3} {
				if len(runes) > increment {
					var gram2 string
					for i := 0; i < len(runes)-increment+1; i++ {
						gram2 = string(runes[i : i+increment])
						if v, ok := j.Freq[gram2]; ok && v > 0.0 {
							result <- gram2
						}
					}
				}
			}
			result <- word
		}
		close(result)
	}()
	return result
}
