package jiebago_test

import (
	"fmt"

	"github.com/wangbin/jiebago"
)

func Example() {
	var seg jiebago.Segmenter
	seg.LoadDictionary("dict.txt")

	print := func(ch <-chan string) {
		for word := range ch {
			fmt.Printf(" %s /", word)
		}
		fmt.Println()
	}

	fmt.Print("【全模式】：")
	print(seg.CutAll("我来到北京清华大学"))

	fmt.Print("【精确模式】：")
	print(seg.Cut("我来到北京清华大学", false))

	fmt.Print("【新词识别】：")
	print(seg.Cut("他来到了网易杭研大厦", true))

	fmt.Print("【搜索引擎模式】：")
	print(seg.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造", true))
	// Output:
	// 【全模式】： 我 / 来到 / 北京 / 清华 / 清华大学 / 华大 / 大学 /
	// 【精确模式】： 我 / 来到 / 北京 / 清华大学 /
	// 【新词识别】： 他 / 来到 / 了 / 网易 / 杭研 / 大厦 /
	// 【搜索引擎模式】： 小明 / 硕士 / 毕业 / 于 / 中国 / 科学 / 学院 / 科学院 / 中国科学院 / 计算 / 计算所 / ， / 后 / 在 / 日本 / 京都 / 大学 / 日本京都大学 / 深造 /
}

func Example_loadUserDictionary() {
	var seg jiebago.Segmenter
	seg.LoadDictionary("dict.txt")

	print := func(ch <-chan string) {
		for word := range ch {
			fmt.Printf(" %s /", word)
		}
		fmt.Println()
	}
	sentence := "李小福是创新办主任也是云计算方面的专家"
	fmt.Print("Before:")
	print(seg.Cut(sentence, true))

	seg.LoadUserDictionary("userdict.txt")

	fmt.Print("After:")
	print(seg.Cut(sentence, true))
	// Output:
	// Before: 李小福 / 是 / 创新 / 办 / 主任 / 也 / 是 / 云 / 计算 / 方面 / 的 / 专家 /
	// After: 李小福 / 是 / 创新办 / 主任 / 也 / 是 / 云计算 / 方面 / 的 / 专家 /
}

func Example_tokenize() {
	var seg jiebago.Segmenter
	seg.LoadDictionary("dict.txt")

	sentence := []byte("永和服装饰品有限公司")

	// default mode
	tokenizer, _ := jiebago.NewJiebaTokenizer("dict.txt", true, false)
	fmt.Println("Default Mode:")
	for _, token := range tokenizer.Tokenize(sentence) {
		fmt.Printf(
			"Term: %s Start: %d End: %d Position: %d Type: %d\n",
			token.Term, token.Start, token.End, token.Position, token.Type)
	}

	//search mode
	tokenizer, _ = jiebago.NewJiebaTokenizer("dict.txt", true, true)
	fmt.Println("Search Mode:")
	for _, token := range tokenizer.Tokenize(sentence) {
		fmt.Printf(
			"Term: %s Start: %d End: %d Position: %d Type: %d\n",
			token.Term, token.Start, token.End, token.Position, token.Type)
	}
	// Output:
	// Default Mode:
	// Term: 永和 Start: 0 End: 6 Position: 1 Type: 1
	// Term: 服装 Start: 6 End: 12 Position: 2 Type: 1
	// Term: 饰品 Start: 12 End: 18 Position: 3 Type: 1
	// Term: 有限公司 Start: 18 End: 30 Position: 4 Type: 1
	// Search Mode:
	// Term: 永和 Start: 0 End: 6 Position: 1 Type: 1
	// Term: 服装 Start: 6 End: 12 Position: 2 Type: 1
	// Term: 饰品 Start: 12 End: 18 Position: 3 Type: 1
	// Term: 有限 Start: 18 End: 24 Position: 4 Type: 1
	// Term: 公司 Start: 24 End: 30 Position: 5 Type: 1
	// Term: 有限公司 Start: 18 End: 30 Position: 6 Type: 1
}
