package analyse

import (
	"sync"

	"github.com/wangbin/jiebago/dictionary"
)

var DefaultStopWordMap = map[string]int{
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

type StopWord struct {
	stopWordMap map[string]int
	sync.RWMutex
}

func (s *StopWord) AddToken(token dictionary.Token) {
	s.Lock()
	s.stopWordMap[token.Text()] = 1
	s.Unlock()
}

func NewStopWord() *StopWord {
	s := new(StopWord)
	s.stopWordMap = DefaultStopWordMap
	return s
}

func (s StopWord) IsStopWord(word string) bool {
	s.RLock()
	_, ok := s.stopWordMap[word]
	s.RUnlock()
	return ok
}

func (s *StopWord) Load(ch <-chan dictionary.Token) {
	s.Lock()
	for token := range ch {
		s.stopWordMap[token.Text()] = 1
	}
	s.Unlock()
}

func (s *StopWord) loadDictionary(fileName string) error {
	return dictionary.LoadDictionary(s, fileName)
}
