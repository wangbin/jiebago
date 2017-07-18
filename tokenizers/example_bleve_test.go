package tokenizers_test

import (
	"fmt"
	"log"
	"os"

	"github.com/blevesearch/bleve"
	_ "github.com/wangbin/jiebago/tokenizers"
)

func Example_beleveSearch() {
	// open a new index
	indexMapping := bleve.NewIndexMapping()

	err := indexMapping.AddCustomTokenizer("jieba",
		map[string]interface{}{
			"file": "../dict.txt",
			"type": "jieba",
		})
	if err != nil {
		log.Fatal(err)
	}

	// create a custom analyzer
	err = indexMapping.AddCustomAnalyzer("jieba",
		map[string]interface{}{
			"type":      "custom",
			"tokenizer": "jieba",
			"token_filters": []string{
				"possessive_en",
				"to_lower",
				"stop_en",
			},
		})

	if err != nil {
		log.Fatal(err)
	}

	indexMapping.DefaultAnalyzer = "jieba"
	cacheDir := "jieba.beleve"
	os.RemoveAll(cacheDir)
	index, err := bleve.New(cacheDir, indexMapping)

	if err != nil {
		log.Fatal(err)
	}

	docs := []struct {
		Title string
		Name  string
	}{
		{
			Title: "Doc 1",
			Name:  "This is the first document we’ve added",
		},
		{
			Title: "Doc 2",
			Name:  "The second one 你 中文测试中文 is even more interesting! 吃水果",
		},
		{
			Title: "Doc 3",
			Name:  "买水果然后来世博园。",
		},
		{
			Title: "Doc 4",
			Name:  "工信处女干事每月经过下属科室都要亲口交代24口交换机等技术性器件的安装工作",
		},
		{
			Title: "Doc 5",
			Name:  "咱俩交换一下吧。",
		},
	}
	// index docs
	for _, doc := range docs {
		index.Index(doc.Title, doc)
	}

	// search for some text
	for _, keyword := range []string{"水果世博园", "你", "first", "中文", "交换机", "交换"} {
		query := bleve.NewQueryStringQuery(keyword)
		search := bleve.NewSearchRequest(query)
		search.Highlight = bleve.NewHighlight()
		searchResults, err := index.Search(search)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Result of \"%s\": %d matches:\n", keyword, searchResults.Total)
		for i, hit := range searchResults.Hits {
			rv := fmt.Sprintf("%d. %s, (%f)\n", i+searchResults.Request.From+1, hit.ID, hit.Score)
			for fragmentField, fragments := range hit.Fragments {
				rv += fmt.Sprintf("%s: ", fragmentField)
				for _, fragment := range fragments {
					rv += fmt.Sprintf("%s", fragment)
				}
			}
			fmt.Printf("%s\n", rv)
		}
	}
	// Output:
	// Result of "水果世博园": 2 matches:
	// 1. Doc 3, (1.099550)
	// Name: 买<mark>水果</mark>然后来<mark>世博</mark>园。
	// 2. Doc 2, (0.031941)
	// Name: The second one 你 中文测试中文 is even more interesting! 吃<mark>水果</mark>
	// Result of "你": 1 matches:
	// 1. Doc 2, (0.391161)
	// Name: The second one <mark>你</mark> 中文测试中文 is even more interesting! 吃水果
	// Result of "first": 1 matches:
	// 1. Doc 1, (0.512150)
	// Name: This is the <mark>first</mark> document we’ve added
	// Result of "中文": 1 matches:
	// 1. Doc 2, (0.553186)
	// Name: The second one 你 <mark>中文</mark>测试<mark>中文</mark> is even more interesting! 吃水果
	// Result of "交换机": 2 matches:
	// 1. Doc 4, (0.608495)
	// Name: 工信处女干事每月经过下属科室都要亲口交代24口<mark>交换机</mark>等技术性器件的安装工作
	// 2. Doc 5, (0.086700)
	// Name: 咱俩<mark>交换</mark>一下吧。
	// Result of "交换": 2 matches:
	// 1. Doc 5, (0.534158)
	// Name: 咱俩<mark>交换</mark>一下吧。
	// 2. Doc 4, (0.296297)
	// Name: 工信处女干事每月经过下属科室都要亲口交代24口<mark>交换</mark>机等技术性器件的安装工作
}
