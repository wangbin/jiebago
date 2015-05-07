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
