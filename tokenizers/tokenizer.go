package tokenizers

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/registry"
	"github.com/wangbin/jiebago"
)

// Name is the jieba tokenizer name.
const Name = "jieba"

var ideographRegexp = regexp.MustCompile(`\p{Han}+`)

// JiebaTokenizer is the beleve tokenizer for jiebago.
type JiebaTokenizer struct {
	seg             jiebago.Segmenter
	hmm, searchMode bool
}

/*
NewJiebaTokenizer creates a new JiebaTokenizer.

Parameters:

    dictFilePath: path of the dictioanry file.

    hmm: whether to use Hidden Markov Model to cut unknown words,
    i.e. not found in dictionary. For example word "安卓" (means "Android" in
    English) not in the dictionary file. If hmm is set to false, it will be
    cutted into two single words "安" and "卓", if hmm is set to true, it will
    be traded as one single word because Jieba using Hidden Markov Model with
    Viterbi algorithm to guess the best possibility.

    searchMode: whether to further cut long words into serveral short words.
    In Chinese, some long words may contains other words, for example "交换机"
    is a Chinese word for "Switcher", if sechMode is false, it will trade
    "交换机" as a single word. If searchMode is true, it will further split
    this word into "交换", "换机", which are valid Chinese words.
*/
func NewJiebaTokenizer(dictFilePath string, hmm, searchMode bool) (analysis.Tokenizer, error) {
	var seg jiebago.Segmenter
	err := seg.LoadDictionary(dictFilePath)
	return &JiebaTokenizer{
		seg:        seg,
		hmm:        hmm,
		searchMode: searchMode,
	}, err
}

// Tokenize cuts input into bleve token stream.
func (jt *JiebaTokenizer) Tokenize(input []byte) analysis.TokenStream {
	rv := make(analysis.TokenStream, 0)
	runeStart := 0
	start := 0
	end := 0
	pos := 1
	var width int
	var gram string
	for word := range jt.seg.Cut(string(input), jt.hmm) {
		if jt.searchMode {
			runes := []rune(word)
			width = len(runes)
			for _, step := range [2]int{2, 3} {
				if width <= step {
					continue
				}
				for i := 0; i < width-step+1; i++ {
					gram = string(runes[i : i+step])
					gramLen := len(gram)
					if frequency, ok := jt.seg.Frequency(gram); ok && frequency > 0 {
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

/*
JiebaTokenizerConstructor creates a JiebaTokenizer.

Parameter config should contains at least one parameter:

    file: the path of the dictionary file.

    hmm: optional, specify whether to use Hidden Markov Model, see NewJiebaTokenizer for details.

    search: optional, speficy whether to use search mode, see NewJiebaTokenizer for details.
*/
func JiebaTokenizerConstructor(config map[string]interface{}, cache *registry.Cache) (
	analysis.Tokenizer, error) {
	dictFilePath, ok := config["file"].(string)
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

	return NewJiebaTokenizer(dictFilePath, hmm, searchMode)
}

func detectTokenType(term string) analysis.TokenType {
	if ideographRegexp.MatchString(term) {
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
