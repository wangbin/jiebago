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

	defaultCutResult = [][]Pair{[]Pair{Pair{"这", "r"}, Pair{"是", "v"}, Pair{"一个", "m"}, Pair{"伸手不见五指", "i"}, Pair{"的", "uj"}, Pair{"黑夜", "n"}, Pair{"。", "x"}, Pair{"我", "r"}, Pair{"叫", "v"}, Pair{"孙悟空", "nr"}, Pair{"，", "x"}, Pair{"我", "r"}, Pair{"爱", "v"}, Pair{"北京", "ns"}, Pair{"，", "x"}, Pair{"我", "r"}, Pair{"爱", "v"}, Pair{"Python", "eng"}, Pair{"和", "c"}, Pair{"C++", "nz"}, Pair{"。", "x"}},
		[]Pair{Pair{"我", "r"}, Pair{"不", "d"}, Pair{"喜欢", "v"}, Pair{"日本", "ns"}, Pair{"和服", "nz"}, Pair{"。", "x"}},
		[]Pair{Pair{"雷猴", "n"}, Pair{"回归", "v"}, Pair{"人间", "n"}, Pair{"。", "x"}},
		[]Pair{Pair{"工信处", "n"}, Pair{"女干事", "n"}, Pair{"每月", "r"}, Pair{"经过", "p"}, Pair{"下属", "v"}, Pair{"科室", "n"}, Pair{"都", "d"}, Pair{"要", "v"}, Pair{"亲口", "n"}, Pair{"交代", "n"}, Pair{"24", "m"}, Pair{"口", "n"}, Pair{"交换机", "n"}, Pair{"等", "u"}, Pair{"技术性", "n"}, Pair{"器件", "n"}, Pair{"的", "uj"}, Pair{"安装", "v"}, Pair{"工作", "vn"}},
		[]Pair{Pair{"我", "r"}, Pair{"需要", "v"}, Pair{"廉租房", "n"}},
		[]Pair{Pair{"永和", "nz"}, Pair{"服装", "vn"}, Pair{"饰品", "n"}, Pair{"有限公司", "n"}},
		[]Pair{Pair{"我", "r"}, Pair{"爱", "v"}, Pair{"北京", "ns"}, Pair{"天安门", "ns"}},
		[]Pair{Pair{"abc", "eng"}},
		[]Pair{Pair{"隐", "n"}, Pair{"马尔可夫", "nr"}},
		[]Pair{Pair{"雷猴", "n"}, Pair{"是", "v"}, Pair{"个", "q"}, Pair{"好", "a"}, Pair{"网站", "n"}},
		[]Pair{Pair{"“", "x"}, Pair{"Microsoft", "eng"}, Pair{"”", "x"}, Pair{"一", "m"}, Pair{"词", "n"}, Pair{"由", "p"}, Pair{"“", "x"}, Pair{"MICROcomputer", "eng"}, Pair{"（", "x"}, Pair{"微型", "b"}, Pair{"计算机", "n"}, Pair{"）", "x"}, Pair{"”", "x"}, Pair{"和", "c"}, Pair{"“", "x"}, Pair{"SOFTware", "eng"}, Pair{"（", "x"}, Pair{"软件", "n"}, Pair{"）", "x"}, Pair{"”", "x"}, Pair{"两", "m"}, Pair{"部分", "n"}, Pair{"组成", "v"}},
		[]Pair{Pair{"草泥马", "n"}, Pair{"和", "c"}, Pair{"欺实", "v"}, Pair{"马", "n"}, Pair{"是", "v"}, Pair{"今年", "t"}, Pair{"的", "uj"}, Pair{"流行", "v"}, Pair{"词汇", "n"}},
		[]Pair{Pair{"伊藤", "nr"}, Pair{"洋华堂", "n"}, Pair{"总府", "n"}, Pair{"店", "n"}},
		[]Pair{Pair{"中国科学院计算技术研究所", "nt"}},
		[]Pair{Pair{"罗密欧", "nr"}, Pair{"与", "p"}, Pair{"朱丽叶", "nr"}},
		[]Pair{Pair{"我", "r"}, Pair{"购买", "v"}, Pair{"了", "ul"}, Pair{"道具", "n"}, Pair{"和", "c"}, Pair{"服装", "vn"}},
		[]Pair{Pair{"PS", "eng"}, Pair{":", "x"}, Pair{" ", "x"}, Pair{"我", "r"}, Pair{"觉得", "v"}, Pair{"开源", "n"}, Pair{"有", "v"}, Pair{"一个", "m"}, Pair{"好处", "d"}, Pair{"，", "x"}, Pair{"就是", "d"}, Pair{"能够", "v"}, Pair{"敦促", "v"}, Pair{"自己", "r"}, Pair{"不断改进", "l"}, Pair{"，", "x"}, Pair{"避免", "v"}, Pair{"敞", "v"}, Pair{"帚", "ng"}, Pair{"自珍", "b"}},
		[]Pair{Pair{"湖北省", "ns"}, Pair{"石首市", "ns"}},
		[]Pair{Pair{"湖北省", "ns"}, Pair{"十堰市", "ns"}},
		[]Pair{Pair{"总经理", "n"}, Pair{"完成", "v"}, Pair{"了", "ul"}, Pair{"这件", "mq"}, Pair{"事情", "n"}},
		[]Pair{Pair{"电脑", "n"}, Pair{"修好", "v"}, Pair{"了", "ul"}},
		[]Pair{Pair{"做好", "v"}, Pair{"了", "ul"}, Pair{"这件", "mq"}, Pair{"事情", "n"}, Pair{"就", "d"}, Pair{"一了百了", "l"}, Pair{"了", "ul"}},
		[]Pair{Pair{"人们", "n"}, Pair{"审美", "vn"}, Pair{"的", "uj"}, Pair{"观点", "n"}, Pair{"是", "v"}, Pair{"不同", "a"}, Pair{"的", "uj"}},
		[]Pair{Pair{"我们", "r"}, Pair{"买", "v"}, Pair{"了", "ul"}, Pair{"一个", "m"}, Pair{"美的", "nr"}, Pair{"空调", "n"}},
		[]Pair{Pair{"线程", "n"}, Pair{"初始化", "l"}, Pair{"时", "n"}, Pair{"我们", "r"}, Pair{"要", "v"}, Pair{"注意", "v"}},
		[]Pair{Pair{"一个", "m"}, Pair{"分子", "n"}, Pair{"是", "v"}, Pair{"由", "p"}, Pair{"好多", "m"}, Pair{"原子", "n"}, Pair{"组织", "v"}, Pair{"成", "v"}, Pair{"的", "uj"}},
		[]Pair{Pair{"祝", "v"}, Pair{"你", "r"}, Pair{"马到功成", "i"}},
		[]Pair{Pair{"他", "r"}, Pair{"掉", "v"}, Pair{"进", "v"}, Pair{"了", "ul"}, Pair{"无底洞", "ns"}, Pair{"里", "f"}},
		[]Pair{Pair{"中国", "ns"}, Pair{"的", "uj"}, Pair{"首都", "d"}, Pair{"是", "v"}, Pair{"北京", "ns"}},
		[]Pair{Pair{"孙君意", "nr"}},
		[]Pair{Pair{"外交部", "nt"}, Pair{"发言人", "l"}, Pair{"马朝旭", "nr"}},
		[]Pair{Pair{"领导人", "n"}, Pair{"会议", "n"}, Pair{"和", "c"}, Pair{"第四届", "m"}, Pair{"东亚", "ns"}, Pair{"峰会", "n"}},
		[]Pair{Pair{"在", "p"}, Pair{"过去", "t"}, Pair{"的", "uj"}, Pair{"这", "r"}, Pair{"五年", "t"}},
		[]Pair{Pair{"还", "d"}, Pair{"需要", "v"}, Pair{"很", "d"}, Pair{"长", "a"}, Pair{"的", "uj"}, Pair{"路", "n"}, Pair{"要", "v"}, Pair{"走", "v"}},
		[]Pair{Pair{"60", "m"}, Pair{"周年", "t"}, Pair{"首都", "d"}, Pair{"阅兵", "v"}},
		[]Pair{Pair{"你好", "l"}, Pair{"人们", "n"}, Pair{"审美", "vn"}, Pair{"的", "uj"}, Pair{"观点", "n"}, Pair{"是", "v"}, Pair{"不同", "a"}, Pair{"的", "uj"}},
		[]Pair{Pair{"买", "v"}, Pair{"水果", "n"}, Pair{"然后", "c"}, Pair{"来", "v"}, Pair{"世博园", "nr"}},
		[]Pair{Pair{"买", "v"}, Pair{"水果", "n"}, Pair{"然后", "c"}, Pair{"去", "v"}, Pair{"世博园", "nr"}},
		[]Pair{Pair{"但是", "c"}, Pair{"后来", "t"}, Pair{"我", "r"}, Pair{"才", "d"}, Pair{"知道", "v"}, Pair{"你", "r"}, Pair{"是", "v"}, Pair{"对", "p"}, Pair{"的", "uj"}},
		[]Pair{Pair{"存在", "v"}, Pair{"即", "v"}, Pair{"合理", "vn"}},
		[]Pair{Pair{"的的", "u"}, Pair{"的的", "u"}, Pair{"的", "uj"}, Pair{"在的", "u"}, Pair{"的的", "u"}, Pair{"的", "uj"}, Pair{"就", "d"}, Pair{"以", "p"}, Pair{"和和", "nz"}, Pair{"和", "c"}},
		[]Pair{Pair{"I", "x"}, Pair{" ", "x"}, Pair{"love", "eng"}, Pair{"你", "r"}, Pair{"，", "x"}, Pair{"不以为耻", "i"}, Pair{"，", "x"}, Pair{"反", "zg"}, Pair{"以为", "c"}, Pair{"rong", "eng"}},
		[]Pair{Pair{"因", "p"}},
		[]Pair{},
		[]Pair{Pair{"hello", "eng"}, Pair{"你好", "l"}, Pair{"人们", "n"}, Pair{"审美", "vn"}, Pair{"的", "uj"}, Pair{"观点", "n"}, Pair{"是", "v"}, Pair{"不同", "a"}, Pair{"的", "uj"}},
		[]Pair{Pair{"很好", "a"}, Pair{"但", "c"}, Pair{"主要", "b"}, Pair{"是", "v"}, Pair{"基于", "p"}, Pair{"网页", "n"}, Pair{"形式", "n"}},
		[]Pair{Pair{"hello", "eng"}, Pair{"你好", "l"}, Pair{"人们", "n"}, Pair{"审美", "vn"}, Pair{"的", "uj"}, Pair{"观点", "n"}, Pair{"是", "v"}, Pair{"不同", "a"}, Pair{"的", "uj"}},
		[]Pair{Pair{"为什么", "r"}, Pair{"我", "r"}, Pair{"不能", "v"}, Pair{"拥有", "v"}, Pair{"想要", "v"}, Pair{"的", "uj"}, Pair{"生活", "vn"}},
		[]Pair{Pair{"后来", "t"}, Pair{"我", "r"}, Pair{"才", "d"}},
		[]Pair{Pair{"此次", "r"}, Pair{"来", "v"}, Pair{"中国", "ns"}, Pair{"是", "v"}, Pair{"为了", "p"}},
		[]Pair{Pair{"使用", "v"}, Pair{"了", "ul"}, Pair{"它", "r"}, Pair{"就", "d"}, Pair{"可以", "c"}, Pair{"解决", "v"}, Pair{"一些", "m"}, Pair{"问题", "n"}},
		[]Pair{Pair{",", "x"}, Pair{"使用", "v"}, Pair{"了", "ul"}, Pair{"它", "r"}, Pair{"就", "d"}, Pair{"可以", "c"}, Pair{"解决", "v"}, Pair{"一些", "m"}, Pair{"问题", "n"}},
		[]Pair{Pair{"其实", "d"}, Pair{"使用", "v"}, Pair{"了", "ul"}, Pair{"它", "r"}, Pair{"就", "d"}, Pair{"可以", "c"}, Pair{"解决", "v"}, Pair{"一些", "m"}, Pair{"问题", "n"}},
		[]Pair{Pair{"好人", "n"}, Pair{"使用", "v"}, Pair{"了", "ul"}, Pair{"它", "r"}, Pair{"就", "d"}, Pair{"可以", "c"}, Pair{"解决", "v"}, Pair{"一些", "m"}, Pair{"问题", "n"}},
		[]Pair{Pair{"是因为", "c"}, Pair{"和", "c"}, Pair{"国家", "n"}},
		[]Pair{Pair{"老年", "t"}, Pair{"搜索", "v"}, Pair{"还", "d"}, Pair{"支持", "v"}},
		[]Pair{Pair{"干脆", "d"}, Pair{"就", "d"}, Pair{"把", "p"}, Pair{"那部", "r"}, Pair{"蒙人", "n"}, Pair{"的", "uj"}, Pair{"闲法", "n"}, Pair{"给", "p"}, Pair{"废", "v"}, Pair{"了", "ul"}, Pair{"拉倒", "v"}, Pair{"！", "x"}, Pair{"RT", "eng"}, Pair{" ", "x"}, Pair{"@", "x"}, Pair{"laoshipukong", "eng"}, Pair{" ", "x"}, Pair{":", "x"}, Pair{" ", "x"}, Pair{"27", "m"}, Pair{"日", "m"}, Pair{"，", "x"}, Pair{"全国人大常委会", "nt"}, Pair{"第三次", "m"}, Pair{"审议", "v"}, Pair{"侵权", "v"}, Pair{"责任法", "n"}, Pair{"草案", "n"}, Pair{"，", "x"}, Pair{"删除", "v"}, Pair{"了", "ul"}, Pair{"有关", "vn"}, Pair{"医疗", "n"}, Pair{"损害", "v"}, Pair{"责任", "n"}, Pair{"“", "x"}, Pair{"举证", "v"}, Pair{"倒置", "v"}, Pair{"”", "x"}, Pair{"的", "uj"}, Pair{"规定", "n"}, Pair{"。", "x"}, Pair{"在", "p"}, Pair{"医患", "n"}, Pair{"纠纷", "n"}, Pair{"中本", "ns"}, Pair{"已", "d"}, Pair{"处于", "v"}, Pair{"弱势", "n"}, Pair{"地位", "n"}, Pair{"的", "uj"}, Pair{"消费者", "n"}, Pair{"由此", "c"}, Pair{"将", "d"}, Pair{"陷入", "v"}, Pair{"万劫不复", "i"}, Pair{"的", "uj"}, Pair{"境地", "s"}, Pair{"。", "x"}, Pair{" ", "x"}},
		[]Pair{Pair{"大", "a"}},
		[]Pair{},
		[]Pair{Pair{"他", "r"}, Pair{"说", "v"}, Pair{"的", "uj"}, Pair{"确实", "ad"}, Pair{"在", "p"}, Pair{"理", "n"}},
		[]Pair{Pair{"长春", "ns"}, Pair{"市长", "n"}, Pair{"春节", "t"}, Pair{"讲话", "n"}},
		[]Pair{Pair{"结婚", "v"}, Pair{"的", "uj"}, Pair{"和", "c"}, Pair{"尚未", "d"}, Pair{"结婚", "v"}, Pair{"的", "uj"}},
		[]Pair{Pair{"结合", "v"}, Pair{"成", "n"}, Pair{"分子", "n"}, Pair{"时", "n"}},
		[]Pair{Pair{"旅游", "vn"}, Pair{"和", "c"}, Pair{"服务", "vn"}, Pair{"是", "v"}, Pair{"最好", "a"}, Pair{"的", "uj"}},
		[]Pair{Pair{"这件", "mq"}, Pair{"事情", "n"}, Pair{"的确", "d"}, Pair{"是", "v"}, Pair{"我", "r"}, Pair{"的", "uj"}, Pair{"错", "n"}},
		[]Pair{Pair{"供", "v"}, Pair{"大家", "n"}, Pair{"参考", "v"}, Pair{"指正", "v"}},
		[]Pair{Pair{"哈尔滨", "ns"}, Pair{"政府", "n"}, Pair{"公布", "v"}, Pair{"塌", "v"}, Pair{"桥", "n"}, Pair{"原因", "n"}},
		[]Pair{Pair{"我", "r"}, Pair{"在", "p"}, Pair{"机场", "n"}, Pair{"入口处", "i"}},
		[]Pair{Pair{"邢永臣", "nr"}, Pair{"摄影", "n"}, Pair{"报道", "v"}},
		[]Pair{Pair{"BP", "eng"}, Pair{"神经网络", "n"}, Pair{"如何", "r"}, Pair{"训练", "vn"}, Pair{"才能", "v"}, Pair{"在", "p"}, Pair{"分类", "n"}, Pair{"时", "n"}, Pair{"增加", "v"}, Pair{"区分度", "n"}, Pair{"？", "x"}},
		[]Pair{Pair{"南京市", "ns"}, Pair{"长江大桥", "ns"}},
		[]Pair{Pair{"应", "v"}, Pair{"一些", "m"}, Pair{"使用者", "n"}, Pair{"的", "uj"}, Pair{"建议", "n"}, Pair{"，", "x"}, Pair{"也", "d"}, Pair{"为了", "p"}, Pair{"便于", "v"}, Pair{"利用", "n"}, Pair{"NiuTrans", "eng"}, Pair{"用于", "v"}, Pair{"SMT", "eng"}, Pair{"研究", "vn"}},
		[]Pair{Pair{"长春市", "ns"}, Pair{"长春", "ns"}, Pair{"药店", "n"}},
		[]Pair{Pair{"邓颖超", "nr"}, Pair{"生前", "t"}, Pair{"最", "d"}, Pair{"喜欢", "v"}, Pair{"的", "uj"}, Pair{"衣服", "n"}},
		[]Pair{Pair{"胡锦涛", "nr"}, Pair{"是", "v"}, Pair{"热爱", "a"}, Pair{"世界", "n"}, Pair{"和平", "nz"}, Pair{"的", "uj"}, Pair{"政治局", "n"}, Pair{"常委", "j"}},
		[]Pair{Pair{"程序员", "n"}, Pair{"祝", "v"}, Pair{"海林", "nz"}, Pair{"和", "c"}, Pair{"朱会震", "nr"}, Pair{"是", "v"}, Pair{"在", "p"}, Pair{"孙健", "nr"}, Pair{"的", "uj"}, Pair{"左面", "f"}, Pair{"和", "c"}, Pair{"右面", "f"}, Pair{",", "x"}, Pair{" ", "x"}, Pair{"范凯", "nr"}, Pair{"在", "p"}, Pair{"最", "a"}, Pair{"右面", "f"}, Pair{".", "m"}, Pair{"再往", "d"}, Pair{"左", "f"}, Pair{"是", "v"}, Pair{"李松洪", "nr"}},
		[]Pair{Pair{"一次性", "d"}, Pair{"交", "v"}, Pair{"多少", "m"}, Pair{"钱", "n"}},
		[]Pair{Pair{"两块", "m"}, Pair{"五", "m"}, Pair{"一套", "m"}, Pair{"，", "x"}, Pair{"三块", "m"}, Pair{"八", "m"}, Pair{"一斤", "m"}, Pair{"，", "x"}, Pair{"四块", "m"}, Pair{"七", "m"}, Pair{"一本", "m"}, Pair{"，", "x"}, Pair{"五块", "m"}, Pair{"六", "m"}, Pair{"一条", "m"}},
		[]Pair{Pair{"小", "a"}, Pair{"和尚", "nr"}, Pair{"留", "v"}, Pair{"了", "ul"}, Pair{"一个", "m"}, Pair{"像", "v"}, Pair{"大", "a"}, Pair{"和尚", "nr"}, Pair{"一样", "r"}, Pair{"的", "uj"}, Pair{"和尚头", "nr"}},
		[]Pair{Pair{"我", "r"}, Pair{"是", "v"}, Pair{"中华人民共和国", "ns"}, Pair{"公民", "n"}, Pair{";", "x"}, Pair{"我", "r"}, Pair{"爸爸", "n"}, Pair{"是", "v"}, Pair{"共和党", "nt"}, Pair{"党员", "n"}, Pair{";", "x"}, Pair{" ", "x"}, Pair{"地铁", "n"}, Pair{"和平门", "ns"}, Pair{"站", "v"}},
		[]Pair{Pair{"张晓梅", "nr"}, Pair{"去", "v"}, Pair{"人民", "n"}, Pair{"医院", "n"}, Pair{"做", "v"}, Pair{"了", "ul"}, Pair{"个", "q"}, Pair{"B超", "n"}, Pair{"然后", "c"}, Pair{"去", "v"}, Pair{"买", "v"}, Pair{"了", "ul"}, Pair{"件", "q"}, Pair{"T恤", "n"}},
		[]Pair{Pair{"AT&T", "nz"}, Pair{"是", "v"}, Pair{"一件", "m"}, Pair{"不错", "a"}, Pair{"的", "uj"}, Pair{"公司", "n"}, Pair{"，", "x"}, Pair{"给", "p"}, Pair{"你", "r"}, Pair{"发", "v"}, Pair{"offer", "eng"}, Pair{"了", "ul"}, Pair{"吗", "y"}, Pair{"？", "x"}},
		[]Pair{Pair{"C++", "nz"}, Pair{"和", "c"}, Pair{"c#", "nz"}, Pair{"是", "v"}, Pair{"什么", "r"}, Pair{"关系", "n"}, Pair{"？", "x"}, Pair{"11", "m"}, Pair{"+", "x"}, Pair{"122", "m"}, Pair{"=", "x"}, Pair{"133", "m"}, Pair{"，", "x"}, Pair{"是", "v"}, Pair{"吗", "y"}, Pair{"？", "x"}, Pair{"PI", "eng"}, Pair{"=", "x"}, Pair{"3.14159", "m"}},
		[]Pair{Pair{"你", "r"}, Pair{"认识", "v"}, Pair{"那个", "r"}, Pair{"和", "c"}, Pair{"主席", "n"}, Pair{"握手", "v"}, Pair{"的", "uj"}, Pair{"的哥", "n"}, Pair{"吗", "y"}, Pair{"？", "x"}, Pair{"他", "r"}, Pair{"开", "v"}, Pair{"一辆", "m"}, Pair{"黑色", "n"}, Pair{"的士", "n"}, Pair{"。", "x"}},
		[]Pair{Pair{"枪杆子", "n"}, Pair{"中", "f"}, Pair{"出", "v"}, Pair{"政权", "n"}},
	}
	noHMMCutResult = [][]Pair{
		[]Pair{Pair{"这", "r"}, Pair{"是", "v"}, Pair{"一个", "m"}, Pair{"伸手不见五指", "i"}, Pair{"的", "uj"}, Pair{"黑夜", "n"}, Pair{"。", "x"}, Pair{"我", "r"}, Pair{"叫", "v"}, Pair{"孙悟空", "nr"}, Pair{"，", "x"}, Pair{"我", "r"}, Pair{"爱", "v"}, Pair{"北京", "ns"}, Pair{"，", "x"}, Pair{"我", "r"}, Pair{"爱", "v"}, Pair{"Python", "eng"}, Pair{"和", "c"}, Pair{"C++", "nz"}, Pair{"。", "x"}},
		[]Pair{Pair{"我", "r"}, Pair{"不", "d"}, Pair{"喜欢", "v"}, Pair{"日本", "ns"}, Pair{"和服", "nz"}, Pair{"。", "x"}},
		[]Pair{Pair{"雷猴", "n"}, Pair{"回归", "v"}, Pair{"人间", "n"}, Pair{"。", "x"}},
		[]Pair{Pair{"工信处", "n"}, Pair{"女干事", "n"}, Pair{"每月", "r"}, Pair{"经过", "p"}, Pair{"下属", "v"}, Pair{"科室", "n"}, Pair{"都", "d"}, Pair{"要", "v"}, Pair{"亲口", "n"}, Pair{"交代", "n"}, Pair{"24", "eng"}, Pair{"口", "q"}, Pair{"交换机", "n"}, Pair{"等", "u"}, Pair{"技术性", "n"}, Pair{"器件", "n"}, Pair{"的", "uj"}, Pair{"安装", "v"}, Pair{"工作", "vn"}},
		[]Pair{Pair{"我", "r"}, Pair{"需要", "v"}, Pair{"廉租房", "n"}},
		[]Pair{Pair{"永和", "nz"}, Pair{"服装", "vn"}, Pair{"饰品", "n"}, Pair{"有限公司", "n"}},
		[]Pair{Pair{"我", "r"}, Pair{"爱", "v"}, Pair{"北京", "ns"}, Pair{"天安门", "ns"}},
		[]Pair{Pair{"abc", "eng"}},
		[]Pair{Pair{"隐", "n"}, Pair{"马尔可夫", "nr"}},
		[]Pair{Pair{"雷猴", "n"}, Pair{"是", "v"}, Pair{"个", "q"}, Pair{"好", "a"}, Pair{"网站", "n"}},
		[]Pair{Pair{"“", "x"}, Pair{"Microsoft", "eng"}, Pair{"”", "x"}, Pair{"一", "m"}, Pair{"词", "n"}, Pair{"由", "p"}, Pair{"“", "x"}, Pair{"MICROcomputer", "eng"}, Pair{"（", "x"}, Pair{"微型", "b"}, Pair{"计算机", "n"}, Pair{"）", "x"}, Pair{"”", "x"}, Pair{"和", "c"}, Pair{"“", "x"}, Pair{"SOFTware", "eng"}, Pair{"（", "x"}, Pair{"软件", "n"}, Pair{"）", "x"}, Pair{"”", "x"}, Pair{"两", "m"}, Pair{"部分", "n"}, Pair{"组成", "v"}},
		[]Pair{Pair{"草泥马", "n"}, Pair{"和", "c"}, Pair{"欺", "vn"}, Pair{"实", "n"}, Pair{"马", "n"}, Pair{"是", "v"}, Pair{"今年", "t"}, Pair{"的", "uj"}, Pair{"流行", "v"}, Pair{"词汇", "n"}},
		[]Pair{Pair{"伊", "ns"}, Pair{"藤", "nr"}, Pair{"洋华堂", "n"}, Pair{"总府", "n"}, Pair{"店", "n"}},
		[]Pair{Pair{"中国科学院计算技术研究所", "nt"}},
		[]Pair{Pair{"罗密欧", "nr"}, Pair{"与", "p"}, Pair{"朱丽叶", "nr"}},
		[]Pair{Pair{"我", "r"}, Pair{"购买", "v"}, Pair{"了", "ul"}, Pair{"道具", "n"}, Pair{"和", "c"}, Pair{"服装", "vn"}},
		[]Pair{Pair{"PS", "eng"}, Pair{":", "x"}, Pair{" ", "x"}, Pair{"我", "r"}, Pair{"觉得", "v"}, Pair{"开源", "n"}, Pair{"有", "v"}, Pair{"一个", "m"}, Pair{"好处", "d"}, Pair{"，", "x"}, Pair{"就是", "d"}, Pair{"能够", "v"}, Pair{"敦促", "v"}, Pair{"自己", "r"}, Pair{"不断改进", "l"}, Pair{"，", "x"}, Pair{"避免", "v"}, Pair{"敞", "v"}, Pair{"帚", "ng"}, Pair{"自珍", "b"}},
		[]Pair{Pair{"湖北省", "ns"}, Pair{"石首市", "ns"}},
		[]Pair{Pair{"湖北省", "ns"}, Pair{"十堰市", "ns"}},
		[]Pair{Pair{"总经理", "n"}, Pair{"完成", "v"}, Pair{"了", "ul"}, Pair{"这件", "mq"}, Pair{"事情", "n"}},
		[]Pair{Pair{"电脑", "n"}, Pair{"修好", "v"}, Pair{"了", "ul"}},
		[]Pair{Pair{"做好", "v"}, Pair{"了", "ul"}, Pair{"这件", "mq"}, Pair{"事情", "n"}, Pair{"就", "d"}, Pair{"一了百了", "l"}, Pair{"了", "ul"}},
		[]Pair{Pair{"人们", "n"}, Pair{"审美", "vn"}, Pair{"的", "uj"}, Pair{"观点", "n"}, Pair{"是", "v"}, Pair{"不同", "a"}, Pair{"的", "uj"}},
		[]Pair{Pair{"我们", "r"}, Pair{"买", "v"}, Pair{"了", "ul"}, Pair{"一个", "m"}, Pair{"美的", "nr"}, Pair{"空调", "n"}},
		[]Pair{Pair{"线程", "n"}, Pair{"初始化", "l"}, Pair{"时", "n"}, Pair{"我们", "r"}, Pair{"要", "v"}, Pair{"注意", "v"}},
		[]Pair{Pair{"一个", "m"}, Pair{"分子", "n"}, Pair{"是", "v"}, Pair{"由", "p"}, Pair{"好多", "m"}, Pair{"原子", "n"}, Pair{"组织", "v"}, Pair{"成", "n"}, Pair{"的", "uj"}},
		[]Pair{Pair{"祝", "v"}, Pair{"你", "r"}, Pair{"马到功成", "i"}},
		[]Pair{Pair{"他", "r"}, Pair{"掉", "zg"}, Pair{"进", "v"}, Pair{"了", "ul"}, Pair{"无底洞", "ns"}, Pair{"里", "f"}},
		[]Pair{Pair{"中国", "ns"}, Pair{"的", "uj"}, Pair{"首都", "d"}, Pair{"是", "v"}, Pair{"北京", "ns"}},
		[]Pair{Pair{"孙", "zg"}, Pair{"君", "nz"}, Pair{"意", "n"}},
		[]Pair{Pair{"外交部", "nt"}, Pair{"发言人", "l"}, Pair{"马朝旭", "nr"}},
		[]Pair{Pair{"领导人", "n"}, Pair{"会议", "n"}, Pair{"和", "c"}, Pair{"第四届", "m"}, Pair{"东亚", "ns"}, Pair{"峰会", "n"}},
		[]Pair{Pair{"在", "p"}, Pair{"过去", "t"}, Pair{"的", "uj"}, Pair{"这", "r"}, Pair{"五年", "t"}},
		[]Pair{Pair{"还", "d"}, Pair{"需要", "v"}, Pair{"很", "zg"}, Pair{"长", "a"}, Pair{"的", "uj"}, Pair{"路", "n"}, Pair{"要", "v"}, Pair{"走", "v"}},
		[]Pair{Pair{"60", "eng"}, Pair{"周年", "t"}, Pair{"首都", "d"}, Pair{"阅兵", "v"}},
		[]Pair{Pair{"你好", "l"}, Pair{"人们", "n"}, Pair{"审美", "vn"}, Pair{"的", "uj"}, Pair{"观点", "n"}, Pair{"是", "v"}, Pair{"不同", "a"}, Pair{"的", "uj"}},
		[]Pair{Pair{"买", "v"}, Pair{"水果", "n"}, Pair{"然后", "c"}, Pair{"来", "v"}, Pair{"世博园", "nr"}},
		[]Pair{Pair{"买", "v"}, Pair{"水果", "n"}, Pair{"然后", "c"}, Pair{"去", "v"}, Pair{"世博园", "nr"}},
		[]Pair{Pair{"但是", "c"}, Pair{"后来", "t"}, Pair{"我", "r"}, Pair{"才", "d"}, Pair{"知道", "v"}, Pair{"你", "r"}, Pair{"是", "v"}, Pair{"对", "p"}, Pair{"的", "uj"}},
		[]Pair{Pair{"存在", "v"}, Pair{"即", "v"}, Pair{"合理", "vn"}},
		[]Pair{Pair{"的", "uj"}, Pair{"的", "uj"}, Pair{"的", "uj"}, Pair{"的", "uj"}, Pair{"的", "uj"}, Pair{"在", "p"}, Pair{"的", "uj"}, Pair{"的", "uj"}, Pair{"的", "uj"}, Pair{"的", "uj"}, Pair{"就", "d"}, Pair{"以", "p"}, Pair{"和", "c"}, Pair{"和", "c"}, Pair{"和", "c"}},
		[]Pair{Pair{"I", "eng"}, Pair{" ", "x"}, Pair{"love", "eng"}, Pair{"你", "r"}, Pair{"，", "x"}, Pair{"不以为耻", "i"}, Pair{"，", "x"}, Pair{"反", "zg"}, Pair{"以为", "c"}, Pair{"rong", "eng"}},
		[]Pair{Pair{"因", "p"}},
		[]Pair{},
		[]Pair{Pair{"hello", "eng"}, Pair{"你好", "l"}, Pair{"人们", "n"}, Pair{"审美", "vn"}, Pair{"的", "uj"}, Pair{"观点", "n"}, Pair{"是", "v"}, Pair{"不同", "a"}, Pair{"的", "uj"}},
		[]Pair{Pair{"很", "zg"}, Pair{"好", "a"}, Pair{"但", "c"}, Pair{"主要", "b"}, Pair{"是", "v"}, Pair{"基于", "p"}, Pair{"网页", "n"}, Pair{"形式", "n"}},
		[]Pair{Pair{"hello", "eng"}, Pair{"你好", "l"}, Pair{"人们", "n"}, Pair{"审美", "vn"}, Pair{"的", "uj"}, Pair{"观点", "n"}, Pair{"是", "v"}, Pair{"不同", "a"}, Pair{"的", "uj"}},
		[]Pair{Pair{"为什么", "r"}, Pair{"我", "r"}, Pair{"不能", "v"}, Pair{"拥有", "v"}, Pair{"想要", "v"}, Pair{"的", "uj"}, Pair{"生活", "vn"}},
		[]Pair{Pair{"后来", "t"}, Pair{"我", "r"}, Pair{"才", "d"}},
		[]Pair{Pair{"此次", "r"}, Pair{"来", "v"}, Pair{"中国", "ns"}, Pair{"是", "v"}, Pair{"为了", "p"}},
		[]Pair{Pair{"使用", "v"}, Pair{"了", "ul"}, Pair{"它", "r"}, Pair{"就", "d"}, Pair{"可以", "c"}, Pair{"解决", "v"}, Pair{"一些", "m"}, Pair{"问题", "n"}},
		[]Pair{Pair{",", "x"}, Pair{"使用", "v"}, Pair{"了", "ul"}, Pair{"它", "r"}, Pair{"就", "d"}, Pair{"可以", "c"}, Pair{"解决", "v"}, Pair{"一些", "m"}, Pair{"问题", "n"}},
		[]Pair{Pair{"其实", "d"}, Pair{"使用", "v"}, Pair{"了", "ul"}, Pair{"它", "r"}, Pair{"就", "d"}, Pair{"可以", "c"}, Pair{"解决", "v"}, Pair{"一些", "m"}, Pair{"问题", "n"}},
		[]Pair{Pair{"好人", "n"}, Pair{"使用", "v"}, Pair{"了", "ul"}, Pair{"它", "r"}, Pair{"就", "d"}, Pair{"可以", "c"}, Pair{"解决", "v"}, Pair{"一些", "m"}, Pair{"问题", "n"}},
		[]Pair{Pair{"是因为", "c"}, Pair{"和", "c"}, Pair{"国家", "n"}},
		[]Pair{Pair{"老年", "t"}, Pair{"搜索", "v"}, Pair{"还", "d"}, Pair{"支持", "v"}},
		[]Pair{Pair{"干脆", "d"}, Pair{"就", "d"}, Pair{"把", "p"}, Pair{"那", "r"}, Pair{"部", "n"}, Pair{"蒙", "v"}, Pair{"人", "n"}, Pair{"的", "uj"}, Pair{"闲", "n"}, Pair{"法", "j"}, Pair{"给", "p"}, Pair{"废", "v"}, Pair{"了", "ul"}, Pair{"拉倒", "v"}, Pair{"！", "x"}, Pair{"RT", "eng"}, Pair{" ", "x"}, Pair{"@", "x"}, Pair{"laoshipukong", "eng"}, Pair{" ", "x"}, Pair{":", "x"}, Pair{" ", "x"}, Pair{"27", "eng"}, Pair{"日", "m"}, Pair{"，", "x"}, Pair{"全国人大常委会", "nt"}, Pair{"第三次", "m"}, Pair{"审议", "v"}, Pair{"侵权", "v"}, Pair{"责任法", "n"}, Pair{"草案", "n"}, Pair{"，", "x"}, Pair{"删除", "v"}, Pair{"了", "ul"}, Pair{"有关", "vn"}, Pair{"医疗", "n"}, Pair{"损害", "v"}, Pair{"责任", "n"}, Pair{"“", "x"}, Pair{"举证", "v"}, Pair{"倒置", "v"}, Pair{"”", "x"}, Pair{"的", "uj"}, Pair{"规定", "n"}, Pair{"。", "x"}, Pair{"在", "p"}, Pair{"医患", "n"}, Pair{"纠纷", "n"}, Pair{"中", "f"}, Pair{"本", "r"}, Pair{"已", "d"}, Pair{"处于", "v"}, Pair{"弱势", "n"}, Pair{"地位", "n"}, Pair{"的", "uj"}, Pair{"消费者", "n"}, Pair{"由此", "c"}, Pair{"将", "d"}, Pair{"陷入", "v"}, Pair{"万劫不复", "i"}, Pair{"的", "uj"}, Pair{"境地", "s"}, Pair{"。", "x"}, Pair{" ", "x"}},
		[]Pair{Pair{"大", "a"}},
		[]Pair{},
		[]Pair{Pair{"他", "r"}, Pair{"说", "v"}, Pair{"的", "uj"}, Pair{"确实", "ad"}, Pair{"在", "p"}, Pair{"理", "n"}},
		[]Pair{Pair{"长春", "ns"}, Pair{"市长", "n"}, Pair{"春节", "t"}, Pair{"讲话", "n"}},
		[]Pair{Pair{"结婚", "v"}, Pair{"的", "uj"}, Pair{"和", "c"}, Pair{"尚未", "d"}, Pair{"结婚", "v"}, Pair{"的", "uj"}},
		[]Pair{Pair{"结合", "v"}, Pair{"成", "n"}, Pair{"分子", "n"}, Pair{"时", "n"}},
		[]Pair{Pair{"旅游", "vn"}, Pair{"和", "c"}, Pair{"服务", "vn"}, Pair{"是", "v"}, Pair{"最好", "a"}, Pair{"的", "uj"}},
		[]Pair{Pair{"这件", "mq"}, Pair{"事情", "n"}, Pair{"的确", "d"}, Pair{"是", "v"}, Pair{"我", "r"}, Pair{"的", "uj"}, Pair{"错", "v"}},
		[]Pair{Pair{"供", "v"}, Pair{"大家", "n"}, Pair{"参考", "v"}, Pair{"指正", "v"}},
		[]Pair{Pair{"哈尔滨", "ns"}, Pair{"政府", "n"}, Pair{"公布", "v"}, Pair{"塌", "v"}, Pair{"桥", "n"}, Pair{"原因", "n"}},
		[]Pair{Pair{"我", "r"}, Pair{"在", "p"}, Pair{"机场", "n"}, Pair{"入口处", "i"}},
		[]Pair{Pair{"邢", "nr"}, Pair{"永", "ns"}, Pair{"臣", "n"}, Pair{"摄影", "n"}, Pair{"报道", "v"}},
		[]Pair{Pair{"BP", "eng"}, Pair{"神经网络", "n"}, Pair{"如何", "r"}, Pair{"训练", "vn"}, Pair{"才能", "v"}, Pair{"在", "p"}, Pair{"分类", "n"}, Pair{"时", "n"}, Pair{"增加", "v"}, Pair{"区分度", "n"}, Pair{"？", "x"}},
		[]Pair{Pair{"南京市", "ns"}, Pair{"长江大桥", "ns"}},
		[]Pair{Pair{"应", "v"}, Pair{"一些", "m"}, Pair{"使用者", "n"}, Pair{"的", "uj"}, Pair{"建议", "n"}, Pair{"，", "x"}, Pair{"也", "d"}, Pair{"为了", "p"}, Pair{"便于", "v"}, Pair{"利用", "n"}, Pair{"NiuTrans", "eng"}, Pair{"用于", "v"}, Pair{"SMT", "eng"}, Pair{"研究", "vn"}},
		[]Pair{Pair{"长春市", "ns"}, Pair{"长春", "ns"}, Pair{"药店", "n"}},
		[]Pair{Pair{"邓颖超", "nr"}, Pair{"生前", "t"}, Pair{"最", "d"}, Pair{"喜欢", "v"}, Pair{"的", "uj"}, Pair{"衣服", "n"}},
		[]Pair{Pair{"胡锦涛", "nr"}, Pair{"是", "v"}, Pair{"热爱", "a"}, Pair{"世界", "n"}, Pair{"和平", "nz"}, Pair{"的", "uj"}, Pair{"政治局", "n"}, Pair{"常委", "j"}},
		[]Pair{Pair{"程序员", "n"}, Pair{"祝", "v"}, Pair{"海林", "nz"}, Pair{"和", "c"}, Pair{"朱", "nr"}, Pair{"会", "v"}, Pair{"震", "v"}, Pair{"是", "v"}, Pair{"在", "p"}, Pair{"孙", "zg"}, Pair{"健", "a"}, Pair{"的", "uj"}, Pair{"左面", "f"}, Pair{"和", "c"}, Pair{"右面", "f"}, Pair{",", "x"}, Pair{" ", "x"}, Pair{"范", "nr"}, Pair{"凯", "nr"}, Pair{"在", "p"}, Pair{"最", "d"}, Pair{"右面", "f"}, Pair{".", "x"}, Pair{"再", "d"}, Pair{"往", "zg"}, Pair{"左", "m"}, Pair{"是", "v"}, Pair{"李", "nr"}, Pair{"松", "v"}, Pair{"洪", "nr"}},
		[]Pair{Pair{"一次性", "d"}, Pair{"交", "v"}, Pair{"多少", "m"}, Pair{"钱", "n"}},
		[]Pair{Pair{"两块", "m"}, Pair{"五", "m"}, Pair{"一套", "m"}, Pair{"，", "x"}, Pair{"三块", "m"}, Pair{"八", "m"}, Pair{"一斤", "m"}, Pair{"，", "x"}, Pair{"四块", "m"}, Pair{"七", "m"}, Pair{"一本", "m"}, Pair{"，", "x"}, Pair{"五块", "m"}, Pair{"六", "m"}, Pair{"一条", "m"}},
		[]Pair{Pair{"小", "a"}, Pair{"和尚", "nr"}, Pair{"留", "v"}, Pair{"了", "ul"}, Pair{"一个", "m"}, Pair{"像", "v"}, Pair{"大", "a"}, Pair{"和尚", "nr"}, Pair{"一样", "r"}, Pair{"的", "uj"}, Pair{"和尚头", "nr"}},
		[]Pair{Pair{"我", "r"}, Pair{"是", "v"}, Pair{"中华人民共和国", "ns"}, Pair{"公民", "n"}, Pair{";", "x"}, Pair{"我", "r"}, Pair{"爸爸", "n"}, Pair{"是", "v"}, Pair{"共和党", "nt"}, Pair{"党员", "n"}, Pair{";", "x"}, Pair{" ", "x"}, Pair{"地铁", "n"}, Pair{"和平门", "ns"}, Pair{"站", "v"}},
		[]Pair{Pair{"张晓梅", "nr"}, Pair{"去", "v"}, Pair{"人民", "n"}, Pair{"医院", "n"}, Pair{"做", "v"}, Pair{"了", "ul"}, Pair{"个", "q"}, Pair{"B超", "n"}, Pair{"然后", "c"}, Pair{"去", "v"}, Pair{"买", "v"}, Pair{"了", "ul"}, Pair{"件", "zg"}, Pair{"T恤", "n"}},
		[]Pair{Pair{"AT&T", "nz"}, Pair{"是", "v"}, Pair{"一件", "m"}, Pair{"不错", "a"}, Pair{"的", "uj"}, Pair{"公司", "n"}, Pair{"，", "x"}, Pair{"给", "p"}, Pair{"你", "r"}, Pair{"发", "v"}, Pair{"offer", "eng"}, Pair{"了", "ul"}, Pair{"吗", "y"}, Pair{"？", "x"}},
		[]Pair{Pair{"C++", "nz"}, Pair{"和", "c"}, Pair{"c#", "nz"}, Pair{"是", "v"}, Pair{"什么", "r"}, Pair{"关系", "n"}, Pair{"？", "x"}, Pair{"11", "eng"}, Pair{"+", "x"}, Pair{"122", "eng"}, Pair{"=", "x"}, Pair{"133", "eng"}, Pair{"，", "x"}, Pair{"是", "v"}, Pair{"吗", "y"}, Pair{"？", "x"}, Pair{"PI", "eng"}, Pair{"=", "x"}, Pair{"3", "eng"}, Pair{".", "x"}, Pair{"14159", "eng"}},
		[]Pair{Pair{"你", "r"}, Pair{"认识", "v"}, Pair{"那个", "r"}, Pair{"和", "c"}, Pair{"主席", "n"}, Pair{"握手", "v"}, Pair{"的", "uj"}, Pair{"的哥", "n"}, Pair{"吗", "y"}, Pair{"？", "x"}, Pair{"他", "r"}, Pair{"开", "v"}, Pair{"一辆", "m"}, Pair{"黑色", "n"}, Pair{"的士", "n"}, Pair{"。", "x"}},
		[]Pair{Pair{"枪杆子", "n"}, Pair{"中", "f"}, Pair{"出", "v"}, Pair{"政权", "n"}},
	}
)

