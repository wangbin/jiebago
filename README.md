结巴分词Go版 jiebago
===================

[![Build Status](https://travis-ci.org/wangbin/jiebago.png?branch=master)](https://travis-ci.org/wangbin/jiebago)

[结巴分词](https://github.com/fxsjy/jieba)是[@fxsjy](https://github.com/fxsjy)用Python编写的中文分词组件，jiebago是结巴分词的Go语言实现，目前已经实现的功能包括：三种模式分词、自定义词典、关键词提取和词性标注。


安装
=====

	go get github.com/wangbin/jiebago
	
分词
=====

    package main

    import (
        "fmt"
        "github.com/wangbin/jiebago"
        "strings"
    )

    var sentence = "我来到北京清华大学"

    func main() {
        jiebago.SetDictionary("/Path/to/default/dictionary/file") // 设定字典
        fmt.Printf("【全模式】: %s\n\n", strings.Join(jiebago.Cut(sentence, true, true), "/ "))
        fmt.Printf("【精确模式】: %s\n\n", strings.Join(jiebago.Cut(sentence, false, true), "/ "))
        fmt.Printf("【新词识别】：%s\n\n", strings.Join(jiebago.Cut("他来到了网易杭研大厦", false, true), ", "))
        fmt.Printf("【搜索引擎模式】：%s\n\n", strings.Join(jiebago.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造", true), ", "))
    }
    
使用结巴分词自带的[词典文件](https://github.com/fxsjy/jieba/blob/master/jieba/dict.txt)，输出结果如下：

    【全模式】: 我/ 来到/ 北京/ 清华/ 清华大学/ 华大/ 大学

    【精确模式】: 我/ 来到/ 北京/ 清华大学

    【新词识别】：他, 来到, 了, 网易, 杭研, 大厦

    【搜索引擎模式】：小明, 硕士, 毕业, 于, 中国, 科学, 学院, 科学院, 中国科学院, 计算, 计算所, ，, 后, 在, 日本, 京都, 大学, 日本京都大学, 深造
    
添加自定义词典
=============

    var sentence = "李小福是创新办主任也是云计算方面的专家"
    jiebago.SetDictionary("/Path/to/default/dictionary/file")
    fmt.Printf("Before: %s\n\n", strings.Join(jiebago.Cut(sentence, false, true), "/ "))
    jiebago.LoadUserDict("/Path/to/user/dictionary/file")
    fmt.Printf("After: %s\n\n", strings.Join(jiebago.Cut(sentence, false, true), "/ "))

使用结巴分词自带的[词典文件](https://github.com/fxsjy/jieba/blob/master/jieba/dict.txt)和[用户自定义词典文件](https://github.com/fxsjy/jieba/blob/master/test/userdict.txt)，结果输出如下：

    Before: 李小福/ 是/ 创新/ 办/ 主任/ 也/ 是/ 云/ 计算/ 方面/ 的/ 专家

    After: 李小福/ 是/ 创新办/ 主任/ 也/ 是/ 云计算/ 方面/ 的/ 专家
    
关键词提取
========

需要先安装analyse模块：

    go get github.com/wangbin/jiebago/analyse
    
示例代码：

    package main

    import (
        "fmt"
        "github.com/wangbin/jiebago"
        "github.com/wangbin/jiebago/analyse"
        "strings"
    )

    var sentence = "这是一个伸手不见五指的黑夜。我叫孙悟空，我爱北京，我爱Python和C++。"

    func main() {
        jiebago.SetDictionary("/Path/to/default/dictionary/file")
        analyse.SetIdf("/Path/to/idf/file")
        fmt.Println(strings.Join(analyse.ExtractTags(sentence, 20), "/ "))
    }
    
输出：

    Python/ C++/ 伸手不见五指/ 孙悟空/ 黑夜/ 北京/ 这是/ 一个
    
词性标注
=======

需要先安装posseg模块：

    go get github.com/wangbin/jiebago/posseg
    
示例代码：

    package main

    import (
        "fmt"
        "github.com/wangbin/jiebago"
        "github.com/wangbin/jiebago/posseg"
    )

    var sentence = "我爱北京天安门"

    func main() {
        posseg.SetDictionary("/Path/to/default/dictionary/file")
        for _, wt := range posseg.Cut(sentence, true) {
            fmt.Printf("%s %s\n", wt.Word, wt.Tag)
        }
    }
    
输出：

    我 r
    爱 v
    北京 ns
    天安门 ns
    
并行分词
=======

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
            ch <- jiebago.Cut(line, false, true)
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


Tokenize
=========

    var sentence = "永和服装饰品有限公司"
    // 默认模式
    for _, token := range jiebago.Tokenize(sentence, "default", true) {
        fmt.Printf("word %s\t\t start: %d \t\t end:%d\n", token.Word, token.Start, token.End)
    }
    // 搜索模式
    for _, token := range jiebago.Tokenize(sentence, "search", true) {
        fmt.Printf("word %s\t\t start: %d \t\t end:%d\n", token.Word, token.Start, token.End)
    }

输出结果：

    word 永和                start: 0                end:2
    word 服装                start: 2                end:4
    word 饰品                start: 4                end:6
    word 有限公司            start: 6                end:10

    word 永和                start: 0                end:2
    word 服装                start: 0                end:2
    word 饰品                start: 0                end:2
    word 有限                start: 0                end:2
    word 公司                start: 2                end:4
    word 有限公司            start: 0                end:4

分词速度
=======

 - 2MB / Second in Full Mode
 - 700KB / Second in Default Mode
 - Test Env: AMD Phenom(tm) II X6 1055T CPU @ 2.8GHz; 《金庸全集》 

许可证
======
MIT: http://wangbin.mit-license.org
