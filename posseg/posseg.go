package posseg

import (
	"github.com/wangbin/jiebago"
	"regexp"
)

var (
	wordTagMap     = make(map[string]string)
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

func SetDictionary(dictFileName string) error {
	err := jiebago.SetDictionary(dictFileName)
	if err != nil {
		return err
	}
	dictFilePath, err := jiebago.DictPath(dictFileName)
	if err != nil {
		return err
	}
	wtfs, err := jiebago.ParseDictFile(dictFilePath)

	for _, wtf := range wtfs {
		wordTagMap[wtf.Word] = wtf.Tag
	}
	return nil
}

func cutDetailInternal(sentence string) chan WordTag {
	result := make(chan WordTag)

	go func() {
		runes := []rune(sentence)
		_, posList := Viterbi(runes)
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

func cutDetail(sentence string) chan WordTag {
	result := make(chan WordTag)

	go func() {
		blocks := jiebago.RegexpSplit(reHanDetail, sentence)
		for _, blk := range blocks {
			if reHanDetail.MatchString(blk) {
				for wordTag := range cutDetailInternal(blk) {
					result <- wordTag
				}
			} else {
				for _, x := range jiebago.RegexpSplit(reSkipDetail, blk) {
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

func cutDAG(sentence string) chan WordTag {
	result := make(chan WordTag)

	go func() {
		dag := jiebago.DAG(sentence)
		routes := jiebago.Calc(sentence, dag)
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
						if tag, ok := wordTagMap[sbuf]; ok {
							result <- WordTag{sbuf, tag}
						} else {
							result <- WordTag{sbuf, "x"}
						}
						buf = make([]rune, 0)
					} else {
						bufString := string(buf)
						if v, ok := jiebago.Trie.Freq[bufString]; !ok || v == 0.0 {
							for t := range cutDetail(bufString) {
								result <- t
							}
						} else {
							for _, elem := range buf {
								selem := string(elem)
								if tag, ok := wordTagMap[selem]; ok {
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
				if tag, ok := wordTagMap[sl_word]; ok {
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
				if tag, ok := wordTagMap[sbuf]; ok {
					result <- WordTag{sbuf, tag}
				} else {
					result <- WordTag{sbuf, "x"}
				}
			} else {
				bufString := string(buf)
				if v, ok := jiebago.Trie.Freq[bufString]; !ok || v == 0.0 {
					for t := range cutDetail(bufString) {
						result <- t
					}
				} else {
					for _, elem := range buf {
						selem := string(elem)
						if tag, ok := wordTagMap[selem]; ok {
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

func cutDAGNoHMM(sentence string) chan WordTag {
	result := make(chan WordTag)

	go func() {
		dag := jiebago.DAG(sentence)
		routes := jiebago.Calc(sentence, dag)
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
				if tag, ok := wordTagMap[sl_word]; ok {
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

func Cut(sentence string, HMM bool) chan WordTag {
	for key := range jiebago.UserWordTagTab {
		wordTagMap[key] = jiebago.UserWordTagTab[key]
		delete(jiebago.UserWordTagTab, key)
	}
	result := make(chan WordTag)
	blocks := jiebago.RegexpSplit(reHanInternal, sentence)
	var cut cutFunc
	if HMM {
		cut = cutDAG
	} else {
		cut = cutDAGNoHMM
	}
	go func() {
		for _, blk := range blocks {
			if reHanInternal.MatchString(blk) {
				for wordTag := range cut(blk) {
					result <- wordTag
				}
			} else {
				for _, x := range jiebago.RegexpSplit(reSkipInternal, blk) {
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