func chanToArray(ch chan Pair) []Pair {
	result := make([]Pair, 0)
	for word := range ch {
		result = append(result, word)
	}
	return result
}

func TestCut(t *testing.T) {
	p, err := Open("../dict.txt")
	if err != nil {
		t.Fatal(err)
	}
	for index, content := range test_contents {
		result := chanToArray(p.Cut(content, true))
		if len(defaultCutResult[index]) != len(result) {
			t.Fatal(content)
		}
		for i, _ := range result {
			if result[i] != defaultCutResult[index][i] {
				t.Fatalf("expect %s, got %s", defaultCutResult[index][i], result[i])
			}
		}
		result = chanToArray(p.Cut(content, false))
		if len(noHMMCutResult[index]) != len(result) {
			t.Fatal(content)
		}
		for i, _ := range result {
			if result[i] != noHMMCutResult[index][i] {
				t.Fatal(content)
			}
		}

	}
}

func TestBug132(t *testing.T) {
	/*
		https://github.com/fxsjy/jieba/issues/132
	*/
	p, _ := Open("../dict.txt")
	sentence := "又跛又啞"
	cutResult := []Pair{
		Pair{"又", "d"},
		Pair{"跛", "a"},
		Pair{"又", "d"},
		Pair{"啞", "v"},
	}
	result := chanToArray(p.Cut(sentence, true))
	if len(cutResult) != len(result) {
		t.Fatal(result)
	}
	for i, _ := range result {
		if result[i] != cutResult[i] {
			t.Fatal(result[i])
		}
	}
}

