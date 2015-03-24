package analyse

import (
	"github.com/wangbin/jiebago"
	"sort"
)

type IDFLoader struct {
	IDFFreq map[string]float64
	Median  float64
}

func NewIDFLoader(IDFFileName string) (*IDFLoader, error) {
	IDFFilePath, err := jiebago.DictPath(IDFFileName)
	if err != nil {
		return nil, err
	}
	wtfs, err := jiebago.ParseDictFile(IDFFilePath)
	if err != nil {
		return nil, err
	}

	freqs := make([]float64, len(wtfs))
	loader := &IDFLoader{make(map[string]float64), 0.0}
	for index, wtf := range wtfs {
		loader.IDFFreq[wtf.Word] = wtf.Freq
		freqs[index] = wtf.Freq
	}
	sort.Float64s(freqs)
	loader.Median = freqs[len(freqs)/2]
	return loader, nil
}
