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

	defaultCutResult = [][]WordTag{[]WordTag{WorTag{"这", "r"}, WorTag{"是", "v"}, WorTag{"一个", "m"}, WorTag{"伸手不见五指", "i"}, WorTag{"的", "uj"}, WorTag{"黑夜", "n"}, WorTag{"。", "x"}, WorTag{"我", "r"}, WorTag{"叫", "v"}, WorTag{"孙悟空", "nr"}, WorTag{"，", "x"}, WorTag{"我", "r"}, WorTag{"爱", "v"}, WorTag{"北京", "ns"}, WorTag{"，", "x"}, WorTag{"我", "r"}, WorTag{"爱", "v"}, WorTag{"Python", "eng"}, WorTag{"和", "c"}, WorTag{"C++", "nz"}, WorTag{"。", "x"}},
		[]WordTag{WorTag{"我", "r"}, WorTag{"不", "d"}, WorTag{"喜欢", "v"}, WorTag{"日本", "ns"}, WorTag{"和服", "nz"}, WorTag{"。", "x"}},
		[]WordTag{WorTag{"雷猴", "n"}, WorTag{"回归", "v"}, WorTag{"人间", "n"}, WorTag{"。", "x"}},
		[]WordTag{WorTag{"工信处", "n"}, WorTag{"女干事", "n"}, WorTag{"每月", "r"}, WorTag{"经过", "p"}, WorTag{"下属", "v"}, WorTag{"科室", "n"}, WorTag{"都", "d"}, WorTag{"要", "v"}, WorTag{"亲口", "n"}, WorTag{"交代", "n"}, WorTag{"24", "m"}, WorTag{"口", "n"}, WorTag{"交换机", "n"}, WorTag{"等", "u"}, WorTag{"技术性", "n"}, WorTag{"器件", "n"}, WorTag{"的", "uj"}, WorTag{"安装", "v"}, WorTag{"工作", "vn"}},
		[]WordTag{WorTag{"我", "r"}, WorTag{"需要", "v"}, WorTag{"廉租房", "n"}},
		[]WordTag{WorTag{"永和", "nz"}, WorTag{"服装", "vn"}, WorTag{"饰品", "n"}, WorTag{"有限公司", "n"}},
		[]WordTag{WorTag{"我", "r"}, WorTag{"爱", "v"}, WorTag{"北京", "ns"}, WorTag{"天安门", "ns"}},
		[]WordTag{WorTag{"abc", "eng"}},
		[]WordTag{WorTag{"隐", "n"}, WorTag{"马尔可夫", "nr"}},
		[]WordTag{WorTag{"雷猴", "n"}, WorTag{"是", "v"}, WorTag{"个", "q"}, WorTag{"好", "a"}, WorTag{"网站", "n"}},
		[]WordTag{WorTag{"“", "x"}, WorTag{"Microsoft", "eng"}, WorTag{"”", "x"}, WorTag{"一", "m"}, WorTag{"词", "n"}, WorTag{"由", "p"}, WorTag{"“", "x"}, WorTag{"MICROcomputer", "eng"}, WorTag{"（", "x"}, WorTag{"微型", "b"}, WorTag{"计算机", "n"}, WorTag{"）", "x"}, WorTag{"”", "x"}, WorTag{"和", "c"}, WorTag{"“", "x"}, WorTag{"SOFTware", "eng"}, WorTag{"（", "x"}, WorTag{"软件", "n"}, WorTag{"）", "x"}, WorTag{"”", "x"}, WorTag{"两", "m"}, WorTag{"部分", "n"}, WorTag{"组成", "v"}},
		[]WordTag{WorTag{"草泥马", "n"}, WorTag{"和", "c"}, WorTag{"欺实", "v"}, WorTag{"马", "n"}, WorTag{"是", "v"}, WorTag{"今年", "t"}, WorTag{"的", "uj"}, WorTag{"流行", "v"}, WorTag{"词汇", "n"}},
		[]WordTag{WorTag{"伊藤", "nr"}, WorTag{"洋华堂", "n"}, WorTag{"总府", "n"}, WorTag{"店", "n"}},
		[]WordTag{WorTag{"中国科学院计算技术研究所", "nt"}},
		[]WordTag{WorTag{"罗密欧", "nr"}, WorTag{"与", "p"}, WorTag{"朱丽叶", "nr"}},
		[]WordTag{WorTag{"我", "r"}, WorTag{"购买", "v"}, WorTag{"了", "ul"}, WorTag{"道具", "n"}, WorTag{"和", "c"}, WorTag{"服装", "vn"}},
		[]WordTag{WorTag{"PS", "eng"}, WorTag{":", "x"}, WorTag{" ", "x"}, WorTag{"我", "r"}, WorTag{"觉得", "v"}, WorTag{"开源", "n"}, WorTag{"有", "v"}, WorTag{"一个", "m"}, WorTag{"好处", "d"}, WorTag{"，", "x"}, WorTag{"就是", "d"}, WorTag{"能够", "v"}, WorTag{"敦促", "v"}, WorTag{"自己", "r"}, WorTag{"不断改进", "l"}, WorTag{"，", "x"}, WorTag{"避免", "v"}, WorTag{"敞", "v"}, WorTag{"帚", "ng"}, WorTag{"自珍", "b"}},
		[]WordTag{WorTag{"湖北省", "ns"}, WorTag{"石首市", "ns"}},
		[]WordTag{WorTag{"湖北省", "ns"}, WorTag{"十堰市", "ns"}},
		[]WordTag{WorTag{"总经理", "n"}, WorTag{"完成", "v"}, WorTag{"了", "ul"}, WorTag{"这件", "mq"}, WorTag{"事情", "n"}},
		[]WordTag{WorTag{"电脑", "n"}, WorTag{"修好", "v"}, WorTag{"了", "ul"}},
		[]WordTag{WorTag{"做好", "v"}, WorTag{"了", "ul"}, WorTag{"这件", "mq"}, WorTag{"事情", "n"}, WorTag{"就", "d"}, WorTag{"一了百了", "l"}, WorTag{"了", "ul"}},
		[]WordTag{WorTag{"人们", "n"}, WorTag{"审美", "vn"}, WorTag{"的", "uj"}, WorTag{"观点", "n"}, WorTag{"是", "v"}, WorTag{"不同", "a"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"我们", "r"}, WorTag{"买", "v"}, WorTag{"了", "ul"}, WorTag{"一个", "m"}, WorTag{"美的", "nr"}, WorTag{"空调", "n"}},
		[]WordTag{WorTag{"线程", "n"}, WorTag{"初始化", "l"}, WorTag{"时", "n"}, WorTag{"我们", "r"}, WorTag{"要", "v"}, WorTag{"注意", "v"}},
		[]WordTag{WorTag{"一个", "m"}, WorTag{"分子", "n"}, WorTag{"是", "v"}, WorTag{"由", "p"}, WorTag{"好多", "m"}, WorTag{"原子", "n"}, WorTag{"组织", "v"}, WorTag{"成", "v"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"祝", "v"}, WorTag{"你", "r"}, WorTag{"马到功成", "i"}},
		[]WordTag{WorTag{"他", "r"}, WorTag{"掉", "v"}, WorTag{"进", "v"}, WorTag{"了", "ul"}, WorTag{"无底洞", "ns"}, WorTag{"里", "f"}},
		[]WordTag{WorTag{"中国", "ns"}, WorTag{"的", "uj"}, WorTag{"首都", "d"}, WorTag{"是", "v"}, WorTag{"北京", "ns"}},
		[]WordTag{WorTag{"孙君意", "nr"}},
		[]WordTag{WorTag{"外交部", "nt"}, WorTag{"发言人", "l"}, WorTag{"马朝旭", "nr"}},
		[]WordTag{WorTag{"领导人", "n"}, WorTag{"会议", "n"}, WorTag{"和", "c"}, WorTag{"第四届", "m"}, WorTag{"东亚", "ns"}, WorTag{"峰会", "n"}},
		[]WordTag{WorTag{"在", "p"}, WorTag{"过去", "t"}, WorTag{"的", "uj"}, WorTag{"这", "r"}, WorTag{"五年", "t"}},
		[]WordTag{WorTag{"还", "d"}, WorTag{"需要", "v"}, WorTag{"很", "d"}, WorTag{"长", "a"}, WorTag{"的", "uj"}, WorTag{"路", "n"}, WorTag{"要", "v"}, WorTag{"走", "v"}},
		[]WordTag{WorTag{"60", "m"}, WorTag{"周年", "t"}, WorTag{"首都", "d"}, WorTag{"阅兵", "v"}},
		[]WordTag{WorTag{"你好", "l"}, WorTag{"人们", "n"}, WorTag{"审美", "vn"}, WorTag{"的", "uj"}, WorTag{"观点", "n"}, WorTag{"是", "v"}, WorTag{"不同", "a"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"买", "v"}, WorTag{"水果", "n"}, WorTag{"然后", "c"}, WorTag{"来", "v"}, WorTag{"世博园", "nr"}},
		[]WordTag{WorTag{"买", "v"}, WorTag{"水果", "n"}, WorTag{"然后", "c"}, WorTag{"去", "v"}, WorTag{"世博园", "nr"}},
		[]WordTag{WorTag{"但是", "c"}, WorTag{"后来", "t"}, WorTag{"我", "r"}, WorTag{"才", "d"}, WorTag{"知道", "v"}, WorTag{"你", "r"}, WorTag{"是", "v"}, WorTag{"对", "p"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"存在", "v"}, WorTag{"即", "v"}, WorTag{"合理", "vn"}},
		[]WordTag{WorTag{"的的", "u"}, WorTag{"的的", "u"}, WorTag{"的", "uj"}, WorTag{"在的", "u"}, WorTag{"的的", "u"}, WorTag{"的", "uj"}, WorTag{"就", "d"}, WorTag{"以", "p"}, WorTag{"和和", "nz"}, WorTag{"和", "c"}},
		[]WordTag{WorTag{"I", "x"}, WorTag{" ", "x"}, WorTag{"love", "eng"}, WorTag{"你", "r"}, WorTag{"，", "x"}, WorTag{"不以为耻", "i"}, WorTag{"，", "x"}, WorTag{"反", "zg"}, WorTag{"以为", "c"}, WorTag{"rong", "eng"}},
		[]WordTag{WorTag{"因", "p"}},
		[]WordTag{},
		[]WordTag{WorTag{"hello", "eng"}, WorTag{"你好", "l"}, WorTag{"人们", "n"}, WorTag{"审美", "vn"}, WorTag{"的", "uj"}, WorTag{"观点", "n"}, WorTag{"是", "v"}, WorTag{"不同", "a"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"很好", "a"}, WorTag{"但", "c"}, WorTag{"主要", "b"}, WorTag{"是", "v"}, WorTag{"基于", "p"}, WorTag{"网页", "n"}, WorTag{"形式", "n"}},
		[]WordTag{WorTag{"hello", "eng"}, WorTag{"你好", "l"}, WorTag{"人们", "n"}, WorTag{"审美", "vn"}, WorTag{"的", "uj"}, WorTag{"观点", "n"}, WorTag{"是", "v"}, WorTag{"不同", "a"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"为什么", "r"}, WorTag{"我", "r"}, WorTag{"不能", "v"}, WorTag{"拥有", "v"}, WorTag{"想要", "v"}, WorTag{"的", "uj"}, WorTag{"生活", "vn"}},
		[]WordTag{WorTag{"后来", "t"}, WorTag{"我", "r"}, WorTag{"才", "d"}},
		[]WordTag{WorTag{"此次", "r"}, WorTag{"来", "v"}, WorTag{"中国", "ns"}, WorTag{"是", "v"}, WorTag{"为了", "p"}},
		[]WordTag{WorTag{"使用", "v"}, WorTag{"了", "ul"}, WorTag{"它", "r"}, WorTag{"就", "d"}, WorTag{"可以", "c"}, WorTag{"解决", "v"}, WorTag{"一些", "m"}, WorTag{"问题", "n"}},
		[]WordTag{WorTag{",", "x"}, WorTag{"使用", "v"}, WorTag{"了", "ul"}, WorTag{"它", "r"}, WorTag{"就", "d"}, WorTag{"可以", "c"}, WorTag{"解决", "v"}, WorTag{"一些", "m"}, WorTag{"问题", "n"}},
		[]WordTag{WorTag{"其实", "d"}, WorTag{"使用", "v"}, WorTag{"了", "ul"}, WorTag{"它", "r"}, WorTag{"就", "d"}, WorTag{"可以", "c"}, WorTag{"解决", "v"}, WorTag{"一些", "m"}, WorTag{"问题", "n"}},
		[]WordTag{WorTag{"好人", "n"}, WorTag{"使用", "v"}, WorTag{"了", "ul"}, WorTag{"它", "r"}, WorTag{"就", "d"}, WorTag{"可以", "c"}, WorTag{"解决", "v"}, WorTag{"一些", "m"}, WorTag{"问题", "n"}},
		[]WordTag{WorTag{"是因为", "c"}, WorTag{"和", "c"}, WorTag{"国家", "n"}},
		[]WordTag{WorTag{"老年", "t"}, WorTag{"搜索", "v"}, WorTag{"还", "d"}, WorTag{"支持", "v"}},
		[]WordTag{WorTag{"干脆", "d"}, WorTag{"就", "d"}, WorTag{"把", "p"}, WorTag{"那部", "r"}, WorTag{"蒙人", "n"}, WorTag{"的", "uj"}, WorTag{"闲法", "n"}, WorTag{"给", "p"}, WorTag{"废", "v"}, WorTag{"了", "ul"}, WorTag{"拉倒", "v"}, WorTag{"！", "x"}, WorTag{"RT", "eng"}, WorTag{" ", "x"}, WorTag{"@", "x"}, WorTag{"laoshipukong", "eng"}, WorTag{" ", "x"}, WorTag{":", "x"}, WorTag{" ", "x"}, WorTag{"27", "m"}, WorTag{"日", "m"}, WorTag{"，", "x"}, WorTag{"全国人大常委会", "nt"}, WorTag{"第三次", "m"}, WorTag{"审议", "v"}, WorTag{"侵权", "v"}, WorTag{"责任法", "n"}, WorTag{"草案", "n"}, WorTag{"，", "x"}, WorTag{"删除", "v"}, WorTag{"了", "ul"}, WorTag{"有关", "vn"}, WorTag{"医疗", "n"}, WorTag{"损害", "v"}, WorTag{"责任", "n"}, WorTag{"“", "x"}, WorTag{"举证", "v"}, WorTag{"倒置", "v"}, WorTag{"”", "x"}, WorTag{"的", "uj"}, WorTag{"规定", "n"}, WorTag{"。", "x"}, WorTag{"在", "p"}, WorTag{"医患", "n"}, WorTag{"纠纷", "n"}, WorTag{"中本", "ns"}, WorTag{"已", "d"}, WorTag{"处于", "v"}, WorTag{"弱势", "n"}, WorTag{"地位", "n"}, WorTag{"的", "uj"}, WorTag{"消费者", "n"}, WorTag{"由此", "c"}, WorTag{"将", "d"}, WorTag{"陷入", "v"}, WorTag{"万劫不复", "i"}, WorTag{"的", "uj"}, WorTag{"境地", "s"}, WorTag{"。", "x"}, WorTag{" ", "x"}},
		[]WordTag{WorTag{"大", "a"}},
		[]WordTag{},
		[]WordTag{WorTag{"他", "r"}, WorTag{"说", "v"}, WorTag{"的", "uj"}, WorTag{"确实", "ad"}, WorTag{"在", "p"}, WorTag{"理", "n"}},
		[]WordTag{WorTag{"长春", "ns"}, WorTag{"市长", "n"}, WorTag{"春节", "t"}, WorTag{"讲话", "n"}},
		[]WordTag{WorTag{"结婚", "v"}, WorTag{"的", "uj"}, WorTag{"和", "c"}, WorTag{"尚未", "d"}, WorTag{"结婚", "v"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"结合", "v"}, WorTag{"成", "n"}, WorTag{"分子", "n"}, WorTag{"时", "n"}},
		[]WordTag{WorTag{"旅游", "vn"}, WorTag{"和", "c"}, WorTag{"服务", "vn"}, WorTag{"是", "v"}, WorTag{"最好", "a"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"这件", "mq"}, WorTag{"事情", "n"}, WorTag{"的确", "d"}, WorTag{"是", "v"}, WorTag{"我", "r"}, WorTag{"的", "uj"}, WorTag{"错", "n"}},
		[]WordTag{WorTag{"供", "v"}, WorTag{"大家", "n"}, WorTag{"参考", "v"}, WorTag{"指正", "v"}},
		[]WordTag{WorTag{"哈尔滨", "ns"}, WorTag{"政府", "n"}, WorTag{"公布", "v"}, WorTag{"塌", "v"}, WorTag{"桥", "n"}, WorTag{"原因", "n"}},
		[]WordTag{WorTag{"我", "r"}, WorTag{"在", "p"}, WorTag{"机场", "n"}, WorTag{"入口处", "i"}},
		[]WordTag{WorTag{"邢永臣", "nr"}, WorTag{"摄影", "n"}, WorTag{"报道", "v"}},
		[]WordTag{WorTag{"BP", "eng"}, WorTag{"神经网络", "n"}, WorTag{"如何", "r"}, WorTag{"训练", "vn"}, WorTag{"才能", "v"}, WorTag{"在", "p"}, WorTag{"分类", "n"}, WorTag{"时", "n"}, WorTag{"增加", "v"}, WorTag{"区分度", "n"}, WorTag{"？", "x"}},
		[]WordTag{WorTag{"南京市", "ns"}, WorTag{"长江大桥", "ns"}},
		[]WordTag{WorTag{"应", "v"}, WorTag{"一些", "m"}, WorTag{"使用者", "n"}, WorTag{"的", "uj"}, WorTag{"建议", "n"}, WorTag{"，", "x"}, WorTag{"也", "d"}, WorTag{"为了", "p"}, WorTag{"便于", "v"}, WorTag{"利用", "n"}, WorTag{"NiuTrans", "eng"}, WorTag{"用于", "v"}, WorTag{"SMT", "eng"}, WorTag{"研究", "vn"}},
		[]WordTag{WorTag{"长春市", "ns"}, WorTag{"长春", "ns"}, WorTag{"药店", "n"}},
		[]WordTag{WorTag{"邓颖超", "nr"}, WorTag{"生前", "t"}, WorTag{"最", "d"}, WorTag{"喜欢", "v"}, WorTag{"的", "uj"}, WorTag{"衣服", "n"}},
		[]WordTag{WorTag{"胡锦涛", "nr"}, WorTag{"是", "v"}, WorTag{"热爱", "a"}, WorTag{"世界", "n"}, WorTag{"和平", "nz"}, WorTag{"的", "uj"}, WorTag{"政治局", "n"}, WorTag{"常委", "j"}},
		[]WordTag{WorTag{"程序员", "n"}, WorTag{"祝", "v"}, WorTag{"海林", "nz"}, WorTag{"和", "c"}, WorTag{"朱会震", "nr"}, WorTag{"是", "v"}, WorTag{"在", "p"}, WorTag{"孙健", "nr"}, WorTag{"的", "uj"}, WorTag{"左面", "f"}, WorTag{"和", "c"}, WorTag{"右面", "f"}, WorTag{",", "x"}, WorTag{" ", "x"}, WorTag{"范凯", "nr"}, WorTag{"在", "p"}, WorTag{"最", "a"}, WorTag{"右面", "f"}, WorTag{".", "m"}, WorTag{"再往", "d"}, WorTag{"左", "f"}, WorTag{"是", "v"}, WorTag{"李松洪", "nr"}},
		[]WordTag{WorTag{"一次性", "d"}, WorTag{"交", "v"}, WorTag{"多少", "m"}, WorTag{"钱", "n"}},
		[]WordTag{WorTag{"两块", "m"}, WorTag{"五", "m"}, WorTag{"一套", "m"}, WorTag{"，", "x"}, WorTag{"三块", "m"}, WorTag{"八", "m"}, WorTag{"一斤", "m"}, WorTag{"，", "x"}, WorTag{"四块", "m"}, WorTag{"七", "m"}, WorTag{"一本", "m"}, WorTag{"，", "x"}, WorTag{"五块", "m"}, WorTag{"六", "m"}, WorTag{"一条", "m"}},
		[]WordTag{WorTag{"小", "a"}, WorTag{"和尚", "nr"}, WorTag{"留", "v"}, WorTag{"了", "ul"}, WorTag{"一个", "m"}, WorTag{"像", "v"}, WorTag{"大", "a"}, WorTag{"和尚", "nr"}, WorTag{"一样", "r"}, WorTag{"的", "uj"}, WorTag{"和尚头", "nr"}},
		[]WordTag{WorTag{"我", "r"}, WorTag{"是", "v"}, WorTag{"中华人民共和国", "ns"}, WorTag{"公民", "n"}, WorTag{";", "x"}, WorTag{"我", "r"}, WorTag{"爸爸", "n"}, WorTag{"是", "v"}, WorTag{"共和党", "nt"}, WorTag{"党员", "n"}, WorTag{";", "x"}, WorTag{" ", "x"}, WorTag{"地铁", "n"}, WorTag{"和平门", "ns"}, WorTag{"站", "v"}},
		[]WordTag{WorTag{"张晓梅", "nr"}, WorTag{"去", "v"}, WorTag{"人民", "n"}, WorTag{"医院", "n"}, WorTag{"做", "v"}, WorTag{"了", "ul"}, WorTag{"个", "q"}, WorTag{"B超", "n"}, WorTag{"然后", "c"}, WorTag{"去", "v"}, WorTag{"买", "v"}, WorTag{"了", "ul"}, WorTag{"件", "q"}, WorTag{"T恤", "n"}},
		[]WordTag{WorTag{"AT&T", "nz"}, WorTag{"是", "v"}, WorTag{"一件", "m"}, WorTag{"不错", "a"}, WorTag{"的", "uj"}, WorTag{"公司", "n"}, WorTag{"，", "x"}, WorTag{"给", "p"}, WorTag{"你", "r"}, WorTag{"发", "v"}, WorTag{"offer", "eng"}, WorTag{"了", "ul"}, WorTag{"吗", "y"}, WorTag{"？", "x"}},
		[]WordTag{WorTag{"C++", "nz"}, WorTag{"和", "c"}, WorTag{"c#", "nz"}, WorTag{"是", "v"}, WorTag{"什么", "r"}, WorTag{"关系", "n"}, WorTag{"？", "x"}, WorTag{"11", "m"}, WorTag{"+", "x"}, WorTag{"122", "m"}, WorTag{"=", "x"}, WorTag{"133", "m"}, WorTag{"，", "x"}, WorTag{"是", "v"}, WorTag{"吗", "y"}, WorTag{"？", "x"}, WorTag{"PI", "eng"}, WorTag{"=", "x"}, WorTag{"3.14159", "m"}},
		[]WordTag{WorTag{"你", "r"}, WorTag{"认识", "v"}, WorTag{"那个", "r"}, WorTag{"和", "c"}, WorTag{"主席", "n"}, WorTag{"握手", "v"}, WorTag{"的", "uj"}, WorTag{"的哥", "n"}, WorTag{"吗", "y"}, WorTag{"？", "x"}, WorTag{"他", "r"}, WorTag{"开", "v"}, WorTag{"一辆", "m"}, WorTag{"黑色", "n"}, WorTag{"的士", "n"}, WorTag{"。", "x"}},
		[]WordTag{WorTag{"枪杆子", "n"}, WorTag{"中", "f"}, WorTag{"出", "v"}, WorTag{"政权", "n"}},
	}
	noHMMCutResult = [][]WordTag{
		[]WordTag{WorTag{"这", "r"}, WorTag{"是", "v"}, WorTag{"一个", "m"}, WorTag{"伸手不见五指", "i"}, WorTag{"的", "uj"}, WorTag{"黑夜", "n"}, WorTag{"。", "x"}, WorTag{"我", "r"}, WorTag{"叫", "v"}, WorTag{"孙悟空", "nr"}, WorTag{"，", "x"}, WorTag{"我", "r"}, WorTag{"爱", "v"}, WorTag{"北京", "ns"}, WorTag{"，", "x"}, WorTag{"我", "r"}, WorTag{"爱", "v"}, WorTag{"Python", "eng"}, WorTag{"和", "c"}, WorTag{"C++", "nz"}, WorTag{"。", "x"}},
		[]WordTag{WorTag{"我", "r"}, WorTag{"不", "d"}, WorTag{"喜欢", "v"}, WorTag{"日本", "ns"}, WorTag{"和服", "nz"}, WorTag{"。", "x"}},
		[]WordTag{WorTag{"雷猴", "n"}, WorTag{"回归", "v"}, WorTag{"人间", "n"}, WorTag{"。", "x"}},
		[]WordTag{WorTag{"工信处", "n"}, WorTag{"女干事", "n"}, WorTag{"每月", "r"}, WorTag{"经过", "p"}, WorTag{"下属", "v"}, WorTag{"科室", "n"}, WorTag{"都", "d"}, WorTag{"要", "v"}, WorTag{"亲口", "n"}, WorTag{"交代", "n"}, WorTag{"24", "eng"}, WorTag{"口", "q"}, WorTag{"交换机", "n"}, WorTag{"等", "u"}, WorTag{"技术性", "n"}, WorTag{"器件", "n"}, WorTag{"的", "uj"}, WorTag{"安装", "v"}, WorTag{"工作", "vn"}},
		[]WordTag{WorTag{"我", "r"}, WorTag{"需要", "v"}, WorTag{"廉租房", "n"}},
		[]WordTag{WorTag{"永和", "nz"}, WorTag{"服装", "vn"}, WorTag{"饰品", "n"}, WorTag{"有限公司", "n"}},
		[]WordTag{WorTag{"我", "r"}, WorTag{"爱", "v"}, WorTag{"北京", "ns"}, WorTag{"天安门", "ns"}},
		[]WordTag{WorTag{"abc", "eng"}},
		[]WordTag{WorTag{"隐", "n"}, WorTag{"马尔可夫", "nr"}},
		[]WordTag{WorTag{"雷猴", "n"}, WorTag{"是", "v"}, WorTag{"个", "q"}, WorTag{"好", "a"}, WorTag{"网站", "n"}},
		[]WordTag{WorTag{"“", "x"}, WorTag{"Microsoft", "eng"}, WorTag{"”", "x"}, WorTag{"一", "m"}, WorTag{"词", "n"}, WorTag{"由", "p"}, WorTag{"“", "x"}, WorTag{"MICROcomputer", "eng"}, WorTag{"（", "x"}, WorTag{"微型", "b"}, WorTag{"计算机", "n"}, WorTag{"）", "x"}, WorTag{"”", "x"}, WorTag{"和", "c"}, WorTag{"“", "x"}, WorTag{"SOFTware", "eng"}, WorTag{"（", "x"}, WorTag{"软件", "n"}, WorTag{"）", "x"}, WorTag{"”", "x"}, WorTag{"两", "m"}, WorTag{"部分", "n"}, WorTag{"组成", "v"}},
		[]WordTag{WorTag{"草泥马", "n"}, WorTag{"和", "c"}, WorTag{"欺", "vn"}, WorTag{"实", "n"}, WorTag{"马", "n"}, WorTag{"是", "v"}, WorTag{"今年", "t"}, WorTag{"的", "uj"}, WorTag{"流行", "v"}, WorTag{"词汇", "n"}},
		[]WordTag{WorTag{"伊", "ns"}, WorTag{"藤", "nr"}, WorTag{"洋华堂", "n"}, WorTag{"总府", "n"}, WorTag{"店", "n"}},
		[]WordTag{WorTag{"中国科学院计算技术研究所", "nt"}},
		[]WordTag{WorTag{"罗密欧", "nr"}, WorTag{"与", "p"}, WorTag{"朱丽叶", "nr"}},
		[]WordTag{WorTag{"我", "r"}, WorTag{"购买", "v"}, WorTag{"了", "ul"}, WorTag{"道具", "n"}, WorTag{"和", "c"}, WorTag{"服装", "vn"}},
		[]WordTag{WorTag{"PS", "eng"}, WorTag{":", "x"}, WorTag{" ", "x"}, WorTag{"我", "r"}, WorTag{"觉得", "v"}, WorTag{"开源", "n"}, WorTag{"有", "v"}, WorTag{"一个", "m"}, WorTag{"好处", "d"}, WorTag{"，", "x"}, WorTag{"就是", "d"}, WorTag{"能够", "v"}, WorTag{"敦促", "v"}, WorTag{"自己", "r"}, WorTag{"不断改进", "l"}, WorTag{"，", "x"}, WorTag{"避免", "v"}, WorTag{"敞", "v"}, WorTag{"帚", "ng"}, WorTag{"自珍", "b"}},
		[]WordTag{WorTag{"湖北省", "ns"}, WorTag{"石首市", "ns"}},
		[]WordTag{WorTag{"湖北省", "ns"}, WorTag{"十堰市", "ns"}},
		[]WordTag{WorTag{"总经理", "n"}, WorTag{"完成", "v"}, WorTag{"了", "ul"}, WorTag{"这件", "mq"}, WorTag{"事情", "n"}},
		[]WordTag{WorTag{"电脑", "n"}, WorTag{"修好", "v"}, WorTag{"了", "ul"}},
		[]WordTag{WorTag{"做好", "v"}, WorTag{"了", "ul"}, WorTag{"这件", "mq"}, WorTag{"事情", "n"}, WorTag{"就", "d"}, WorTag{"一了百了", "l"}, WorTag{"了", "ul"}},
		[]WordTag{WorTag{"人们", "n"}, WorTag{"审美", "vn"}, WorTag{"的", "uj"}, WorTag{"观点", "n"}, WorTag{"是", "v"}, WorTag{"不同", "a"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"我们", "r"}, WorTag{"买", "v"}, WorTag{"了", "ul"}, WorTag{"一个", "m"}, WorTag{"美的", "nr"}, WorTag{"空调", "n"}},
		[]WordTag{WorTag{"线程", "n"}, WorTag{"初始化", "l"}, WorTag{"时", "n"}, WorTag{"我们", "r"}, WorTag{"要", "v"}, WorTag{"注意", "v"}},
		[]WordTag{WorTag{"一个", "m"}, WorTag{"分子", "n"}, WorTag{"是", "v"}, WorTag{"由", "p"}, WorTag{"好多", "m"}, WorTag{"原子", "n"}, WorTag{"组织", "v"}, WorTag{"成", "n"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"祝", "v"}, WorTag{"你", "r"}, WorTag{"马到功成", "i"}},
		[]WordTag{WorTag{"他", "r"}, WorTag{"掉", "zg"}, WorTag{"进", "v"}, WorTag{"了", "ul"}, WorTag{"无底洞", "ns"}, WorTag{"里", "f"}},
		[]WordTag{WorTag{"中国", "ns"}, WorTag{"的", "uj"}, WorTag{"首都", "d"}, WorTag{"是", "v"}, WorTag{"北京", "ns"}},
		[]WordTag{WorTag{"孙", "zg"}, WorTag{"君", "nz"}, WorTag{"意", "n"}},
		[]WordTag{WorTag{"外交部", "nt"}, WorTag{"发言人", "l"}, WorTag{"马朝旭", "nr"}},
		[]WordTag{WorTag{"领导人", "n"}, WorTag{"会议", "n"}, WorTag{"和", "c"}, WorTag{"第四届", "m"}, WorTag{"东亚", "ns"}, WorTag{"峰会", "n"}},
		[]WordTag{WorTag{"在", "p"}, WorTag{"过去", "t"}, WorTag{"的", "uj"}, WorTag{"这", "r"}, WorTag{"五年", "t"}},
		[]WordTag{WorTag{"还", "d"}, WorTag{"需要", "v"}, WorTag{"很", "zg"}, WorTag{"长", "a"}, WorTag{"的", "uj"}, WorTag{"路", "n"}, WorTag{"要", "v"}, WorTag{"走", "v"}},
		[]WordTag{WorTag{"60", "eng"}, WorTag{"周年", "t"}, WorTag{"首都", "d"}, WorTag{"阅兵", "v"}},
		[]WordTag{WorTag{"你好", "l"}, WorTag{"人们", "n"}, WorTag{"审美", "vn"}, WorTag{"的", "uj"}, WorTag{"观点", "n"}, WorTag{"是", "v"}, WorTag{"不同", "a"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"买", "v"}, WorTag{"水果", "n"}, WorTag{"然后", "c"}, WorTag{"来", "v"}, WorTag{"世博园", "nr"}},
		[]WordTag{WorTag{"买", "v"}, WorTag{"水果", "n"}, WorTag{"然后", "c"}, WorTag{"去", "v"}, WorTag{"世博园", "nr"}},
		[]WordTag{WorTag{"但是", "c"}, WorTag{"后来", "t"}, WorTag{"我", "r"}, WorTag{"才", "d"}, WorTag{"知道", "v"}, WorTag{"你", "r"}, WorTag{"是", "v"}, WorTag{"对", "p"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"存在", "v"}, WorTag{"即", "v"}, WorTag{"合理", "vn"}},
		[]WordTag{WorTag{"的", "uj"}, WorTag{"的", "uj"}, WorTag{"的", "uj"}, WorTag{"的", "uj"}, WorTag{"的", "uj"}, WorTag{"在", "p"}, WorTag{"的", "uj"}, WorTag{"的", "uj"}, WorTag{"的", "uj"}, WorTag{"的", "uj"}, WorTag{"就", "d"}, WorTag{"以", "p"}, WorTag{"和", "c"}, WorTag{"和", "c"}, WorTag{"和", "c"}},
		[]WordTag{WorTag{"I", "eng"}, WorTag{" ", "x"}, WorTag{"love", "eng"}, WorTag{"你", "r"}, WorTag{"，", "x"}, WorTag{"不以为耻", "i"}, WorTag{"，", "x"}, WorTag{"反", "zg"}, WorTag{"以为", "c"}, WorTag{"rong", "eng"}},
		[]WordTag{WorTag{"因", "p"}},
		[]WordTag{},
		[]WordTag{WorTag{"hello", "eng"}, WorTag{"你好", "l"}, WorTag{"人们", "n"}, WorTag{"审美", "vn"}, WorTag{"的", "uj"}, WorTag{"观点", "n"}, WorTag{"是", "v"}, WorTag{"不同", "a"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"很", "zg"}, WorTag{"好", "a"}, WorTag{"但", "c"}, WorTag{"主要", "b"}, WorTag{"是", "v"}, WorTag{"基于", "p"}, WorTag{"网页", "n"}, WorTag{"形式", "n"}},
		[]WordTag{WorTag{"hello", "eng"}, WorTag{"你好", "l"}, WorTag{"人们", "n"}, WorTag{"审美", "vn"}, WorTag{"的", "uj"}, WorTag{"观点", "n"}, WorTag{"是", "v"}, WorTag{"不同", "a"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"为什么", "r"}, WorTag{"我", "r"}, WorTag{"不能", "v"}, WorTag{"拥有", "v"}, WorTag{"想要", "v"}, WorTag{"的", "uj"}, WorTag{"生活", "vn"}},
		[]WordTag{WorTag{"后来", "t"}, WorTag{"我", "r"}, WorTag{"才", "d"}},
		[]WordTag{WorTag{"此次", "r"}, WorTag{"来", "v"}, WorTag{"中国", "ns"}, WorTag{"是", "v"}, WorTag{"为了", "p"}},
		[]WordTag{WorTag{"使用", "v"}, WorTag{"了", "ul"}, WorTag{"它", "r"}, WorTag{"就", "d"}, WorTag{"可以", "c"}, WorTag{"解决", "v"}, WorTag{"一些", "m"}, WorTag{"问题", "n"}},
		[]WordTag{WorTag{",", "x"}, WorTag{"使用", "v"}, WorTag{"了", "ul"}, WorTag{"它", "r"}, WorTag{"就", "d"}, WorTag{"可以", "c"}, WorTag{"解决", "v"}, WorTag{"一些", "m"}, WorTag{"问题", "n"}},
		[]WordTag{WorTag{"其实", "d"}, WorTag{"使用", "v"}, WorTag{"了", "ul"}, WorTag{"它", "r"}, WorTag{"就", "d"}, WorTag{"可以", "c"}, WorTag{"解决", "v"}, WorTag{"一些", "m"}, WorTag{"问题", "n"}},
		[]WordTag{WorTag{"好人", "n"}, WorTag{"使用", "v"}, WorTag{"了", "ul"}, WorTag{"它", "r"}, WorTag{"就", "d"}, WorTag{"可以", "c"}, WorTag{"解决", "v"}, WorTag{"一些", "m"}, WorTag{"问题", "n"}},
		[]WordTag{WorTag{"是因为", "c"}, WorTag{"和", "c"}, WorTag{"国家", "n"}},
		[]WordTag{WorTag{"老年", "t"}, WorTag{"搜索", "v"}, WorTag{"还", "d"}, WorTag{"支持", "v"}},
		[]WordTag{WorTag{"干脆", "d"}, WorTag{"就", "d"}, WorTag{"把", "p"}, WorTag{"那", "r"}, WorTag{"部", "n"}, WorTag{"蒙", "v"}, WorTag{"人", "n"}, WorTag{"的", "uj"}, WorTag{"闲", "n"}, WorTag{"法", "j"}, WorTag{"给", "p"}, WorTag{"废", "v"}, WorTag{"了", "ul"}, WorTag{"拉倒", "v"}, WorTag{"！", "x"}, WorTag{"RT", "eng"}, WorTag{" ", "x"}, WorTag{"@", "x"}, WorTag{"laoshipukong", "eng"}, WorTag{" ", "x"}, WorTag{":", "x"}, WorTag{" ", "x"}, WorTag{"27", "eng"}, WorTag{"日", "m"}, WorTag{"，", "x"}, WorTag{"全国人大常委会", "nt"}, WorTag{"第三次", "m"}, WorTag{"审议", "v"}, WorTag{"侵权", "v"}, WorTag{"责任法", "n"}, WorTag{"草案", "n"}, WorTag{"，", "x"}, WorTag{"删除", "v"}, WorTag{"了", "ul"}, WorTag{"有关", "vn"}, WorTag{"医疗", "n"}, WorTag{"损害", "v"}, WorTag{"责任", "n"}, WorTag{"“", "x"}, WorTag{"举证", "v"}, WorTag{"倒置", "v"}, WorTag{"”", "x"}, WorTag{"的", "uj"}, WorTag{"规定", "n"}, WorTag{"。", "x"}, WorTag{"在", "p"}, WorTag{"医患", "n"}, WorTag{"纠纷", "n"}, WorTag{"中", "f"}, WorTag{"本", "r"}, WorTag{"已", "d"}, WorTag{"处于", "v"}, WorTag{"弱势", "n"}, WorTag{"地位", "n"}, WorTag{"的", "uj"}, WorTag{"消费者", "n"}, WorTag{"由此", "c"}, WorTag{"将", "d"}, WorTag{"陷入", "v"}, WorTag{"万劫不复", "i"}, WorTag{"的", "uj"}, WorTag{"境地", "s"}, WorTag{"。", "x"}, WorTag{" ", "x"}},
		[]WordTag{WorTag{"大", "a"}},
		[]WordTag{},
		[]WordTag{WorTag{"他", "r"}, WorTag{"说", "v"}, WorTag{"的", "uj"}, WorTag{"确实", "ad"}, WorTag{"在", "p"}, WorTag{"理", "n"}},
		[]WordTag{WorTag{"长春", "ns"}, WorTag{"市长", "n"}, WorTag{"春节", "t"}, WorTag{"讲话", "n"}},
		[]WordTag{WorTag{"结婚", "v"}, WorTag{"的", "uj"}, WorTag{"和", "c"}, WorTag{"尚未", "d"}, WorTag{"结婚", "v"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"结合", "v"}, WorTag{"成", "n"}, WorTag{"分子", "n"}, WorTag{"时", "n"}},
		[]WordTag{WorTag{"旅游", "vn"}, WorTag{"和", "c"}, WorTag{"服务", "vn"}, WorTag{"是", "v"}, WorTag{"最好", "a"}, WorTag{"的", "uj"}},
		[]WordTag{WorTag{"这件", "mq"}, WorTag{"事情", "n"}, WorTag{"的确", "d"}, WorTag{"是", "v"}, WorTag{"我", "r"}, WorTag{"的", "uj"}, WorTag{"错", "v"}},
		[]WordTag{WorTag{"供", "v"}, WorTag{"大家", "n"}, WorTag{"参考", "v"}, WorTag{"指正", "v"}},
		[]WordTag{WorTag{"哈尔滨", "ns"}, WorTag{"政府", "n"}, WorTag{"公布", "v"}, WorTag{"塌", "v"}, WorTag{"桥", "n"}, WorTag{"原因", "n"}},
		[]WordTag{WorTag{"我", "r"}, WorTag{"在", "p"}, WorTag{"机场", "n"}, WorTag{"入口处", "i"}},
		[]WordTag{WorTag{"邢", "nr"}, WorTag{"永", "ns"}, WorTag{"臣", "n"}, WorTag{"摄影", "n"}, WorTag{"报道", "v"}},
		[]WordTag{WorTag{"BP", "eng"}, WorTag{"神经网络", "n"}, WorTag{"如何", "r"}, WorTag{"训练", "vn"}, WorTag{"才能", "v"}, WorTag{"在", "p"}, WorTag{"分类", "n"}, WorTag{"时", "n"}, WorTag{"增加", "v"}, WorTag{"区分度", "n"}, WorTag{"？", "x"}},
		[]WordTag{WorTag{"南京市", "ns"}, WorTag{"长江大桥", "ns"}},
		[]WordTag{WorTag{"应", "v"}, WorTag{"一些", "m"}, WorTag{"使用者", "n"}, WorTag{"的", "uj"}, WorTag{"建议", "n"}, WorTag{"，", "x"}, WorTag{"也", "d"}, WorTag{"为了", "p"}, WorTag{"便于", "v"}, WorTag{"利用", "n"}, WorTag{"NiuTrans", "eng"}, WorTag{"用于", "v"}, WorTag{"SMT", "eng"}, WorTag{"研究", "vn"}},
		[]WordTag{WorTag{"长春市", "ns"}, WorTag{"长春", "ns"}, WorTag{"药店", "n"}},
		[]WordTag{WorTag{"邓颖超", "nr"}, WorTag{"生前", "t"}, WorTag{"最", "d"}, WorTag{"喜欢", "v"}, WorTag{"的", "uj"}, WorTag{"衣服", "n"}},
		[]WordTag{WorTag{"胡锦涛", "nr"}, WorTag{"是", "v"}, WorTag{"热爱", "a"}, WorTag{"世界", "n"}, WorTag{"和平", "nz"}, WorTag{"的", "uj"}, WorTag{"政治局", "n"}, WorTag{"常委", "j"}},
		[]WordTag{WorTag{"程序员", "n"}, WorTag{"祝", "v"}, WorTag{"海林", "nz"}, WorTag{"和", "c"}, WorTag{"朱", "nr"}, WorTag{"会", "v"}, WorTag{"震", "v"}, WorTag{"是", "v"}, WorTag{"在", "p"}, WorTag{"孙", "zg"}, WorTag{"健", "a"}, WorTag{"的", "uj"}, WorTag{"左面", "f"}, WorTag{"和", "c"}, WorTag{"右面", "f"}, WorTag{",", "x"}, WorTag{" ", "x"}, WorTag{"范", "nr"}, WorTag{"凯", "nr"}, WorTag{"在", "p"}, WorTag{"最", "d"}, WorTag{"右面", "f"}, WorTag{".", "x"}, WorTag{"再", "d"}, WorTag{"往", "zg"}, WorTag{"左", "m"}, WorTag{"是", "v"}, WorTag{"李", "nr"}, WorTag{"松", "v"}, WorTag{"洪", "nr"}},
		[]WordTag{WorTag{"一次性", "d"}, WorTag{"交", "v"}, WorTag{"多少", "m"}, WorTag{"钱", "n"}},
		[]WordTag{WorTag{"两块", "m"}, WorTag{"五", "m"}, WorTag{"一套", "m"}, WorTag{"，", "x"}, WorTag{"三块", "m"}, WorTag{"八", "m"}, WorTag{"一斤", "m"}, WorTag{"，", "x"}, WorTag{"四块", "m"}, WorTag{"七", "m"}, WorTag{"一本", "m"}, WorTag{"，", "x"}, WorTag{"五块", "m"}, WorTag{"六", "m"}, WorTag{"一条", "m"}},
		[]WordTag{WorTag{"小", "a"}, WorTag{"和尚", "nr"}, WorTag{"留", "v"}, WorTag{"了", "ul"}, WorTag{"一个", "m"}, WorTag{"像", "v"}, WorTag{"大", "a"}, WorTag{"和尚", "nr"}, WorTag{"一样", "r"}, WorTag{"的", "uj"}, WorTag{"和尚头", "nr"}},
		[]WordTag{WorTag{"我", "r"}, WorTag{"是", "v"}, WorTag{"中华人民共和国", "ns"}, WorTag{"公民", "n"}, WorTag{";", "x"}, WorTag{"我", "r"}, WorTag{"爸爸", "n"}, WorTag{"是", "v"}, WorTag{"共和党", "nt"}, WorTag{"党员", "n"}, WorTag{";", "x"}, WorTag{" ", "x"}, WorTag{"地铁", "n"}, WorTag{"和平门", "ns"}, WorTag{"站", "v"}},
		[]WordTag{WorTag{"张晓梅", "nr"}, WorTag{"去", "v"}, WorTag{"人民", "n"}, WorTag{"医院", "n"}, WorTag{"做", "v"}, WorTag{"了", "ul"}, WorTag{"个", "q"}, WorTag{"B超", "n"}, WorTag{"然后", "c"}, WorTag{"去", "v"}, WorTag{"买", "v"}, WorTag{"了", "ul"}, WorTag{"件", "zg"}, WorTag{"T恤", "n"}},
		[]WordTag{WorTag{"AT&T", "nz"}, WorTag{"是", "v"}, WorTag{"一件", "m"}, WorTag{"不错", "a"}, WorTag{"的", "uj"}, WorTag{"公司", "n"}, WorTag{"，", "x"}, WorTag{"给", "p"}, WorTag{"你", "r"}, WorTag{"发", "v"}, WorTag{"offer", "eng"}, WorTag{"了", "ul"}, WorTag{"吗", "y"}, WorTag{"？", "x"}},
		[]WordTag{WorTag{"C++", "nz"}, WorTag{"和", "c"}, WorTag{"c#", "nz"}, WorTag{"是", "v"}, WorTag{"什么", "r"}, WorTag{"关系", "n"}, WorTag{"？", "x"}, WorTag{"11", "eng"}, WorTag{"+", "x"}, WorTag{"122", "eng"}, WorTag{"=", "x"}, WorTag{"133", "eng"}, WorTag{"，", "x"}, WorTag{"是", "v"}, WorTag{"吗", "y"}, WorTag{"？", "x"}, WorTag{"PI", "eng"}, WorTag{"=", "x"}, WorTag{"3", "eng"}, WorTag{".", "x"}, WorTag{"14159", "eng"}},
		[]WordTag{WorTag{"你", "r"}, WorTag{"认识", "v"}, WorTag{"那个", "r"}, WorTag{"和", "c"}, WorTag{"主席", "n"}, WorTag{"握手", "v"}, WorTag{"的", "uj"}, WorTag{"的哥", "n"}, WorTag{"吗", "y"}, WorTag{"？", "x"}, WorTag{"他", "r"}, WorTag{"开", "v"}, WorTag{"一辆", "m"}, WorTag{"黑色", "n"}, WorTag{"的士", "n"}, WorTag{"。", "x"}},
		[]WordTag{WorTag{"枪杆子", "n"}, WorTag{"中", "f"}, WorTag{"出", "v"}, WorTag{"政权", "n"}},
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

func TestBug132(t *testing.T) {
	/*
		https://github.com/fxsjy/jieba/issues/132
	*/
	jiebago.SetDictionary("../dict.txt")
	sentence := "又跛又啞"
	cutResult := []WordTag{
		WordTag{"又", "d"},
		WordTag{"跛", "a"},
		WordTag{"又", "d"},
		WordTag{"啞", "v"},
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

func TestBug137(t *testing.T) {
	/*
		https://github.com/fxsjy/jieba/issues/137
	*/
	jiebago.SetDictionary("../dict.txt")
	sentence := "前港督衛奕信在八八年十月宣布成立中央政策研究組"
	cutResult := []WordTag{
		[]WordTag{WorTag{"前", "f"},
			WorTag{"港督", "n"},
			WorTag{"衛奕", "z"},
			WorTag{"信", "n"},
			WorTag{"在", "p"},
			WorTag{"八八年", "m"},
			WorTag{"十月", "t"},
			WorTag{"宣布", "v"},
			WorTag{"成立", "v"},
			WorTag{"中央", "n"},
			WorTag{"政策", "n"},
			WorTag{"研究", "vn"},
			WorTag{"組", "x"}},
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

func TestUserDict(t *testing.T) {
	jiebago.SetDictionary("../dict.txt")
	jiebago.LoadUserDict("../userdict.txt")
	sentence := "李小福是创新办主任也是云计算方面的专家; 什么是八一双鹿例如我输入一个带“韩玉赏鉴”的标题，在自定义词库中也增加了此词为N类型"

	cutResult := []WordTag{
		WorTag{"李小福", "nr"},
		WorTag{"是", "v"},
		WorTag{"创新办", "i"},
		WorTag{"主任", "b"},
		WorTag{"也", "d"},
		WorTag{"是", "v"},
		WorTag{"云计算", "x"},
		WorTag{"方面", "n"},
		WorTag{"的", "uj"},
		WorTag{"专家", "n"},
		WorTag{";", "x"},
		WorTag{" ", "x"},
		WorTag{"什么", "r"},
		WorTag{"是", "v"},
		WorTag{"八一双鹿", "nz"},
		WorTag{"例如", "v"},
		WorTag{"我", "r"},
		WorTag{"输入", "v"},
		WorTag{"一个", "m"},
		WorTag{"带", "v"},
		WorTag{"“", "x"},
		WorTag{"韩玉赏鉴", "nz"},
		WorTag{"”", "x"},
		WorTag{"的", "uj"},
		WorTag{"标题", "n"},
		WorTag{"，", "x"},
		WorTag{"在", "p"},
		WorTag{"自定义词", "n"},
		WorTag{"库中", "nrt"},
		WorTag{"也", "d"},
		WorTag{"增加", "v"},
		WorTag{"了", "ul"},
		WorTag{"此", "r"},
		WorTag{"词", "n"},
		WorTag{"为", "p"},
		WorTag{"N", "eng"},
		WorTag{"类型", "n"}}

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
