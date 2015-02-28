package analyse

import (
	"github.com/wangbin/jiebago"
	"sort"
)

var (
	loader *idfLoader
)

func init() {
	loader = newIDFLoader()
}

type idfLoader struct {
	Path   string
	Freq   map[string]float64
	Median float64
}

func newIDFLoader() *idfLoader {
	loader := new(idfLoader)
	loader.Freq = make(map[string]float64)
	return loader
}

func (loader *idfLoader) newPath(idfFilePath string) error {
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

// Set the IDF file path, could be absolute path of IDF file, or IDF file
// name in current directory.
func SetIdf(idfFileName string) error {
	idfFilePath, err := jiebago.DictPath(idfFileName)
	if err != nil {
		return err
	}
	return loader.newPath(idfFilePath)
}
