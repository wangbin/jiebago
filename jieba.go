// Golang implemention of jieba (Python Chinese word segmentation module).
package jiebago

import (
	"errors"
	"github.com/wangbin/jiebago/finalseg"
	"regexp"
	"sort"
)

var (
	ErrInitialized = errors.New("already initialized")
	reEng          = regexp.MustCompile(`[[:alnum:]]`)
	reHanCutAll    = regexp.MustCompile(`(\p{Han}+)`)
	reSkipCutAll   = regexp.MustCompile(`[^[:alnum:]+#\n]`)
	reHanDefault   = regexp.MustCompile(`([\p{Han}+[:alnum:]+#&\._]+)`)
	reSkipDefault  = regexp.MustCompile(`(\r\n|\s)`)
)

// RegexpSplit split slices s into substrings separated by the expression and
// returns a slice of the substrings between those expression matches.
// If capturing parentheses are used in expression, then the text of all groups
// in the expression are also returned as part of the resulting slice.
//
// This function acts consistent with Python's re.split function.
func RegexpSplit(re *regexp.Regexp, s string, n int) []string {
	if n == 0 {
		return nil
	}

	if len(re.String()) > 0 && len(s) == 0 {
		return []string{""}
	}

	var matches [][]int
	if len(re.SubexpNames()) > 1 {
		matches = re.FindAllStringSubmatchIndex(s, n)
	} else {
		matches = re.FindAllStringIndex(s, n)
	}
	strings := make([]string, 0, len(matches))

	beg := 0
	end := 0
	for _, match := range matches {
		if n > 0 && len(strings) >= n-1 {
			break
		}

		end = match[0]
		if match[1] != 0 {
			strings = append(strings, s[beg:end])
		}
		beg = match[1]
		if len(re.SubexpNames()) > 1 {
			strings = append(strings, s[match[0]:match[1]])
		}
	}

	if end != len(s) {
		strings = append(strings, s[beg:])
	}

	return strings
}

type Segmenter interface {
	Freq(string) (float64, bool)
	Total() float64
}

type Jieba struct {
	total   float64
	freqMap map[string]float64
}

func (j Jieba) Freq(key string) (float64, bool) {
	freq, ok := j.freqMap[key]
	return freq, ok
}

func (j Jieba) Total() float64 {
	return j.total
}

func (j *Jieba) AddEntry(entry Entry) {
	j.Add(entry.Word, entry.Freq)
}

func (j *Jieba) Add(word string, freq float64) {
	j.freqMap[word] = freq
	j.total += freq
	runes := []rune(word)
	for i := 0; i < len(runes); i++ {
		frag := string(runes[0 : i+1])
		if _, ok := j.Freq(frag); !ok {
			j.freqMap[frag] = 0.0
		}
	}
}

// Load user specified dictionary file.
func (j *Jieba) LoadUserDict(dictFileName string) error {
	return LoadDict(j, dictFileName, false)
}

func (j *Jieba) SetDict(dictFileName string) error {
	if len(j.freqMap) > 0 || j.total > 0.0 {
		return ErrInitialized
	}
	return LoadDict(j, dictFileName, false)
}

func New() *Jieba {
	return &Jieba{total: 0.0, freqMap: make(map[string]float64)}
}

// Set the dictionary, could be absolute path of dictionary file, or dictionary
// name in current directory. This function must be called before cut any
// sentence.
func Open(dictFileName string) (*Jieba, error) {
	j := New()
	err := LoadDict(j, dictFileName, false)
	return j, err
}

type cutFunc func(sentence string) chan string

func (j *Jieba) cutDAG(sentence string) chan string {
	result := make(chan string)
	go func() {
		runes := []rune(sentence)
		dag := DAG(j, runes)
		routes := Routes(j, runes, dag)
		x := 0
		var y int
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
						if v, ok := j.Freq(bufString); !ok || v == 0.0 {
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
				if v, ok := j.Freq(bufString); !ok || v == 0.0 {
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
		runes := []rune(sentence)
		dag := DAG(j, runes)
		routes := Routes(j, runes, dag)
		x := 0
		var y int
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
		dag := DAG(j, runes)
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
func (j *Jieba) Cut(sentence string, hmm bool) chan string {
	result := make(chan string)
	var cut cutFunc
	if hmm {
		cut = j.cutDAG
	} else {
		cut = j.cutDAGNoHMM
	}

	go func() {
		for _, block := range RegexpSplit(reHanDefault, sentence, -1) {
			if len(block) == 0 {
				continue
			}
			if reHanDefault.MatchString(block) {
				for x := range cut(block) {
					result <- x
				}
			} else {
				for _, subBlock := range RegexpSplit(reSkipDefault, block, -1) {
					if reSkipDefault.MatchString(subBlock) {
						result <- subBlock
					} else {
						for _, r := range subBlock {
							result <- string(r)
						}
					}
				}
			}
		}
		close(result)
	}()
	return result
}

func (j *Jieba) CutAll(sentence string) chan string {
	result := make(chan string)
	go func() {
		for _, block := range RegexpSplit(reHanCutAll, sentence, -1) {
			if len(block) == 0 {
				continue
			}
			if reHanCutAll.MatchString(block) {
				for x := range j.cutAll(block) {
					result <- x
				}
			} else {
				for _, subBlock := range reSkipCutAll.Split(block, -1) {
					result <- subBlock
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
		for word := range j.Cut(sentence, hmm) {
			runes := []rune(word)
			for _, increment := range []int{2, 3} {
				if len(runes) > increment {
					var gram2 string
					for i := 0; i < len(runes)-increment+1; i++ {
						gram2 = string(runes[i : i+increment])
						if v, ok := j.Freq(gram2); ok && v > 0.0 {
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
