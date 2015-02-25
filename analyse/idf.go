package analyse

import (
	"bufio"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
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
	idfFile, err := os.Open(idfFilePath)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(idfFile)
	freqs := make([]float64, 0)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		word, freqStr := words[0], words[1]
		freq, err := strconv.ParseFloat(freqStr, 64)
		if err != nil {
			continue
		}
		loader.Freq[word] = freq
		freqs = append(freqs, freq)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	sort.Float64s(freqs)
	loader.Median = freqs[len(freqs)/2]
	return nil

}

func SetIdf(idfFilePath string) error {
	if !filepath.IsAbs(idfFilePath) {
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		idfFilePath = filepath.Clean(filepath.Join(pwd, idfFilePath))
	}
	return idfLoader.newPath(idfFilePath)
}
