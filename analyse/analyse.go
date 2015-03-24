package analyse

import (
	"fmt"
	"github.com/wangbin/jiebago"
	"sort"
	"strings"
	"unicode/utf8"
)

type wordWeight struct {
	Word   string
	Weight float64
}

func (w wordWeight) String() string {
	return fmt.Sprintf("{%s: %f}", w.Word, w.Weight)
}

type wordWeights []wordWeight

func (ws wordWeights) Len() int {
	return len(ws)
}

func (ws wordWeights) Less(i, j int) bool {
	if ws[i].Weight == ws[j].Weight {
		return ws[i].Word < ws[j].Word
	}

	return ws[i].Weight < ws[j].Weight
}

func (ws wordWeights) Swap(i, j int) {
	ws[i], ws[j] = ws[j], ws[i]
}

type TagExtracter struct {
	*jiebago.Jieba
	*IDFLoader
	stopWords map[string]int
}

func NewTagExtracter(dictFileName, IDFFileName string) (*TagExtracter, error) {
	j, err := jiebago.NewJieba(dictFileName)
	if err != nil {
		return nil, err
	}
	i, err := NewIDFLoader(IDFFileName)
	if err != nil {
		return nil, err
	}
	return &TagExtracter{j, i, StopWords}, nil
}

// Set the stop words file path, could be absolute path of stop words file, or
// file name in current directory.
func (t *TagExtracter) SetStopWords(stopWordsFileName string) error {
	stopWordsFilePath, err := jiebago.DictPath(stopWordsFileName)
	if err != nil {
		return err
	}

	wtfs, err := jiebago.ParseDictFile(stopWordsFilePath)
	for _, wtf := range wtfs {
		t.stopWords[wtf.Word] = 1
	}
	return nil
}

// Keyword extraction.
func (t *TagExtracter) ExtractTags(sentence string, topK int) (tags wordWeights) {
	freq := make(map[string]float64)

	for w := range t.Cut(sentence, false, true) {
		w = strings.TrimSpace(w)
		if utf8.RuneCountInString(w) < 2 {
			continue
		}
		if _, ok := t.stopWords[w]; ok {
			continue
		}
		if f, ok := freq[w]; ok {
			freq[w] = f + 1.0
		} else {
			freq[w] = 1.0
		}
	}
	total := 0.0
	for _, f := range freq {
		total += f
	}
	for k, v := range freq {
		freq[k] = v / total
	}
	ws := make(wordWeights, 0)
	for k, v := range freq {
		var ti wordWeight
		if freq_, ok := t.IDFFreq[k]; ok {
			ti = wordWeight{Word: k, Weight: freq_ * v}
		} else {
			ti = wordWeight{Word: k, Weight: t.Median * v}
		}
		ws = append(ws, ti)
	}
	sort.Sort(sort.Reverse(ws))
	if len(ws) > topK {
		tags = ws[:topK]
	} else {
		tags = ws
	}
	return tags
}
