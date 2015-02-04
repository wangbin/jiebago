package analyse

import (
	"github.com/wangbin/jiebago"
	"math"
	"testing"
)

var (
	sentence = "此外，公司拟对全资子公司吉林欧亚置业有限公司增资4.3亿元，增资后，吉林欧亚置业注册资本由7000万元增加到5亿元。吉林欧亚置业主要经营范围为房地产开发及百货零售等业务。目前在建吉林欧亚城市商业综合体项目。2013年，实现营业收入0万元，实现净利润-139.13万元。"

	tagRanks = TfIdfs{
		TfIdf{Word: "吉林", Freq: 1.000000},
		TfIdf{Word: "欧亚", Freq: 0.864834},
		TfIdf{Word: "置业", Freq: 0.553466},
		TfIdf{Word: "实现", Freq: 0.520661},
		TfIdf{Word: "收入", Freq: 0.379700},
		TfIdf{Word: "增资", Freq: 0.355086},
		TfIdf{Word: "子公司", Freq: 0.349758},
		TfIdf{Word: "全资", Freq: 0.308537},
		TfIdf{Word: "城市", Freq: 0.306104},
		TfIdf{Word: "商业", Freq: 0.304837},
	}
)

func TesTextRank(t *testing.T) {
	jiebago.SetDictionary("../dict.txt")
	SetIdf("idf.txt")

	results := TextRank(sentence, 10)
	for index, tw := range results {
		if tw.Word != tagRanks[index].Word || math.Abs(tw.Freq-tagRanks[index].Freq) > 1e-6 {
			t.Errorf("%v != %v", tw, tagRanks[index])
		}
	}
}
