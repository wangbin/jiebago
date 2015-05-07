package posseg

import (
	"testing"
)

var (
	seg          Segmenter
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

	defaultCutResult = [][]Segment{[]Segment{Segment{"这", "r"}, Segment{"是", "v"}, Segment{"一个", "m"}, Segment{"伸手不见五指", "i"}, Segment{"的", "uj"}, Segment{"黑夜", "n"}, Segment{"。", "x"}, Segment{"我", "r"}, Segment{"叫", "v"}, Segment{"孙悟空", "nr"}, Segment{"，", "x"}, Segment{"我", "r"}, Segment{"爱", "v"}, Segment{"北京", "ns"}, Segment{"，", "x"}, Segment{"我", "r"}, Segment{"爱", "v"}, Segment{"Python", "eng"}, Segment{"和", "c"}, Segment{"C++", "nz"}, Segment{"。", "x"}},
		[]Segment{Segment{"我", "r"}, Segment{"不", "d"}, Segment{"喜欢", "v"}, Segment{"日本", "ns"}, Segment{"和服", "nz"}, Segment{"。", "x"}},
		[]Segment{Segment{"雷猴", "n"}, Segment{"回归", "v"}, Segment{"人间", "n"}, Segment{"。", "x"}},
		[]Segment{Segment{"工信处", "n"}, Segment{"女干事", "n"}, Segment{"每月", "r"}, Segment{"经过", "p"}, Segment{"下属", "v"}, Segment{"科室", "n"}, Segment{"都", "d"}, Segment{"要", "v"}, Segment{"亲口", "n"}, Segment{"交代", "n"}, Segment{"24", "m"}, Segment{"口", "n"}, Segment{"交换机", "n"}, Segment{"等", "u"}, Segment{"技术性", "n"}, Segment{"器件", "n"}, Segment{"的", "uj"}, Segment{"安装", "v"}, Segment{"工作", "vn"}},
		[]Segment{Segment{"我", "r"}, Segment{"需要", "v"}, Segment{"廉租房", "n"}},
		[]Segment{Segment{"永和", "nz"}, Segment{"服装", "vn"}, Segment{"饰品", "n"}, Segment{"有限公司", "n"}},
		[]Segment{Segment{"我", "r"}, Segment{"爱", "v"}, Segment{"北京", "ns"}, Segment{"天安门", "ns"}},
		[]Segment{Segment{"abc", "eng"}},
		[]Segment{Segment{"隐", "n"}, Segment{"马尔可夫", "nr"}},
		[]Segment{Segment{"雷猴", "n"}, Segment{"是", "v"}, Segment{"个", "q"}, Segment{"好", "a"}, Segment{"网站", "n"}},
		[]Segment{Segment{"“", "x"}, Segment{"Microsoft", "eng"}, Segment{"”", "x"}, Segment{"一", "m"}, Segment{"词", "n"}, Segment{"由", "p"}, Segment{"“", "x"}, Segment{"MICROcomputer", "eng"}, Segment{"（", "x"}, Segment{"微型", "b"}, Segment{"计算机", "n"}, Segment{"）", "x"}, Segment{"”", "x"}, Segment{"和", "c"}, Segment{"“", "x"}, Segment{"SOFTware", "eng"}, Segment{"（", "x"}, Segment{"软件", "n"}, Segment{"）", "x"}, Segment{"”", "x"}, Segment{"两", "m"}, Segment{"部分", "n"}, Segment{"组成", "v"}},
		[]Segment{Segment{"草泥马", "n"}, Segment{"和", "c"}, Segment{"欺实", "v"}, Segment{"马", "n"}, Segment{"是", "v"}, Segment{"今年", "t"}, Segment{"的", "uj"}, Segment{"流行", "v"}, Segment{"词汇", "n"}},
		[]Segment{Segment{"伊藤", "nr"}, Segment{"洋华堂", "n"}, Segment{"总府", "n"}, Segment{"店", "n"}},
		[]Segment{Segment{"中国科学院计算技术研究所", "nt"}},
		[]Segment{Segment{"罗密欧", "nr"}, Segment{"与", "p"}, Segment{"朱丽叶", "nr"}},
		[]Segment{Segment{"我", "r"}, Segment{"购买", "v"}, Segment{"了", "ul"}, Segment{"道具", "n"}, Segment{"和", "c"}, Segment{"服装", "vn"}},
		[]Segment{Segment{"PS", "eng"}, Segment{":", "x"}, Segment{" ", "x"}, Segment{"我", "r"}, Segment{"觉得", "v"}, Segment{"开源", "n"}, Segment{"有", "v"}, Segment{"一个", "m"}, Segment{"好处", "d"}, Segment{"，", "x"}, Segment{"就是", "d"}, Segment{"能够", "v"}, Segment{"敦促", "v"}, Segment{"自己", "r"}, Segment{"不断改进", "l"}, Segment{"，", "x"}, Segment{"避免", "v"}, Segment{"敞", "v"}, Segment{"帚", "ng"}, Segment{"自珍", "b"}},
		[]Segment{Segment{"湖北省", "ns"}, Segment{"石首市", "ns"}},
		[]Segment{Segment{"湖北省", "ns"}, Segment{"十堰市", "ns"}},
		[]Segment{Segment{"总经理", "n"}, Segment{"完成", "v"}, Segment{"了", "ul"}, Segment{"这件", "mq"}, Segment{"事情", "n"}},
		[]Segment{Segment{"电脑", "n"}, Segment{"修好", "v"}, Segment{"了", "ul"}},
		[]Segment{Segment{"做好", "v"}, Segment{"了", "ul"}, Segment{"这件", "mq"}, Segment{"事情", "n"}, Segment{"就", "d"}, Segment{"一了百了", "l"}, Segment{"了", "ul"}},
		[]Segment{Segment{"人们", "n"}, Segment{"审美", "vn"}, Segment{"的", "uj"}, Segment{"观点", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"我们", "r"}, Segment{"买", "v"}, Segment{"了", "ul"}, Segment{"一个", "m"}, Segment{"美的", "nr"}, Segment{"空调", "n"}},
		[]Segment{Segment{"线程", "n"}, Segment{"初始化", "l"}, Segment{"时", "n"}, Segment{"我们", "r"}, Segment{"要", "v"}, Segment{"注意", "v"}},
		[]Segment{Segment{"一个", "m"}, Segment{"分子", "n"}, Segment{"是", "v"}, Segment{"由", "p"}, Segment{"好多", "m"}, Segment{"原子", "n"}, Segment{"组织", "v"}, Segment{"成", "v"}, Segment{"的", "uj"}},
		[]Segment{Segment{"祝", "v"}, Segment{"你", "r"}, Segment{"马到功成", "i"}},
		[]Segment{Segment{"他", "r"}, Segment{"掉", "v"}, Segment{"进", "v"}, Segment{"了", "ul"}, Segment{"无底洞", "ns"}, Segment{"里", "f"}},
		[]Segment{Segment{"中国", "ns"}, Segment{"的", "uj"}, Segment{"首都", "d"}, Segment{"是", "v"}, Segment{"北京", "ns"}},
		[]Segment{Segment{"孙君意", "nr"}},
		[]Segment{Segment{"外交部", "nt"}, Segment{"发言人", "l"}, Segment{"马朝旭", "nr"}},
		[]Segment{Segment{"领导人", "n"}, Segment{"会议", "n"}, Segment{"和", "c"}, Segment{"第四届", "m"}, Segment{"东亚", "ns"}, Segment{"峰会", "n"}},
		[]Segment{Segment{"在", "p"}, Segment{"过去", "t"}, Segment{"的", "uj"}, Segment{"这", "r"}, Segment{"五年", "t"}},
		[]Segment{Segment{"还", "d"}, Segment{"需要", "v"}, Segment{"很", "d"}, Segment{"长", "a"}, Segment{"的", "uj"}, Segment{"路", "n"}, Segment{"要", "v"}, Segment{"走", "v"}},
		[]Segment{Segment{"60", "m"}, Segment{"周年", "t"}, Segment{"首都", "d"}, Segment{"阅兵", "v"}},
		[]Segment{Segment{"你好", "l"}, Segment{"人们", "n"}, Segment{"审美", "vn"}, Segment{"的", "uj"}, Segment{"观点", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"买", "v"}, Segment{"水果", "n"}, Segment{"然后", "c"}, Segment{"来", "v"}, Segment{"世博园", "nr"}},
		[]Segment{Segment{"买", "v"}, Segment{"水果", "n"}, Segment{"然后", "c"}, Segment{"去", "v"}, Segment{"世博园", "nr"}},
		[]Segment{Segment{"但是", "c"}, Segment{"后来", "t"}, Segment{"我", "r"}, Segment{"才", "d"}, Segment{"知道", "v"}, Segment{"你", "r"}, Segment{"是", "v"}, Segment{"对", "p"}, Segment{"的", "uj"}},
		[]Segment{Segment{"存在", "v"}, Segment{"即", "v"}, Segment{"合理", "vn"}},
		[]Segment{Segment{"的的", "u"}, Segment{"的的", "u"}, Segment{"的", "uj"}, Segment{"在的", "u"}, Segment{"的的", "u"}, Segment{"的", "uj"}, Segment{"就", "d"}, Segment{"以", "p"}, Segment{"和和", "nz"}, Segment{"和", "c"}},
		[]Segment{Segment{"I", "x"}, Segment{" ", "x"}, Segment{"love", "eng"}, Segment{"你", "r"}, Segment{"，", "x"}, Segment{"不以为耻", "i"}, Segment{"，", "x"}, Segment{"反", "zg"}, Segment{"以为", "c"}, Segment{"rong", "eng"}},
		[]Segment{Segment{"因", "p"}},
		[]Segment{},
		[]Segment{Segment{"hello", "eng"}, Segment{"你好", "l"}, Segment{"人们", "n"}, Segment{"审美", "vn"}, Segment{"的", "uj"}, Segment{"观点", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"很好", "a"}, Segment{"但", "c"}, Segment{"主要", "b"}, Segment{"是", "v"}, Segment{"基于", "p"}, Segment{"网页", "n"}, Segment{"形式", "n"}},
		[]Segment{Segment{"hello", "eng"}, Segment{"你好", "l"}, Segment{"人们", "n"}, Segment{"审美", "vn"}, Segment{"的", "uj"}, Segment{"观点", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"为什么", "r"}, Segment{"我", "r"}, Segment{"不能", "v"}, Segment{"拥有", "v"}, Segment{"想要", "v"}, Segment{"的", "uj"}, Segment{"生活", "vn"}},
		[]Segment{Segment{"后来", "t"}, Segment{"我", "r"}, Segment{"才", "d"}},
		[]Segment{Segment{"此次", "r"}, Segment{"来", "v"}, Segment{"中国", "ns"}, Segment{"是", "v"}, Segment{"为了", "p"}},
		[]Segment{Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解决", "v"}, Segment{"一些", "m"}, Segment{"问题", "n"}},
		[]Segment{Segment{",", "x"}, Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解决", "v"}, Segment{"一些", "m"}, Segment{"问题", "n"}},
		[]Segment{Segment{"其实", "d"}, Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解决", "v"}, Segment{"一些", "m"}, Segment{"问题", "n"}},
		[]Segment{Segment{"好人", "n"}, Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解决", "v"}, Segment{"一些", "m"}, Segment{"问题", "n"}},
		[]Segment{Segment{"是因为", "c"}, Segment{"和", "c"}, Segment{"国家", "n"}},
		[]Segment{Segment{"老年", "t"}, Segment{"搜索", "v"}, Segment{"还", "d"}, Segment{"支持", "v"}},
		[]Segment{Segment{"干脆", "d"}, Segment{"就", "d"}, Segment{"把", "p"}, Segment{"那部", "r"}, Segment{"蒙人", "n"}, Segment{"的", "uj"}, Segment{"闲法", "n"}, Segment{"给", "p"}, Segment{"废", "v"}, Segment{"了", "ul"}, Segment{"拉倒", "v"}, Segment{"！", "x"}, Segment{"RT", "eng"}, Segment{" ", "x"}, Segment{"@", "x"}, Segment{"laoshipukong", "eng"}, Segment{" ", "x"}, Segment{":", "x"}, Segment{" ", "x"}, Segment{"27", "m"}, Segment{"日", "m"}, Segment{"，", "x"}, Segment{"全国人大常委会", "nt"}, Segment{"第三次", "m"}, Segment{"审议", "v"}, Segment{"侵权", "v"}, Segment{"责任法", "n"}, Segment{"草案", "n"}, Segment{"，", "x"}, Segment{"删除", "v"}, Segment{"了", "ul"}, Segment{"有关", "vn"}, Segment{"医疗", "n"}, Segment{"损害", "v"}, Segment{"责任", "n"}, Segment{"“", "x"}, Segment{"举证", "v"}, Segment{"倒置", "v"}, Segment{"”", "x"}, Segment{"的", "uj"}, Segment{"规定", "n"}, Segment{"。", "x"}, Segment{"在", "p"}, Segment{"医患", "n"}, Segment{"纠纷", "n"}, Segment{"中本", "ns"}, Segment{"已", "d"}, Segment{"处于", "v"}, Segment{"弱势", "n"}, Segment{"地位", "n"}, Segment{"的", "uj"}, Segment{"消费者", "n"}, Segment{"由此", "c"}, Segment{"将", "d"}, Segment{"陷入", "v"}, Segment{"万劫不复", "i"}, Segment{"的", "uj"}, Segment{"境地", "s"}, Segment{"。", "x"}, Segment{" ", "x"}},
		[]Segment{Segment{"大", "a"}},
		[]Segment{},
		[]Segment{Segment{"他", "r"}, Segment{"说", "v"}, Segment{"的", "uj"}, Segment{"确实", "ad"}, Segment{"在", "p"}, Segment{"理", "n"}},
		[]Segment{Segment{"长春", "ns"}, Segment{"市长", "n"}, Segment{"春节", "t"}, Segment{"讲话", "n"}},
		[]Segment{Segment{"结婚", "v"}, Segment{"的", "uj"}, Segment{"和", "c"}, Segment{"尚未", "d"}, Segment{"结婚", "v"}, Segment{"的", "uj"}},
		[]Segment{Segment{"结合", "v"}, Segment{"成", "n"}, Segment{"分子", "n"}, Segment{"时", "n"}},
		[]Segment{Segment{"旅游", "vn"}, Segment{"和", "c"}, Segment{"服务", "vn"}, Segment{"是", "v"}, Segment{"最好", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"这件", "mq"}, Segment{"事情", "n"}, Segment{"的确", "d"}, Segment{"是", "v"}, Segment{"我", "r"}, Segment{"的", "uj"}, Segment{"错", "n"}},
		[]Segment{Segment{"供", "v"}, Segment{"大家", "n"}, Segment{"参考", "v"}, Segment{"指正", "v"}},
		[]Segment{Segment{"哈尔滨", "ns"}, Segment{"政府", "n"}, Segment{"公布", "v"}, Segment{"塌", "v"}, Segment{"桥", "n"}, Segment{"原因", "n"}},
		[]Segment{Segment{"我", "r"}, Segment{"在", "p"}, Segment{"机场", "n"}, Segment{"入口处", "i"}},
		[]Segment{Segment{"邢永臣", "nr"}, Segment{"摄影", "n"}, Segment{"报道", "v"}},
		[]Segment{Segment{"BP", "eng"}, Segment{"神经网络", "n"}, Segment{"如何", "r"}, Segment{"训练", "vn"}, Segment{"才能", "v"}, Segment{"在", "p"}, Segment{"分类", "n"}, Segment{"时", "n"}, Segment{"增加", "v"}, Segment{"区分度", "n"}, Segment{"？", "x"}},
		[]Segment{Segment{"南京市", "ns"}, Segment{"长江大桥", "ns"}},
		[]Segment{Segment{"应", "v"}, Segment{"一些", "m"}, Segment{"使用者", "n"}, Segment{"的", "uj"}, Segment{"建议", "n"}, Segment{"，", "x"}, Segment{"也", "d"}, Segment{"为了", "p"}, Segment{"便于", "v"}, Segment{"利用", "n"}, Segment{"NiuTrans", "eng"}, Segment{"用于", "v"}, Segment{"SMT", "eng"}, Segment{"研究", "vn"}},
		[]Segment{Segment{"长春市", "ns"}, Segment{"长春", "ns"}, Segment{"药店", "n"}},
		[]Segment{Segment{"邓颖超", "nr"}, Segment{"生前", "t"}, Segment{"最", "d"}, Segment{"喜欢", "v"}, Segment{"的", "uj"}, Segment{"衣服", "n"}},
		[]Segment{Segment{"胡锦涛", "nr"}, Segment{"是", "v"}, Segment{"热爱", "a"}, Segment{"世界", "n"}, Segment{"和平", "nz"}, Segment{"的", "uj"}, Segment{"政治局", "n"}, Segment{"常委", "j"}},
		[]Segment{Segment{"程序员", "n"}, Segment{"祝", "v"}, Segment{"海林", "nz"}, Segment{"和", "c"}, Segment{"朱会震", "nr"}, Segment{"是", "v"}, Segment{"在", "p"}, Segment{"孙健", "nr"}, Segment{"的", "uj"}, Segment{"左面", "f"}, Segment{"和", "c"}, Segment{"右面", "f"}, Segment{",", "x"}, Segment{" ", "x"}, Segment{"范凯", "nr"}, Segment{"在", "p"}, Segment{"最", "a"}, Segment{"右面", "f"}, Segment{".", "m"}, Segment{"再往", "d"}, Segment{"左", "f"}, Segment{"是", "v"}, Segment{"李松洪", "nr"}},
		[]Segment{Segment{"一次性", "d"}, Segment{"交", "v"}, Segment{"多少", "m"}, Segment{"钱", "n"}},
		[]Segment{Segment{"两块", "m"}, Segment{"五", "m"}, Segment{"一套", "m"}, Segment{"，", "x"}, Segment{"三块", "m"}, Segment{"八", "m"}, Segment{"一斤", "m"}, Segment{"，", "x"}, Segment{"四块", "m"}, Segment{"七", "m"}, Segment{"一本", "m"}, Segment{"，", "x"}, Segment{"五块", "m"}, Segment{"六", "m"}, Segment{"一条", "m"}},
		[]Segment{Segment{"小", "a"}, Segment{"和尚", "nr"}, Segment{"留", "v"}, Segment{"了", "ul"}, Segment{"一个", "m"}, Segment{"像", "v"}, Segment{"大", "a"}, Segment{"和尚", "nr"}, Segment{"一样", "r"}, Segment{"的", "uj"}, Segment{"和尚头", "nr"}},
		[]Segment{Segment{"我", "r"}, Segment{"是", "v"}, Segment{"中华人民共和国", "ns"}, Segment{"公民", "n"}, Segment{";", "x"}, Segment{"我", "r"}, Segment{"爸爸", "n"}, Segment{"是", "v"}, Segment{"共和党", "nt"}, Segment{"党员", "n"}, Segment{";", "x"}, Segment{" ", "x"}, Segment{"地铁", "n"}, Segment{"和平门", "ns"}, Segment{"站", "v"}},
		[]Segment{Segment{"张晓梅", "nr"}, Segment{"去", "v"}, Segment{"人民", "n"}, Segment{"医院", "n"}, Segment{"做", "v"}, Segment{"了", "ul"}, Segment{"个", "q"}, Segment{"B超", "n"}, Segment{"然后", "c"}, Segment{"去", "v"}, Segment{"买", "v"}, Segment{"了", "ul"}, Segment{"件", "q"}, Segment{"T恤", "n"}},
		[]Segment{Segment{"AT&T", "nz"}, Segment{"是", "v"}, Segment{"一件", "m"}, Segment{"不错", "a"}, Segment{"的", "uj"}, Segment{"公司", "n"}, Segment{"，", "x"}, Segment{"给", "p"}, Segment{"你", "r"}, Segment{"发", "v"}, Segment{"offer", "eng"}, Segment{"了", "ul"}, Segment{"吗", "y"}, Segment{"？", "x"}},
		[]Segment{Segment{"C++", "nz"}, Segment{"和", "c"}, Segment{"c#", "nz"}, Segment{"是", "v"}, Segment{"什么", "r"}, Segment{"关系", "n"}, Segment{"？", "x"}, Segment{"11", "m"}, Segment{"+", "x"}, Segment{"122", "m"}, Segment{"=", "x"}, Segment{"133", "m"}, Segment{"，", "x"}, Segment{"是", "v"}, Segment{"吗", "y"}, Segment{"？", "x"}, Segment{"PI", "eng"}, Segment{"=", "x"}, Segment{"3.14159", "m"}},
		[]Segment{Segment{"你", "r"}, Segment{"认识", "v"}, Segment{"那个", "r"}, Segment{"和", "c"}, Segment{"主席", "n"}, Segment{"握手", "v"}, Segment{"的", "uj"}, Segment{"的哥", "n"}, Segment{"吗", "y"}, Segment{"？", "x"}, Segment{"他", "r"}, Segment{"开", "v"}, Segment{"一辆", "m"}, Segment{"黑色", "n"}, Segment{"的士", "n"}, Segment{"。", "x"}},
		[]Segment{Segment{"枪杆子", "n"}, Segment{"中", "f"}, Segment{"出", "v"}, Segment{"政权", "n"}},
	}
	noHMMCutResult = [][]Segment{
		[]Segment{Segment{"这", "r"}, Segment{"是", "v"}, Segment{"一个", "m"}, Segment{"伸手不见五指", "i"}, Segment{"的", "uj"}, Segment{"黑夜", "n"}, Segment{"。", "x"}, Segment{"我", "r"}, Segment{"叫", "v"}, Segment{"孙悟空", "nr"}, Segment{"，", "x"}, Segment{"我", "r"}, Segment{"爱", "v"}, Segment{"北京", "ns"}, Segment{"，", "x"}, Segment{"我", "r"}, Segment{"爱", "v"}, Segment{"Python", "eng"}, Segment{"和", "c"}, Segment{"C++", "nz"}, Segment{"。", "x"}},
		[]Segment{Segment{"我", "r"}, Segment{"不", "d"}, Segment{"喜欢", "v"}, Segment{"日本", "ns"}, Segment{"和服", "nz"}, Segment{"。", "x"}},
		[]Segment{Segment{"雷猴", "n"}, Segment{"回归", "v"}, Segment{"人间", "n"}, Segment{"。", "x"}},
		[]Segment{Segment{"工信处", "n"}, Segment{"女干事", "n"}, Segment{"每月", "r"}, Segment{"经过", "p"}, Segment{"下属", "v"}, Segment{"科室", "n"}, Segment{"都", "d"}, Segment{"要", "v"}, Segment{"亲口", "n"}, Segment{"交代", "n"}, Segment{"24", "eng"}, Segment{"口", "q"}, Segment{"交换机", "n"}, Segment{"等", "u"}, Segment{"技术性", "n"}, Segment{"器件", "n"}, Segment{"的", "uj"}, Segment{"安装", "v"}, Segment{"工作", "vn"}},
		[]Segment{Segment{"我", "r"}, Segment{"需要", "v"}, Segment{"廉租房", "n"}},
		[]Segment{Segment{"永和", "nz"}, Segment{"服装", "vn"}, Segment{"饰品", "n"}, Segment{"有限公司", "n"}},
		[]Segment{Segment{"我", "r"}, Segment{"爱", "v"}, Segment{"北京", "ns"}, Segment{"天安门", "ns"}},
		[]Segment{Segment{"abc", "eng"}},
		[]Segment{Segment{"隐", "n"}, Segment{"马尔可夫", "nr"}},
		[]Segment{Segment{"雷猴", "n"}, Segment{"是", "v"}, Segment{"个", "q"}, Segment{"好", "a"}, Segment{"网站", "n"}},
		[]Segment{Segment{"“", "x"}, Segment{"Microsoft", "eng"}, Segment{"”", "x"}, Segment{"一", "m"}, Segment{"词", "n"}, Segment{"由", "p"}, Segment{"“", "x"}, Segment{"MICROcomputer", "eng"}, Segment{"（", "x"}, Segment{"微型", "b"}, Segment{"计算机", "n"}, Segment{"）", "x"}, Segment{"”", "x"}, Segment{"和", "c"}, Segment{"“", "x"}, Segment{"SOFTware", "eng"}, Segment{"（", "x"}, Segment{"软件", "n"}, Segment{"）", "x"}, Segment{"”", "x"}, Segment{"两", "m"}, Segment{"部分", "n"}, Segment{"组成", "v"}},
		[]Segment{Segment{"草泥马", "n"}, Segment{"和", "c"}, Segment{"欺", "vn"}, Segment{"实", "n"}, Segment{"马", "n"}, Segment{"是", "v"}, Segment{"今年", "t"}, Segment{"的", "uj"}, Segment{"流行", "v"}, Segment{"词汇", "n"}},
		[]Segment{Segment{"伊", "ns"}, Segment{"藤", "nr"}, Segment{"洋华堂", "n"}, Segment{"总府", "n"}, Segment{"店", "n"}},
		[]Segment{Segment{"中国科学院计算技术研究所", "nt"}},
		[]Segment{Segment{"罗密欧", "nr"}, Segment{"与", "p"}, Segment{"朱丽叶", "nr"}},
		[]Segment{Segment{"我", "r"}, Segment{"购买", "v"}, Segment{"了", "ul"}, Segment{"道具", "n"}, Segment{"和", "c"}, Segment{"服装", "vn"}},
		[]Segment{Segment{"PS", "eng"}, Segment{":", "x"}, Segment{" ", "x"}, Segment{"我", "r"}, Segment{"觉得", "v"}, Segment{"开源", "n"}, Segment{"有", "v"}, Segment{"一个", "m"}, Segment{"好处", "d"}, Segment{"，", "x"}, Segment{"就是", "d"}, Segment{"能够", "v"}, Segment{"敦促", "v"}, Segment{"自己", "r"}, Segment{"不断改进", "l"}, Segment{"，", "x"}, Segment{"避免", "v"}, Segment{"敞", "v"}, Segment{"帚", "ng"}, Segment{"自珍", "b"}},
		[]Segment{Segment{"湖北省", "ns"}, Segment{"石首市", "ns"}},
		[]Segment{Segment{"湖北省", "ns"}, Segment{"十堰市", "ns"}},
		[]Segment{Segment{"总经理", "n"}, Segment{"完成", "v"}, Segment{"了", "ul"}, Segment{"这件", "mq"}, Segment{"事情", "n"}},
		[]Segment{Segment{"电脑", "n"}, Segment{"修好", "v"}, Segment{"了", "ul"}},
		[]Segment{Segment{"做好", "v"}, Segment{"了", "ul"}, Segment{"这件", "mq"}, Segment{"事情", "n"}, Segment{"就", "d"}, Segment{"一了百了", "l"}, Segment{"了", "ul"}},
		[]Segment{Segment{"人们", "n"}, Segment{"审美", "vn"}, Segment{"的", "uj"}, Segment{"观点", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"我们", "r"}, Segment{"买", "v"}, Segment{"了", "ul"}, Segment{"一个", "m"}, Segment{"美的", "nr"}, Segment{"空调", "n"}},
		[]Segment{Segment{"线程", "n"}, Segment{"初始化", "l"}, Segment{"时", "n"}, Segment{"我们", "r"}, Segment{"要", "v"}, Segment{"注意", "v"}},
		[]Segment{Segment{"一个", "m"}, Segment{"分子", "n"}, Segment{"是", "v"}, Segment{"由", "p"}, Segment{"好多", "m"}, Segment{"原子", "n"}, Segment{"组织", "v"}, Segment{"成", "n"}, Segment{"的", "uj"}},
		[]Segment{Segment{"祝", "v"}, Segment{"你", "r"}, Segment{"马到功成", "i"}},
		[]Segment{Segment{"他", "r"}, Segment{"掉", "zg"}, Segment{"进", "v"}, Segment{"了", "ul"}, Segment{"无底洞", "ns"}, Segment{"里", "f"}},
		[]Segment{Segment{"中国", "ns"}, Segment{"的", "uj"}, Segment{"首都", "d"}, Segment{"是", "v"}, Segment{"北京", "ns"}},
		[]Segment{Segment{"孙", "zg"}, Segment{"君", "nz"}, Segment{"意", "n"}},
		[]Segment{Segment{"外交部", "nt"}, Segment{"发言人", "l"}, Segment{"马朝旭", "nr"}},
		[]Segment{Segment{"领导人", "n"}, Segment{"会议", "n"}, Segment{"和", "c"}, Segment{"第四届", "m"}, Segment{"东亚", "ns"}, Segment{"峰会", "n"}},
		[]Segment{Segment{"在", "p"}, Segment{"过去", "t"}, Segment{"的", "uj"}, Segment{"这", "r"}, Segment{"五年", "t"}},
		[]Segment{Segment{"还", "d"}, Segment{"需要", "v"}, Segment{"很", "zg"}, Segment{"长", "a"}, Segment{"的", "uj"}, Segment{"路", "n"}, Segment{"要", "v"}, Segment{"走", "v"}},
		[]Segment{Segment{"60", "eng"}, Segment{"周年", "t"}, Segment{"首都", "d"}, Segment{"阅兵", "v"}},
		[]Segment{Segment{"你好", "l"}, Segment{"人们", "n"}, Segment{"审美", "vn"}, Segment{"的", "uj"}, Segment{"观点", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"买", "v"}, Segment{"水果", "n"}, Segment{"然后", "c"}, Segment{"来", "v"}, Segment{"世博园", "nr"}},
		[]Segment{Segment{"买", "v"}, Segment{"水果", "n"}, Segment{"然后", "c"}, Segment{"去", "v"}, Segment{"世博园", "nr"}},
		[]Segment{Segment{"但是", "c"}, Segment{"后来", "t"}, Segment{"我", "r"}, Segment{"才", "d"}, Segment{"知道", "v"}, Segment{"你", "r"}, Segment{"是", "v"}, Segment{"对", "p"}, Segment{"的", "uj"}},
		[]Segment{Segment{"存在", "v"}, Segment{"即", "v"}, Segment{"合理", "vn"}},
		[]Segment{Segment{"的", "uj"}, Segment{"的", "uj"}, Segment{"的", "uj"}, Segment{"的", "uj"}, Segment{"的", "uj"}, Segment{"在", "p"}, Segment{"的", "uj"}, Segment{"的", "uj"}, Segment{"的", "uj"}, Segment{"的", "uj"}, Segment{"就", "d"}, Segment{"以", "p"}, Segment{"和", "c"}, Segment{"和", "c"}, Segment{"和", "c"}},
		[]Segment{Segment{"I", "eng"}, Segment{" ", "x"}, Segment{"love", "eng"}, Segment{"你", "r"}, Segment{"，", "x"}, Segment{"不以为耻", "i"}, Segment{"，", "x"}, Segment{"反", "zg"}, Segment{"以为", "c"}, Segment{"rong", "eng"}},
		[]Segment{Segment{"因", "p"}},
		[]Segment{},
		[]Segment{Segment{"hello", "eng"}, Segment{"你好", "l"}, Segment{"人们", "n"}, Segment{"审美", "vn"}, Segment{"的", "uj"}, Segment{"观点", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"很", "zg"}, Segment{"好", "a"}, Segment{"但", "c"}, Segment{"主要", "b"}, Segment{"是", "v"}, Segment{"基于", "p"}, Segment{"网页", "n"}, Segment{"形式", "n"}},
		[]Segment{Segment{"hello", "eng"}, Segment{"你好", "l"}, Segment{"人们", "n"}, Segment{"审美", "vn"}, Segment{"的", "uj"}, Segment{"观点", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"为什么", "r"}, Segment{"我", "r"}, Segment{"不能", "v"}, Segment{"拥有", "v"}, Segment{"想要", "v"}, Segment{"的", "uj"}, Segment{"生活", "vn"}},
		[]Segment{Segment{"后来", "t"}, Segment{"我", "r"}, Segment{"才", "d"}},
		[]Segment{Segment{"此次", "r"}, Segment{"来", "v"}, Segment{"中国", "ns"}, Segment{"是", "v"}, Segment{"为了", "p"}},
		[]Segment{Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解决", "v"}, Segment{"一些", "m"}, Segment{"问题", "n"}},
		[]Segment{Segment{",", "x"}, Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解决", "v"}, Segment{"一些", "m"}, Segment{"问题", "n"}},
		[]Segment{Segment{"其实", "d"}, Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解决", "v"}, Segment{"一些", "m"}, Segment{"问题", "n"}},
		[]Segment{Segment{"好人", "n"}, Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解决", "v"}, Segment{"一些", "m"}, Segment{"问题", "n"}},
		[]Segment{Segment{"是因为", "c"}, Segment{"和", "c"}, Segment{"国家", "n"}},
		[]Segment{Segment{"老年", "t"}, Segment{"搜索", "v"}, Segment{"还", "d"}, Segment{"支持", "v"}},
		[]Segment{Segment{"干脆", "d"}, Segment{"就", "d"}, Segment{"把", "p"}, Segment{"那", "r"}, Segment{"部", "n"}, Segment{"蒙", "v"}, Segment{"人", "n"}, Segment{"的", "uj"}, Segment{"闲", "n"}, Segment{"法", "j"}, Segment{"给", "p"}, Segment{"废", "v"}, Segment{"了", "ul"}, Segment{"拉倒", "v"}, Segment{"！", "x"}, Segment{"RT", "eng"}, Segment{" ", "x"}, Segment{"@", "x"}, Segment{"laoshipukong", "eng"}, Segment{" ", "x"}, Segment{":", "x"}, Segment{" ", "x"}, Segment{"27", "eng"}, Segment{"日", "m"}, Segment{"，", "x"}, Segment{"全国人大常委会", "nt"}, Segment{"第三次", "m"}, Segment{"审议", "v"}, Segment{"侵权", "v"}, Segment{"责任法", "n"}, Segment{"草案", "n"}, Segment{"，", "x"}, Segment{"删除", "v"}, Segment{"了", "ul"}, Segment{"有关", "vn"}, Segment{"医疗", "n"}, Segment{"损害", "v"}, Segment{"责任", "n"}, Segment{"“", "x"}, Segment{"举证", "v"}, Segment{"倒置", "v"}, Segment{"”", "x"}, Segment{"的", "uj"}, Segment{"规定", "n"}, Segment{"。", "x"}, Segment{"在", "p"}, Segment{"医患", "n"}, Segment{"纠纷", "n"}, Segment{"中", "f"}, Segment{"本", "r"}, Segment{"已", "d"}, Segment{"处于", "v"}, Segment{"弱势", "n"}, Segment{"地位", "n"}, Segment{"的", "uj"}, Segment{"消费者", "n"}, Segment{"由此", "c"}, Segment{"将", "d"}, Segment{"陷入", "v"}, Segment{"万劫不复", "i"}, Segment{"的", "uj"}, Segment{"境地", "s"}, Segment{"。", "x"}, Segment{" ", "x"}},
		[]Segment{Segment{"大", "a"}},
		[]Segment{},
		[]Segment{Segment{"他", "r"}, Segment{"说", "v"}, Segment{"的", "uj"}, Segment{"确实", "ad"}, Segment{"在", "p"}, Segment{"理", "n"}},
		[]Segment{Segment{"长春", "ns"}, Segment{"市长", "n"}, Segment{"春节", "t"}, Segment{"讲话", "n"}},
		[]Segment{Segment{"结婚", "v"}, Segment{"的", "uj"}, Segment{"和", "c"}, Segment{"尚未", "d"}, Segment{"结婚", "v"}, Segment{"的", "uj"}},
		[]Segment{Segment{"结合", "v"}, Segment{"成", "n"}, Segment{"分子", "n"}, Segment{"时", "n"}},
		[]Segment{Segment{"旅游", "vn"}, Segment{"和", "c"}, Segment{"服务", "vn"}, Segment{"是", "v"}, Segment{"最好", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"这件", "mq"}, Segment{"事情", "n"}, Segment{"的确", "d"}, Segment{"是", "v"}, Segment{"我", "r"}, Segment{"的", "uj"}, Segment{"错", "v"}},
		[]Segment{Segment{"供", "v"}, Segment{"大家", "n"}, Segment{"参考", "v"}, Segment{"指正", "v"}},
		[]Segment{Segment{"哈尔滨", "ns"}, Segment{"政府", "n"}, Segment{"公布", "v"}, Segment{"塌", "v"}, Segment{"桥", "n"}, Segment{"原因", "n"}},
		[]Segment{Segment{"我", "r"}, Segment{"在", "p"}, Segment{"机场", "n"}, Segment{"入口处", "i"}},
		[]Segment{Segment{"邢", "nr"}, Segment{"永", "ns"}, Segment{"臣", "n"}, Segment{"摄影", "n"}, Segment{"报道", "v"}},
		[]Segment{Segment{"BP", "eng"}, Segment{"神经网络", "n"}, Segment{"如何", "r"}, Segment{"训练", "vn"}, Segment{"才能", "v"}, Segment{"在", "p"}, Segment{"分类", "n"}, Segment{"时", "n"}, Segment{"增加", "v"}, Segment{"区分度", "n"}, Segment{"？", "x"}},
		[]Segment{Segment{"南京市", "ns"}, Segment{"长江大桥", "ns"}},
		[]Segment{Segment{"应", "v"}, Segment{"一些", "m"}, Segment{"使用者", "n"}, Segment{"的", "uj"}, Segment{"建议", "n"}, Segment{"，", "x"}, Segment{"也", "d"}, Segment{"为了", "p"}, Segment{"便于", "v"}, Segment{"利用", "n"}, Segment{"NiuTrans", "eng"}, Segment{"用于", "v"}, Segment{"SMT", "eng"}, Segment{"研究", "vn"}},
		[]Segment{Segment{"长春市", "ns"}, Segment{"长春", "ns"}, Segment{"药店", "n"}},
		[]Segment{Segment{"邓颖超", "nr"}, Segment{"生前", "t"}, Segment{"最", "d"}, Segment{"喜欢", "v"}, Segment{"的", "uj"}, Segment{"衣服", "n"}},
		[]Segment{Segment{"胡锦涛", "nr"}, Segment{"是", "v"}, Segment{"热爱", "a"}, Segment{"世界", "n"}, Segment{"和平", "nz"}, Segment{"的", "uj"}, Segment{"政治局", "n"}, Segment{"常委", "j"}},
		[]Segment{Segment{"程序员", "n"}, Segment{"祝", "v"}, Segment{"海林", "nz"}, Segment{"和", "c"}, Segment{"朱", "nr"}, Segment{"会", "v"}, Segment{"震", "v"}, Segment{"是", "v"}, Segment{"在", "p"}, Segment{"孙", "zg"}, Segment{"健", "a"}, Segment{"的", "uj"}, Segment{"左面", "f"}, Segment{"和", "c"}, Segment{"右面", "f"}, Segment{",", "x"}, Segment{" ", "x"}, Segment{"范", "nr"}, Segment{"凯", "nr"}, Segment{"在", "p"}, Segment{"最", "d"}, Segment{"右面", "f"}, Segment{".", "x"}, Segment{"再", "d"}, Segment{"往", "zg"}, Segment{"左", "m"}, Segment{"是", "v"}, Segment{"李", "nr"}, Segment{"松", "v"}, Segment{"洪", "nr"}},
		[]Segment{Segment{"一次性", "d"}, Segment{"交", "v"}, Segment{"多少", "m"}, Segment{"钱", "n"}},
		[]Segment{Segment{"两块", "m"}, Segment{"五", "m"}, Segment{"一套", "m"}, Segment{"，", "x"}, Segment{"三块", "m"}, Segment{"八", "m"}, Segment{"一斤", "m"}, Segment{"，", "x"}, Segment{"四块", "m"}, Segment{"七", "m"}, Segment{"一本", "m"}, Segment{"，", "x"}, Segment{"五块", "m"}, Segment{"六", "m"}, Segment{"一条", "m"}},
		[]Segment{Segment{"小", "a"}, Segment{"和尚", "nr"}, Segment{"留", "v"}, Segment{"了", "ul"}, Segment{"一个", "m"}, Segment{"像", "v"}, Segment{"大", "a"}, Segment{"和尚", "nr"}, Segment{"一样", "r"}, Segment{"的", "uj"}, Segment{"和尚头", "nr"}},
		[]Segment{Segment{"我", "r"}, Segment{"是", "v"}, Segment{"中华人民共和国", "ns"}, Segment{"公民", "n"}, Segment{";", "x"}, Segment{"我", "r"}, Segment{"爸爸", "n"}, Segment{"是", "v"}, Segment{"共和党", "nt"}, Segment{"党员", "n"}, Segment{";", "x"}, Segment{" ", "x"}, Segment{"地铁", "n"}, Segment{"和平门", "ns"}, Segment{"站", "v"}},
		[]Segment{Segment{"张晓梅", "nr"}, Segment{"去", "v"}, Segment{"人民", "n"}, Segment{"医院", "n"}, Segment{"做", "v"}, Segment{"了", "ul"}, Segment{"个", "q"}, Segment{"B超", "n"}, Segment{"然后", "c"}, Segment{"去", "v"}, Segment{"买", "v"}, Segment{"了", "ul"}, Segment{"件", "zg"}, Segment{"T恤", "n"}},
		[]Segment{Segment{"AT&T", "nz"}, Segment{"是", "v"}, Segment{"一件", "m"}, Segment{"不错", "a"}, Segment{"的", "uj"}, Segment{"公司", "n"}, Segment{"，", "x"}, Segment{"给", "p"}, Segment{"你", "r"}, Segment{"发", "v"}, Segment{"offer", "eng"}, Segment{"了", "ul"}, Segment{"吗", "y"}, Segment{"？", "x"}},
		[]Segment{Segment{"C++", "nz"}, Segment{"和", "c"}, Segment{"c#", "nz"}, Segment{"是", "v"}, Segment{"什么", "r"}, Segment{"关系", "n"}, Segment{"？", "x"}, Segment{"11", "eng"}, Segment{"+", "x"}, Segment{"122", "eng"}, Segment{"=", "x"}, Segment{"133", "eng"}, Segment{"，", "x"}, Segment{"是", "v"}, Segment{"吗", "y"}, Segment{"？", "x"}, Segment{"PI", "eng"}, Segment{"=", "x"}, Segment{"3", "eng"}, Segment{".", "x"}, Segment{"14159", "eng"}},
		[]Segment{Segment{"你", "r"}, Segment{"认识", "v"}, Segment{"那个", "r"}, Segment{"和", "c"}, Segment{"主席", "n"}, Segment{"握手", "v"}, Segment{"的", "uj"}, Segment{"的哥", "n"}, Segment{"吗", "y"}, Segment{"？", "x"}, Segment{"他", "r"}, Segment{"开", "v"}, Segment{"一辆", "m"}, Segment{"黑色", "n"}, Segment{"的士", "n"}, Segment{"。", "x"}},
		[]Segment{Segment{"枪杆子", "n"}, Segment{"中", "f"}, Segment{"出", "v"}, Segment{"政权", "n"}},
	}
)

func init() {
	seg.LoadDictionary("../dict.txt")
}

func chanToArray(ch <-chan Segment) []Segment {
	var result []Segment
	for word := range ch {
		result = append(result, word)
	}
	return result
}

func TestCut(t *testing.T) {
	for index, content := range testContents {
		result := chanToArray(seg.Cut(content, true))
		if len(defaultCutResult[index]) != len(result) {
			t.Errorf("default cut for %s length should be %d not %d\n",
				content, len(defaultCutResult[index]), len(result))
			t.Errorf("expect: %v\n", defaultCutResult[index])
			t.Fatalf("got: %v\n", result)
		}
		for i := range result {
			if result[i] != defaultCutResult[index][i] {
				t.Fatalf("expect %s, got %s", defaultCutResult[index][i], result[i])
			}
		}
		result = chanToArray(seg.Cut(content, false))
		if len(noHMMCutResult[index]) != len(result) {
			t.Fatal(content)
		}
		for i := range result {
			if result[i] != noHMMCutResult[index][i] {
				t.Fatal(content)
			}
		}

	}
}

// https://github.com/fxsjy/jieba/issues/132
func TestBug132(t *testing.T) {
	sentence := "又跛又啞"
	cutResult := []Segment{
		Segment{"又", "d"},
		Segment{"跛", "a"},
		Segment{"又", "d"},
		Segment{"啞", "v"},
	}
	result := chanToArray(seg.Cut(sentence, true))
	if len(cutResult) != len(result) {
		t.Fatal(result)
	}
	for i := range result {
		if result[i] != cutResult[i] {
			t.Fatal(result[i])
		}
	}
}

// https://github.com/fxsjy/jieba/issues/137
func TestBug137(t *testing.T) {
	sentence := "前港督衛奕信在八八年十月宣布成立中央政策研究組"
	cutResult := []Segment{
		Segment{"前", "f"},
		Segment{"港督", "n"},
		Segment{"衛奕", "z"},
		Segment{"信", "n"},
		Segment{"在", "p"},
		Segment{"八八年", "m"},
		Segment{"十月", "t"},
		Segment{"宣布", "v"},
		Segment{"成立", "v"},
		Segment{"中央", "n"},
		Segment{"政策", "n"},
		Segment{"研究", "vn"},
		Segment{"組", "x"},
	}
	result := chanToArray(seg.Cut(sentence, true))
	if len(cutResult) != len(result) {
		t.Fatal(result)
	}
	for i := range result {
		if result[i] != cutResult[i] {
			t.Fatal(result[i])
		}
	}
}

func TestUserDict(t *testing.T) {
	seg.LoadUserDictionary("../userdict.txt")
	defer seg.LoadDictionary("../dict.txt")
	sentence := "李小福是创新办主任也是云计算方面的专家; 什么是八一双鹿例如我输入一个带“韩玉赏鉴”的标题，在自定义词库中也增加了此词为N类型"

	cutResult := []Segment{
		Segment{"李小福", "nr"},
		Segment{"是", "v"},
		Segment{"创新办", "i"},
		Segment{"主任", "b"},
		Segment{"也", "d"},
		Segment{"是", "v"},
		Segment{"云计算", "x"},
		Segment{"方面", "n"},
		Segment{"的", "uj"},
		Segment{"专家", "n"},
		Segment{";", "x"},
		Segment{" ", "x"},
		Segment{"什么", "r"},
		Segment{"是", "v"},
		Segment{"八一双鹿", "nz"},
		Segment{"例如", "v"},
		Segment{"我", "r"},
		Segment{"输入", "v"},
		Segment{"一个", "m"},
		Segment{"带", "v"},
		Segment{"“", "x"},
		Segment{"韩玉赏鉴", "nz"},
		Segment{"”", "x"},
		Segment{"的", "uj"},
		Segment{"标题", "n"},
		Segment{"，", "x"},
		Segment{"在", "p"},
		Segment{"自定义词", "n"},
		Segment{"库中", "nrt"},
		Segment{"也", "d"},
		Segment{"增加", "v"},
		Segment{"了", "ul"},
		Segment{"此", "r"},
		Segment{"词", "n"},
		Segment{"为", "p"},
		Segment{"N", "eng"},
		Segment{"类型", "n"}}

	result := chanToArray(seg.Cut(sentence, true))
	if len(cutResult) != len(result) {
		t.Fatal(result)
	}
	for i := range result {
		if result[i] != cutResult[i] {
			t.Fatal(result[i])
		}
	}
}

func BenchmarkCutNoHMM(b *testing.B) {
	sentence := "工信处女干事每月经过下属科室都要亲口交代24口交换机等技术性器件的安装工作"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		chanToArray(seg.Cut(sentence, false))
	}
}

func BenchmarkCut(b *testing.B) {
	sentence := "工信处女干事每月经过下属科室都要亲口交代24口交换机等技术性器件的安装工作"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		chanToArray(seg.Cut(sentence, true))
	}
}
