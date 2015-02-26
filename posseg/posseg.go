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

func __cut(sentence string) []WordTag {
	result := make([]WordTag, 0)
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
			result = append(result, WordTag{string(runes[begin : i+1]), posList[i].Tag})
			next = i + 1
		case 'S':
			result = append(result, WordTag{string(char), posList[i].Tag})
			next = i + 1
		}
	}
	if next < len(runes) {
		result = append(result, WordTag{string(runes[next:]), posList[next].Tag})
	}
	return result
}

func cutDetail(sentence string) []WordTag {
	result := make([]WordTag, 0)
	blocks := jiebago.RegexpSplit(reHanDetail, sentence)
	for _, blk := range blocks {
		if reHanDetail.MatchString(blk) {
			for _, wordTag := range __cut(blk) {
				result = append(result, wordTag)
			}
		} else {
			for _, x := range jiebago.RegexpSplit(reSkipDetail, blk) {
				if len(x) == 0 {
					continue
				}
				switch {
				case reNum.MatchString(x):
					result = append(result, WordTag{x, "m"})
				case reEng.MatchString(x):
					result = append(result, WordTag{x, "eng"})
				default:
					result = append(result, WordTag{x, "x"})
				}
			}
		}
	}

	return result
}

type cutAction func(sentence string) []WordTag

func cut_DAG(sentence string) []WordTag {
	dag := jiebago.GetDAG(sentence)
	routes := jiebago.Calc(sentence, dag)
	x := 0
	var y int
	runes := []rune(sentence)
	length := len(runes)
	result := make([]WordTag, 0)
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
						result = append(result, WordTag{sbuf, tag})
					} else {
						result = append(result, WordTag{sbuf, "x"})
					}
					buf = make([]rune, 0)
				} else {
					bufString := string(buf)
					if v, ok := jiebago.T.Freq[bufString]; !ok || v == 0.0 {
						recognized := cutDetail(bufString)
						for _, t := range recognized {
							result = append(result, t)
						}
					} else {
						for _, elem := range buf {
							selem := string(elem)
							if tag, ok := wordTagMap[selem]; ok {
								result = append(result, WordTag{string(elem), tag})
							} else {
								result = append(result, WordTag{string(elem), "x"})
							}

						}
					}
					buf = make([]rune, 0)
				}
			}
			sl_word := string(l_word)
			if tag, ok := wordTagMap[sl_word]; ok {
				result = append(result, WordTag{sl_word, tag})
			} else {
				result = append(result, WordTag{sl_word, "x"})
			}
		}
		x = y
	}

	if len(buf) > 0 {
		if len(buf) == 1 {
			sbuf := string(buf)
			if tag, ok := wordTagMap[sbuf]; ok {
				result = append(result, WordTag{sbuf, tag})
			} else {
				result = append(result, WordTag{sbuf, "x"})
			}
		} else {
			bufString := string(buf)
			if v, ok := jiebago.T.Freq[bufString]; !ok || v == 0.0 {
				recognized := cutDetail(bufString)
				for _, t := range recognized {
					result = append(result, t)
				}
			} else {
				for _, elem := range buf {
					selem := string(elem)
					if tag, ok := wordTagMap[selem]; ok {
						result = append(result, WordTag{selem, tag})
					} else {
						result = append(result, WordTag{selem, "x"})
					}
				}
			}
		}
	}
	return result
}

func cut_DAG_NO_HMM(sentence string) []WordTag {
	result := make([]WordTag, 0)
	dag := jiebago.GetDAG(sentence)
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
				result = append(result, WordTag{string(buf), "eng"})
				buf = make([]rune, 0)
			}
			sl_word := string(l_word)
			if tag, ok := wordTagMap[sl_word]; ok {
				result = append(result, WordTag{sl_word, tag})
			} else {
				result = append(result, WordTag{sl_word, "x"})
			}
			x = y
		}
	}
	if len(buf) > 0 {
		result = append(result, WordTag{string(buf), "eng"})
		buf = make([]rune, 0)
	}
	return result
}

func cut(sentence string, HMM bool) []WordTag {
	result := make([]WordTag, 0)
	blocks := jiebago.RegexpSplit(reHanInternal, sentence)
	var cut_block cutAction
	if HMM {
		cut_block = cut_DAG
	} else {
		cut_block = cut_DAG_NO_HMM
	}
	for _, blk := range blocks {
		if reHanInternal.MatchString(blk) {
			for _, wordTag := range cut_block(blk) {
				result = append(result, wordTag)
			}
		} else {
			for _, x := range jiebago.RegexpSplit(reSkipInternal, blk) {
				if reSkipInternal.MatchString(x) {
					result = append(result, WordTag{x, "x"})
				} else {
					for _, xx := range x {
						s := string(xx)
						switch {
						case reNum.MatchString(s):
							result = append(result, WordTag{s, "m"})
						case reEng.MatchString(x):
							result = append(result, WordTag{x, "eng"})
							break
						default:
							result = append(result, WordTag{s, "x"})
						}
					}
				}
			}
		}
	}
	return result
}

func Cut(sentence string, HMM bool) []WordTag {
	for key := range jiebago.UserWordTagTab {
		wordTagMap[key] = jiebago.UserWordTagTab[key]
		delete(jiebago.UserWordTagTab, key)
	}
	return cut(sentence, HMM)
}
