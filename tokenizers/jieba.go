package tokenizers

import (
	"fmt"
	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/registry"
	"github.com/wangbin/jiebago"
	"regexp"
	"strconv"
)

const Name = "jieba"

var IdeographRegexp = regexp.MustCompile(`\p{Han}+`)

type JiebaTokenizer struct {
	dictFileName    string
	hmm, searchMode bool
}

func NewJiebaTokenizer(dictFileName string, hmm, searchMode bool) (analysis.Tokenizer, error) {
	err := jiebago.SetDictionary(dictFileName)
	return &JiebaTokenizer{
		dictFileName: dictFileName,
		hmm:          hmm,
		searchMode:   searchMode,
	}, err
}

func (jt *JiebaTokenizer) Tokenize(input []byte) analysis.TokenStream {
	rv := make(analysis.TokenStream, 0)
	runeStart := 0
	start := 0
	end := 0
	pos := 1
	var width int
	var gram string
	for word := range jiebago.Cut(string(input), false, jt.hmm) {
		if jt.searchMode {
			runes := []rune(word)
			width = len(runes)
			for _, step := range [2]int{2, 3} {
				if width > step {
					for i := 0; i < width-step+1; i++ {
						gram = string(runes[i : i+step])
						gramLen := len(gram)
						if value, ok := jiebago.Trie.Freq[gram]; ok && value > 0 {
							gramStart := start + len(string(runes[:i]))
							token := analysis.Token{
								Term:     []byte(gram),
								Start:    gramStart,
								End:      gramStart + gramLen,
								Position: pos,
								Type:     detectTokenType(gram),
							}
							rv = append(rv, &token)
							pos++
						}
					}
				}
			}
		}
		end = start + len(word)
		token := analysis.Token{
			Term:     []byte(word),
			Start:    start,
			End:      end,
			Position: pos,
			Type:     detectTokenType(word),
		}
		rv = append(rv, &token)
		pos++
		runeStart += width
		start = end
	}
	return rv
}

func JiebaTokenizerConstructor(config map[string]interface{}, cache *registry.Cache) (
	analysis.Tokenizer, error) {
	dictFileName, ok := config["file"].(string)
	if !ok {
		return nil, fmt.Errorf("must specify dictionary file path")
	}
	hmm, ok := config["hmm"].(bool)
	if !ok {
		hmm = true
	}
	searchMode, ok := config["search"].(bool)
	if !ok {
		searchMode = true
	}

	return NewJiebaTokenizer(dictFileName, hmm, searchMode)
}

func detectTokenType(term string) analysis.TokenType {
	if IdeographRegexp.MatchString(term) {
		return analysis.Ideographic
	}
	_, err := strconv.ParseFloat(term, 64)
	if err == nil {
		return analysis.Numeric
	}
	return analysis.AlphaNumeric
}

func init() {
	registry.RegisterTokenizer(Name, JiebaTokenizerConstructor)
}
