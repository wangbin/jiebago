package analyse

import (
	"github.com/wangbin/jiebago"
)

var stopWords map[string]int

func init() {
	stopWords = map[string]int{
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
}

// Set the stop words file path, could be absolute path of stop words file, or
// file name in current directory.
func SetStopWords(stopWordsFileName string) error {
	stopWordsFilePath, err := jiebago.DictPath(stopWordsFileName)
	if err != nil {
		return err
	}

	wtfs, err := jiebago.ParseDictFile(stopWordsFilePath)
	for _, wtf := range wtfs {
		stopWords[wtf.Word] = 1
	}
	return nil
}
