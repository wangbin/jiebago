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

// Keyword extraction.
func ExtractTags(sentence string, topK int) (tags wordWeights) {
	freq := make(map[string]float64)

	for w := range jiebago.Cut(sentence, false, true) {
		w = strings.TrimSpace(w)
		if utf8.RuneCountInString(w) < 2 {
			continue
		}
		if _, ok := stopWords[w]; ok {
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
		if freq_, ok := loader.Freq[k]; ok {
			ti = wordWeight{Word: k, Weight: freq_ * v}
		} else {
			ti = wordWeight{Word: k, Weight: loader.Median * v}
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