func TestBug137(t *testing.T) {
	/*
		https://github.com/fxsjy/jieba/issues/137
	*/
	p, _ := Open("../dict.txt")
	sentence := "前港督衛奕信在八八年十月宣布成立中央政策研究組"
	cutResult := []Pair{
		Pair{"前", "f"},
		Pair{"港督", "n"},
		Pair{"衛奕", "z"},
		Pair{"信", "n"},
		Pair{"在", "p"},
		Pair{"八八年", "m"},
		Pair{"十月", "t"},
		Pair{"宣布", "v"},
		Pair{"成立", "v"},
		Pair{"中央", "n"},
		Pair{"政策", "n"},
		Pair{"研究", "vn"},
		Pair{"組", "x"},
	}
	result := chanToArray(p.Cut(sentence, true))
	if len(cutResult) != len(result) {
		t.Fatal(result)
	}
	for i, _ := range result {
		if result[i] != cutResult[i] {
			t.Fatal(result[i])
		}
	}
}

func TestUserDict(t *testing.T) {
	p, _ := Open("../dict.txt")
	p.LoadUserDict("../userdict.txt")
	sentence := "李小福是创新办主任也是云计算方面的专家; 什么是八一双鹿例如我输入一个带“韩玉赏鉴”的标题，在自定义词库中也增加了此词为N类型"

	cutResult := []Pair{
		Pair{"李小福", "nr"},
		Pair{"是", "v"},
		Pair{"创新办", "i"},
		Pair{"主任", "b"},
		Pair{"也", "d"},
		Pair{"是", "v"},
		Pair{"云计算", "x"},
		Pair{"方面", "n"},
		Pair{"的", "uj"},
		Pair{"专家", "n"},
		Pair{";", "x"},
		Pair{" ", "x"},
		Pair{"什么", "r"},
		Pair{"是", "v"},
		Pair{"八一双鹿", "nz"},
		Pair{"例如", "v"},
		Pair{"我", "r"},
		Pair{"输入", "v"},
		Pair{"一个", "m"},
		Pair{"带", "v"},
		Pair{"“", "x"},
		Pair{"韩玉赏鉴", "nz"},
		Pair{"”", "x"},
		Pair{"的", "uj"},
		Pair{"标题", "n"},
		Pair{"，", "x"},
		Pair{"在", "p"},
		Pair{"自定义词", "n"},
		Pair{"库中", "nrt"},
		Pair{"也", "d"},
		Pair{"增加", "v"},
		Pair{"了", "ul"},
		Pair{"此", "r"},
		Pair{"词", "n"},
		Pair{"为", "p"},
		Pair{"N", "eng"},
		Pair{"类型", "n"}}

	result := chanToArray(p.Cut(sentence, true))
	if len(cutResult) != len(result) {
		t.Fatal(result)
	}
	for i, _ := range result {
		if result[i] != cutResult[i] {
			t.Fatal(result[i])
		}
	}
}

func BenchmarkCutNoHMM(b *testing.B) {
	p, _ := Open("dict.txt")
	sentence := "工信处女干事每月经过下属科室都要亲口交代24口交换机等技术性器件的安装工作"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		chanToArray(p.Cut(sentence, false))
	}
}

func BenchmarkCut(b *testing.B) {
	p, _ := Open("dict.txt")
	sentence := "工信处女干事每月经过下属科室都要亲口交代24口交换机等技术性器件的安装工作"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		chanToArray(p.Cut(sentence, true))
	}
}
