package analyse

import (
	"sort"
	"sync"

	"github.com/wangbin/jiebago/dictionary"
)

type Idf struct {
	freqMap map[string]float64
	median  float64
	freqs   []float64
	sync.RWMutex
}

func (i *Idf) AddToken(token dictionary.Token) {
	i.Lock()
	i.freqMap[token.Text()] = token.Frequency()
	i.freqs = append(i.freqs, token.Frequency())
	sort.Float64s(i.freqs)
	i.median = i.freqs[len(i.freqs)/2]
	i.Unlock()
}

func (i *Idf) Load(ch <-chan dictionary.Token) {
	i.Lock()
	for token := range ch {
		i.freqMap[token.Text()] = token.Frequency()
		i.freqs = append(i.freqs, token.Frequency())
	}
	sort.Float64s(i.freqs)
	i.median = i.freqs[len(i.freqs)/2]
	i.Unlock()
}

func (i *Idf) loadDictionary(fileName string) error {
	return dictionary.LoadDictionary(i, fileName)
}

func (i Idf) Frequency(key string) (float64, bool) {
	i.RLock()
	freq, ok := i.freqMap[key]
	i.RUnlock()
	return freq, ok
}

func NewIdf() *Idf {
	return &Idf{freqMap: make(map[string]float64), freqs: make([]float64, 0)}
}
