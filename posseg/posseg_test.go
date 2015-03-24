package posseg

import (
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

	defaultCutResult = [][]WordTag{[]WordTag{WordTag{"这", "r"}, WordTag{"是", "v"}, WordTag{"一个", "m"}, WordTag{"伸手不见五指", "i"}, WordTag{"的", "uj"}, WordTag{"黑夜", "n"}, WordTag{"。", "x"}, WordTag{"我", "r"}, WordTag{"叫", "v"}, WordTag{"孙悟空", "nr"}, WordTag{"，", "x"}, WordTag{"我", "r"}, WordTag{"爱", "v"}, WordTag{"北京", "ns"}, WordTag{"，", "x"}, WordTag{"我", "r"}, WordTag{"爱", "v"}, WordTag{"Python", "eng"}, WordTag{"和", "c"}, WordTag{"C++", "nz"}, WordTag{"。", "x"}},
		[]WordTag{WordTag{"我", "r"}, WordTag{"不", "d"}, WordTag{"喜欢", "v"}, WordTag{"日本", "ns"}, WordTag{"和服", "nz"}, WordTag{"。", "x"}},
		[]WordTag{WordTag{"雷猴", "n"}, WordTag{"回归", "v"}, WordTag{"人间", "n"}, WordTag{"。", "x"}},
		[]WordTag{WordTag{"工信处", "n"}, WordTag{"女干事", "n"}, WordTag{"每月", "r"}, WordTag{"经过", "p"}, WordTag{"下属", "v"}, WordTag{"科室", "n"}, WordTag{"都", "d"}, WordTag{"要", "v"}, WordTag{"亲口", "n"}, WordTag{"交代", "n"}, WordTag{"24", "m"}, WordTag{"口", "n"}, WordTag{"交换机", "n"}, WordTag{"等", "u"}, WordTag{"技术性", "n"}, WordTag{"器件", "n"}, WordTag{"的", "uj"}, WordTag{"安装", "v"}, WordTag{"工作", "vn"}},
		[]WordTag{WordTag{"我", "r"}, WordTag{"需要", "v"}, WordTag{"廉租房", "n"}},
		[]WordTag{WordTag{"永和", "nz"}, WordTag{"服装", "vn"}, WordTag{"饰品", "n"}, WordTag{"有限公司", "n"}},
		[]WordTag{WordTag{"我", "r"}, WordTag{"爱", "v"}, WordTag{"北京", "ns"}, WordTag{"天安门", "ns"}},
		[]WordTag{WordTag{"abc", "eng"}},
		[]WordTag{WordTag{"隐", "n"}, WordTag{"马尔可夫", "nr"}},
		[]WordTag{WordTag{"雷猴", "n"}, WordTag{"是", "v"}, WordTag{"个", "q"}, WordTag{"好", "a"}, WordTag{"网站", "n"}},
		[]WordTag{WordTag{"“", "x"}, WordTag{"Microsoft", "eng"}, WordTag{"”", "x"}, WordTag{"一", "m"}, WordTag{"词", "n"}, WordTag{"由", "p"}, WordTag{"“", "x"}, WordTag{"MICROcomputer", "eng"}, WordTag{"（", "x"}, WordTag{"微型", "b"}, WordTag{"计算机", "n"}, WordTag{"）", "x"}, WordTag{"”", "x"}, WordTag{"和", "c"}, WordTag{"“", "x"}, WordTag{"SOFTware", "eng"}, WordTag{"（", "x"}, WordTag{"软件", "n"}, WordTag{"）", "x"}, WordTag{"”", "x"}, WordTag{"两", "m"}, WordTag{"部分", "n"}, WordTag{"组成", "v"}},
		[]WordTag{WordTag{"草泥马", "n"}, WordTag{"和", "c"}, WordTag{"欺实", "v"}, WordTag{"马", "n"}, WordTag{"是", "v"}, WordTag{"今年", "t"}, WordTag{"的", "uj"}, WordTag{"流行", "v"}, WordTag{"词汇", "n"}},
		[]WordTag{WordTag{"伊藤", "nr"}, WordTag{"洋华堂", "n"}, WordTag{"总府", "n"}, WordTag{"店", "n"}},
		[]WordTag{WordTag{"中国科学院计算技术研究所", "nt"}},
		[]WordTag{WordTag{"罗密欧", "nr"}, WordTag{"与", "p"}, WordTag{"朱丽叶", "nr"}},
		[]WordTag{WordTag{"我", "r"}, WordTag{"购买", "v"}, WordTag{"了", "ul"}, WordTag{"道具", "n"}, WordTag{"和", "c"}, WordTag{"服装", "vn"}},
		[]WordTag{WordTag{"PS", "eng"}, WordTag{":", "x"}, WordTag{" ", "x"}, WordTag{"我", "r"}, WordTag{"觉得", "v"}, WordTag{"开源", "n"}, WordTag{"有", "v"}, WordTag{"一个", "m"}, WordTag{"好处", "d"}, WordTag{"，", "x"}, WordTag{"就是", "d"}, WordTag{"能够", "v"}, WordTag{"敦促", "v"}, WordTag{"自己", "r"}, WordTag{"不断改进", "l"}, WordTag{"，", "x"}, WordTag{"避免", "v"}, WordTag{"敞", "v"}, WordTag{"帚", "ng"}, WordTag{"自珍", "b"}},
		[]WordTag{WordTag{"湖北省", "ns"}, WordTag{"石首市", "ns"}},
		[]WordTag{WordTag{"湖北省", "ns"}, WordTag{"十堰市", "ns"}},
		[]WordTag{WordTag{"总经理", "n"}, WordTag{"完成", "v"}, WordTag{"了", "ul"}, WordTag{"这件", "mq"}, WordTag{"事情", "n"}},
		[]WordTag{WordTag{"电脑", "n"}, WordTag{"修好", "v"}, WordTag{"了", "ul"}},
		[]WordTag{WordTag{"做好", "v"}, WordTag{"了", "ul"}, WordTag{"这件", "mq"}, WordTag{"事情", "n"}, WordTag{"就", "d"}, WordTag{"一了百了", "l"}, WordTag{"了", "ul"}},
		[]WordTag{WordTag{"人们", "n"}, WordTag{"审美", "vn"}, WordTag{"的", "uj"}, WordTag{"观点", "n"}, WordTag{"是", "v"}, WordTag{"不同", "a"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"我们", "r"}, WordTag{"买", "v"}, WordTag{"了", "ul"}, WordTag{"一个", "m"}, WordTag{"美的", "nr"}, WordTag{"空调", "n"}},
		[]WordTag{WordTag{"线程", "n"}, WordTag{"初始化", "l"}, WordTag{"时", "n"}, WordTag{"我们", "r"}, WordTag{"要", "v"}, WordTag{"注意", "v"}},
		[]WordTag{WordTag{"一个", "m"}, WordTag{"分子", "n"}, WordTag{"是", "v"}, WordTag{"由", "p"}, WordTag{"好多", "m"}, WordTag{"原子", "n"}, WordTag{"组织", "v"}, WordTag{"成", "v"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"祝", "v"}, WordTag{"你", "r"}, WordTag{"马到功成", "i"}},
		[]WordTag{WordTag{"他", "r"}, WordTag{"掉", "v"}, WordTag{"进", "v"}, WordTag{"了", "ul"}, WordTag{"无底洞", "ns"}, WordTag{"里", "f"}},
		[]WordTag{WordTag{"中国", "ns"}, WordTag{"的", "uj"}, WordTag{"首都", "d"}, WordTag{"是", "v"}, WordTag{"北京", "ns"}},
		[]WordTag{WordTag{"孙君意", "nr"}},
		[]WordTag{WordTag{"外交部", "nt"}, WordTag{"发言人", "l"}, WordTag{"马朝旭", "nr"}},
		[]WordTag{WordTag{"领导人", "n"}, WordTag{"会议", "n"}, WordTag{"和", "c"}, WordTag{"第四届", "m"}, WordTag{"东亚", "ns"}, WordTag{"峰会", "n"}},
		[]WordTag{WordTag{"在", "p"}, WordTag{"过去", "t"}, WordTag{"的", "uj"}, WordTag{"这", "r"}, WordTag{"五年", "t"}},
		[]WordTag{WordTag{"还", "d"}, WordTag{"需要", "v"}, WordTag{"很", "d"}, WordTag{"长", "a"}, WordTag{"的", "uj"}, WordTag{"路", "n"}, WordTag{"要", "v"}, WordTag{"走", "v"}},
		[]WordTag{WordTag{"60", "m"}, WordTag{"周年", "t"}, WordTag{"首都", "d"}, WordTag{"阅兵", "v"}},
		[]WordTag{WordTag{"你好", "l"}, WordTag{"人们", "n"}, WordTag{"审美", "vn"}, WordTag{"的", "uj"}, WordTag{"观点", "n"}, WordTag{"是", "v"}, WordTag{"不同", "a"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"买", "v"}, WordTag{"水果", "n"}, WordTag{"然后", "c"}, WordTag{"来", "v"}, WordTag{"世博园", "nr"}},
		[]WordTag{WordTag{"买", "v"}, WordTag{"水果", "n"}, WordTag{"然后", "c"}, WordTag{"去", "v"}, WordTag{"世博园", "nr"}},
		[]WordTag{WordTag{"但是", "c"}, WordTag{"后来", "t"}, WordTag{"我", "r"}, WordTag{"才", "d"}, WordTag{"知道", "v"}, WordTag{"你", "r"}, WordTag{"是", "v"}, WordTag{"对", "p"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"存在", "v"}, WordTag{"即", "v"}, WordTag{"合理", "vn"}},
		[]WordTag{WordTag{"的的", "u"}, WordTag{"的的", "u"}, WordTag{"的", "uj"}, WordTag{"在的", "u"}, WordTag{"的的", "u"}, WordTag{"的", "uj"}, WordTag{"就", "d"}, WordTag{"以", "p"}, WordTag{"和和", "nz"}, WordTag{"和", "c"}},
		[]WordTag{WordTag{"I", "x"}, WordTag{" ", "x"}, WordTag{"love", "eng"}, WordTag{"你", "r"}, WordTag{"，", "x"}, WordTag{"不以为耻", "i"}, WordTag{"，", "x"}, WordTag{"反", "zg"}, WordTag{"以为", "c"}, WordTag{"rong", "eng"}},
		[]WordTag{WordTag{"因", "p"}},
		[]WordTag{},
		[]WordTag{WordTag{"hello", "eng"}, WordTag{"你好", "l"}, WordTag{"人们", "n"}, WordTag{"审美", "vn"}, WordTag{"的", "uj"}, WordTag{"观点", "n"}, WordTag{"是", "v"}, WordTag{"不同", "a"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"很好", "a"}, WordTag{"但", "c"}, WordTag{"主要", "b"}, WordTag{"是", "v"}, WordTag{"基于", "p"}, WordTag{"网页", "n"}, WordTag{"形式", "n"}},
		[]WordTag{WordTag{"hello", "eng"}, WordTag{"你好", "l"}, WordTag{"人们", "n"}, WordTag{"审美", "vn"}, WordTag{"的", "uj"}, WordTag{"观点", "n"}, WordTag{"是", "v"}, WordTag{"不同", "a"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"为什么", "r"}, WordTag{"我", "r"}, WordTag{"不能", "v"}, WordTag{"拥有", "v"}, WordTag{"想要", "v"}, WordTag{"的", "uj"}, WordTag{"生活", "vn"}},
		[]WordTag{WordTag{"后来", "t"}, WordTag{"我", "r"}, WordTag{"才", "d"}},
		[]WordTag{WordTag{"此次", "r"}, WordTag{"来", "v"}, WordTag{"中国", "ns"}, WordTag{"是", "v"}, WordTag{"为了", "p"}},
		[]WordTag{WordTag{"使用", "v"}, WordTag{"了", "ul"}, WordTag{"它", "r"}, WordTag{"就", "d"}, WordTag{"可以", "c"}, WordTag{"解决", "v"}, WordTag{"一些", "m"}, WordTag{"问题", "n"}},
		[]WordTag{WordTag{",", "x"}, WordTag{"使用", "v"}, WordTag{"了", "ul"}, WordTag{"它", "r"}, WordTag{"就", "d"}, WordTag{"可以", "c"}, WordTag{"解决", "v"}, WordTag{"一些", "m"}, WordTag{"问题", "n"}},
		[]WordTag{WordTag{"其实", "d"}, WordTag{"使用", "v"}, WordTag{"了", "ul"}, WordTag{"它", "r"}, WordTag{"就", "d"}, WordTag{"可以", "c"}, WordTag{"解决", "v"}, WordTag{"一些", "m"}, WordTag{"问题", "n"}},
		[]WordTag{WordTag{"好人", "n"}, WordTag{"使用", "v"}, WordTag{"了", "ul"}, WordTag{"它", "r"}, WordTag{"就", "d"}, WordTag{"可以", "c"}, WordTag{"解决", "v"}, WordTag{"一些", "m"}, WordTag{"问题", "n"}},
		[]WordTag{WordTag{"是因为", "c"}, WordTag{"和", "c"}, WordTag{"国家", "n"}},
		[]WordTag{WordTag{"老年", "t"}, WordTag{"搜索", "v"}, WordTag{"还", "d"}, WordTag{"支持", "v"}},
		[]WordTag{WordTag{"干脆", "d"}, WordTag{"就", "d"}, WordTag{"把", "p"}, WordTag{"那部", "r"}, WordTag{"蒙人", "n"}, WordTag{"的", "uj"}, WordTag{"闲法", "n"}, WordTag{"给", "p"}, WordTag{"废", "v"}, WordTag{"了", "ul"}, WordTag{"拉倒", "v"}, WordTag{"！", "x"}, WordTag{"RT", "eng"}, WordTag{" ", "x"}, WordTag{"@", "x"}, WordTag{"laoshipukong", "eng"}, WordTag{" ", "x"}, WordTag{":", "x"}, WordTag{" ", "x"}, WordTag{"27", "m"}, WordTag{"日", "m"}, WordTag{"，", "x"}, WordTag{"全国人大常委会", "nt"}, WordTag{"第三次", "m"}, WordTag{"审议", "v"}, WordTag{"侵权", "v"}, WordTag{"责任法", "n"}, WordTag{"草案", "n"}, WordTag{"，", "x"}, WordTag{"删除", "v"}, WordTag{"了", "ul"}, WordTag{"有关", "vn"}, WordTag{"医疗", "n"}, WordTag{"损害", "v"}, WordTag{"责任", "n"}, WordTag{"“", "x"}, WordTag{"举证", "v"}, WordTag{"倒置", "v"}, WordTag{"”", "x"}, WordTag{"的", "uj"}, WordTag{"规定", "n"}, WordTag{"。", "x"}, WordTag{"在", "p"}, WordTag{"医患", "n"}, WordTag{"纠纷", "n"}, WordTag{"中本", "ns"}, WordTag{"已", "d"}, WordTag{"处于", "v"}, WordTag{"弱势", "n"}, WordTag{"地位", "n"}, WordTag{"的", "uj"}, WordTag{"消费者", "n"}, WordTag{"由此", "c"}, WordTag{"将", "d"}, WordTag{"陷入", "v"}, WordTag{"万劫不复", "i"}, WordTag{"的", "uj"}, WordTag{"境地", "s"}, WordTag{"。", "x"}, WordTag{" ", "x"}},
		[]WordTag{WordTag{"大", "a"}},
		[]WordTag{},
		[]WordTag{WordTag{"他", "r"}, WordTag{"说", "v"}, WordTag{"的", "uj"}, WordTag{"确实", "ad"}, WordTag{"在", "p"}, WordTag{"理", "n"}},
		[]WordTag{WordTag{"长春", "ns"}, WordTag{"市长", "n"}, WordTag{"春节", "t"}, WordTag{"讲话", "n"}},
		[]WordTag{WordTag{"结婚", "v"}, WordTag{"的", "uj"}, WordTag{"和", "c"}, WordTag{"尚未", "d"}, WordTag{"结婚", "v"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"结合", "v"}, WordTag{"成", "n"}, WordTag{"分子", "n"}, WordTag{"时", "n"}},
		[]WordTag{WordTag{"旅游", "vn"}, WordTag{"和", "c"}, WordTag{"服务", "vn"}, WordTag{"是", "v"}, WordTag{"最好", "a"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"这件", "mq"}, WordTag{"事情", "n"}, WordTag{"的确", "d"}, WordTag{"是", "v"}, WordTag{"我", "r"}, WordTag{"的", "uj"}, WordTag{"错", "n"}},
		[]WordTag{WordTag{"供", "v"}, WordTag{"大家", "n"}, WordTag{"参考", "v"}, WordTag{"指正", "v"}},
		[]WordTag{WordTag{"哈尔滨", "ns"}, WordTag{"政府", "n"}, WordTag{"公布", "v"}, WordTag{"塌", "v"}, WordTag{"桥", "n"}, WordTag{"原因", "n"}},
		[]WordTag{WordTag{"我", "r"}, WordTag{"在", "p"}, WordTag{"机场", "n"}, WordTag{"入口处", "i"}},
		[]WordTag{WordTag{"邢永臣", "nr"}, WordTag{"摄影", "n"}, WordTag{"报道", "v"}},
		[]WordTag{WordTag{"BP", "eng"}, WordTag{"神经网络", "n"}, WordTag{"如何", "r"}, WordTag{"训练", "vn"}, WordTag{"才能", "v"}, WordTag{"在", "p"}, WordTag{"分类", "n"}, WordTag{"时", "n"}, WordTag{"增加", "v"}, WordTag{"区分度", "n"}, WordTag{"？", "x"}},
		[]WordTag{WordTag{"南京市", "ns"}, WordTag{"长江大桥", "ns"}},
		[]WordTag{WordTag{"应", "v"}, WordTag{"一些", "m"}, WordTag{"使用者", "n"}, WordTag{"的", "uj"}, WordTag{"建议", "n"}, WordTag{"，", "x"}, WordTag{"也", "d"}, WordTag{"为了", "p"}, WordTag{"便于", "v"}, WordTag{"利用", "n"}, WordTag{"NiuTrans", "eng"}, WordTag{"用于", "v"}, WordTag{"SMT", "eng"}, WordTag{"研究", "vn"}},
		[]WordTag{WordTag{"长春市", "ns"}, WordTag{"长春", "ns"}, WordTag{"药店", "n"}},
		[]WordTag{WordTag{"邓颖超", "nr"}, WordTag{"生前", "t"}, WordTag{"最", "d"}, WordTag{"喜欢", "v"}, WordTag{"的", "uj"}, WordTag{"衣服", "n"}},
		[]WordTag{WordTag{"胡锦涛", "nr"}, WordTag{"是", "v"}, WordTag{"热爱", "a"}, WordTag{"世界", "n"}, WordTag{"和平", "nz"}, WordTag{"的", "uj"}, WordTag{"政治局", "n"}, WordTag{"常委", "j"}},
		[]WordTag{WordTag{"程序员", "n"}, WordTag{"祝", "v"}, WordTag{"海林", "nz"}, WordTag{"和", "c"}, WordTag{"朱会震", "nr"}, WordTag{"是", "v"}, WordTag{"在", "p"}, WordTag{"孙健", "nr"}, WordTag{"的", "uj"}, WordTag{"左面", "f"}, WordTag{"和", "c"}, WordTag{"右面", "f"}, WordTag{",", "x"}, WordTag{" ", "x"}, WordTag{"范凯", "nr"}, WordTag{"在", "p"}, WordTag{"最", "a"}, WordTag{"右面", "f"}, WordTag{".", "m"}, WordTag{"再往", "d"}, WordTag{"左", "f"}, WordTag{"是", "v"}, WordTag{"李松洪", "nr"}},
		[]WordTag{WordTag{"一次性", "d"}, WordTag{"交", "v"}, WordTag{"多少", "m"}, WordTag{"钱", "n"}},
		[]WordTag{WordTag{"两块", "m"}, WordTag{"五", "m"}, WordTag{"一套", "m"}, WordTag{"，", "x"}, WordTag{"三块", "m"}, WordTag{"八", "m"}, WordTag{"一斤", "m"}, WordTag{"，", "x"}, WordTag{"四块", "m"}, WordTag{"七", "m"}, WordTag{"一本", "m"}, WordTag{"，", "x"}, WordTag{"五块", "m"}, WordTag{"六", "m"}, WordTag{"一条", "m"}},
		[]WordTag{WordTag{"小", "a"}, WordTag{"和尚", "nr"}, WordTag{"留", "v"}, WordTag{"了", "ul"}, WordTag{"一个", "m"}, WordTag{"像", "v"}, WordTag{"大", "a"}, WordTag{"和尚", "nr"}, WordTag{"一样", "r"}, WordTag{"的", "uj"}, WordTag{"和尚头", "nr"}},
		[]WordTag{WordTag{"我", "r"}, WordTag{"是", "v"}, WordTag{"中华人民共和国", "ns"}, WordTag{"公民", "n"}, WordTag{";", "x"}, WordTag{"我", "r"}, WordTag{"爸爸", "n"}, WordTag{"是", "v"}, WordTag{"共和党", "nt"}, WordTag{"党员", "n"}, WordTag{";", "x"}, WordTag{" ", "x"}, WordTag{"地铁", "n"}, WordTag{"和平门", "ns"}, WordTag{"站", "v"}},
		[]WordTag{WordTag{"张晓梅", "nr"}, WordTag{"去", "v"}, WordTag{"人民", "n"}, WordTag{"医院", "n"}, WordTag{"做", "v"}, WordTag{"了", "ul"}, WordTag{"个", "q"}, WordTag{"B超", "n"}, WordTag{"然后", "c"}, WordTag{"去", "v"}, WordTag{"买", "v"}, WordTag{"了", "ul"}, WordTag{"件", "q"}, WordTag{"T恤", "n"}},
		[]WordTag{WordTag{"AT&T", "nz"}, WordTag{"是", "v"}, WordTag{"一件", "m"}, WordTag{"不错", "a"}, WordTag{"的", "uj"}, WordTag{"公司", "n"}, WordTag{"，", "x"}, WordTag{"给", "p"}, WordTag{"你", "r"}, WordTag{"发", "v"}, WordTag{"offer", "eng"}, WordTag{"了", "ul"}, WordTag{"吗", "y"}, WordTag{"？", "x"}},
		[]WordTag{WordTag{"C++", "nz"}, WordTag{"和", "c"}, WordTag{"c#", "nz"}, WordTag{"是", "v"}, WordTag{"什么", "r"}, WordTag{"关系", "n"}, WordTag{"？", "x"}, WordTag{"11", "m"}, WordTag{"+", "x"}, WordTag{"122", "m"}, WordTag{"=", "x"}, WordTag{"133", "m"}, WordTag{"，", "x"}, WordTag{"是", "v"}, WordTag{"吗", "y"}, WordTag{"？", "x"}, WordTag{"PI", "eng"}, WordTag{"=", "x"}, WordTag{"3.14159", "m"}},
		[]WordTag{WordTag{"你", "r"}, WordTag{"认识", "v"}, WordTag{"那个", "r"}, WordTag{"和", "c"}, WordTag{"主席", "n"}, WordTag{"握手", "v"}, WordTag{"的", "uj"}, WordTag{"的哥", "n"}, WordTag{"吗", "y"}, WordTag{"？", "x"}, WordTag{"他", "r"}, WordTag{"开", "v"}, WordTag{"一辆", "m"}, WordTag{"黑色", "n"}, WordTag{"的士", "n"}, WordTag{"。", "x"}},
		[]WordTag{WordTag{"枪杆子", "n"}, WordTag{"中", "f"}, WordTag{"出", "v"}, WordTag{"政权", "n"}},
	}
	noHMMCutResult = [][]WordTag{
		[]WordTag{WordTag{"这", "r"}, WordTag{"是", "v"}, WordTag{"一个", "m"}, WordTag{"伸手不见五指", "i"}, WordTag{"的", "uj"}, WordTag{"黑夜", "n"}, WordTag{"。", "x"}, WordTag{"我", "r"}, WordTag{"叫", "v"}, WordTag{"孙悟空", "nr"}, WordTag{"，", "x"}, WordTag{"我", "r"}, WordTag{"爱", "v"}, WordTag{"北京", "ns"}, WordTag{"，", "x"}, WordTag{"我", "r"}, WordTag{"爱", "v"}, WordTag{"Python", "eng"}, WordTag{"和", "c"}, WordTag{"C++", "nz"}, WordTag{"。", "x"}},
		[]WordTag{WordTag{"我", "r"}, WordTag{"不", "d"}, WordTag{"喜欢", "v"}, WordTag{"日本", "ns"}, WordTag{"和服", "nz"}, WordTag{"。", "x"}},
		[]WordTag{WordTag{"雷猴", "n"}, WordTag{"回归", "v"}, WordTag{"人间", "n"}, WordTag{"。", "x"}},
		[]WordTag{WordTag{"工信处", "n"}, WordTag{"女干事", "n"}, WordTag{"每月", "r"}, WordTag{"经过", "p"}, WordTag{"下属", "v"}, WordTag{"科室", "n"}, WordTag{"都", "d"}, WordTag{"要", "v"}, WordTag{"亲口", "n"}, WordTag{"交代", "n"}, WordTag{"24", "eng"}, WordTag{"口", "q"}, WordTag{"交换机", "n"}, WordTag{"等", "u"}, WordTag{"技术性", "n"}, WordTag{"器件", "n"}, WordTag{"的", "uj"}, WordTag{"安装", "v"}, WordTag{"工作", "vn"}},
		[]WordTag{WordTag{"我", "r"}, WordTag{"需要", "v"}, WordTag{"廉租房", "n"}},
		[]WordTag{WordTag{"永和", "nz"}, WordTag{"服装", "vn"}, WordTag{"饰品", "n"}, WordTag{"有限公司", "n"}},
		[]WordTag{WordTag{"我", "r"}, WordTag{"爱", "v"}, WordTag{"北京", "ns"}, WordTag{"天安门", "ns"}},
		[]WordTag{WordTag{"abc", "eng"}},
		[]WordTag{WordTag{"隐", "n"}, WordTag{"马尔可夫", "nr"}},
		[]WordTag{WordTag{"雷猴", "n"}, WordTag{"是", "v"}, WordTag{"个", "q"}, WordTag{"好", "a"}, WordTag{"网站", "n"}},
		[]WordTag{WordTag{"“", "x"}, WordTag{"Microsoft", "eng"}, WordTag{"”", "x"}, WordTag{"一", "m"}, WordTag{"词", "n"}, WordTag{"由", "p"}, WordTag{"“", "x"}, WordTag{"MICROcomputer", "eng"}, WordTag{"（", "x"}, WordTag{"微型", "b"}, WordTag{"计算机", "n"}, WordTag{"）", "x"}, WordTag{"”", "x"}, WordTag{"和", "c"}, WordTag{"“", "x"}, WordTag{"SOFTware", "eng"}, WordTag{"（", "x"}, WordTag{"软件", "n"}, WordTag{"）", "x"}, WordTag{"”", "x"}, WordTag{"两", "m"}, WordTag{"部分", "n"}, WordTag{"组成", "v"}},
		[]WordTag{WordTag{"草泥马", "n"}, WordTag{"和", "c"}, WordTag{"欺", "vn"}, WordTag{"实", "n"}, WordTag{"马", "n"}, WordTag{"是", "v"}, WordTag{"今年", "t"}, WordTag{"的", "uj"}, WordTag{"流行", "v"}, WordTag{"词汇", "n"}},
		[]WordTag{WordTag{"伊", "ns"}, WordTag{"藤", "nr"}, WordTag{"洋华堂", "n"}, WordTag{"总府", "n"}, WordTag{"店", "n"}},
		[]WordTag{WordTag{"中国科学院计算技术研究所", "nt"}},
		[]WordTag{WordTag{"罗密欧", "nr"}, WordTag{"与", "p"}, WordTag{"朱丽叶", "nr"}},
		[]WordTag{WordTag{"我", "r"}, WordTag{"购买", "v"}, WordTag{"了", "ul"}, WordTag{"道具", "n"}, WordTag{"和", "c"}, WordTag{"服装", "vn"}},
		[]WordTag{WordTag{"PS", "eng"}, WordTag{":", "x"}, WordTag{" ", "x"}, WordTag{"我", "r"}, WordTag{"觉得", "v"}, WordTag{"开源", "n"}, WordTag{"有", "v"}, WordTag{"一个", "m"}, WordTag{"好处", "d"}, WordTag{"，", "x"}, WordTag{"就是", "d"}, WordTag{"能够", "v"}, WordTag{"敦促", "v"}, WordTag{"自己", "r"}, WordTag{"不断改进", "l"}, WordTag{"，", "x"}, WordTag{"避免", "v"}, WordTag{"敞", "v"}, WordTag{"帚", "ng"}, WordTag{"自珍", "b"}},
		[]WordTag{WordTag{"湖北省", "ns"}, WordTag{"石首市", "ns"}},
		[]WordTag{WordTag{"湖北省", "ns"}, WordTag{"十堰市", "ns"}},
		[]WordTag{WordTag{"总经理", "n"}, WordTag{"完成", "v"}, WordTag{"了", "ul"}, WordTag{"这件", "mq"}, WordTag{"事情", "n"}},
		[]WordTag{WordTag{"电脑", "n"}, WordTag{"修好", "v"}, WordTag{"了", "ul"}},
		[]WordTag{WordTag{"做好", "v"}, WordTag{"了", "ul"}, WordTag{"这件", "mq"}, WordTag{"事情", "n"}, WordTag{"就", "d"}, WordTag{"一了百了", "l"}, WordTag{"了", "ul"}},
		[]WordTag{WordTag{"人们", "n"}, WordTag{"审美", "vn"}, WordTag{"的", "uj"}, WordTag{"观点", "n"}, WordTag{"是", "v"}, WordTag{"不同", "a"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"我们", "r"}, WordTag{"买", "v"}, WordTag{"了", "ul"}, WordTag{"一个", "m"}, WordTag{"美的", "nr"}, WordTag{"空调", "n"}},
		[]WordTag{WordTag{"线程", "n"}, WordTag{"初始化", "l"}, WordTag{"时", "n"}, WordTag{"我们", "r"}, WordTag{"要", "v"}, WordTag{"注意", "v"}},
		[]WordTag{WordTag{"一个", "m"}, WordTag{"分子", "n"}, WordTag{"是", "v"}, WordTag{"由", "p"}, WordTag{"好多", "m"}, WordTag{"原子", "n"}, WordTag{"组织", "v"}, WordTag{"成", "n"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"祝", "v"}, WordTag{"你", "r"}, WordTag{"马到功成", "i"}},
		[]WordTag{WordTag{"他", "r"}, WordTag{"掉", "zg"}, WordTag{"进", "v"}, WordTag{"了", "ul"}, WordTag{"无底洞", "ns"}, WordTag{"里", "f"}},
		[]WordTag{WordTag{"中国", "ns"}, WordTag{"的", "uj"}, WordTag{"首都", "d"}, WordTag{"是", "v"}, WordTag{"北京", "ns"}},
		[]WordTag{WordTag{"孙", "zg"}, WordTag{"君", "nz"}, WordTag{"意", "n"}},
		[]WordTag{WordTag{"外交部", "nt"}, WordTag{"发言人", "l"}, WordTag{"马朝旭", "nr"}},
		[]WordTag{WordTag{"领导人", "n"}, WordTag{"会议", "n"}, WordTag{"和", "c"}, WordTag{"第四届", "m"}, WordTag{"东亚", "ns"}, WordTag{"峰会", "n"}},
		[]WordTag{WordTag{"在", "p"}, WordTag{"过去", "t"}, WordTag{"的", "uj"}, WordTag{"这", "r"}, WordTag{"五年", "t"}},
		[]WordTag{WordTag{"还", "d"}, WordTag{"需要", "v"}, WordTag{"很", "zg"}, WordTag{"长", "a"}, WordTag{"的", "uj"}, WordTag{"路", "n"}, WordTag{"要", "v"}, WordTag{"走", "v"}},
		[]WordTag{WordTag{"60", "eng"}, WordTag{"周年", "t"}, WordTag{"首都", "d"}, WordTag{"阅兵", "v"}},
		[]WordTag{WordTag{"你好", "l"}, WordTag{"人们", "n"}, WordTag{"审美", "vn"}, WordTag{"的", "uj"}, WordTag{"观点", "n"}, WordTag{"是", "v"}, WordTag{"不同", "a"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"买", "v"}, WordTag{"水果", "n"}, WordTag{"然后", "c"}, WordTag{"来", "v"}, WordTag{"世博园", "nr"}},
		[]WordTag{WordTag{"买", "v"}, WordTag{"水果", "n"}, WordTag{"然后", "c"}, WordTag{"去", "v"}, WordTag{"世博园", "nr"}},
		[]WordTag{WordTag{"但是", "c"}, WordTag{"后来", "t"}, WordTag{"我", "r"}, WordTag{"才", "d"}, WordTag{"知道", "v"}, WordTag{"你", "r"}, WordTag{"是", "v"}, WordTag{"对", "p"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"存在", "v"}, WordTag{"即", "v"}, WordTag{"合理", "vn"}},
		[]WordTag{WordTag{"的", "uj"}, WordTag{"的", "uj"}, WordTag{"的", "uj"}, WordTag{"的", "uj"}, WordTag{"的", "uj"}, WordTag{"在", "p"}, WordTag{"的", "uj"}, WordTag{"的", "uj"}, WordTag{"的", "uj"}, WordTag{"的", "uj"}, WordTag{"就", "d"}, WordTag{"以", "p"}, WordTag{"和", "c"}, WordTag{"和", "c"}, WordTag{"和", "c"}},
		[]WordTag{WordTag{"I", "eng"}, WordTag{" ", "x"}, WordTag{"love", "eng"}, WordTag{"你", "r"}, WordTag{"，", "x"}, WordTag{"不以为耻", "i"}, WordTag{"，", "x"}, WordTag{"反", "zg"}, WordTag{"以为", "c"}, WordTag{"rong", "eng"}},
		[]WordTag{WordTag{"因", "p"}},
		[]WordTag{},
		[]WordTag{WordTag{"hello", "eng"}, WordTag{"你好", "l"}, WordTag{"人们", "n"}, WordTag{"审美", "vn"}, WordTag{"的", "uj"}, WordTag{"观点", "n"}, WordTag{"是", "v"}, WordTag{"不同", "a"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"很", "zg"}, WordTag{"好", "a"}, WordTag{"但", "c"}, WordTag{"主要", "b"}, WordTag{"是", "v"}, WordTag{"基于", "p"}, WordTag{"网页", "n"}, WordTag{"形式", "n"}},
		[]WordTag{WordTag{"hello", "eng"}, WordTag{"你好", "l"}, WordTag{"人们", "n"}, WordTag{"审美", "vn"}, WordTag{"的", "uj"}, WordTag{"观点", "n"}, WordTag{"是", "v"}, WordTag{"不同", "a"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"为什么", "r"}, WordTag{"我", "r"}, WordTag{"不能", "v"}, WordTag{"拥有", "v"}, WordTag{"想要", "v"}, WordTag{"的", "uj"}, WordTag{"生活", "vn"}},
		[]WordTag{WordTag{"后来", "t"}, WordTag{"我", "r"}, WordTag{"才", "d"}},
		[]WordTag{WordTag{"此次", "r"}, WordTag{"来", "v"}, WordTag{"中国", "ns"}, WordTag{"是", "v"}, WordTag{"为了", "p"}},
		[]WordTag{WordTag{"使用", "v"}, WordTag{"了", "ul"}, WordTag{"它", "r"}, WordTag{"就", "d"}, WordTag{"可以", "c"}, WordTag{"解决", "v"}, WordTag{"一些", "m"}, WordTag{"问题", "n"}},
		[]WordTag{WordTag{",", "x"}, WordTag{"使用", "v"}, WordTag{"了", "ul"}, WordTag{"它", "r"}, WordTag{"就", "d"}, WordTag{"可以", "c"}, WordTag{"解决", "v"}, WordTag{"一些", "m"}, WordTag{"问题", "n"}},
		[]WordTag{WordTag{"其实", "d"}, WordTag{"使用", "v"}, WordTag{"了", "ul"}, WordTag{"它", "r"}, WordTag{"就", "d"}, WordTag{"可以", "c"}, WordTag{"解决", "v"}, WordTag{"一些", "m"}, WordTag{"问题", "n"}},
		[]WordTag{WordTag{"好人", "n"}, WordTag{"使用", "v"}, WordTag{"了", "ul"}, WordTag{"它", "r"}, WordTag{"就", "d"}, WordTag{"可以", "c"}, WordTag{"解决", "v"}, WordTag{"一些", "m"}, WordTag{"问题", "n"}},
		[]WordTag{WordTag{"是因为", "c"}, WordTag{"和", "c"}, WordTag{"国家", "n"}},
		[]WordTag{WordTag{"老年", "t"}, WordTag{"搜索", "v"}, WordTag{"还", "d"}, WordTag{"支持", "v"}},
		[]WordTag{WordTag{"干脆", "d"}, WordTag{"就", "d"}, WordTag{"把", "p"}, WordTag{"那", "r"}, WordTag{"部", "n"}, WordTag{"蒙", "v"}, WordTag{"人", "n"}, WordTag{"的", "uj"}, WordTag{"闲", "n"}, WordTag{"法", "j"}, WordTag{"给", "p"}, WordTag{"废", "v"}, WordTag{"了", "ul"}, WordTag{"拉倒", "v"}, WordTag{"！", "x"}, WordTag{"RT", "eng"}, WordTag{" ", "x"}, WordTag{"@", "x"}, WordTag{"laoshipukong", "eng"}, WordTag{" ", "x"}, WordTag{":", "x"}, WordTag{" ", "x"}, WordTag{"27", "eng"}, WordTag{"日", "m"}, WordTag{"，", "x"}, WordTag{"全国人大常委会", "nt"}, WordTag{"第三次", "m"}, WordTag{"审议", "v"}, WordTag{"侵权", "v"}, WordTag{"责任法", "n"}, WordTag{"草案", "n"}, WordTag{"，", "x"}, WordTag{"删除", "v"}, WordTag{"了", "ul"}, WordTag{"有关", "vn"}, WordTag{"医疗", "n"}, WordTag{"损害", "v"}, WordTag{"责任", "n"}, WordTag{"“", "x"}, WordTag{"举证", "v"}, WordTag{"倒置", "v"}, WordTag{"”", "x"}, WordTag{"的", "uj"}, WordTag{"规定", "n"}, WordTag{"。", "x"}, WordTag{"在", "p"}, WordTag{"医患", "n"}, WordTag{"纠纷", "n"}, WordTag{"中", "f"}, WordTag{"本", "r"}, WordTag{"已", "d"}, WordTag{"处于", "v"}, WordTag{"弱势", "n"}, WordTag{"地位", "n"}, WordTag{"的", "uj"}, WordTag{"消费者", "n"}, WordTag{"由此", "c"}, WordTag{"将", "d"}, WordTag{"陷入", "v"}, WordTag{"万劫不复", "i"}, WordTag{"的", "uj"}, WordTag{"境地", "s"}, WordTag{"。", "x"}, WordTag{" ", "x"}},
		[]WordTag{WordTag{"大", "a"}},
		[]WordTag{},
		[]WordTag{WordTag{"他", "r"}, WordTag{"说", "v"}, WordTag{"的", "uj"}, WordTag{"确实", "ad"}, WordTag{"在", "p"}, WordTag{"理", "n"}},
		[]WordTag{WordTag{"长春", "ns"}, WordTag{"市长", "n"}, WordTag{"春节", "t"}, WordTag{"讲话", "n"}},
		[]WordTag{WordTag{"结婚", "v"}, WordTag{"的", "uj"}, WordTag{"和", "c"}, WordTag{"尚未", "d"}, WordTag{"结婚", "v"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"结合", "v"}, WordTag{"成", "n"}, WordTag{"分子", "n"}, WordTag{"时", "n"}},
		[]WordTag{WordTag{"旅游", "vn"}, WordTag{"和", "c"}, WordTag{"服务", "vn"}, WordTag{"是", "v"}, WordTag{"最好", "a"}, WordTag{"的", "uj"}},
		[]WordTag{WordTag{"这件", "mq"}, WordTag{"事情", "n"}, WordTag{"的确", "d"}, WordTag{"是", "v"}, WordTag{"我", "r"}, WordTag{"的", "uj"}, WordTag{"错", "v"}},
		[]WordTag{WordTag{"供", "v"}, WordTag{"大家", "n"}, WordTag{"参考", "v"}, WordTag{"指正", "v"}},
		[]WordTag{WordTag{"哈尔滨", "ns"}, WordTag{"政府", "n"}, WordTag{"公布", "v"}, WordTag{"塌", "v"}, WordTag{"桥", "n"}, WordTag{"原因", "n"}},
		[]WordTag{WordTag{"我", "r"}, WordTag{"在", "p"}, WordTag{"机场", "n"}, WordTag{"入口处", "i"}},
		[]WordTag{WordTag{"邢", "nr"}, WordTag{"永", "ns"}, WordTag{"臣", "n"}, WordTag{"摄影", "n"}, WordTag{"报道", "v"}},
		[]WordTag{WordTag{"BP", "eng"}, WordTag{"神经网络", "n"}, WordTag{"如何", "r"}, WordTag{"训练", "vn"}, WordTag{"才能", "v"}, WordTag{"在", "p"}, WordTag{"分类", "n"}, WordTag{"时", "n"}, WordTag{"增加", "v"}, WordTag{"区分度", "n"}, WordTag{"？", "x"}},
		[]WordTag{WordTag{"南京市", "ns"}, WordTag{"长江大桥", "ns"}},
		[]WordTag{WordTag{"应", "v"}, WordTag{"一些", "m"}, WordTag{"使用者", "n"}, WordTag{"的", "uj"}, WordTag{"建议", "n"}, WordTag{"，", "x"}, WordTag{"也", "d"}, WordTag{"为了", "p"}, WordTag{"便于", "v"}, WordTag{"利用", "n"}, WordTag{"NiuTrans", "eng"}, WordTag{"用于", "v"}, WordTag{"SMT", "eng"}, WordTag{"研究", "vn"}},
		[]WordTag{WordTag{"长春市", "ns"}, WordTag{"长春", "ns"}, WordTag{"药店", "n"}},
		[]WordTag{WordTag{"邓颖超", "nr"}, WordTag{"生前", "t"}, WordTag{"最", "d"}, WordTag{"喜欢", "v"}, WordTag{"的", "uj"}, WordTag{"衣服", "n"}},
		[]WordTag{WordTag{"胡锦涛", "nr"}, WordTag{"是", "v"}, WordTag{"热爱", "a"}, WordTag{"世界", "n"}, WordTag{"和平", "nz"}, WordTag{"的", "uj"}, WordTag{"政治局", "n"}, WordTag{"常委", "j"}},
		[]WordTag{WordTag{"程序员", "n"}, WordTag{"祝", "v"}, WordTag{"海林", "nz"}, WordTag{"和", "c"}, WordTag{"朱", "nr"}, WordTag{"会", "v"}, WordTag{"震", "v"}, WordTag{"是", "v"}, WordTag{"在", "p"}, WordTag{"孙", "zg"}, WordTag{"健", "a"}, WordTag{"的", "uj"}, WordTag{"左面", "f"}, WordTag{"和", "c"}, WordTag{"右面", "f"}, WordTag{",", "x"}, WordTag{" ", "x"}, WordTag{"范", "nr"}, WordTag{"凯", "nr"}, WordTag{"在", "p"}, WordTag{"最", "d"}, WordTag{"右面", "f"}, WordTag{".", "x"}, WordTag{"再", "d"}, WordTag{"往", "zg"}, WordTag{"左", "m"}, WordTag{"是", "v"}, WordTag{"李", "nr"}, WordTag{"松", "v"}, WordTag{"洪", "nr"}},
		[]WordTag{WordTag{"一次性", "d"}, WordTag{"交", "v"}, WordTag{"多少", "m"}, WordTag{"钱", "n"}},
		[]WordTag{WordTag{"两块", "m"}, WordTag{"五", "m"}, WordTag{"一套", "m"}, WordTag{"，", "x"}, WordTag{"三块", "m"}, WordTag{"八", "m"}, WordTag{"一斤", "m"}, WordTag{"，", "x"}, WordTag{"四块", "m"}, WordTag{"七", "m"}, WordTag{"一本", "m"}, WordTag{"，", "x"}, WordTag{"五块", "m"}, WordTag{"六", "m"}, WordTag{"一条", "m"}},
		[]WordTag{WordTag{"小", "a"}, WordTag{"和尚", "nr"}, WordTag{"留", "v"}, WordTag{"了", "ul"}, WordTag{"一个", "m"}, WordTag{"像", "v"}, WordTag{"大", "a"}, WordTag{"和尚", "nr"}, WordTag{"一样", "r"}, WordTag{"的", "uj"}, WordTag{"和尚头", "nr"}},
		[]WordTag{WordTag{"我", "r"}, WordTag{"是", "v"}, WordTag{"中华人民共和国", "ns"}, WordTag{"公民", "n"}, WordTag{";", "x"}, WordTag{"我", "r"}, WordTag{"爸爸", "n"}, WordTag{"是", "v"}, WordTag{"共和党", "nt"}, WordTag{"党员", "n"}, WordTag{";", "x"}, WordTag{" ", "x"}, WordTag{"地铁", "n"}, WordTag{"和平门", "ns"}, WordTag{"站", "v"}},
		[]WordTag{WordTag{"张晓梅", "nr"}, WordTag{"去", "v"}, WordTag{"人民", "n"}, WordTag{"医院", "n"}, WordTag{"做", "v"}, WordTag{"了", "ul"}, WordTag{"个", "q"}, WordTag{"B超", "n"}, WordTag{"然后", "c"}, WordTag{"去", "v"}, WordTag{"买", "v"}, WordTag{"了", "ul"}, WordTag{"件", "zg"}, WordTag{"T恤", "n"}},
		[]WordTag{WordTag{"AT&T", "nz"}, WordTag{"是", "v"}, WordTag{"一件", "m"}, WordTag{"不错", "a"}, WordTag{"的", "uj"}, WordTag{"公司", "n"}, WordTag{"，", "x"}, WordTag{"给", "p"}, WordTag{"你", "r"}, WordTag{"发", "v"}, WordTag{"offer", "eng"}, WordTag{"了", "ul"}, WordTag{"吗", "y"}, WordTag{"？", "x"}},
		[]WordTag{WordTag{"C++", "nz"}, WordTag{"和", "c"}, WordTag{"c#", "nz"}, WordTag{"是", "v"}, WordTag{"什么", "r"}, WordTag{"关系", "n"}, WordTag{"？", "x"}, WordTag{"11", "eng"}, WordTag{"+", "x"}, WordTag{"122", "eng"}, WordTag{"=", "x"}, WordTag{"133", "eng"}, WordTag{"，", "x"}, WordTag{"是", "v"}, WordTag{"吗", "y"}, WordTag{"？", "x"}, WordTag{"PI", "eng"}, WordTag{"=", "x"}, WordTag{"3", "eng"}, WordTag{".", "x"}, WordTag{"14159", "eng"}},
		[]WordTag{WordTag{"你", "r"}, WordTag{"认识", "v"}, WordTag{"那个", "r"}, WordTag{"和", "c"}, WordTag{"主席", "n"}, WordTag{"握手", "v"}, WordTag{"的", "uj"}, WordTag{"的哥", "n"}, WordTag{"吗", "y"}, WordTag{"？", "x"}, WordTag{"他", "r"}, WordTag{"开", "v"}, WordTag{"一辆", "m"}, WordTag{"黑色", "n"}, WordTag{"的士", "n"}, WordTag{"。", "x"}},
		[]WordTag{WordTag{"枪杆子", "n"}, WordTag{"中", "f"}, WordTag{"出", "v"}, WordTag{"政权", "n"}},
	}
)

