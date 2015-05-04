package analyse

import (
	"math"
	"testing"
)

var (
	sentence = "此外，公司拟对全资子公司吉林欧亚置业有限公司增资4.3亿元，增资后，吉林欧亚置业注册资本由7000万元增加到5亿元。吉林欧亚置业主要经营范围为房地产开发及百货零售等业务。目前在建吉林欧亚城市商业综合体项目。2013年，实现营业收入0万元，实现净利润-139.13万元。"

	tagRanks = Segments{
		Segment{text: "吉林", weight: 1.0},
		Segment{text: "欧亚", weight: 0.87807810644},
		Segment{text: "置业", weight: 0.562048250306},
		Segment{text: "实现", weight: 0.520905743929},
		Segment{text: "收入", weight: 0.384283870648},
		Segment{text: "增资", weight: 0.360590945312},
		Segment{text: "子公司", weight: 0.353131980904},
		Segment{text: "城市", weight: 0.307509449283},
		Segment{text: "全资", weight: 0.306324426665},
		Segment{text: "商业", weight: 0.306138241063},
	}
)

func TestTextRank(t *testing.T) {
	var tr TextRanker
	tr.LoadDictionary("../dict.txt")
	results := tr.TextRank(sentence, 10)
	for index, tw := range results {
		if tw.text != tagRanks[index].text || math.Abs(tw.weight-tagRanks[index].weight) > 1e-6 {
			t.Fatalf("%v != %v", tw, tagRanks[index])
		}
	}
}
