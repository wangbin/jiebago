package analyse

import (
	"github.com/wangbin/jiebago"
	"sort"
)

type idf struct {
	freqMap map[string]float64
	median  float64
	freqs   []float64
}

func (l *IDFLoader) AddEntry(entry jiebago.Entry) {
	l.IDFFreq[entry.Word] = entry.Freq
	l.freqs = append(l.freqs, entry.Freq)
}

func NewIDFLoader(IDFFileName string) (*IDFLoader, error) {
	loader := &IDFLoader{make(map[string]float64), 0.0, make([]float64, 0)}
	err := jiebago.LoadDict(loader, IDFFileName, false)
	if err != nil {
		return nil, err
	}

	sort.Float64s(loader.freqs)
	loader.Median = loader.freqs[len(loader.freqs)/2]
	loader.freqs = []float64{}
	return loader, nil
}