func chanToArray(ch chan WordTag) []WordTag {
	result := make([]WordTag, 0)
	for word := range ch {
		result = append(result, word)
	}
	return result
}

func TestCut(t *testing.T) {
	p, err := NewPosseg("../dict.txt")
	if err != nil {
		t.Fatal(err)
	}
	for index, content := range test_contents {
		result := chanToArray(p.Cut(content, true))
		if len(defaultCutResult[index]) != len(result) {
			t.Error(content)
		}
		for i, _ := range result {
			if result[i] != defaultCutResult[index][i] {
				t.Errorf("expect %s, got %s", defaultCutResult[index][i], result[i])
			}
		}
		result = chanToArray(p.Cut(content, false))
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
	p, _ := NewPosseg("../dict.txt")
	sentence := "又跛又啞"
	cutResult := []WordTag{
		WordTag{"又", "d"},
		WordTag{"跛", "a"},
		WordTag{"又", "d"},
		WordTag{"啞", "v"},
	}
	result := chanToArray(p.Cut(sentence, true))
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
	p, _ := NewPosseg("../dict.txt")
	sentence := "前港督衛奕信在八八年十月宣布成立中央政策研究組"
	cutResult := []WordTag{
		WordTag{"前", "f"},
		WordTag{"港督", "n"},
		WordTag{"衛奕", "z"},
		WordTag{"信", "n"},
		WordTag{"在", "p"},
		WordTag{"八八年", "m"},
		WordTag{"十月", "t"},
		WordTag{"宣布", "v"},
		WordTag{"成立", "v"},
		WordTag{"中央", "n"},
		WordTag{"政策", "n"},
		WordTag{"研究", "vn"},
		WordTag{"組", "x"},
	}
	result := chanToArray(p.Cut(sentence, true))
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
	p, _ := NewPosseg("../dict.txt")
	p.LoadUserDict("../userdict.txt")
	sentence := "李小福是创新办主任也是云计算方面的专家; 什么是八一双鹿例如我输入一个带“韩玉赏鉴”的标题，在自定义词库中也增加了此词为N类型"

	cutResult := []WordTag{
		WordTag{"李小福", "nr"},
		WordTag{"是", "v"},
		WordTag{"创新办", "i"},
		WordTag{"主任", "b"},
		WordTag{"也", "d"},
		WordTag{"是", "v"},
		WordTag{"云计算", "x"},
		WordTag{"方面", "n"},
		WordTag{"的", "uj"},
		WordTag{"专家", "n"},
		WordTag{";", "x"},
		WordTag{" ", "x"},
		WordTag{"什么", "r"},
		WordTag{"是", "v"},
		WordTag{"八一双鹿", "nz"},
		WordTag{"例如", "v"},
		WordTag{"我", "r"},
		WordTag{"输入", "v"},
		WordTag{"一个", "m"},
		WordTag{"带", "v"},
		WordTag{"“", "x"},
		WordTag{"韩玉赏鉴", "nz"},
		WordTag{"”", "x"},
		WordTag{"的", "uj"},
		WordTag{"标题", "n"},
		WordTag{"，", "x"},
		WordTag{"在", "p"},
		WordTag{"自定义词", "n"},
		WordTag{"库中", "nrt"},
		WordTag{"也", "d"},
		WordTag{"增加", "v"},
		WordTag{"了", "ul"},
		WordTag{"此", "r"},
		WordTag{"词", "n"},
		WordTag{"为", "p"},
		WordTag{"N", "eng"},
		WordTag{"类型", "n"}}

	result := chanToArray(p.Cut(sentence, true))
	if len(cutResult) != len(result) {
		t.Error(result)
	}
	for i, _ := range result {
		if result[i] != cutResult[i] {
			t.Error(result[i])
		}
	}
}
