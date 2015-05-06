package posseg

import (
	"math"
	"regexp"

	"github.com/wangbin/jiebago/util"
)

var (
	reHanDetail    = regexp.MustCompile(`(\p{Han}+)`)
	reSkipDetail   = regexp.MustCompile(`([[\.[:digit:]]+|[:alnum:]]+)`)
	reEng          = regexp.MustCompile(`[[:alnum:]]`)
	reNum          = regexp.MustCompile(`[\.[:digit:]]+`)
	reEng1         = regexp.MustCompile(`[[:alnum:]]$`)
	reHanInternal  = regexp.MustCompile(`([\p{Han}+[:alnum:]+#&\._]+)`)
	reSkipInternal = regexp.MustCompile(`(\r\n|\s)`)
)

// Segment represents a word with it's POS
type Segment struct {
	text, pos string
}

// Text returns the Segment's text.
func (s Segment) Text() string {
	return s.text
}

// Pos returns the Segment's POS.
func (s Segment) Pos() string {
	return s.pos
}

// Segmenter is a Chinese words segmentation struct.
type Segmenter struct {
	dict *Dictionary
}

// LoadDictionary loads dictionary from given file name.
// Everytime LoadDictionary is called, previously loaded dictionary will be cleard.
func (seg *Segmenter) LoadDictionary(fileName string) error {
	seg.dict = &Dictionary{freqMap: make(map[string]float64), posMap: make(map[string]string)}
	return seg.dict.loadDictionary(fileName)
}

// LoadUserDictionary loads a user specified dictionary, it must be called
// after LoadDictionary, and it will not clear any previous loaded dictionary,
// instead it will override exist entries.
func (seg *Segmenter) LoadUserDictionary(fileName string) error {
	return seg.dict.loadDictionary(fileName)
}

func (seg *Segmenter) cutDetailInternal(sentence string) <-chan Segment {
	result := make(chan Segment)

	go func() {
		runes := []rune(sentence)
		posList := viterbi(runes)
		begin := 0
		next := 0
		for i, char := range runes {
			pos := posList[i]
			switch pos.position() {
			case "B":
				begin = i
			case "E":
				result <- Segment{string(runes[begin : i+1]), pos.pos()}
				next = i + 1
			case "S":
				result <- Segment{string(char), pos.pos()}
				next = i + 1
			}
		}
		if next < len(runes) {
			result <- Segment{string(runes[next:]), posList[next].pos()}
		}
		close(result)
	}()
	return result
}

