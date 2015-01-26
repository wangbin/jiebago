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
	stopWords map[string]string
	idfLoader *IDFLoader
)

func init() {
	idfLoader = NewIDFLoader()
	stopWords = map[string]string{
		"the":   "the",
		"of":    "of",
		"is":    "is",
		"and":   "and",
		"to":    "to",
		"in":    "in",
		"that":  "that",
		"we":    "we",
		"for":   "for",
		"an":    "an",
		"are":   "are",
		"by":    "bye",
		"be":    "be",
		"as":    "as",
		"on":    "on",
		"with":  "with",
		"can":   "can",
		"if":    "of",
		"from":  "from",
		"which": "which",
		"you":   "you",
		"it":    "it",
		"this":  "this",
		"then":  "then",
		"at":    "at",
		"have":  "have",
		"all":   "all",
		"not":   "not",
		"one":   "one",
		"has":   "has",
		"or":    "or",
	}
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

func (loader *IDFLoader) NewPath(idfFilePath string) error {
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
	return idfLoader.NewPath(idfFilePath)
}

func SetStopWords(stopWordsFilePath string) error {
	if !filepath.IsAbs(stopWordsFilePath) {
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		stopWordsFilePath = filepath.Clean(filepath.Join(pwd, stopWordsFilePath))
	}
	stopWordsFile, err := os.Open(stopWordsFilePath)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(stopWordsFile)
	for scanner.Scan() {
		stopWord := scanner.Text()
		stopWord = strings.TrimSpace(stopWord)
		stopWords[stopWord] = stopWord
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
