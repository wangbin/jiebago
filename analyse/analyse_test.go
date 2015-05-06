package analyse

import (
	"math"
	"testing"
)

var (
	testContents = []string{
		"这是一个伸手不见五指的黑夜。我叫孙悟空，我爱北京，我爱Python和C++。",
		"我不喜欢日本和服。",
		"雷猴回归人间。",
		"工信处女干事每月经过下属科室都要亲口交代24口交换机等技术性器件的安装工作",
		"我需要廉租房",
		"永和服装饰品有限公司",
		"我爱北京天安门",
		"abc",
		"隐马尔可夫",
		"雷猴是个好网站",
		"“Microsoft”一词由“MICROcomputer（微型计算机）”和“SOFTware（软件）”两部分组成",
		"草泥马和欺实马是今年的流行词汇",
		"伊藤洋华堂总府店",
		"中国科学院计算技术研究所",
		"罗密欧与朱丽叶",
		"我购买了道具和服装",
		"PS: 我觉得开源有一个好处，就是能够敦促自己不断改进，避免敞帚自珍",
		"湖北省石首市",
		"湖北省十堰市",
		"总经理完成了这件事情",
		"电脑修好了",
		"做好了这件事情就一了百了了",
		"人们审美的观点是不同的",
		"我们买了一个美的空调",
		"线程初始化时我们要注意",
		"一个分子是由好多原子组织成的",
		"祝你马到功成",
		"他掉进了无底洞里",
		"中国的首都是北京",
		"孙君意",
		"外交部发言人马朝旭",
		"领导人会议和第四届东亚峰会",
		"在过去的这五年",
		"还需要很长的路要走",
		"60周年首都阅兵",
		"你好人们审美的观点是不同的",
		"买水果然后来世博园",
		"买水果然后去世博园",
		"但是后来我才知道你是对的",
		"存在即合理",
		"的的的的的在的的的的就以和和和",
		"I love你，不以为耻，反以为rong",
		"因",
		"",
		"hello你好人们审美的观点是不同的",
		"很好但主要是基于网页形式",
		"hello你好人们审美的观点是不同的",
		"为什么我不能拥有想要的生活",
		"后来我才",
		"此次来中国是为了",
		"使用了它就可以解决一些问题",
		",使用了它就可以解决一些问题",
		"其实使用了它就可以解决一些问题",
		"好人使用了它就可以解决一些问题",
		"是因为和国家",
		"老年搜索还支持",
		"干脆就把那部蒙人的闲法给废了拉倒！RT @laoshipukong : 27日，全国人大常委会第三次审议侵权责任法草案，删除了有关医疗损害责任“举证倒置”的规定。在医患纠纷中本已处于弱势地位的消费者由此将陷入万劫不复的境地。 ",
		"大",
		"",
		"他说的确实在理",
		"长春市长春节讲话",
		"结婚的和尚未结婚的",
		"结合成分子时",
		"旅游和服务是最好的",
		"这件事情的确是我的错",
		"供大家参考指正",
		"哈尔滨政府公布塌桥原因",
		"我在机场入口处",
		"邢永臣摄影报道",
		"BP神经网络如何训练才能在分类时增加区分度？",
		"南京市长江大桥",
		"应一些使用者的建议，也为了便于利用NiuTrans用于SMT研究",
		"长春市长春药店",
		"邓颖超生前最喜欢的衣服",
		"胡锦涛是热爱世界和平的政治局常委",
		"程序员祝海林和朱会震是在孙健的左面和右面, 范凯在最右面.再往左是李松洪",
		"一次性交多少钱",
		"两块五一套，三块八一斤，四块七一本，五块六一条",
		"小和尚留了一个像大和尚一样的和尚头",
		"我是中华人民共和国公民;我爸爸是共和党党员; 地铁和平门站",
		"张晓梅去人民医院做了个B超然后去买了件T恤",
		"AT&T是一件不错的公司，给你发offer了吗？",
		"C++和c#是什么关系？11+122=133，是吗？PI=3.14159",
		"你认识那个和主席握手的的哥吗？他开一辆黑色的士。",
		"枪杆子中出政权"}

	Tags = [][]string{
		[]string{"Python", "C++", "伸手不见五指", "孙悟空", "黑夜", "北京", "这是", "一个"},
		[]string{"和服", "喜欢", "日本"},
		[]string{"雷猴", "人间", "回归"},
		[]string{"工信处", "女干事", "24", "交换机", "科室", "亲口", "器件", "技术性", "下属", "交代", "每月", "安装", "经过", "工作"},
		[]string{"廉租房", "需要"},
		[]string{"饰品", "永和", "服装", "有限公司"},
		[]string{"天安门", "北京"},
		[]string{"abc"},
		[]string{"马尔可夫"},
		[]string{"雷猴", "网站"},
		[]string{"SOFTware", "Microsoft", "MICROcomputer", "微型", "一词", "软件", "计算机", "组成", "部分"},
		[]string{"草泥马", "欺实", "词汇", "流行", "今年"},
		[]string{"洋华堂", "总府", "伊藤"},
		[]string{"中国科学院计算技术研究所"},
		[]string{"朱丽叶", "罗密欧"},
		[]string{"道具", "服装", "购买"},
		[]string{"自珍", "敞帚", "PS", "开源", "不断改进", "敦促", "好处", "避免", "能够", "觉得", "就是", "自己", "一个"},
		[]string{"石首市", "湖北省"},
		[]string{"十堰市", "湖北省"},
		[]string{"总经理", "这件", "完成", "事情"},
		[]string{"修好", "电脑"},
		[]string{"一了百了", "做好", "这件", "事情"},
		[]string{"审美", "观点", "人们", "不同"},
		[]string{"美的", "空调", "我们", "一个"},
		[]string{"线程", "初始化", "注意", "我们"},
		[]string{"好多", "原子", "分子", "组织", "一个"},
		[]string{"马到功成"},
		[]string{"无底洞"},
		[]string{"首都", "北京", "中国"},
		[]string{"孙君意"},
		[]string{"马朝旭", "外交部", "发言人"},
		[]string{"第四届", "东亚", "峰会", "领导人", "会议"},
		[]string{"五年", "过去"},
		[]string{"很长", "需要"},
		[]string{"60", "阅兵", "周年", "首都"},
		[]string{"审美", "你好", "观点", "人们", "不同"},
		[]string{"世博园", "水果", "然后"},
		[]string{"世博园", "水果", "然后"},
		[]string{"后来", "但是", "知道"},
		[]string{"合理", "存在"},
		[]string{},
		[]string{"rong", "love", "不以为耻", "以为"},
		[]string{},
		[]string{},
		[]string{"hello", "审美", "你好", "观点", "人们", "不同"},
		[]string{"网页", "基于", "形式", "主要"},
		[]string{"hello", "审美", "你好", "观点", "人们", "不同"},
		[]string{"想要", "拥有", "为什么", "生活", "不能"},
		[]string{"后来"},
		[]string{"此次", "为了", "中国"},
		[]string{"解决", "使用", "一些", "问题", "可以"},
		[]string{"解决", "使用", "一些", "问题", "可以"},
		[]string{"解决", "其实", "使用", "一些", "问题", "可以"},
		[]string{"好人", "解决", "使用", "一些", "问题", "可以"},
		[]string{"是因为", "国家"},
		[]string{"老年", "搜索", "支持"},
		[]string{"闲法", "中本", "laoshipukong", "RT", "27", "责任法", "蒙人", "万劫不复", "举证", "倒置", "医患", "那部", "拉倒", "侵权", "全国人大常委会", "草案", "境地", "纠纷", "删除", "弱势"},
		[]string{},
		[]string{},
		[]string{"在理", "确实"},
		[]string{"长春", "春节", "讲话", "市长"},
		[]string{"结婚", "尚未"},
		[]string{"分子", "结合"},
		[]string{"旅游", "最好", "服务"},
		[]string{"的确", "这件", "事情"},
		[]string{"指正", "参考", "大家"},
		[]string{"塌桥", "哈尔滨", "公布", "原因", "政府"},
		[]string{"入口处", "机场"},
		[]string{"邢永臣", "摄影", "报道"},
		[]string{"区分度", "BP", "神经网络", "训练", "分类", "才能", "如何", "增加"},
		[]string{"长江大桥", "南京市"},
		[]string{"SMT", "NiuTrans", "使用者", "便于", "用于", "建议", "利用", "为了", "研究", "一些"},
		[]string{"长春市", "药店", "长春"},
		[]string{"邓颖超", "生前", "衣服", "喜欢"},
		[]string{"政治局", "热爱", "常委", "胡锦涛", "和平", "世界"},
		[]string{"右面", "孙健", "范凯", "李松洪", "朱会震", "海林", "左面", "程序员", "再往"},
		[]string{"一次性", "多少"},
		[]string{"四块", "五块", "三块", "一斤", "两块", "一本", "一套", "一条"},
		[]string{"和尚", "和尚头", "一样", "一个"},
		[]string{"和平门", "共和党", "地铁", "党员", "公民", "爸爸", "中华人民共和国"},
		[]string{"张晓梅", "T恤", "B超", "医院", "人民", "然后"},
		[]string{"offer", "AT&T", "不错", "一件", "公司"},
		[]string{"c#", "PI", "C++", "3.14159", "133", "122", "11", "关系", "什么"},
		[]string{"的士", "的哥", "他开", "握手", "一辆", "黑色", "主席", "认识", "那个"},
		[]string{"枪杆子", "政权"},
	}

	Lyric = `
我沒有心
我沒有真實的自我
我只有消瘦的臉孔
所謂軟弱
所謂的順從一向是我
的座右銘

而我
沒有那海洋的寬闊
我只要熱情的撫摸
所謂空洞
所謂不安全感是我
的墓誌銘

而你
是否和我一般怯懦
是否和我一般矯作
和我一般囉唆

而你
是否和我一般退縮
是否和我一般肌迫
一般地困惑

我沒有力
我沒有滿腔的熱火
我只有滿肚的如果
所謂勇氣
所謂的認同感是我
隨便說說

而你
是否和我一般怯懦
是否和我一般矯作
是否對你來說
只是一場遊戲
雖然沒有把握

而你
是否和我一般退縮
是否和我一般肌迫
是否對你來說
只是逼不得已
雖然沒有藉口
`
	LyciWeight = Segments{
		Segment{text: "所謂", weight: 1.010262},
		Segment{text: "是否", weight: 0.738650},
		Segment{text: "一般", weight: 0.607600},
		Segment{text: "雖然", weight: 0.336754},
		Segment{text: "退縮", weight: 0.336754},
		Segment{text: "肌迫", weight: 0.336754},
		Segment{text: "矯作", weight: 0.336754},
		Segment{text: "沒有", weight: 0.336754},
		Segment{text: "怯懦", weight: 0.271099},
		Segment{text: "隨便", weight: 0.168377},
	}

	LyciWeight2 = Segments{
		Segment{text: "所謂", weight: 1.215739},
		Segment{text: "一般", weight: 0.731179},
		Segment{text: "雖然", weight: 0.405246},
		Segment{text: "退縮", weight: 0.405246},
		Segment{text: "肌迫", weight: 0.405246},
		Segment{text: "矯作", weight: 0.405246},
		Segment{text: "怯懦", weight: 0.326238},
		Segment{text: "逼不得已", weight: 0.202623},
		Segment{text: "右銘", weight: 0.202623},
		Segment{text: "寬闊", weight: 0.202623},
	}
)

