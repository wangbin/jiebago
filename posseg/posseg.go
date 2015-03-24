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

type WordTag struct {
	Word, Tag string
}

type Posseg struct {
	*jiebago.Jieba
	Flag map[string]string
}

func (p *Posseg) Add(wtf *jiebago.WordTagFreq) {
	if len(wtf.Tag) > 0 {
		p.Flag[wtf.Word] = strings.TrimSpace(wtf.Tag)
	}
	p.AddWord(wtf)
}

// Set dictionary, it could be absolute path of dictionary file, or dictionary
// name in current diectory.
func NewPosseg(dictFileName string) (*Posseg, error) {
	j := &jiebago.Jieba{Total: 0.0, Freq: make(map[string]float64)}
	p := &Posseg{j, make(map[string]string)}
	dictFilePath, err := jiebago.DictPath(dictFileName)
	if err != nil {
		return nil, err
	}
	wtfs, err := jiebago.ParseDictFile(dictFilePath)

	for _, wtf := range wtfs {
		p.Add(wtf)
	}
	return p, nil
}

// Load user specified dictionary file.
func (p *Posseg) LoadUserDict(dictFilePath string) error {
	wtfs, err := jiebago.ParseDictFile(dictFilePath)
	if err != nil {
		return err
	}
	for _, wtf := range wtfs {
		p.Add(wtf)
	}
	return nil
}

func (p *Posseg) cutDetailInternal(sentence string) chan WordTag {
	result := make(chan WordTag)

	go func() {
		runes := []rune(sentence)
		_, posList := viterbi(runes)
		begin := 0
		next := 0
		for i, char := range runes {
			pos := posList[i].State
			switch pos {
			case 'B':
				begin = i
			case 'E':
				result <- WordTag{string(runes[begin : i+1]), posList[i].Tag}
				next = i + 1
			case 'S':
				result <- WordTag{string(char), posList[i].Tag}
				next = i + 1
			}
		}
		if next < len(runes) {
			result <- WordTag{string(runes[next:]), posList[next].Tag}
		}
		close(result)
	}()
	return result
}

func (p *Posseg) cutDetail(sentence string) chan WordTag {
	result := make(chan WordTag)

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
						result <- WordTag{x, "m"}
					case reEng.MatchString(x):
						result <- WordTag{x, "eng"}
					default:
						result <- WordTag{x, "x"}
					}
				}
			}
		}
		close(result)
	}()
	return result
}

type cutFunc func(sentence string) chan WordTag

func (p *Posseg) cutDAG(sentence string) chan WordTag {
	result := make(chan WordTag)

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
			if y-x == 1 {
				buf = append(buf, l_word...)
			} else {
				if len(buf) > 0 {
					if len(buf) == 1 {
						sbuf := string(buf)
						if tag, ok := p.Flag[sbuf]; ok {
							result <- WordTag{sbuf, tag}
						} else {
							result <- WordTag{sbuf, "x"}
						}
						buf = make([]rune, 0)
					} else {
						bufString := string(buf)
						if v, ok := p.Freq[bufString]; !ok || v == 0.0 {
							for t := range p.cutDetail(bufString) {
								result <- t
							}
						} else {
							for _, elem := range buf {
								selem := string(elem)
								if tag, ok := p.Flag[selem]; ok {
									result <- WordTag{string(elem), tag}
								} else {
									result <- WordTag{string(elem), "x"}
								}

							}
						}
						buf = make([]rune, 0)
					}
				}
				sl_word := string(l_word)
				if tag, ok := p.Flag[sl_word]; ok {
					result <- WordTag{sl_word, tag}
				} else {
					result <- WordTag{sl_word, "x"}
				}
			}
			x = y
		}

		if len(buf) > 0 {
			if len(buf) == 1 {
				sbuf := string(buf)
				if tag, ok := p.Flag[sbuf]; ok {
					result <- WordTag{sbuf, tag}
				} else {
					result <- WordTag{sbuf, "x"}
				}
			} else {
				bufString := string(buf)
				if v, ok := p.Freq[bufString]; !ok || v == 0.0 {
					for t := range p.cutDetail(bufString) {
						result <- t
					}
				} else {
					for _, elem := range buf {
						selem := string(elem)
						if tag, ok := p.Flag[selem]; ok {
							result <- WordTag{selem, tag}
						} else {
							result <- WordTag{selem, "x"}
						}
					}
				}
			}
		}
		close(result)
	}()
	return result
}

func (p *Posseg) cutDAGNoHMM(sentence string) chan WordTag {
	result := make(chan WordTag)

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
					result <- WordTag{string(buf), "eng"}
					buf = make([]rune, 0)
				}
				sl_word := string(l_word)
				if tag, ok := p.Flag[sl_word]; ok {
					result <- WordTag{sl_word, tag}
				} else {
					result <- WordTag{sl_word, "x"}
				}
				x = y
			}
		}
		if len(buf) > 0 {
			result <- WordTag{string(buf), "eng"}
			buf = make([]rune, 0)
		}
		close(result)
	}()
	return result
}

// Tags the POS of each word after segmentation, using labels compatible with
// ictclas.
func (p *Posseg) Cut(sentence string, HMM bool) chan WordTag {
	result := make(chan WordTag)
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
						result <- WordTag{x, "x"}
					} else {
						for _, xx := range x {
							s := string(xx)
							switch {
							case reNum.MatchString(s):
								result <- WordTag{s, "m"}
							case reEng.MatchString(x):
								result <- WordTag{x, "eng"}
								break
							default:
								result <- WordTag{s, "x"}
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
