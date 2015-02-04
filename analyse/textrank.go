package analyse

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/wangbin/jiebago/posseg"
	"math"
	"sort"
)

const (
	DampingFactor = 0.85
)

var (
	defaultAllowPOS = []string{"ns", "n", "vn", "v"}
)

type edge struct {
	start  string
	end    string
	weight float64
}

type undirectWeightedGraph struct {
	graph map[string][]edge
}

func (u *undirectWeightedGraph) addEdge(start, end string, weight float64) {
	if _, ok := u.graph[start]; !ok {
		u.graph[start] = []edge{edge{start: start, end: end, weight: weight}}
	} else {
		u.graph[start] = append(u.graph[start], edge{start: start, end: end, weight: weight})
	}

	if _, ok := u.graph[end]; !ok {
		u.graph[start] = []edge{edge{start: end, end: start, weight: weight}}
	} else {
		u.graph[start] = append(u.graph[start], edge{start: end, end: start, weight: weight})
	}

}

func (u *undirectWeightedGraph) rank() TfIdfs {
	ws := make(map[string]float64)
	outSum := make(map[string]float64)

	wsdef := 1.0
	if len(u.graph) > 0 {
		wsdef /= float64(len(u.graph))
	}

	for n, out := range u.graph {
		ws[n] = wsdef
		sum := 0.0
		for _, e := range out {
			sum += e.weight
		}
		outSum[n] = sum
	}

	for x := 0; x < 10; x++ {
		for n, inedges := range u.graph {
			s := 0.0
			for _, e := range inedges {
				s += e.weight / outSum[e.end] * ws[e.start]
			}
			ws[n] = (1 - DampingFactor) + DampingFactor*s
		}
	}

	minRank := math.MaxFloat64
	maxRank := math.SmallestNonzeroFloat64
	for _, w := range ws {
		if w < minRank {
			minRank = w
		} else if w > maxRank {
			maxRank = w
		}
	}
	result := make(TfIdfs, 0)
	for n, w := range ws {
		result = append(result, TfIdf{Word: n, Freq: (w - minRank/10.0) / (maxRank - minRank/10.0)})
	}
	sort.Sort(sort.Reverse(result))
	return result
}

func TextRankWithPOS(sentence string, topK int, allowPOS []string) TfIdfs {
	posFilt := mapset.NewSet()
	for _, pos := range allowPOS {
		posFilt.Add(pos)
	}
	g := new(undirectWeightedGraph)
	cm := make(map[[2]string]float64)
	span := 5
	wordTags := posseg.Cut(sentence, true)
	for i := range wordTags {
		if posFilt.Contains(wordTags[i].Tag) {
			for j := i + 1; j < i+span; i++ {
				if j > len(wordTags) {
					break
				}
				if !posFilt.Contains(wordTags[j].Tag) {
					continue
				}
				if _, ok := cm[[2]string{wordTags[i].Word, wordTags[j].Word}]; !ok {
					cm[[2]string{wordTags[i].Word, wordTags[j].Word}] = 1.0
				} else {
					cm[[2]string{wordTags[i].Word, wordTags[j].Word}] += 1.0
				}
			}
		}
	}
	for startEnd, weight := range cm {
		g.addEdge(startEnd[0], startEnd[1], weight)
	}
	tags := g.rank()
	if topK > 0 && len(tags) > topK {
		tags = tags[:topK]
	}
	return tags
}

func TextRank(sentence string, topK int) TfIdfs {
	return TextRankWithPOS(sentence, topK, defaultAllowPOS)
}