func (seg *Segmenter) cutDetail(sentence string) <-chan Segment {
	result := make(chan Segment)
	go func() {
		for _, blk := range util.RegexpSplit(reHanDetail, sentence, -1) {
			if reHanDetail.MatchString(blk) {
				for segment := range seg.cutDetailInternal(blk) {
					result <- segment
				}
				continue
			}
			for _, x := range util.RegexpSplit(reSkipDetail, blk, -1) {
				if len(x) == 0 {
					continue
				}
				switch {
				case reNum.MatchString(x):
					result <- Segment{x, "m"}
				case reEng.MatchString(x):
					result <- Segment{x, "eng"}
				default:
					result <- Segment{x, "x"}
				}
			}
		}
		close(result)
	}()
	return result
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
			freq, ok := seg.dict.Frequency(string(frag))
			if !ok {
				break
			}
			if freq > 0.0 {
				dag[k] = append(dag[k], i)
			}
			i++
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
	var r route
	for idx := n - 1; idx >= 0; idx-- {
		for _, i := range dag[idx] {
			if freq, ok := seg.dict.Frequency(string(runes[idx : i+1])); ok {
				r = route{frequency: math.Log(freq) - seg.dict.logTotal + rs[i+1].frequency, index: i}
			} else {
				r = route{frequency: math.Log(1.0) - seg.dict.logTotal + rs[i+1].frequency, index: i}
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

type cutFunc func(sentence string) <-chan Segment

func (seg *Segmenter) cutDAG(sentence string) <-chan Segment {
	result := make(chan Segment)

	go func() {
		runes := []rune(sentence)
		routes := seg.calc(runes)
		var y int
		length := len(runes)
		var buf []rune
		for x := 0; x < length; {
			y = routes[x].index + 1
			frag := runes[x:y]
			if y-x == 1 {
				buf = append(buf, frag...)
			} else {
				if len(buf) > 0 {
					bufString := string(buf)
					if len(buf) == 1 {
						if tag, ok := seg.dict.Pos(bufString); ok {
							result <- Segment{bufString, tag}
						} else {
							result <- Segment{bufString, "x"}
						}
						buf = make([]rune, 0)
					} else {
						if v, ok := seg.dict.Frequency(bufString); !ok || v == 0.0 {
							for t := range seg.cutDetail(bufString) {
								result <- t
							}
						} else {
							for _, elem := range buf {
								selem := string(elem)
								if tag, ok := seg.dict.Pos(selem); ok {
									result <- Segment{selem, tag}
								} else {
									result <- Segment{selem, "x"}
								}

							}
						}
						buf = make([]rune, 0)
					}
				}
				word := string(frag)
				if tag, ok := seg.dict.Pos(word); ok {
					result <- Segment{word, tag}
				} else {
					result <- Segment{word, "x"}
				}
			}
			x = y
		}

		if len(buf) > 0 {
			bufString := string(buf)
			if len(buf) == 1 {
				if tag, ok := seg.dict.Pos(bufString); ok {
					result <- Segment{bufString, tag}
				} else {
					result <- Segment{bufString, "x"}
				}
			} else {
				if v, ok := seg.dict.Frequency(bufString); !ok || v == 0.0 {
					for t := range seg.cutDetail(bufString) {
						result <- t
					}
				} else {
					for _, elem := range buf {
						selem := string(elem)
						if tag, ok := seg.dict.Pos(selem); ok {
							result <- Segment{selem, tag}
						} else {
							result <- Segment{selem, "x"}
						}
					}
				}
			}
		}
		close(result)
	}()
	return result
}

func (seg *Segmenter) cutDAGNoHMM(sentence string) <-chan Segment {
	result := make(chan Segment)

	go func() {
		runes := []rune(sentence)
		routes := seg.calc(runes)
		var y int
		length := len(runes)
		var buf []rune
		for x := 0; x < length; {
			y = routes[x].index + 1
			frag := runes[x:y]
			if reEng1.MatchString(string(frag)) && len(frag) == 1 {
				buf = append(buf, frag...)
				x = y
			} else {
				if len(buf) > 0 {
					result <- Segment{string(buf), "eng"}
					buf = make([]rune, 0)
				}
				word := string(frag)
				if tag, ok := seg.dict.Pos(word); ok {
					result <- Segment{word, tag}
				} else {
					result <- Segment{word, "x"}
				}
				x = y
			}
		}
		if len(buf) > 0 {
			result <- Segment{string(buf), "eng"}
			buf = make([]rune, 0)
		}
		close(result)
	}()
	return result
}

// Cut cuts a sentence into words.
// Parameter hmm controls whether to use the Hidden Markov Model.
func (seg *Segmenter) Cut(sentence string, hmm bool) <-chan Segment {
	result := make(chan Segment)
	var cut cutFunc
	if hmm {
		cut = seg.cutDAG
	} else {
		cut = seg.cutDAGNoHMM
	}
	go func() {
		for _, blk := range util.RegexpSplit(reHanInternal, sentence, -1) {
			if reHanInternal.MatchString(blk) {
				for wordTag := range cut(blk) {
					result <- wordTag
				}
			} else {
				for _, x := range util.RegexpSplit(reSkipInternal, blk, -1) {
					if reSkipInternal.MatchString(x) {
						result <- Segment{x, "x"}
					} else {
						for _, xx := range x {
							s := string(xx)
							switch {
							case reNum.MatchString(s):
								result <- Segment{s, "m"}
							case reEng.MatchString(x):
								result <- Segment{x, "eng"}
							default:
								result <- Segment{s, "x"}
							}
						}
					}
				}
			}
		}
		close(result)
	}()
	return result
}
