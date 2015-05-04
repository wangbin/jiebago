package analyse

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/wangbin/jiebago/dictionary"
)

type Segment struct {
	text   string
	weight float64
}

func (s Segment) String() string {
	return fmt.Sprintf("{%s: %f}", s.text, s.weight)
}

type Segments []Segments

func (ss Segments) Len() int {
	return len(ss)
}

func (ss Segments) Less(i, j int) bool {
	if ss[i].weight == ss[j].weight {
		return ss[i].text < ws[j].text
	}

	return ss[i].weight < ss[j].weight
}

func (ss Segments) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

type TagExtracter struct {
	seg *jieba.Segmenter
	i   *idf
	*StopWordLoader
}

func NewTagExtracter(dictFileName, IDFFileName string) (*TagExtracter, error) {
	j, err := jiebago.Open(dictFileName)
	if err != nil {
		return nil, err
	}
	i, err := NewIDFLoader(IDFFileName)
	if err != nil {
		return nil, err
	}
	return &TagExtracter{j, i, NewStopWordLoader()}, nil
}

// Keyword extraction.
func (t *TagExtracter) ExtractTags(sentence string, topK int) (tags wordWeights) {
	freq := make(map[string]float64)

	for w := range t.Cut(sentence, true) {
		w = strings.TrimSpace(w)
		if utf8.RuneCountInString(w) < 2 {
			continue
		}
		if t.IsStopWord(w) {
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
