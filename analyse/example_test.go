package analyse_test

import (
	"fmt"

	"github.com/wangbin/jiebago/analyse"
)

func ExampleExtractTags() {
	var t analyse.TagExtracter
	t.LoadDictionary("../dict.txt")
	t.LoadIdf("idf.txt")

	sentence := "这是一个伸手不见五指的黑夜。我叫孙悟空，我爱北京，我爱Python和C++。"
	segments := t.ExtractTags(sentence, 5)
	fmt.Printf("Top %d tags:", len(segments))
	for _, segment := range segments {
		fmt.Printf(" %s /", segment.Text())
	}
	// Output:
	// Top 5 tags: Python / C++ / 伸手不见五指 / 孙悟空 / 黑夜 /
}

func ExampleTextRank() {
	var t analyse.TextRanker
	t.LoadDictionary("../dict.txt")
	sentence := "此外，公司拟对全资子公司吉林欧亚置业有限公司增资4.3亿元，增资后，吉林欧亚置业注册资本由7000万元增加到5亿元。吉林欧亚置业主要经营范围为房地产开发及百货零售等业务。目前在建吉林欧亚城市商业综合体项目。2013年，实现营业收入0万元，实现净利润-139.13万元。"

	result := t.TextRank(sentence, 10)
	for _, segment := range result {
		fmt.Printf("%s %f\n", segment.Text(), segment.Weight())
	}
	// Output:
	// 吉林 1.000000
	// 欧亚 0.878078
	// 置业 0.562048
	// 实现 0.520906
	// 收入 0.384284
	// 增资 0.360591
	// 子公司 0.353132
	// 城市 0.307509
	// 全资 0.306324
	// 商业 0.306138
}
