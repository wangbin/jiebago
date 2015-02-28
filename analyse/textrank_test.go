package analyse

import (
	"math"
	"testing"
)

var (
	sentence = "此外，公司拟对全资子公司吉林欧亚置业有限公司增资4.3亿元，增资后，吉林欧亚置业注册资本由7000万元增加到5亿元。吉林欧亚置业主要经营范围为房地产开发及百货零售等业务。目前在建吉林欧亚城市商业综合体项目。2013年，实现营业收入0万元，实现净利润-139.13万元。"

	tagRanks = wordWeights{
		wordWeight{Word: "吉林", Weight: 1.0},
		wordWeight{Word: "欧亚", Weight: 0.87807810644},
		wordWeight{Word: "置业", Weight: 0.562048250306},
		wordWeight{Word: "实现", Weight: 0.520905743929},
		wordWeight{Word: "收入", Weight: 0.384283870648},
		wordWeight{Word: "增资", Weight: 0.360590945312},
		wordWeight{Word: "子公司", Weight: 0.353131980904},
		wordWeight{Word: "城市", Weight: 0.307509449283},
		wordWeight{Word: "全资", Weight: 0.306324426665},
		wordWeight{Word: "商业", Weight: 0.306138241063},
	}
)

func TestTextRank(t *testing.T) {
	SetDictionary("../dict.txt")
	results := TextRank(sentence, 10)
	for index, tw := range results {
		if tw.Word != tagRanks[index].Word || math.Abs(tw.Weight-tagRanks[index].Weight) > 1e-6 {
			t.Errorf("%v != %v", tw, tagRanks[index])
		}
	}
}
