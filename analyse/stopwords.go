package analyse

import "github.com/wangbin/jiebago"

var defaultStopWords = map[string]int{
	"the":   1,
	"of":    1,
	"is":    1,
	"and":   1,
	"to":    1,
	"in":    1,
	"that":  1,
	"we":    1,
	"for":   1,
	"an":    1,
	"are":   1,
	"by":    1,
	"be":    1,
	"as":    1,
	"on":    1,
	"with":  1,
	"can":   1,
	"if":    1,
	"from":  1,
	"which": 1,
	"you":   1,
	"it":    1,
	"this":  1,
	"then":  1,
	"at":    1,
	"have":  1,
	"all":   1,
	"not":   1,
	"one":   1,
	"has":   1,
	"or":    1,
}

type StopWordLoader struct {
	stopWords map[string]int
}

func (s *StopWordLoader) AddEntry(entry *jiebago.Entry) {
	s.stopWords[entry.Word] = 1
}

func NewStopWordLoader() *StopWordLoader {
	s := new(StopWordLoader)
	s.stopWords = defaultStopWords
	return s
}

// Set the stop words file path, could be absolute path of stop words file, or
// file name in current directory.
func (s *StopWordLoader) SetStopWords(stopWordsFileName string) error {
	return jiebago.LoadDict(s, stopWordsFileName, false)
}

func (s StopWordLoader) IsStopWord(word string) bool {
	_, ok := s.stopWords[word]
	return ok
}
