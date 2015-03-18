#结巴分词 Go 语言版：jiebago


[![Build Status](https://travis-ci.org/wangbin/jiebago.png?branch=master)](https://travis-ci.org/wangbin/jiebago)

[结巴分词](https://github.com/fxsjy/jieba)是[@fxsjy](https://github.com/fxsjy)用Python编写的中文分词组件，jiebago是结巴分词的Go语言实现，目前已经实现的功能包括：三种模式分词、自定义词典、关键词提取和词性标注。


## 安装


	go get github.com/wangbin/jiebago/...
	
## 分词


    package main

    import (
        "fmt"
        "github.com/wangbin/jiebago"
    )

    var sentence = "我来到北京清华大学"

    func print(ch chan string) {
        for word := range ch {
            fmt.Printf("%s / ", word)
        }
        fmt.Println()
        fmt.Println()
    }

    func main() {
        jiebago.SetDictionary("/Path/to/dictionary/file") // 设定字典
        fmt.Print("【全模式】： ")
        print(jiebago.Cut(sentence, true, true))
        fmt.Print("【精确模式】： ")
        print(jiebago.Cut(sentence, false, true))
        fmt.Print("【新词识别】：")
        print(jiebago.Cut("他来到了网易杭研大厦", false, true))
        fmt.Print("【搜索引擎模式】：")
        print(jiebago.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造", true))
    }
    
使用结巴分词自带的[词典文件](https://github.com/fxsjy/jieba/blob/master/jieba/dict.txt)，输出结果如下：

    【全模式】： 我 / 来到 / 北京 / 清华 / 清华大学 / 华大 / 大学 /

    【精确模式】： 我 / 来到 / 北京 / 清华大学 /

    【新词识别】：他 / 来到 / 了 / 网易 / 杭研 / 大厦 /

    【搜索引擎模式】：小明 / 硕士 / 毕业 / 于 / 中国 / 科学 / 学院 / 科学院 / 中国科学院 / 计算 / 计算所 / ， / 后 / 在 / 日本 / 京都 / 大学 / 日本京都大学 / 深造 /

## 添加自定义词典


    var sentence = "李小福是创新办主任也是云计算方面的专家"
	fmt.Print("Before: ")
	print(jiebago.Cut(sentence, false, true))
    jiebago.LoadUserDict("/Path/to/user/dictionary/file")
	fmt.Print("After: ")
	print(jiebago.Cut(sentence, false, true))

使用结巴分词自带的[词典文件](https://github.com/fxsjy/jieba/blob/master/jieba/dict.txt)和[用户自定义词典文件](https://github.com/fxsjy/jieba/blob/master/test/userdict.txt)，结果输出如下：

    Before: 李小福 / 是 / 创新 / 办 / 主任 / 也 / 是 / 云 / 计算 / 方面 / 的 / 专家 /

    After: 李小福 / 是 / 创新办 / 主任 / 也 / 是 / 云计算 / 方面 / 的 / 专家 /

## 关键词提取
    
示例代码：

    package main

    import (
        "fmt"
        "github.com/wangbin/jiebago/analyse"
    )

    var sentence = "这是一个伸手不见五指的黑夜。我叫孙悟空，我爱北京，我爱Python和C++。"

    func main() {
        analyse.SetDictionary("/Path/to/dictionary/file")
        analyse.SetIdf("/Path/to/idf/file")
        for _, ww := range analyse.ExtractTags(sentence, 20) {
           fmt.Printf("%s / ", ww.Word)
        }
    }
    
输出：

    Python / C++ / 伸手不见五指 / 孙悟空 / 黑夜 / 北京 / 这是 / 一个 /

## 基于TextRank算法的关键词抽取实现

示例代码：

    package main

    import (
        "fmt"
        "github.com/wangbin/jiebago/analyse"
    )

    func main() {
        sentence := "此外，公司拟对全资子公司吉林欧亚置业有限公司增资4.3亿元，增资后，吉林欧亚     置业注册资本由7000万元增加到5亿元。吉林欧亚置业主要经营范围为房地产开发及百货零售等业务。目前在建吉林欧亚城市商业综合体项目。2013年，实现营业收入0万元，实现净利润-139.13万元。"

        analyse.SetDictionary("/Path/to/dictionary/file")
        result := analyse.TextRank(sentence, 10)
        for _, wt := range result {
            fmt.Printf("%s %f\n", wt.Word, wt.Freq)
        }
    }

输出：

    吉林 1.000000
    欧亚 0.878078
    置业 0.562048
    实现 0.520906
    收入 0.384284
    增资 0.360591
    子公司 0.353132
    城市 0.307509
    全资 0.306324
    商业 0.306138    

## 词性标注
    
示例代码：

    package main

    import (
        "fmt"
        "github.com/wangbin/jiebago"
        "github.com/wangbin/jiebago/posseg"
    )

    var sentence = "我爱北京天安门"

    func main() {
        posseg.SetDictionary("/Path/to/dictionary/file")
        for wt := range posseg.Cut(sentence, true) {
            fmt.Printf("%s %s\n", wt.Word, wt.Tag)
        }
    }
    
输出：

    我 r
    爱 v
    北京 ns
    天安门 ns
    

## 并行分词

因为Go有强大的goroutine特性，并行分词实现起来非常简单，所以并没有内置到jiebaogo中，而是由使用者自己实现，下面是一个简单的例子：

    lineCount := 0
    inputFile, _ := os.Open(FileName)
    defer inputFile.Close()
    scanner := bufio.NewScanner(inputFile)
    ch := make(chan []string, 1)
    for scanner.Scan() {
        line := scanner.Text()
        fileLength += len([]rune(line))
        lineCount += 1
        go func() {
           for word := range jiebago.Cut(line, false, true) {
              ch <- word
           }
        }()
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }
    outputFile, _ := os.OpenFile("parallelCut.log", os.O_CREATE|os.O_WRONLY, 0600)
    defer outputFile.Close()
    writer := bufio.NewWriter(outputFile)
    results := make([]string, 0)
    for {
        if lineCount <= 0 {
            break
        }
        result, ok := <-ch
        if ok {
            results = append(results, result...)
            lineCount -= 1
        }
    }
    writer.WriteString(strings.Join(results, "/ "))
    writer.Flush()


## Tokenize：返回词语在原文的起始位置


注意新版的 Jiebago Tokenizer 实现了 Bleve 的 Tokenizer 接口，跟之前的实现有很大的变化：

1. 接受的参数必须是 []byte。
2. 输出的 Token 的起始和终止位置是 byte 的位置，不是之前的 rune 的位置，所以和 Python 版的 Jieba.tokenize 输出不一致。

```
package main

import (
    "fmt"
    "github.com/wangbin/jiebago/tokenizers"
)

const DictPath = "/path/to/dict.txt"

var sentence = []byte("永和服装饰品有限公司")

func main() {
    // default mode
    tokenizer, _ := tokenizers.NewJiebaTokenizer(DictPath, true, false)     for _, token := range tokenizer.Tokenize(sentence) {
        fmt.Printf(
            "Term: %s\t  Start: %d \t  End: %d\t Position: %d\t Type: %d\n",
            token.Term, token.Start, token.End, token.Position, token.Type)
    }
    
    //search mode
    tokenizer, _ = tokenizers.NewJiebaTokenizer(DictPath, true, true) 
    for _, token := range tokenizer.Tokenize(sentence) {
        fmt.Printf(
            "Term: %s\t  Start: %d \t  End: %d\t Position: %d\t Type: %d\n",
            token.Term, token.Start, token.End, token.Position, token.Type)
    }

}

```
默认模式输出：

```
Term: 永和        Start: 0        End: 6         Position: 1     Type: 1
Term: 服装        Start: 6        End: 12        Position: 2     Type: 1
Term: 饰品        Start: 12       End: 18        Position: 3     Type: 1
Term: 有限公司    Start: 18       End: 30        Position: 4     Type: 1
```
搜索模式输出：

```
Term: 永和        Start: 0        End: 6         Position: 1     Type: 1
Term: 服装        Start: 6        End: 12        Position: 2     Type: 1
Term: 饰品        Start: 12       End: 18        Position: 3     Type: 1
Term: 有限        Start: 18       End: 24        Position: 4     Type: 1
Term: 公司        Start: 24       End: 30        Position: 5     Type: 1
Term: 有限公司    Start: 18       End: 30        Position: 6     Type: 1
```    
### 配合 bleve 进行中文全文检索

[bleve](http://www.blevesearch.com/) 是一个 Go 语言实现的全文索引系统，jiebago 可以配合 bleve 使用实现中文的全文检索。一个简单的用法示例：

```
package main

import (
    "fmt"
    "github.com/blevesearch/bleve"
    _ "github.com/wangbin/jiebago/analyse/tokenizers"
    "log"
)

func main() {
    // open a new index
    indexMapping := bleve.NewIndexMapping()

    err := indexMapping.AddCustomTokenizer("jieba",
        map[string]interface{}{
            "file": "/Users/wangbin/mygo/src/github.com/wangbin/jiebago/dict.txt",
            "type": "jieba",
        })
    if err != nil {
        log.Fatal(err)
    }

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

    index, err := bleve.New("example.bleve", indexMapping)

    if err != nil {
        log.Fatal(err)
    }

    indexMapping.DefaultAnalyzer = "jieba"

    index, err := bleve.New("example.bleve", indexMapping)

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
        query := bleve.NewMatchQuery(keyword)
        search := bleve.NewSearchRequest(query)
        search.Highlight = bleve.NewHighlight()
        searchResults, err := index.Search(search)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Result of %s: %s\n", keyword, searchResults)
    }
}
```
输出结果：

```
Result of 水果世博园: 2 matches, showing 1 through 2, took 377.988µs
    1. Doc 3 (1.099550)
        Name
                买<span class="highlight">水果</span>然后来<span class="highlight">世博</span>园。
    2. Doc 2 (0.031941)
        Name
                The second one 你 中文测试中文 is even more interesting! 吃<span class="highlight">水果</span>

Result of 你: 1 matches, showing 1 through 1, took 103.367µs
    1. Doc 2 (0.391161)
        Name
                The second one <span class="highlight">你</span> 中文测试中文 is even more interesting! 吃水果

Result of first: 1 matches, showing 1 through 1, took 373.317µs
    1. Doc 1 (0.512150)
        Name
                This is the <span class="highlight">first</span> document we’ve added

Result of 中文: 1 matches, showing 1 through 1, took 106.433µs
    1. Doc 2 (0.553186)
        Name
                The second one 你 <span class="highlight">中文</span>测试<span class="highlight">中文</span> is even more interesting! 吃水果

Result of 交换机: 2 matches, showing 1 through 2, took 188.235µs
    1. Doc 4 (0.608495)
        Name
                工信处女干事每月经过下属科室都要亲口交代24口<span class="highlight">交换</span>机等技术性器件的安装工作
    2. Doc 5 (0.086700)
        Name
                咱俩<span class="highlight">交换</span>一下吧。

Result of 交换: 2 matches, showing 1 through 2, took 148.822µs
    1. Doc 5 (0.534158)
        Name
                咱俩<span class="highlight">交换</span>一下吧。
    2. Doc 4 (0.296297)
        Name
                工信处女干事每月经过下属科室都要亲口交代24口<span class="highlight">交换</span>机等技术性器件的安装工作
```

## 分词速度

 - 2MB / Second in Full Mode
 - 700KB / Second in Default Mode
 - Test Env: AMD Phenom(tm) II X6 1055T CPU @ 2.8GHz; 《金庸全集》 

## 许可证

MIT: http://wangbin.mit-license.org
