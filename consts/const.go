package consts

const DataReceivePort  = `8081`
const (
	TYPE_OWN_DEVICE = iota
	TYPE_OTHER_DEVICE
)

const (
	LEVEL1 = iota
	LEVEL2
	LEVEL3
)

const (
	Cycles  = 5
	Res     = 0.001
	Size    = 100
	Nframes = 64
	Delay   = 8
)

const (
	WhiteIndex = 0
	BlackIndex = 1
	GreenIndex = 2
	RedIndex   = 3
	BlueIndex  = 4
	YellowIndex= 5
)

const (
	DetectionIdTag = "DetectionId"  //id生成器的检测订单号的tag
	UserIdTag	   = "UserId"		//id生成器用户号的tag
)

const (
	DataSliceLength = 60 //数据分组长度
)

const (
	NoMark			= 0	 //没有标记
	MarkNomal 		= 1	 //正常标记
	MarkAbnormal 	= 2 //异常标记
)

const (
	Mailuser 	 = ``	//企业邮箱用户名
	Mailpassword = ``   //企业邮箱密码
	Mailhosts	 = `mail3.bupt.edu.cn:25`	//
)