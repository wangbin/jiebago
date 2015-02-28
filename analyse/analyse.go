package analyse

import (
	"fmt"
	"github.com/wangbin/jiebago"
	"sort"
	"strings"
	"unicode/utf8"
)

type TfIdf struct {
	Word string
	Freq float64
}

func (t TfIdf) String() string {
	return fmt.Sprintf("{%s: %f}", t.Word, t.Freq)
}

type TfIdfs []TfIdf

func (tis TfIdfs) Len() int {
	return len(tis)
}

func (tis TfIdfs) Less(i, j int) bool {
	if tis[i].Freq == tis[j].Freq {
		return tis[i].Word < tis[j].Word
	}

	return tis[i].Freq < tis[j].Freq
}

func (tis TfIdfs) Swap(i, j int) {
	tis[i], tis[j] = tis[j], tis[i]
}

func ExtractTags(sentence string, topK int) (tags TfIdfs) {
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
	tis := make(TfIdfs, 0)
	for k, v := range freq {
		var ti TfIdf
		if freq_, ok := loader.Freq[k]; ok {
			ti = TfIdf{Word: k, Freq: freq_ * v}
		} else {
			ti = TfIdf{Word: k, Freq: loader.Median * v}
		}
		tis = append(tis, ti)
	}
	sort.Sort(sort.Reverse(tis))
	if len(tis) > topK {
		tags = tis[:topK]
	} else {
		tags = tis
	}
	return tags
}
