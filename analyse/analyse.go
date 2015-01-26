package analyse

import (
	"github.com/wangbin/jiebago"
	"sort"
	"strings"
	"unicode/utf8"
)

type TfIdf struct {
	word string
	freq float64
}

type TfIdfs []TfIdf

func (tis TfIdfs) Len() int {
	return len(tis)
}

func (tis TfIdfs) Less(i, j int) bool {
	if tis[i].freq == tis[j].freq {
		return tis[i].word < tis[j].word
	}
	return tis[i].freq < tis[j].freq
}

func (tis TfIdfs) Swap(i, j int) {
	tis[i], tis[j] = tis[j], tis[i]
}

func ExtractTags(sentence string, topK int) []string {
	words := jiebago.Cut(sentence, false, true)
	freq := make(map[string]float64)

	for _, w := range words {
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
		if freq_, ok := idfLoader.Freq[k]; ok {
			ti = TfIdf{word: k, freq: freq_ * v}
		} else {
			ti = TfIdf{word: k, freq: idfLoader.Median * v}
		}
		tis = append(tis, ti)
	}
	sort.Sort(sort.Reverse(tis))
	var topTfIdfs TfIdfs
	if len(tis) > topK {
		topTfIdfs = tis[:topK]
	} else {
		topTfIdfs = tis
	}
	tags := make([]string, len(topTfIdfs))
	for index, ti := range topTfIdfs {
		tags[index] = ti.word
	}
	return tags
}
