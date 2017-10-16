package defines

import "time"

const (
	WaitChannelInfinite      = 0
	WaitChannelNormal		 = 10 * time.Second
)

const (

	ChannelTypeDb			= "proxy"
		ChannelLoadUser 		= "loadUser"
		ChannelLoadUserFinish	= "loadUserFinish"
		ChannelCreateAccount 	= "createAccount"
		ChannelCreateAccountFinish = "createAccountFinish"

	ChannelTypeLobby 		= "lobby"
		ChannelCreateRoom 		= "lobbyCreateRoom"
		ChannelCreateRoomFinish = "lobbyCreateRoomFinish"
		ChannelEnterRoom		= "lobbyEnterRoom"
		ChannelEnterRoomFinish  = "lobbyEnterRoomFinish"

	ChannelTypeMall			= "mall"
		ChannelUpdateMall 		= "mallUpdate"
		ChannelLoadMall 		= "mallLoad"

	ChannelTtypeNotice		= "notice"
		ChannelUpdateNotice		= "noticeUpdate"
		ChannelLoadNotice	= "noticeLoad"


)

const (
	MallItemBuyTypeGold 	= 1
	MallItemBuyTypeWechat	= 2
)

var (
	WDServicePort		= "predefined"
	MSServicePort		= "predefined"
	DBSerivcePort 		= "predefined"
	LbServicePort 		= "predefined"
)

var (
	WorkDir 			string
	GlobalConfig 		StartConfigFile
)

const (
	PpDiamond			= 1		//钻石
	PpRoomCard 			= 2		//房卡
	PpGold				= 3		//金币
	PpScore 			= 4		//积分
	PpRoomId 			= 5
	PpItem 				= 6
	PpMax 				= 7
)

const (
	MallItemCategoryGold 	= 1
	MallItemCategoryDiamond = 2
	MallItemCategoryItem = 3
)

const (
	GameModuleXz 			= 1
)

const (
	RankTypeDiamond 		= 1
)

const (
	LoginTypeWechat 		= 1
	LoginTypeGameAccount 	= 2
	LoginTypeGuest			= 3
)

const (
	ClientPatFMaxSize		= 512
	ClientPktBMaxSize		= 1024*1024
)