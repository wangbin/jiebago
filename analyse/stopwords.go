package analyse

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

var stopWords map[string]string

func init() {
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