func TestExtractTags(t *testing.T) {
	var te TagExtracter
	te.LoadDictionary("../dict.txt")
	te.LoadIdf("idf.txt")

	for index, sentence := range testContents {
		result := te.ExtractTags(sentence, 20)
		if len(result) != len(Tags[index]) {
			t.Fatalf("%s = %v", sentence, result)
		}
		for i, tag := range result {
			if tag.text != Tags[index][i] {
				t.Fatalf("%s != %s", tag, Tags[index][i])
			}
		}
	}
}

func TestExtratTagsWithWeight(t *testing.T) {
	var te TagExtracter
	te.LoadDictionary("../dict.txt")
	te.LoadIdf("idf.txt")
	result := te.ExtractTags(Lyric, 10)
	for index, tag := range result {
		if LyciWeight[index].text != tag.text ||
			math.Abs(LyciWeight[index].weight-tag.weight) > 1e-6 {
			t.Fatalf("%v != %v", tag, LyciWeight[index])
		}
	}
}

func TestExtractTagsWithStopWordsFile(t *testing.T) {
	var te TagExtracter
	te.LoadDictionary("../dict.txt")
	te.LoadIdf("idf.txt")
	te.LoadStopWords("stop_words.txt")
	result := te.ExtractTags(Lyric, 7)
	for index, tag := range result {
		if LyciWeight2[index].text != tag.text ||
			math.Abs(LyciWeight2[index].weight-tag.weight) > 1e-6 {
			t.Fatalf("%v != %v", tag, LyciWeight2[index])
		}
	}
}
