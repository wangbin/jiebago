package jieba

import (
	"math"
	"regexp"

	"github.com/wangbin/jiebago/dictionary"
	"github.com/wangbin/jiebago/finalseg"
	"github.com/wangbin/jiebago/util"
)

var (
	reEng         = regexp.MustCompile(`[[:alnum:]]`)
	reHanCutAll   = regexp.MustCompile(`(\p{Han}+)`)
	reSkipCutAll  = regexp.MustCompile(`[^[:alnum:]+#\n]`)
	reHanDefault  = regexp.MustCompile(`([\p{Han}+[:alnum:]+#&\._]+)`)
	reSkipDefault = regexp.MustCompile(`(\r\n|\s)`)
)

type Segmenter struct {
	*dictionary.Dictionary
}

func (seg *Segmenter) dag(runes []rune) map[int][]int {
	dag := make(map[int][]int)
	n := len(runes)
	var frag []rune
	var i int
	for k := 0; k < n; k++ {
		dag[k] = make([]int, 0)
		i = k
		frag = runes[k : k+1]
		for {
			freq, ok := seg.Frequency(string(frag))
			if !ok {
				break
			}
			if freq > 0.0 {
				dag[k] = append(dag[k], i)
			}
			i += 1
			if i >= n {
				break
			}
			frag = runes[k : i+1]
		}
		if len(dag[k]) == 0 {
			dag[k] = append(dag[k], k)
		}
	}
	return dag
}

type route struct {
	frequency float64
	index     int
}

func (seg *Segmenter) calc(runes []rune) map[int]route {
	dag := seg.dag(runes)
	n := len(runes)
	rs := make(map[int]route)
	rs[n] = route{frequency: 0.0, index: 0}
	logTotal := seg.LogTotal()
	var r route
	for idx := n - 1; idx >= 0; idx-- {
		for _, i := range dag[idx] {
			if freq, ok := seg.Frequency(string(runes[idx : i+1])); ok {
				r = route{frequency: math.Log(freq) - logTotal + rs[i+1].frequency, index: i}
			} else {
				r = route{frequency: math.Log(1.0) - logTotal + rs[i+1].frequency, index: i}
			}
			if v, ok := rs[idx]; !ok {
				rs[idx] = r
			} else {
				if v.frequency < r.frequency || (v.frequency == r.frequency && v.index < r.index) {
					rs[idx] = r
				}
			}
		}
	}
	return rs
}

type cutFunc func(sentence string) <-chan string

func (seg *Segmenter) cutDAG(sentence string) <-chan string {
	result := make(chan string)
	go func() {
		runes := []rune(sentence)
		routes := seg.calc(runes)
		var y int
		length := len(runes)
		buf := make([]rune, 0)
		for x := 0; x < length; {
			y = routes[x].index + 1
			frag := runes[x:y]
			if y-x == 1 {
				buf = append(buf, frag...)
			} else {
				if len(buf) > 0 {
					bufString := string(buf)
					if len(buf) == 1 {
						result <- bufString
					} else {
						if v, ok := seg.Frequency(bufString); !ok || v == 0.0 {
							for x := range finalseg.Cut(bufString) {
								result <- x
							}
						} else {
							for _, elem := range buf {
								result <- string(elem)
							}
						}
					}
					buf = make([]rune, 0)
				}
				result <- string(frag)
			}
			x = y
		}

		if len(buf) > 0 {
			bufString := string(buf)
			if len(buf) == 1 {
				result <- bufString
			} else {
				if v, ok := seg.Frequency(bufString); !ok || v == 0.0 {
					for t := range finalseg.Cut(bufString) {
						result <- t
					}
				} else {
					for _, elem := range buf {
						result <- string(elem)
					}
				}
			}
		}
		close(result)
	}()
	return result
}

func (seg *Segmenter) cutDAGNoHMM(sentence string) <-chan string {
	result := make(chan string)

	go func() {
		runes := []rune(sentence)
		routes := seg.calc(runes)
		var y int
		length := len(runes)
		buf := make([]rune, 0)
		for x := 0; x < length; {
			y = routes[x].index + 1
			frag := runes[x:y]
			if reEng.MatchString(string(frag)) && len(frag) == 1 {
				buf = append(buf, frag...)
				x = y
			} else {
				if len(buf) > 0 {
					result <- string(buf)
					buf = make([]rune, 0)
				}
				result <- string(frag)
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

func (seg *Segmenter) Cut(sentence string, hmm bool) <-chan string {
	result := make(chan string)
	var cut cutFunc
	if hmm {
		cut = seg.cutDAG
	} else {
		cut = seg.cutDAGNoHMM
	}

	go func() {
		for _, block := range util.RegexpSplit(reHanDefault, sentence, -1) {
			if len(block) == 0 {
				continue
			}
			if reHanDefault.MatchString(block) {
				for x := range cut(block) {
					result <- x
				}
				continue
			}
			for _, subBlock := range util.RegexpSplit(reSkipDefault, block, -1) {
				if reSkipDefault.MatchString(subBlock) {
					result <- subBlock
					continue
				}
				for _, r := range subBlock {
					result <- string(r)
				}
			}
		}
		close(result)
	}()
	return result
}

func (seg *Segmenter) cutAll(sentence string) <-chan string {
	result := make(chan string)
	go func() {
		runes := []rune(sentence)
		dag := seg.dag(runes)
		start := -1
		ks := make([]int, len(dag))
		for k := range dag {
			ks[k] = k
		}
		var l []int
		for k := range ks {
			l = dag[k]
			if len(l) == 1 && k > start {
				result <- string(runes[k : l[0]+1])
				start = l[0]
				continue
			}
			for _, j := range l {
				if j > k {
					result <- string(runes[k : j+1])
					start = j
				}
			}
		}
		close(result)
	}()
	return result
}

func (seg *Segmenter) CutAll(sentence string) <-chan string {
	result := make(chan string)
	go func() {
		for _, block := range util.RegexpSplit(reHanCutAll, sentence, -1) {
			if len(block) == 0 {
				continue
			}
			if reHanCutAll.MatchString(block) {
				for x := range seg.cutAll(block) {
					result <- x
				}
				continue
			}
			for _, subBlock := range reSkipCutAll.Split(block, -1) {
				result <- subBlock
			}
		}
		close(result)
	}()
	return result
}

func (seg *Segmenter) CutForSearch(sentence string, hmm bool) <-chan string {
	result := make(chan string)
	go func() {
		for word := range seg.Cut(sentence, hmm) {
			runes := []rune(word)
			for _, increment := range []int{2, 3} {
				if len(runes) <= increment {
					continue
				}
				var gram string
				for i := 0; i < len(runes)-increment+1; i++ {
					gram = string(runes[i : i+increment])
					if v, ok := seg.Frequency(gram); ok && v > 0.0 {
						result <- gram
					}
				}
			}
			result <- word
		}
		close(result)
	}()
	return result
}

func New() *Segmenter {
	return &Segmenter{dictionary.New()}
}
