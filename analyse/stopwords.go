package analyse

import (
	"github.com/wangbin/jiebago"
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

func SetStopWords(stopWordsFileName string) error {
	stopWordsFilePath, err := jiebago.DictPath(stopWordsFileName)
	if err != nil {
		return err
	}

	wtfs, err := jiebago.ParseDictFile(stopWordsFilePath)
	for _, wtf := range wtfs {
		stopWords[wtf.Word] = wtf.Word
	}
	return nil
}
