package posseg

import (
	"github.com/wangbin/jiebago"
	"regexp"
	"strings"
)

var (
	reHanDetail    = regexp.MustCompile(`\p{Han}+`)
	reSkipDetail   = regexp.MustCompile(`[[\.[:digit:]]+|[:alnum:]]+`)
	reEng          = regexp.MustCompile(`[[:alnum:]]`)
	reNum          = regexp.MustCompile(`[\.[:digit:]]+`)
	reEng1         = regexp.MustCompile(`[[:alnum:]]$`)
	reHanInternal  = regexp.MustCompile(`([\p{Han}+[:alnum:]+#&\._]+)`)
	reSkipInternal = regexp.MustCompile(`(\r\n|\s)`)
)

type Pair struct {
	Word, Flag string
}

type Posseg struct {
	*jiebago.Jieba
	Flag map[string]string
}

func (p *Posseg) AddEntry(entry jiebago.Entry) {
	if len(entry.Flag) > 0 {
		p.Flag[entry.Word] = strings.TrimSpace(entry.Flag)
	}
	p.Add(entry.Word, entry.Freq)
}

// Set dictionary, it could be absolute path of dictionary file, or dictionary
// name in current diectory.
func NewPosseg(dictFileName string) (*Posseg, error) {
	j := jiebago.New()
	p := &Posseg{j, make(map[string]string)}
	err := jiebago.LoadDict(p, dictFileName, true)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Load user specified dictionary file.
func (p *Posseg) LoadUserDict(dictFilePath string) error {
	return jiebago.LoadDict(p, dictFilePath, true)
}

func (p *Posseg) cutDetailInternal(sentence string) chan Pair {
	result := make(chan Pair)

	go func() {
		runes := []rune(sentence)
		posList := viterbi(runes)
		begin := 0
		next := 0
		for i, char := range runes {
			pos := posList[i]
			switch pos[0] {
			case 'B':
				begin = i
			case 'E':
				result <- Pair{string(runes[begin : i+1]), string(pos[1:])}
				next = i + 1
			case 'S':
				result <- Pair{string(char), string(pos[1:])}
				next = i + 1
			}
		}
		if next < len(runes) {
			result <- Pair{string(runes[next:]), string(posList[next][1:])}
		}
		close(result)
	}()
	return result
}

func (p *Posseg) cutDetail(sentence string) chan Pair {
	result := make(chan Pair)

	go func() {
		for blk := range jiebago.RegexpSplit(reHanDetail, sentence) {
			if reHanDetail.MatchString(blk) {
				for wordTag := range p.cutDetailInternal(blk) {
					result <- wordTag
				}
			} else {
				for x := range jiebago.RegexpSplit(reSkipDetail, blk) {
					if len(x) == 0 {
						continue
					}
					switch {
					case reNum.MatchString(x):
						result <- Pair{x, "m"}
					case reEng.MatchString(x):
						result <- Pair{x, "eng"}
					default:
						result <- Pair{x, "x"}
					}
				}
			}
		}
		close(result)
	}()
	return result
}

type cutFunc func(sentence string) chan Pair

func (p *Posseg) cutDAG(sentence string) chan Pair {
	result := make(chan Pair)

	go func() {
		dag := p.DAG(sentence)
		routes := p.Calc(sentence, dag)
		var y int
		runes := []rune(sentence)
		length := len(runes)
		buf := make([]rune, 0)
		for x := 0; x < length; {
			y = routes[x].Index + 1
			l_word := runes[x:y]
			if y-x == 1 {
				buf = append(buf, l_word...)
			} else {
				if len(buf) > 0 {
					if len(buf) == 1 {
						sbuf := string(buf)
						if tag, ok := p.Flag[sbuf]; ok {
							result <- Pair{sbuf, tag}
						} else {
							result <- Pair{sbuf, "x"}
						}
						buf = make([]rune, 0)
					} else {
						bufString := string(buf)
						if v, ok := p.Freq(bufString); !ok || v == 0.0 {
							for t := range p.cutDetail(bufString) {
								result <- t
							}
						} else {
							for _, elem := range buf {
								selem := string(elem)
								if tag, ok := p.Flag[selem]; ok {
									result <- Pair{string(elem), tag}
								} else {
									result <- Pair{string(elem), "x"}
								}

							}
						}
						buf = make([]rune, 0)
					}
				}
				sl_word := string(l_word)
				if tag, ok := p.Flag[sl_word]; ok {
					result <- Pair{sl_word, tag}
				} else {
					result <- Pair{sl_word, "x"}
				}
			}
			x = y
		}

		if len(buf) > 0 {
			if len(buf) == 1 {
				sbuf := string(buf)
				if tag, ok := p.Flag[sbuf]; ok {
					result <- Pair{sbuf, tag}
				} else {
					result <- Pair{sbuf, "x"}
				}
			} else {
				bufString := string(buf)
				if v, ok := p.Freq(bufString); !ok || v == 0.0 {
					for t := range p.cutDetail(bufString) {
						result <- t
					}
				} else {
					for _, elem := range buf {
						selem := string(elem)
						if tag, ok := p.Flag[selem]; ok {
							result <- Pair{selem, tag}
						} else {
							result <- Pair{selem, "x"}
						}
					}
				}
			}
		}
		close(result)
	}()
	return result
}

func (p *Posseg) cutDAGNoHMM(sentence string) chan Pair {
	result := make(chan Pair)

	go func() {
		dag := p.DAG(sentence)
		routes := p.Calc(sentence, dag)
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
			if reEng1.MatchString(string(l_word)) && len(l_word) == 1 {
				buf = append(buf, l_word...)
				x = y
			} else {
				if len(buf) > 0 {
					result <- Pair{string(buf), "eng"}
					buf = make([]rune, 0)
				}
				sl_word := string(l_word)
				if tag, ok := p.Flag[sl_word]; ok {
					result <- Pair{sl_word, tag}
				} else {
					result <- Pair{sl_word, "x"}
				}
				x = y
			}
		}
		if len(buf) > 0 {
			result <- Pair{string(buf), "eng"}
			buf = make([]rune, 0)
		}
		close(result)
	}()
	return result
}

// Tags the POS of each word after segmentation, using labels compatible with
// ictclas.
func (p *Posseg) Cut(sentence string, HMM bool) chan Pair {
	result := make(chan Pair)
	var cut cutFunc
	if HMM {
		cut = p.cutDAG
	} else {
		cut = p.cutDAGNoHMM
	}
	go func() {
		for blk := range jiebago.RegexpSplit(reHanInternal, sentence) {
			if reHanInternal.MatchString(blk) {
				for wordTag := range cut(blk) {
					result <- wordTag
				}
			} else {
				for x := range jiebago.RegexpSplit(reSkipInternal, blk) {
					if reSkipInternal.MatchString(x) {
						result <- Pair{x, "x"}
					} else {
						for _, xx := range x {
							s := string(xx)
							switch {
							case reNum.MatchString(s):
								result <- Pair{s, "m"}
							case reEng.MatchString(x):
								result <- Pair{x, "eng"}
								break
							default:
								result <- Pair{s, "x"}
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
