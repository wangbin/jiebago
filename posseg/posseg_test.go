package posseg

import (
	"github.com/wangbin/jiebago"
	"testing"
)

var (
	test_contents = []string{
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
	defaultCutResult = [][]WordTag{
		[]WordTag{WordTag{"\u8fd9", "r"}, WordTag{"\u662f", "v"}, WordTag{"\u4e00\u4e2a", "m"}, WordTag{"\u4f38\u624b\u4e0d\u89c1\u4e94\u6307", "i"}, WordTag{"\u7684", "uj"}, WordTag{"\u9ed1\u591c", "n"}, WordTag{"\u3002", "x"}, WordTag{"\u6211", "r"}, WordTag{"\u53eb", "v"}, WordTag{"\u5b59\u609f\u7a7a", "nr"}, WordTag{"\uff0c", "x"}, WordTag{"\u6211", "r"}, WordTag{"\u7231", "v"}, WordTag{"\u5317\u4eac", "ns"}, WordTag{"\uff0c", "x"}, WordTag{"\u6211", "r"}, WordTag{"\u7231", "v"}, WordTag{"Python", "eng"}, WordTag{"\u548c", "c"}, WordTag{"C++", "nz"}, WordTag{"\u3002", "x"}},
		[]WordTag{WordTag{"\u6211", "r"}, WordTag{"\u4e0d", "d"}, WordTag{"\u559c\u6b22", "v"}, WordTag{"\u65e5\u672c", "ns"}, WordTag{"\u548c\u670d", "nz"}, WordTag{"\u3002", "x"}},
		[]WordTag{WordTag{"\u96f7\u7334", "n"}, WordTag{"\u56de\u5f52", "v"}, WordTag{"\u4eba\u95f4", "n"}, WordTag{"\u3002", "x"}},
		[]WordTag{WordTag{"\u5de5\u4fe1\u5904", "n"}, WordTag{"\u5973\u5e72\u4e8b", "n"}, WordTag{"\u6bcf\u6708", "r"}, WordTag{"\u7ecf\u8fc7", "p"}, WordTag{"\u4e0b\u5c5e", "v"}, WordTag{"\u79d1\u5ba4", "n"}, WordTag{"\u90fd", "d"}, WordTag{"\u8981", "v"}, WordTag{"\u4eb2\u53e3", "n"}, WordTag{"\u4ea4\u4ee3", "n"}, WordTag{"24", "m"}, WordTag{"\u53e3", "n"}, WordTag{"\u4ea4\u6362\u673a", "n"}, WordTag{"\u7b49", "u"}, WordTag{"\u6280\u672f\u6027", "n"}, WordTag{"\u5668\u4ef6", "n"}, WordTag{"\u7684", "uj"}, WordTag{"\u5b89\u88c5", "v"}, WordTag{"\u5de5\u4f5c", "vn"}},
		[]WordTag{WordTag{"\u6211", "r"}, WordTag{"\u9700\u8981", "v"}, WordTag{"\u5ec9\u79df\u623f", "n"}},
		[]WordTag{WordTag{"\u6c38\u548c", "nz"}, WordTag{"\u670d\u88c5", "vn"}, WordTag{"\u9970\u54c1", "n"}, WordTag{"\u6709\u9650\u516c\u53f8", "n"}},
		[]WordTag{WordTag{"\u6211", "r"}, WordTag{"\u7231", "v"}, WordTag{"\u5317\u4eac", "ns"}, WordTag{"\u5929\u5b89\u95e8", "ns"}},
		[]WordTag{WordTag{"abc", "eng"}},
		[]WordTag{WordTag{"\u9690", "n"}, WordTag{"\u9a6c\u5c14\u53ef\u592b", "nr"}},
		[]WordTag{WordTag{"\u96f7\u7334", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u4e2a", "q"}, WordTag{"\u597d", "a"}, WordTag{"\u7f51\u7ad9", "n"}},
		[]WordTag{WordTag{"\u201c", "x"}, WordTag{"Microsoft", "eng"}, WordTag{"\u201d", "x"}, WordTag{"\u4e00", "m"}, WordTag{"\u8bcd", "n"}, WordTag{"\u7531", "p"}, WordTag{"\u201c", "x"}, WordTag{"MICROcomputer", "eng"}, WordTag{"\uff08", "x"}, WordTag{"\u5fae\u578b", "b"}, WordTag{"\u8ba1\u7b97\u673a", "n"}, WordTag{"\uff09", "x"}, WordTag{"\u201d", "x"}, WordTag{"\u548c", "c"}, WordTag{"\u201c", "x"}, WordTag{"SOFTware", "eng"}, WordTag{"\uff08", "x"}, WordTag{"\u8f6f\u4ef6", "n"}, WordTag{"\uff09", "x"}, WordTag{"\u201d", "x"}, WordTag{"\u4e24", "m"}, WordTag{"\u90e8\u5206", "n"}, WordTag{"\u7ec4\u6210", "v"}},
		[]WordTag{WordTag{"\u8349\u6ce5\u9a6c", "n"}, WordTag{"\u548c", "c"}, WordTag{"\u6b3a\u5b9e", "v"}, WordTag{"\u9a6c", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u4eca\u5e74", "t"}, WordTag{"\u7684", "uj"}, WordTag{"\u6d41\u884c", "v"}, WordTag{"\u8bcd\u6c47", "n"}},
		[]WordTag{WordTag{"\u4f0a\u85e4", "nr"}, WordTag{"\u6d0b\u534e\u5802", "n"}, WordTag{"\u603b\u5e9c", "n"}, WordTag{"\u5e97", "n"}},
		[]WordTag{WordTag{"\u4e2d\u56fd\u79d1\u5b66\u9662\u8ba1\u7b97\u6280\u672f\u7814\u7a76\u6240", "nt"}},
		[]WordTag{WordTag{"\u7f57\u5bc6\u6b27", "nr"}, WordTag{"\u4e0e", "p"}, WordTag{"\u6731\u4e3d\u53f6", "nr"}},
		[]WordTag{WordTag{"\u6211", "r"}, WordTag{"\u8d2d\u4e70", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u9053\u5177", "n"}, WordTag{"\u548c", "c"}, WordTag{"\u670d\u88c5", "vn"}},
		[]WordTag{WordTag{"PS", "eng"}, WordTag{":", "x"}, WordTag{" ", "x"}, WordTag{"\u6211", "r"}, WordTag{"\u89c9\u5f97", "v"}, WordTag{"\u5f00\u6e90", "n"}, WordTag{"\u6709", "v"}, WordTag{"\u4e00\u4e2a", "m"}, WordTag{"\u597d\u5904", "d"}, WordTag{"\uff0c", "x"}, WordTag{"\u5c31\u662f", "d"}, WordTag{"\u80fd\u591f", "v"}, WordTag{"\u6566\u4fc3", "v"}, WordTag{"\u81ea\u5df1", "r"}, WordTag{"\u4e0d\u65ad\u6539\u8fdb", "l"}, WordTag{"\uff0c", "x"}, WordTag{"\u907f\u514d", "v"}, WordTag{"\u655e", "v"}, WordTag{"\u5e1a", "ng"}, WordTag{"\u81ea\u73cd", "b"}},
		[]WordTag{WordTag{"\u6e56\u5317\u7701", "ns"}, WordTag{"\u77f3\u9996\u5e02", "ns"}},
		[]WordTag{WordTag{"\u6e56\u5317\u7701", "ns"}, WordTag{"\u5341\u5830\u5e02", "ns"}},
		[]WordTag{WordTag{"\u603b\u7ecf\u7406", "n"}, WordTag{"\u5b8c\u6210", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u8fd9\u4ef6", "mq"}, WordTag{"\u4e8b\u60c5", "n"}},
		[]WordTag{WordTag{"\u7535\u8111", "n"}, WordTag{"\u4fee\u597d", "v"}, WordTag{"\u4e86", "ul"}},
		[]WordTag{WordTag{"\u505a\u597d", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u8fd9\u4ef6", "mq"}, WordTag{"\u4e8b\u60c5", "n"}, WordTag{"\u5c31", "d"}, WordTag{"\u4e00\u4e86\u767e\u4e86", "l"}, WordTag{"\u4e86", "ul"}},
		[]WordTag{WordTag{"\u4eba\u4eec", "n"}, WordTag{"\u5ba1\u7f8e", "vn"}, WordTag{"\u7684", "uj"}, WordTag{"\u89c2\u70b9", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u4e0d\u540c", "a"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u6211\u4eec", "r"}, WordTag{"\u4e70", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u4e00\u4e2a", "m"}, WordTag{"\u7f8e\u7684", "nr"}, WordTag{"\u7a7a\u8c03", "n"}},
		[]WordTag{WordTag{"\u7ebf\u7a0b", "n"}, WordTag{"\u521d\u59cb\u5316", "l"}, WordTag{"\u65f6", "n"}, WordTag{"\u6211\u4eec", "r"}, WordTag{"\u8981", "v"}, WordTag{"\u6ce8\u610f", "v"}},
		[]WordTag{WordTag{"\u4e00\u4e2a", "m"}, WordTag{"\u5206\u5b50", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u7531", "p"}, WordTag{"\u597d\u591a", "m"}, WordTag{"\u539f\u5b50", "n"}, WordTag{"\u7ec4\u7ec7", "v"}, WordTag{"\u6210", "v"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u795d", "v"}, WordTag{"\u4f60", "r"}, WordTag{"\u9a6c\u5230\u529f\u6210", "i"}},
		[]WordTag{WordTag{"\u4ed6", "r"}, WordTag{"\u6389", "v"}, WordTag{"\u8fdb", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u65e0\u5e95\u6d1e", "ns"}, WordTag{"\u91cc", "f"}},
		[]WordTag{WordTag{"\u4e2d\u56fd", "ns"}, WordTag{"\u7684", "uj"}, WordTag{"\u9996\u90fd", "d"}, WordTag{"\u662f", "v"}, WordTag{"\u5317\u4eac", "ns"}},
		[]WordTag{WordTag{"\u5b59\u541b\u610f", "nr"}},
		[]WordTag{WordTag{"\u5916\u4ea4\u90e8", "nt"}, WordTag{"\u53d1\u8a00\u4eba", "l"}, WordTag{"\u9a6c\u671d\u65ed", "nr"}},
		[]WordTag{WordTag{"\u9886\u5bfc\u4eba", "n"}, WordTag{"\u4f1a\u8bae", "n"}, WordTag{"\u548c", "c"}, WordTag{"\u7b2c\u56db\u5c4a", "m"}, WordTag{"\u4e1c\u4e9a", "ns"}, WordTag{"\u5cf0\u4f1a", "n"}},
		[]WordTag{WordTag{"\u5728", "p"}, WordTag{"\u8fc7\u53bb", "t"}, WordTag{"\u7684", "uj"}, WordTag{"\u8fd9", "r"}, WordTag{"\u4e94\u5e74", "t"}},
		[]WordTag{WordTag{"\u8fd8", "d"}, WordTag{"\u9700\u8981", "v"}, WordTag{"\u5f88", "d"}, WordTag{"\u957f", "a"}, WordTag{"\u7684", "uj"}, WordTag{"\u8def", "n"}, WordTag{"\u8981", "v"}, WordTag{"\u8d70", "v"}},
		[]WordTag{WordTag{"60", "m"}, WordTag{"\u5468\u5e74", "t"}, WordTag{"\u9996\u90fd", "d"}, WordTag{"\u9605\u5175", "v"}},
		[]WordTag{WordTag{"\u4f60\u597d", "l"}, WordTag{"\u4eba\u4eec", "n"}, WordTag{"\u5ba1\u7f8e", "vn"}, WordTag{"\u7684", "uj"}, WordTag{"\u89c2\u70b9", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u4e0d\u540c", "a"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u4e70", "v"}, WordTag{"\u6c34\u679c", "n"}, WordTag{"\u7136\u540e", "c"}, WordTag{"\u6765", "v"}, WordTag{"\u4e16\u535a\u56ed", "nr"}},
		[]WordTag{WordTag{"\u4e70", "v"}, WordTag{"\u6c34\u679c", "n"}, WordTag{"\u7136\u540e", "c"}, WordTag{"\u53bb", "v"}, WordTag{"\u4e16\u535a\u56ed", "nr"}},
		[]WordTag{WordTag{"\u4f46\u662f", "c"}, WordTag{"\u540e\u6765", "t"}, WordTag{"\u6211", "r"}, WordTag{"\u624d", "d"}, WordTag{"\u77e5\u9053", "v"}, WordTag{"\u4f60", "r"}, WordTag{"\u662f", "v"}, WordTag{"\u5bf9", "p"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u5b58\u5728", "v"}, WordTag{"\u5373", "v"}, WordTag{"\u5408\u7406", "vn"}},
		[]WordTag{WordTag{"\u7684\u7684", "u"}, WordTag{"\u7684\u7684", "u"}, WordTag{"\u7684", "uj"}, WordTag{"\u5728\u7684", "u"}, WordTag{"\u7684\u7684", "u"}, WordTag{"\u7684", "uj"}, WordTag{"\u5c31", "d"}, WordTag{"\u4ee5", "p"}, WordTag{"\u548c\u548c", "nz"}, WordTag{"\u548c", "c"}},
		[]WordTag{WordTag{"I", "x"}, WordTag{" ", "x"}, WordTag{"love", "eng"}, WordTag{"\u4f60", "r"}, WordTag{"\uff0c", "x"}, WordTag{"\u4e0d\u4ee5\u4e3a\u803b", "i"}, WordTag{"\uff0c", "x"}, WordTag{"\u53cd", "zg"}, WordTag{"\u4ee5\u4e3a", "c"}, WordTag{"rong", "eng"}},
		[]WordTag{WordTag{"\u56e0", "p"}},
		[]WordTag{},
		[]WordTag{WordTag{"hello", "eng"}, WordTag{"\u4f60\u597d", "l"}, WordTag{"\u4eba\u4eec", "n"}, WordTag{"\u5ba1\u7f8e", "vn"}, WordTag{"\u7684", "uj"}, WordTag{"\u89c2\u70b9", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u4e0d\u540c", "a"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u5f88\u597d", "a"}, WordTag{"\u4f46", "c"}, WordTag{"\u4e3b\u8981", "b"}, WordTag{"\u662f", "v"}, WordTag{"\u57fa\u4e8e", "p"}, WordTag{"\u7f51\u9875", "n"}, WordTag{"\u5f62\u5f0f", "n"}},
		[]WordTag{WordTag{"hello", "eng"}, WordTag{"\u4f60\u597d", "l"}, WordTag{"\u4eba\u4eec", "n"}, WordTag{"\u5ba1\u7f8e", "vn"}, WordTag{"\u7684", "uj"}, WordTag{"\u89c2\u70b9", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u4e0d\u540c", "a"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u4e3a\u4ec0\u4e48", "r"}, WordTag{"\u6211", "r"}, WordTag{"\u4e0d\u80fd", "v"}, WordTag{"\u62e5\u6709", "v"}, WordTag{"\u60f3\u8981", "v"}, WordTag{"\u7684", "uj"}, WordTag{"\u751f\u6d3b", "vn"}},
		[]WordTag{WordTag{"\u540e\u6765", "t"}, WordTag{"\u6211", "r"}, WordTag{"\u624d", "d"}},
		[]WordTag{WordTag{"\u6b64\u6b21", "r"}, WordTag{"\u6765", "v"}, WordTag{"\u4e2d\u56fd", "ns"}, WordTag{"\u662f", "v"}, WordTag{"\u4e3a\u4e86", "p"}},
		[]WordTag{WordTag{"\u4f7f\u7528", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u5b83", "r"}, WordTag{"\u5c31", "d"}, WordTag{"\u53ef\u4ee5", "c"}, WordTag{"\u89e3\u51b3", "v"}, WordTag{"\u4e00\u4e9b", "m"}, WordTag{"\u95ee\u9898", "n"}},
		[]WordTag{WordTag{",", "x"}, WordTag{"\u4f7f\u7528", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u5b83", "r"}, WordTag{"\u5c31", "d"}, WordTag{"\u53ef\u4ee5", "c"}, WordTag{"\u89e3\u51b3", "v"}, WordTag{"\u4e00\u4e9b", "m"}, WordTag{"\u95ee\u9898", "n"}},
		[]WordTag{WordTag{"\u5176\u5b9e", "d"}, WordTag{"\u4f7f\u7528", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u5b83", "r"}, WordTag{"\u5c31", "d"}, WordTag{"\u53ef\u4ee5", "c"}, WordTag{"\u89e3\u51b3", "v"}, WordTag{"\u4e00\u4e9b", "m"}, WordTag{"\u95ee\u9898", "n"}},
		[]WordTag{WordTag{"\u597d\u4eba", "n"}, WordTag{"\u4f7f\u7528", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u5b83", "r"}, WordTag{"\u5c31", "d"}, WordTag{"\u53ef\u4ee5", "c"}, WordTag{"\u89e3\u51b3", "v"}, WordTag{"\u4e00\u4e9b", "m"}, WordTag{"\u95ee\u9898", "n"}},
		[]WordTag{WordTag{"\u662f\u56e0\u4e3a", "c"}, WordTag{"\u548c", "c"}, WordTag{"\u56fd\u5bb6", "n"}},
		[]WordTag{WordTag{"\u8001\u5e74", "t"}, WordTag{"\u641c\u7d22", "v"}, WordTag{"\u8fd8", "d"}, WordTag{"\u652f\u6301", "v"}},
		[]WordTag{WordTag{"\u5e72\u8106", "d"}, WordTag{"\u5c31", "d"}, WordTag{"\u628a", "p"}, WordTag{"\u90a3\u90e8", "r"}, WordTag{"\u8499\u4eba", "n"}, WordTag{"\u7684", "uj"}, WordTag{"\u95f2\u6cd5", "n"}, WordTag{"\u7ed9", "p"}, WordTag{"\u5e9f", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u62c9\u5012", "v"}, WordTag{"\uff01", "x"}, WordTag{"RT", "eng"}, WordTag{" ", "x"}, WordTag{"@", "x"}, WordTag{"laoshipukong", "eng"}, WordTag{" ", "x"}, WordTag{":", "x"}, WordTag{" ", "x"}, WordTag{"27", "m"}, WordTag{"\u65e5", "m"}, WordTag{"\uff0c", "x"}, WordTag{"\u5168\u56fd\u4eba\u5927\u5e38\u59d4\u4f1a", "nt"}, WordTag{"\u7b2c\u4e09\u6b21", "m"}, WordTag{"\u5ba1\u8bae", "v"}, WordTag{"\u4fb5\u6743", "v"}, WordTag{"\u8d23\u4efb\u6cd5", "n"}, WordTag{"\u8349\u6848", "n"}, WordTag{"\uff0c", "x"}, WordTag{"\u5220\u9664", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u6709\u5173", "vn"}, WordTag{"\u533b\u7597", "n"}, WordTag{"\u635f\u5bb3", "v"}, WordTag{"\u8d23\u4efb", "n"}, WordTag{"\u201c", "x"}, WordTag{"\u4e3e\u8bc1", "v"}, WordTag{"\u5012\u7f6e", "v"}, WordTag{"\u201d", "x"}, WordTag{"\u7684", "uj"}, WordTag{"\u89c4\u5b9a", "n"}, WordTag{"\u3002", "x"}, WordTag{"\u5728", "p"}, WordTag{"\u533b\u60a3", "n"}, WordTag{"\u7ea0\u7eb7", "n"}, WordTag{"\u4e2d\u672c", "ns"}, WordTag{"\u5df2", "d"}, WordTag{"\u5904\u4e8e", "v"}, WordTag{"\u5f31\u52bf", "n"}, WordTag{"\u5730\u4f4d", "n"}, WordTag{"\u7684", "uj"}, WordTag{"\u6d88\u8d39\u8005", "n"}, WordTag{"\u7531\u6b64", "c"}, WordTag{"\u5c06", "d"}, WordTag{"\u9677\u5165", "v"}, WordTag{"\u4e07\u52ab\u4e0d\u590d", "i"}, WordTag{"\u7684", "uj"}, WordTag{"\u5883\u5730", "s"}, WordTag{"\u3002", "x"}, WordTag{" ", "x"}},
		[]WordTag{WordTag{"\u5927", "a"}},
		[]WordTag{},
		[]WordTag{WordTag{"\u4ed6", "r"}, WordTag{"\u8bf4", "v"}, WordTag{"\u7684", "uj"}, WordTag{"\u786e\u5b9e", "ad"}, WordTag{"\u5728", "p"}, WordTag{"\u7406", "n"}},
		[]WordTag{WordTag{"\u957f\u6625", "ns"}, WordTag{"\u5e02\u957f", "n"}, WordTag{"\u6625\u8282", "t"}, WordTag{"\u8bb2\u8bdd", "n"}},
		[]WordTag{WordTag{"\u7ed3\u5a5a", "v"}, WordTag{"\u7684", "uj"}, WordTag{"\u548c", "c"}, WordTag{"\u5c1a\u672a", "d"}, WordTag{"\u7ed3\u5a5a", "v"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u7ed3\u5408", "v"}, WordTag{"\u6210", "n"}, WordTag{"\u5206\u5b50", "n"}, WordTag{"\u65f6", "n"}},
		[]WordTag{WordTag{"\u65c5\u6e38", "vn"}, WordTag{"\u548c", "c"}, WordTag{"\u670d\u52a1", "vn"}, WordTag{"\u662f", "v"}, WordTag{"\u6700\u597d", "a"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u8fd9\u4ef6", "mq"}, WordTag{"\u4e8b\u60c5", "n"}, WordTag{"\u7684\u786e", "d"}, WordTag{"\u662f", "v"}, WordTag{"\u6211", "r"}, WordTag{"\u7684", "uj"}, WordTag{"\u9519", "n"}},
		[]WordTag{WordTag{"\u4f9b", "v"}, WordTag{"\u5927\u5bb6", "n"}, WordTag{"\u53c2\u8003", "v"}, WordTag{"\u6307\u6b63", "v"}},
		[]WordTag{WordTag{"\u54c8\u5c14\u6ee8", "ns"}, WordTag{"\u653f\u5e9c", "n"}, WordTag{"\u516c\u5e03", "v"}, WordTag{"\u584c", "v"}, WordTag{"\u6865", "n"}, WordTag{"\u539f\u56e0", "n"}},
		[]WordTag{WordTag{"\u6211", "r"}, WordTag{"\u5728", "p"}, WordTag{"\u673a\u573a", "n"}, WordTag{"\u5165\u53e3\u5904", "i"}},
		[]WordTag{WordTag{"\u90a2\u6c38\u81e3", "nr"}, WordTag{"\u6444\u5f71", "n"}, WordTag{"\u62a5\u9053", "v"}},
		[]WordTag{WordTag{"BP", "eng"}, WordTag{"\u795e\u7ecf\u7f51\u7edc", "n"}, WordTag{"\u5982\u4f55", "r"}, WordTag{"\u8bad\u7ec3", "vn"}, WordTag{"\u624d\u80fd", "v"}, WordTag{"\u5728", "p"}, WordTag{"\u5206\u7c7b", "n"}, WordTag{"\u65f6", "n"}, WordTag{"\u589e\u52a0", "v"}, WordTag{"\u533a\u5206\u5ea6", "n"}, WordTag{"\uff1f", "x"}},
		[]WordTag{WordTag{"\u5357\u4eac\u5e02", "ns"}, WordTag{"\u957f\u6c5f\u5927\u6865", "ns"}},
		[]WordTag{WordTag{"\u5e94", "v"}, WordTag{"\u4e00\u4e9b", "m"}, WordTag{"\u4f7f\u7528\u8005", "n"}, WordTag{"\u7684", "uj"}, WordTag{"\u5efa\u8bae", "n"}, WordTag{"\uff0c", "x"}, WordTag{"\u4e5f", "d"}, WordTag{"\u4e3a\u4e86", "p"}, WordTag{"\u4fbf\u4e8e", "v"}, WordTag{"\u5229\u7528", "n"}, WordTag{"NiuTrans", "eng"}, WordTag{"\u7528\u4e8e", "v"}, WordTag{"SMT", "eng"}, WordTag{"\u7814\u7a76", "vn"}},
		[]WordTag{WordTag{"\u957f\u6625\u5e02", "ns"}, WordTag{"\u957f\u6625", "ns"}, WordTag{"\u836f\u5e97", "n"}},
		[]WordTag{WordTag{"\u9093\u9896\u8d85", "nr"}, WordTag{"\u751f\u524d", "t"}, WordTag{"\u6700", "d"}, WordTag{"\u559c\u6b22", "v"}, WordTag{"\u7684", "uj"}, WordTag{"\u8863\u670d", "n"}},
		[]WordTag{WordTag{"\u80e1\u9526\u6d9b", "nr"}, WordTag{"\u662f", "v"}, WordTag{"\u70ed\u7231", "a"}, WordTag{"\u4e16\u754c", "n"}, WordTag{"\u548c\u5e73", "nz"}, WordTag{"\u7684", "uj"}, WordTag{"\u653f\u6cbb\u5c40", "n"}, WordTag{"\u5e38\u59d4", "j"}},
		[]WordTag{WordTag{"\u7a0b\u5e8f\u5458", "n"}, WordTag{"\u795d", "v"}, WordTag{"\u6d77\u6797", "nz"}, WordTag{"\u548c", "c"}, WordTag{"\u6731\u4f1a\u9707", "nr"}, WordTag{"\u662f", "v"}, WordTag{"\u5728", "p"}, WordTag{"\u5b59\u5065", "nr"}, WordTag{"\u7684", "uj"}, WordTag{"\u5de6\u9762", "f"}, WordTag{"\u548c", "c"}, WordTag{"\u53f3\u9762", "f"}, WordTag{",", "x"}, WordTag{" ", "x"}, WordTag{"\u8303\u51ef", "nr"}, WordTag{"\u5728", "p"}, WordTag{"\u6700", "a"}, WordTag{"\u53f3\u9762", "f"}, WordTag{".", "m"}, WordTag{"\u518d\u5f80", "d"}, WordTag{"\u5de6", "f"}, WordTag{"\u662f", "v"}, WordTag{"\u674e\u677e\u6d2a", "nr"}},
		[]WordTag{WordTag{"\u4e00\u6b21\u6027", "d"}, WordTag{"\u4ea4", "v"}, WordTag{"\u591a\u5c11", "m"}, WordTag{"\u94b1", "n"}},
		[]WordTag{WordTag{"\u4e24\u5757", "m"}, WordTag{"\u4e94", "m"}, WordTag{"\u4e00\u5957", "m"}, WordTag{"\uff0c", "x"}, WordTag{"\u4e09\u5757", "m"}, WordTag{"\u516b", "m"}, WordTag{"\u4e00\u65a4", "m"}, WordTag{"\uff0c", "x"}, WordTag{"\u56db\u5757", "m"}, WordTag{"\u4e03", "m"}, WordTag{"\u4e00\u672c", "m"}, WordTag{"\uff0c", "x"}, WordTag{"\u4e94\u5757", "m"}, WordTag{"\u516d", "m"}, WordTag{"\u4e00\u6761", "m"}},
		[]WordTag{WordTag{"\u5c0f", "a"}, WordTag{"\u548c\u5c1a", "nr"}, WordTag{"\u7559", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u4e00\u4e2a", "m"}, WordTag{"\u50cf", "v"}, WordTag{"\u5927", "a"}, WordTag{"\u548c\u5c1a", "nr"}, WordTag{"\u4e00\u6837", "r"}, WordTag{"\u7684", "uj"}, WordTag{"\u548c\u5c1a\u5934", "nr"}},
		[]WordTag{WordTag{"\u6211", "r"}, WordTag{"\u662f", "v"}, WordTag{"\u4e2d\u534e\u4eba\u6c11\u5171\u548c\u56fd", "ns"}, WordTag{"\u516c\u6c11", "n"}, WordTag{";", "x"}, WordTag{"\u6211", "r"}, WordTag{"\u7238\u7238", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u5171\u548c\u515a", "nt"}, WordTag{"\u515a\u5458", "n"}, WordTag{";", "x"}, WordTag{" ", "x"}, WordTag{"\u5730\u94c1", "n"}, WordTag{"\u548c\u5e73\u95e8", "ns"}, WordTag{"\u7ad9", "v"}},
		[]WordTag{WordTag{"\u5f20\u6653\u6885", "nr"}, WordTag{"\u53bb", "v"}, WordTag{"\u4eba\u6c11", "n"}, WordTag{"\u533b\u9662", "n"}, WordTag{"\u505a", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u4e2a", "q"}, WordTag{"B\u8d85", "n"}, WordTag{"\u7136\u540e", "c"}, WordTag{"\u53bb", "v"}, WordTag{"\u4e70", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u4ef6", "q"}, WordTag{"T\u6064", "n"}},
		[]WordTag{WordTag{"AT&T", "nz"}, WordTag{"\u662f", "v"}, WordTag{"\u4e00\u4ef6", "m"}, WordTag{"\u4e0d\u9519", "a"}, WordTag{"\u7684", "uj"}, WordTag{"\u516c\u53f8", "n"}, WordTag{"\uff0c", "x"}, WordTag{"\u7ed9", "p"}, WordTag{"\u4f60", "r"}, WordTag{"\u53d1", "v"}, WordTag{"offer", "eng"}, WordTag{"\u4e86", "ul"}, WordTag{"\u5417", "y"}, WordTag{"\uff1f", "x"}},
		[]WordTag{WordTag{"C++", "nz"}, WordTag{"\u548c", "c"}, WordTag{"c#", "nz"}, WordTag{"\u662f", "v"}, WordTag{"\u4ec0\u4e48", "r"}, WordTag{"\u5173\u7cfb", "n"}, WordTag{"\uff1f", "x"}, WordTag{"11", "m"}, WordTag{"+", "x"}, WordTag{"122", "m"}, WordTag{"=", "x"}, WordTag{"133", "m"}, WordTag{"\uff0c", "x"}, WordTag{"\u662f", "v"}, WordTag{"\u5417", "y"}, WordTag{"\uff1f", "x"}, WordTag{"PI", "eng"}, WordTag{"=", "x"}, WordTag{"3.14159", "m"}},
		[]WordTag{WordTag{"\u4f60", "r"}, WordTag{"\u8ba4\u8bc6", "v"}, WordTag{"\u90a3\u4e2a", "r"}, WordTag{"\u548c", "c"}, WordTag{"\u4e3b\u5e2d", "n"}, WordTag{"\u63e1\u624b", "v"}, WordTag{"\u7684", "uj"}, WordTag{"\u7684\u54e5", "n"}, WordTag{"\u5417", "y"}, WordTag{"\uff1f", "x"}, WordTag{"\u4ed6", "r"}, WordTag{"\u5f00", "v"}, WordTag{"\u4e00\u8f86", "m"}, WordTag{"\u9ed1\u8272", "n"}, WordTag{"\u7684\u58eb", "n"}, WordTag{"\u3002", "x"}},
		[]WordTag{WordTag{"\u67aa\u6746\u5b50", "n"}, WordTag{"\u4e2d", "f"}, WordTag{"\u51fa", "v"}, WordTag{"\u653f\u6743", "n"}},
	}
	noHMMCutResult = [][]WordTag{
		[]WordTag{WordTag{"\u8fd9", "r"}, WordTag{"\u662f", "v"}, WordTag{"\u4e00\u4e2a", "m"}, WordTag{"\u4f38\u624b\u4e0d\u89c1\u4e94\u6307", "i"}, WordTag{"\u7684", "uj"}, WordTag{"\u9ed1\u591c", "n"}, WordTag{"\u3002", "x"}, WordTag{"\u6211", "r"}, WordTag{"\u53eb", "v"}, WordTag{"\u5b59\u609f\u7a7a", "nr"}, WordTag{"\uff0c", "x"}, WordTag{"\u6211", "r"}, WordTag{"\u7231", "v"}, WordTag{"\u5317\u4eac", "ns"}, WordTag{"\uff0c", "x"}, WordTag{"\u6211", "r"}, WordTag{"\u7231", "v"}, WordTag{"Python", "eng"}, WordTag{"\u548c", "c"}, WordTag{"C++", "nz"}, WordTag{"\u3002", "x"}},
		[]WordTag{WordTag{"\u6211", "r"}, WordTag{"\u4e0d", "d"}, WordTag{"\u559c\u6b22", "v"}, WordTag{"\u65e5\u672c", "ns"}, WordTag{"\u548c\u670d", "nz"}, WordTag{"\u3002", "x"}},
		[]WordTag{WordTag{"\u96f7\u7334", "n"}, WordTag{"\u56de\u5f52", "v"}, WordTag{"\u4eba\u95f4", "n"}, WordTag{"\u3002", "x"}},
		[]WordTag{WordTag{"\u5de5\u4fe1\u5904", "n"}, WordTag{"\u5973\u5e72\u4e8b", "n"}, WordTag{"\u6bcf\u6708", "r"}, WordTag{"\u7ecf\u8fc7", "p"}, WordTag{"\u4e0b\u5c5e", "v"}, WordTag{"\u79d1\u5ba4", "n"}, WordTag{"\u90fd", "d"}, WordTag{"\u8981", "v"}, WordTag{"\u4eb2\u53e3", "n"}, WordTag{"\u4ea4\u4ee3", "n"}, WordTag{"24", "eng"}, WordTag{"\u53e3", "q"}, WordTag{"\u4ea4\u6362\u673a", "n"}, WordTag{"\u7b49", "u"}, WordTag{"\u6280\u672f\u6027", "n"}, WordTag{"\u5668\u4ef6", "n"}, WordTag{"\u7684", "uj"}, WordTag{"\u5b89\u88c5", "v"}, WordTag{"\u5de5\u4f5c", "vn"}},
		[]WordTag{WordTag{"\u6211", "r"}, WordTag{"\u9700\u8981", "v"}, WordTag{"\u5ec9\u79df\u623f", "n"}},
		[]WordTag{WordTag{"\u6c38\u548c", "nz"}, WordTag{"\u670d\u88c5", "vn"}, WordTag{"\u9970\u54c1", "n"}, WordTag{"\u6709\u9650\u516c\u53f8", "n"}},
		[]WordTag{WordTag{"\u6211", "r"}, WordTag{"\u7231", "v"}, WordTag{"\u5317\u4eac", "ns"}, WordTag{"\u5929\u5b89\u95e8", "ns"}},
		[]WordTag{WordTag{"abc", "eng"}},
		[]WordTag{WordTag{"\u9690", "n"}, WordTag{"\u9a6c\u5c14\u53ef\u592b", "nr"}},
		[]WordTag{WordTag{"\u96f7\u7334", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u4e2a", "q"}, WordTag{"\u597d", "a"}, WordTag{"\u7f51\u7ad9", "n"}},
		[]WordTag{WordTag{"\u201c", "x"}, WordTag{"Microsoft", "eng"}, WordTag{"\u201d", "x"}, WordTag{"\u4e00", "m"}, WordTag{"\u8bcd", "n"}, WordTag{"\u7531", "p"}, WordTag{"\u201c", "x"}, WordTag{"MICROcomputer", "eng"}, WordTag{"\uff08", "x"}, WordTag{"\u5fae\u578b", "b"}, WordTag{"\u8ba1\u7b97\u673a", "n"}, WordTag{"\uff09", "x"}, WordTag{"\u201d", "x"}, WordTag{"\u548c", "c"}, WordTag{"\u201c", "x"}, WordTag{"SOFTware", "eng"}, WordTag{"\uff08", "x"}, WordTag{"\u8f6f\u4ef6", "n"}, WordTag{"\uff09", "x"}, WordTag{"\u201d", "x"}, WordTag{"\u4e24", "m"}, WordTag{"\u90e8\u5206", "n"}, WordTag{"\u7ec4\u6210", "v"}},
		[]WordTag{WordTag{"\u8349\u6ce5\u9a6c", "n"}, WordTag{"\u548c", "c"}, WordTag{"\u6b3a", "vn"}, WordTag{"\u5b9e", "n"}, WordTag{"\u9a6c", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u4eca\u5e74", "t"}, WordTag{"\u7684", "uj"}, WordTag{"\u6d41\u884c", "v"}, WordTag{"\u8bcd\u6c47", "n"}},
		[]WordTag{WordTag{"\u4f0a", "ns"}, WordTag{"\u85e4", "nr"}, WordTag{"\u6d0b\u534e\u5802", "n"}, WordTag{"\u603b\u5e9c", "n"}, WordTag{"\u5e97", "n"}},
		[]WordTag{WordTag{"\u4e2d\u56fd\u79d1\u5b66\u9662\u8ba1\u7b97\u6280\u672f\u7814\u7a76\u6240", "nt"}},
		[]WordTag{WordTag{"\u7f57\u5bc6\u6b27", "nr"}, WordTag{"\u4e0e", "p"}, WordTag{"\u6731\u4e3d\u53f6", "nr"}},
		[]WordTag{WordTag{"\u6211", "r"}, WordTag{"\u8d2d\u4e70", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u9053\u5177", "n"}, WordTag{"\u548c", "c"}, WordTag{"\u670d\u88c5", "vn"}},
		[]WordTag{WordTag{"PS", "eng"}, WordTag{":", "x"}, WordTag{" ", "x"}, WordTag{"\u6211", "r"}, WordTag{"\u89c9\u5f97", "v"}, WordTag{"\u5f00\u6e90", "n"}, WordTag{"\u6709", "v"}, WordTag{"\u4e00\u4e2a", "m"}, WordTag{"\u597d\u5904", "d"}, WordTag{"\uff0c", "x"}, WordTag{"\u5c31\u662f", "d"}, WordTag{"\u80fd\u591f", "v"}, WordTag{"\u6566\u4fc3", "v"}, WordTag{"\u81ea\u5df1", "r"}, WordTag{"\u4e0d\u65ad\u6539\u8fdb", "l"}, WordTag{"\uff0c", "x"}, WordTag{"\u907f\u514d", "v"}, WordTag{"\u655e", "v"}, WordTag{"\u5e1a", "ng"}, WordTag{"\u81ea\u73cd", "b"}},
		[]WordTag{WordTag{"\u6e56\u5317\u7701", "ns"}, WordTag{"\u77f3\u9996\u5e02", "ns"}},
		[]WordTag{WordTag{"\u6e56\u5317\u7701", "ns"}, WordTag{"\u5341\u5830\u5e02", "ns"}},
		[]WordTag{WordTag{"\u603b\u7ecf\u7406", "n"}, WordTag{"\u5b8c\u6210", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u8fd9\u4ef6", "mq"}, WordTag{"\u4e8b\u60c5", "n"}},
		[]WordTag{WordTag{"\u7535\u8111", "n"}, WordTag{"\u4fee\u597d", "v"}, WordTag{"\u4e86", "ul"}},
		[]WordTag{WordTag{"\u505a\u597d", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u8fd9\u4ef6", "mq"}, WordTag{"\u4e8b\u60c5", "n"}, WordTag{"\u5c31", "d"}, WordTag{"\u4e00\u4e86\u767e\u4e86", "l"}, WordTag{"\u4e86", "ul"}},
		[]WordTag{WordTag{"\u4eba\u4eec", "n"}, WordTag{"\u5ba1\u7f8e", "vn"}, WordTag{"\u7684", "uj"}, WordTag{"\u89c2\u70b9", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u4e0d\u540c", "a"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u6211\u4eec", "r"}, WordTag{"\u4e70", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u4e00\u4e2a", "m"}, WordTag{"\u7f8e\u7684", "nr"}, WordTag{"\u7a7a\u8c03", "n"}},
		[]WordTag{WordTag{"\u7ebf\u7a0b", "n"}, WordTag{"\u521d\u59cb\u5316", "l"}, WordTag{"\u65f6", "n"}, WordTag{"\u6211\u4eec", "r"}, WordTag{"\u8981", "v"}, WordTag{"\u6ce8\u610f", "v"}},
		[]WordTag{WordTag{"\u4e00\u4e2a", "m"}, WordTag{"\u5206\u5b50", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u7531", "p"}, WordTag{"\u597d\u591a", "m"}, WordTag{"\u539f\u5b50", "n"}, WordTag{"\u7ec4\u7ec7", "v"}, WordTag{"\u6210", "n"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u795d", "v"}, WordTag{"\u4f60", "r"}, WordTag{"\u9a6c\u5230\u529f\u6210", "i"}},
		[]WordTag{WordTag{"\u4ed6", "r"}, WordTag{"\u6389", "zg"}, WordTag{"\u8fdb", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u65e0\u5e95\u6d1e", "ns"}, WordTag{"\u91cc", "f"}},
		[]WordTag{WordTag{"\u4e2d\u56fd", "ns"}, WordTag{"\u7684", "uj"}, WordTag{"\u9996\u90fd", "d"}, WordTag{"\u662f", "v"}, WordTag{"\u5317\u4eac", "ns"}},
		[]WordTag{WordTag{"\u5b59", "zg"}, WordTag{"\u541b", "nz"}, WordTag{"\u610f", "n"}},
		[]WordTag{WordTag{"\u5916\u4ea4\u90e8", "nt"}, WordTag{"\u53d1\u8a00\u4eba", "l"}, WordTag{"\u9a6c\u671d\u65ed", "nr"}},
		[]WordTag{WordTag{"\u9886\u5bfc\u4eba", "n"}, WordTag{"\u4f1a\u8bae", "n"}, WordTag{"\u548c", "c"}, WordTag{"\u7b2c\u56db\u5c4a", "m"}, WordTag{"\u4e1c\u4e9a", "ns"}, WordTag{"\u5cf0\u4f1a", "n"}},
		[]WordTag{WordTag{"\u5728", "p"}, WordTag{"\u8fc7\u53bb", "t"}, WordTag{"\u7684", "uj"}, WordTag{"\u8fd9", "r"}, WordTag{"\u4e94\u5e74", "t"}},
		[]WordTag{WordTag{"\u8fd8", "d"}, WordTag{"\u9700\u8981", "v"}, WordTag{"\u5f88", "zg"}, WordTag{"\u957f", "a"}, WordTag{"\u7684", "uj"}, WordTag{"\u8def", "n"}, WordTag{"\u8981", "v"}, WordTag{"\u8d70", "v"}},
		[]WordTag{WordTag{"60", "eng"}, WordTag{"\u5468\u5e74", "t"}, WordTag{"\u9996\u90fd", "d"}, WordTag{"\u9605\u5175", "v"}},
		[]WordTag{WordTag{"\u4f60\u597d", "l"}, WordTag{"\u4eba\u4eec", "n"}, WordTag{"\u5ba1\u7f8e", "vn"}, WordTag{"\u7684", "uj"}, WordTag{"\u89c2\u70b9", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u4e0d\u540c", "a"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u4e70", "v"}, WordTag{"\u6c34\u679c", "n"}, WordTag{"\u7136\u540e", "c"}, WordTag{"\u6765", "v"}, WordTag{"\u4e16\u535a\u56ed", "nr"}},
		[]WordTag{WordTag{"\u4e70", "v"}, WordTag{"\u6c34\u679c", "n"}, WordTag{"\u7136\u540e", "c"}, WordTag{"\u53bb", "v"}, WordTag{"\u4e16\u535a\u56ed", "nr"}},
		[]WordTag{WordTag{"\u4f46\u662f", "c"}, WordTag{"\u540e\u6765", "t"}, WordTag{"\u6211", "r"}, WordTag{"\u624d", "d"}, WordTag{"\u77e5\u9053", "v"}, WordTag{"\u4f60", "r"}, WordTag{"\u662f", "v"}, WordTag{"\u5bf9", "p"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u5b58\u5728", "v"}, WordTag{"\u5373", "v"}, WordTag{"\u5408\u7406", "vn"}},
		[]WordTag{WordTag{"\u7684", "uj"}, WordTag{"\u7684", "uj"}, WordTag{"\u7684", "uj"}, WordTag{"\u7684", "uj"}, WordTag{"\u7684", "uj"}, WordTag{"\u5728", "p"}, WordTag{"\u7684", "uj"}, WordTag{"\u7684", "uj"}, WordTag{"\u7684", "uj"}, WordTag{"\u7684", "uj"}, WordTag{"\u5c31", "d"}, WordTag{"\u4ee5", "p"}, WordTag{"\u548c", "c"}, WordTag{"\u548c", "c"}, WordTag{"\u548c", "c"}},
		[]WordTag{WordTag{"I", "eng"}, WordTag{" ", "x"}, WordTag{"love", "eng"}, WordTag{"\u4f60", "r"}, WordTag{"\uff0c", "x"}, WordTag{"\u4e0d\u4ee5\u4e3a\u803b", "i"}, WordTag{"\uff0c", "x"}, WordTag{"\u53cd", "zg"}, WordTag{"\u4ee5\u4e3a", "c"}, WordTag{"rong", "eng"}},
		[]WordTag{WordTag{"\u56e0", "p"}},
		[]WordTag{},
		[]WordTag{WordTag{"hello", "eng"}, WordTag{"\u4f60\u597d", "l"}, WordTag{"\u4eba\u4eec", "n"}, WordTag{"\u5ba1\u7f8e", "vn"}, WordTag{"\u7684", "uj"}, WordTag{"\u89c2\u70b9", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u4e0d\u540c", "a"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u5f88", "zg"}, WordTag{"\u597d", "a"}, WordTag{"\u4f46", "c"}, WordTag{"\u4e3b\u8981", "b"}, WordTag{"\u662f", "v"}, WordTag{"\u57fa\u4e8e", "p"}, WordTag{"\u7f51\u9875", "n"}, WordTag{"\u5f62\u5f0f", "n"}},
		[]WordTag{WordTag{"hello", "eng"}, WordTag{"\u4f60\u597d", "l"}, WordTag{"\u4eba\u4eec", "n"}, WordTag{"\u5ba1\u7f8e", "vn"}, WordTag{"\u7684", "uj"}, WordTag{"\u89c2\u70b9", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u4e0d\u540c", "a"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u4e3a\u4ec0\u4e48", "r"}, WordTag{"\u6211", "r"}, WordTag{"\u4e0d\u80fd", "v"}, WordTag{"\u62e5\u6709", "v"}, WordTag{"\u60f3\u8981", "v"}, WordTag{"\u7684", "uj"}, WordTag{"\u751f\u6d3b", "vn"}},
		[]WordTag{WordTag{"\u540e\u6765", "t"}, WordTag{"\u6211", "r"}, WordTag{"\u624d", "d"}},
		[]WordTag{WordTag{"\u6b64\u6b21", "r"}, WordTag{"\u6765", "v"}, WordTag{"\u4e2d\u56fd", "ns"}, WordTag{"\u662f", "v"}, WordTag{"\u4e3a\u4e86", "p"}},
		[]WordTag{WordTag{"\u4f7f\u7528", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u5b83", "r"}, WordTag{"\u5c31", "d"}, WordTag{"\u53ef\u4ee5", "c"}, WordTag{"\u89e3\u51b3", "v"}, WordTag{"\u4e00\u4e9b", "m"}, WordTag{"\u95ee\u9898", "n"}},
		[]WordTag{WordTag{",", "x"}, WordTag{"\u4f7f\u7528", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u5b83", "r"}, WordTag{"\u5c31", "d"}, WordTag{"\u53ef\u4ee5", "c"}, WordTag{"\u89e3\u51b3", "v"}, WordTag{"\u4e00\u4e9b", "m"}, WordTag{"\u95ee\u9898", "n"}},
		[]WordTag{WordTag{"\u5176\u5b9e", "d"}, WordTag{"\u4f7f\u7528", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u5b83", "r"}, WordTag{"\u5c31", "d"}, WordTag{"\u53ef\u4ee5", "c"}, WordTag{"\u89e3\u51b3", "v"}, WordTag{"\u4e00\u4e9b", "m"}, WordTag{"\u95ee\u9898", "n"}},
		[]WordTag{WordTag{"\u597d\u4eba", "n"}, WordTag{"\u4f7f\u7528", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u5b83", "r"}, WordTag{"\u5c31", "d"}, WordTag{"\u53ef\u4ee5", "c"}, WordTag{"\u89e3\u51b3", "v"}, WordTag{"\u4e00\u4e9b", "m"}, WordTag{"\u95ee\u9898", "n"}},
		[]WordTag{WordTag{"\u662f\u56e0\u4e3a", "c"}, WordTag{"\u548c", "c"}, WordTag{"\u56fd\u5bb6", "n"}},
		[]WordTag{WordTag{"\u8001\u5e74", "t"}, WordTag{"\u641c\u7d22", "v"}, WordTag{"\u8fd8", "d"}, WordTag{"\u652f\u6301", "v"}},
		[]WordTag{WordTag{"\u5e72\u8106", "d"}, WordTag{"\u5c31", "d"}, WordTag{"\u628a", "p"}, WordTag{"\u90a3", "r"}, WordTag{"\u90e8", "n"}, WordTag{"\u8499", "v"}, WordTag{"\u4eba", "n"}, WordTag{"\u7684", "uj"}, WordTag{"\u95f2", "n"}, WordTag{"\u6cd5", "j"}, WordTag{"\u7ed9", "p"}, WordTag{"\u5e9f", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u62c9\u5012", "v"}, WordTag{"\uff01", "x"}, WordTag{"RT", "eng"}, WordTag{" ", "x"}, WordTag{"@", "x"}, WordTag{"laoshipukong", "eng"}, WordTag{" ", "x"}, WordTag{":", "x"}, WordTag{" ", "x"}, WordTag{"27", "eng"}, WordTag{"\u65e5", "m"}, WordTag{"\uff0c", "x"}, WordTag{"\u5168\u56fd\u4eba\u5927\u5e38\u59d4\u4f1a", "nt"}, WordTag{"\u7b2c\u4e09\u6b21", "m"}, WordTag{"\u5ba1\u8bae", "v"}, WordTag{"\u4fb5\u6743", "v"}, WordTag{"\u8d23\u4efb\u6cd5", "n"}, WordTag{"\u8349\u6848", "n"}, WordTag{"\uff0c", "x"}, WordTag{"\u5220\u9664", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u6709\u5173", "vn"}, WordTag{"\u533b\u7597", "n"}, WordTag{"\u635f\u5bb3", "v"}, WordTag{"\u8d23\u4efb", "n"}, WordTag{"\u201c", "x"}, WordTag{"\u4e3e\u8bc1", "v"}, WordTag{"\u5012\u7f6e", "v"}, WordTag{"\u201d", "x"}, WordTag{"\u7684", "uj"}, WordTag{"\u89c4\u5b9a", "n"}, WordTag{"\u3002", "x"}, WordTag{"\u5728", "p"}, WordTag{"\u533b\u60a3", "n"}, WordTag{"\u7ea0\u7eb7", "n"}, WordTag{"\u4e2d", "f"}, WordTag{"\u672c", "r"}, WordTag{"\u5df2", "d"}, WordTag{"\u5904\u4e8e", "v"}, WordTag{"\u5f31\u52bf", "n"}, WordTag{"\u5730\u4f4d", "n"}, WordTag{"\u7684", "uj"}, WordTag{"\u6d88\u8d39\u8005", "n"}, WordTag{"\u7531\u6b64", "c"}, WordTag{"\u5c06", "d"}, WordTag{"\u9677\u5165", "v"}, WordTag{"\u4e07\u52ab\u4e0d\u590d", "i"}, WordTag{"\u7684", "uj"}, WordTag{"\u5883\u5730", "s"}, WordTag{"\u3002", "x"}, WordTag{" ", "x"}},
		[]WordTag{WordTag{"\u5927", "a"}},
		[]WordTag{},
		[]WordTag{WordTag{"\u4ed6", "r"}, WordTag{"\u8bf4", "v"}, WordTag{"\u7684", "uj"}, WordTag{"\u786e\u5b9e", "ad"}, WordTag{"\u5728", "p"}, WordTag{"\u7406", "n"}},
		[]WordTag{WordTag{"\u957f\u6625", "ns"}, WordTag{"\u5e02\u957f", "n"}, WordTag{"\u6625\u8282", "t"}, WordTag{"\u8bb2\u8bdd", "n"}},
		[]WordTag{WordTag{"\u7ed3\u5a5a", "v"}, WordTag{"\u7684", "uj"}, WordTag{"\u548c", "c"}, WordTag{"\u5c1a\u672a", "d"}, WordTag{"\u7ed3\u5a5a", "v"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u7ed3\u5408", "v"}, WordTag{"\u6210", "n"}, WordTag{"\u5206\u5b50", "n"}, WordTag{"\u65f6", "n"}},
		[]WordTag{WordTag{"\u65c5\u6e38", "vn"}, WordTag{"\u548c", "c"}, WordTag{"\u670d\u52a1", "vn"}, WordTag{"\u662f", "v"}, WordTag{"\u6700\u597d", "a"}, WordTag{"\u7684", "uj"}},
		[]WordTag{WordTag{"\u8fd9\u4ef6", "mq"}, WordTag{"\u4e8b\u60c5", "n"}, WordTag{"\u7684\u786e", "d"}, WordTag{"\u662f", "v"}, WordTag{"\u6211", "r"}, WordTag{"\u7684", "uj"}, WordTag{"\u9519", "v"}},
		[]WordTag{WordTag{"\u4f9b", "v"}, WordTag{"\u5927\u5bb6", "n"}, WordTag{"\u53c2\u8003", "v"}, WordTag{"\u6307\u6b63", "v"}},
		[]WordTag{WordTag{"\u54c8\u5c14\u6ee8", "ns"}, WordTag{"\u653f\u5e9c", "n"}, WordTag{"\u516c\u5e03", "v"}, WordTag{"\u584c", "v"}, WordTag{"\u6865", "n"}, WordTag{"\u539f\u56e0", "n"}},
		[]WordTag{WordTag{"\u6211", "r"}, WordTag{"\u5728", "p"}, WordTag{"\u673a\u573a", "n"}, WordTag{"\u5165\u53e3\u5904", "i"}},
		[]WordTag{WordTag{"\u90a2", "nr"}, WordTag{"\u6c38", "ns"}, WordTag{"\u81e3", "n"}, WordTag{"\u6444\u5f71", "n"}, WordTag{"\u62a5\u9053", "v"}},
		[]WordTag{WordTag{"BP", "eng"}, WordTag{"\u795e\u7ecf\u7f51\u7edc", "n"}, WordTag{"\u5982\u4f55", "r"}, WordTag{"\u8bad\u7ec3", "vn"}, WordTag{"\u624d\u80fd", "v"}, WordTag{"\u5728", "p"}, WordTag{"\u5206\u7c7b", "n"}, WordTag{"\u65f6", "n"}, WordTag{"\u589e\u52a0", "v"}, WordTag{"\u533a\u5206\u5ea6", "n"}, WordTag{"\uff1f", "x"}},
		[]WordTag{WordTag{"\u5357\u4eac\u5e02", "ns"}, WordTag{"\u957f\u6c5f\u5927\u6865", "ns"}},
		[]WordTag{WordTag{"\u5e94", "v"}, WordTag{"\u4e00\u4e9b", "m"}, WordTag{"\u4f7f\u7528\u8005", "n"}, WordTag{"\u7684", "uj"}, WordTag{"\u5efa\u8bae", "n"}, WordTag{"\uff0c", "x"}, WordTag{"\u4e5f", "d"}, WordTag{"\u4e3a\u4e86", "p"}, WordTag{"\u4fbf\u4e8e", "v"}, WordTag{"\u5229\u7528", "n"}, WordTag{"NiuTrans", "eng"}, WordTag{"\u7528\u4e8e", "v"}, WordTag{"SMT", "eng"}, WordTag{"\u7814\u7a76", "vn"}},
		[]WordTag{WordTag{"\u957f\u6625\u5e02", "ns"}, WordTag{"\u957f\u6625", "ns"}, WordTag{"\u836f\u5e97", "n"}},
		[]WordTag{WordTag{"\u9093\u9896\u8d85", "nr"}, WordTag{"\u751f\u524d", "t"}, WordTag{"\u6700", "d"}, WordTag{"\u559c\u6b22", "v"}, WordTag{"\u7684", "uj"}, WordTag{"\u8863\u670d", "n"}},
		[]WordTag{WordTag{"\u80e1\u9526\u6d9b", "nr"}, WordTag{"\u662f", "v"}, WordTag{"\u70ed\u7231", "a"}, WordTag{"\u4e16\u754c", "n"}, WordTag{"\u548c\u5e73", "nz"}, WordTag{"\u7684", "uj"}, WordTag{"\u653f\u6cbb\u5c40", "n"}, WordTag{"\u5e38\u59d4", "j"}},
		[]WordTag{WordTag{"\u7a0b\u5e8f\u5458", "n"}, WordTag{"\u795d", "v"}, WordTag{"\u6d77\u6797", "nz"}, WordTag{"\u548c", "c"}, WordTag{"\u6731", "nr"}, WordTag{"\u4f1a", "v"}, WordTag{"\u9707", "v"}, WordTag{"\u662f", "v"}, WordTag{"\u5728", "p"}, WordTag{"\u5b59", "zg"}, WordTag{"\u5065", "a"}, WordTag{"\u7684", "uj"}, WordTag{"\u5de6\u9762", "f"}, WordTag{"\u548c", "c"}, WordTag{"\u53f3\u9762", "f"}, WordTag{",", "x"}, WordTag{" ", "x"}, WordTag{"\u8303", "nr"}, WordTag{"\u51ef", "nr"}, WordTag{"\u5728", "p"}, WordTag{"\u6700", "d"}, WordTag{"\u53f3\u9762", "f"}, WordTag{".", "x"}, WordTag{"\u518d", "d"}, WordTag{"\u5f80", "zg"}, WordTag{"\u5de6", "m"}, WordTag{"\u662f", "v"}, WordTag{"\u674e", "nr"}, WordTag{"\u677e", "v"}, WordTag{"\u6d2a", "nr"}},
		[]WordTag{WordTag{"\u4e00\u6b21\u6027", "d"}, WordTag{"\u4ea4", "v"}, WordTag{"\u591a\u5c11", "m"}, WordTag{"\u94b1", "n"}},
		[]WordTag{WordTag{"\u4e24\u5757", "m"}, WordTag{"\u4e94", "m"}, WordTag{"\u4e00\u5957", "m"}, WordTag{"\uff0c", "x"}, WordTag{"\u4e09\u5757", "m"}, WordTag{"\u516b", "m"}, WordTag{"\u4e00\u65a4", "m"}, WordTag{"\uff0c", "x"}, WordTag{"\u56db\u5757", "m"}, WordTag{"\u4e03", "m"}, WordTag{"\u4e00\u672c", "m"}, WordTag{"\uff0c", "x"}, WordTag{"\u4e94\u5757", "m"}, WordTag{"\u516d", "m"}, WordTag{"\u4e00\u6761", "m"}},
		[]WordTag{WordTag{"\u5c0f", "a"}, WordTag{"\u548c\u5c1a", "nr"}, WordTag{"\u7559", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u4e00\u4e2a", "m"}, WordTag{"\u50cf", "v"}, WordTag{"\u5927", "a"}, WordTag{"\u548c\u5c1a", "nr"}, WordTag{"\u4e00\u6837", "r"}, WordTag{"\u7684", "uj"}, WordTag{"\u548c\u5c1a\u5934", "nr"}},
		[]WordTag{WordTag{"\u6211", "r"}, WordTag{"\u662f", "v"}, WordTag{"\u4e2d\u534e\u4eba\u6c11\u5171\u548c\u56fd", "ns"}, WordTag{"\u516c\u6c11", "n"}, WordTag{";", "x"}, WordTag{"\u6211", "r"}, WordTag{"\u7238\u7238", "n"}, WordTag{"\u662f", "v"}, WordTag{"\u5171\u548c\u515a", "nt"}, WordTag{"\u515a\u5458", "n"}, WordTag{";", "x"}, WordTag{" ", "x"}, WordTag{"\u5730\u94c1", "n"}, WordTag{"\u548c\u5e73\u95e8", "ns"}, WordTag{"\u7ad9", "v"}},
		[]WordTag{WordTag{"\u5f20\u6653\u6885", "nr"}, WordTag{"\u53bb", "v"}, WordTag{"\u4eba\u6c11", "n"}, WordTag{"\u533b\u9662", "n"}, WordTag{"\u505a", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u4e2a", "q"}, WordTag{"B\u8d85", "n"}, WordTag{"\u7136\u540e", "c"}, WordTag{"\u53bb", "v"}, WordTag{"\u4e70", "v"}, WordTag{"\u4e86", "ul"}, WordTag{"\u4ef6", "zg"}, WordTag{"T\u6064", "n"}},
		[]WordTag{WordTag{"AT&T", "nz"}, WordTag{"\u662f", "v"}, WordTag{"\u4e00\u4ef6", "m"}, WordTag{"\u4e0d\u9519", "a"}, WordTag{"\u7684", "uj"}, WordTag{"\u516c\u53f8", "n"}, WordTag{"\uff0c", "x"}, WordTag{"\u7ed9", "p"}, WordTag{"\u4f60", "r"}, WordTag{"\u53d1", "v"}, WordTag{"offer", "eng"}, WordTag{"\u4e86", "ul"}, WordTag{"\u5417", "y"}, WordTag{"\uff1f", "x"}},
		[]WordTag{WordTag{"C++", "nz"}, WordTag{"\u548c", "c"}, WordTag{"c#", "nz"}, WordTag{"\u662f", "v"}, WordTag{"\u4ec0\u4e48", "r"}, WordTag{"\u5173\u7cfb", "n"}, WordTag{"\uff1f", "x"}, WordTag{"11", "eng"}, WordTag{"+", "x"}, WordTag{"122", "eng"}, WordTag{"=", "x"}, WordTag{"133", "eng"}, WordTag{"\uff0c", "x"}, WordTag{"\u662f", "v"}, WordTag{"\u5417", "y"}, WordTag{"\uff1f", "x"}, WordTag{"PI", "eng"}, WordTag{"=", "x"}, WordTag{"3", "eng"}, WordTag{".", "x"}, WordTag{"14159", "eng"}},
		[]WordTag{WordTag{"\u4f60", "r"}, WordTag{"\u8ba4\u8bc6", "v"}, WordTag{"\u90a3\u4e2a", "r"}, WordTag{"\u548c", "c"}, WordTag{"\u4e3b\u5e2d", "n"}, WordTag{"\u63e1\u624b", "v"}, WordTag{"\u7684", "uj"}, WordTag{"\u7684\u54e5", "n"}, WordTag{"\u5417", "y"}, WordTag{"\uff1f", "x"}, WordTag{"\u4ed6", "r"}, WordTag{"\u5f00", "v"}, WordTag{"\u4e00\u8f86", "m"}, WordTag{"\u9ed1\u8272", "n"}, WordTag{"\u7684\u58eb", "n"}, WordTag{"\u3002", "x"}},
		[]WordTag{WordTag{"\u67aa\u6746\u5b50", "n"}, WordTag{"\u4e2d", "f"}, WordTag{"\u51fa", "v"}, WordTag{"\u653f\u6743", "n"}},
	}
)

func TestCut(t *testing.T) {
	jiebago.SetDictionary("../dict.txt")
	for index, content := range test_contents {
		result := Cut(content, true)
		if len(defaultCutResult[index]) != len(result) {
			t.Error(content)
		}
		for i, _ := range result {
			if result[i] != defaultCutResult[index][i] {
				t.Error(content)
			}
		}
		result = Cut(content, false)
		if len(noHMMCutResult[index]) != len(result) {
			t.Error(content)
		}
		for i, _ := range result {
			if result[i] != noHMMCutResult[index][i] {
				t.Error(content)
			}
		}

	}
}

func TestUserDict(t *testing.T) {
	jiebago.SetDictionary("../dict.txt")
	jiebago.LoadUserDict("../userdict.txt")
	sentence := "李小福是创新办主任也是云计算方面的专家; 什么是八一双鹿例如我输入一个带“韩玉赏鉴”的标题，在自定义词库中也增加了此词为N类型"
	cutResult := []WordTag{
		WordTag{"\u674e\u5c0f\u798f", "nr"},
		WordTag{"\u662f", "v"},
		WordTag{"\u521b\u65b0\u529e", "i"},
		WordTag{"\u4e3b\u4efb", "b"},
		WordTag{"\u4e5f", "d"},
		WordTag{"\u662f", "v"},
		WordTag{"\u4e91\u8ba1\u7b97", "x"},
		WordTag{"\u65b9\u9762", "n"},
		WordTag{"\u7684", "uj"},
		WordTag{"\u4e13\u5bb6", "n"},
		WordTag{";", "x"},
		WordTag{" ", "x"},
		WordTag{"\u4ec0\u4e48", "r"},
		WordTag{"\u662f", "v"},
		WordTag{"\u516b\u4e00\u53cc\u9e7f", "nz"},
		WordTag{"\u4f8b\u5982", "v"},
		WordTag{"\u6211", "r"},
		WordTag{"\u8f93\u5165", "v"},
		WordTag{"\u4e00\u4e2a", "m"},
		WordTag{"\u5e26", "v"},
		WordTag{"\u201c", "x"},
		WordTag{"\u97e9\u7389\u8d4f\u9274", "nz"},
		WordTag{"\u201d", "x"},
		WordTag{"\u7684", "uj"},
		WordTag{"\u6807\u9898", "n"},
		WordTag{"\uff0c", "x"},
		WordTag{"\u5728", "p"},
		WordTag{"\u81ea\u5b9a\u4e49\u8bcd", "n"},
		WordTag{"\u5e93\u4e2d", "nrt"},
		WordTag{"\u4e5f", "d"},
		WordTag{"\u589e\u52a0", "v"},
		WordTag{"\u4e86", "ul"},
		WordTag{"\u6b64", "r"},
		WordTag{"\u8bcd", "n"},
		WordTag{"\u4e3a", "p"},
		WordTag{"N", "eng"},
		WordTag{"\u7c7b\u578b", "n"},
	}

	result := Cut(sentence, true)
	if len(cutResult) != len(result) {
		t.Error(result)
	}
	for i, _ := range result {
		if result[i] != cutResult[i] {
			t.Error(result[i])
		}
	}
}

func TestBug132(t *testing.T) {
	/*
		https://github.com/fxsjy/jieba/issues/132
	*/
	jiebago.SetDictionary("../dict.txt")
	sentence := "又跛又啞"
	cutResult := []WordTag{
		WordTag{"\u53c8", "d"},
		WordTag{"\u8ddb", "a"},
		WordTag{"\u53c8", "d"},
		WordTag{"\u555e", "v"},
	}
	result := Cut(sentence, true)
	if len(cutResult) != len(result) {
		t.Error(result)
	}
	for i, _ := range result {
		if result[i] != cutResult[i] {
			t.Error(result[i])
		}
	}
}
