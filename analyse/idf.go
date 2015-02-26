package analyse

import (
	"github.com/wangbin/jiebago"
	"sort"
)

var (
	idfLoader *IDFLoader
)

func init() {
	idfLoader = NewIDFLoader()
}

type IDFLoader struct {
	Path   string
	Freq   map[string]float64
	Median float64
}

func NewIDFLoader() *IDFLoader {
	loader := new(IDFLoader)
	loader.Freq = make(map[string]float64)
	return loader
}

func (loader *IDFLoader) newPath(idfFilePath string) error {
	if loader.Path == idfFilePath {
		return nil
	}
	wtfs, err := jiebago.ParseDictFile(idfFilePath)
	if err != nil {
		return err
	}

	freqs := make([]float64, 0)

	for _, wtf := range wtfs {
		loader.Freq[wtf.Word] = wtf.Freq
		freqs = append(freqs, wtf.Freq)
	}

	sort.Float64s(freqs)
	loader.Median = freqs[len(freqs)/2]
	return nil
}

func SetIdf(idfFileName string) error {
	idfFilePath, err := jiebago.DictPath(idfFileName)
	if err != nil {
		return err
	}
	return idfLoader.newPath(idfFilePath)
}
