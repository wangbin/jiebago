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
