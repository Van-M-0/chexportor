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

const (
	MSServicePort		= ":11344"
	DBSerivcePort 		= ":11345"
	LbServicePort 		= ":11346"
)

const (
	PpDiamond			= 1		//钻石
	PpRoomCard 			= 2		//房卡
	PpGold				= 3		//金币
	PpScore 			= 4		//积分
	PpRoomId 			= 5
	PpMax 				= 6
)

const (
	MallItemCategoryDiamond = 1
	MallItemCategoryGold 	= 2
	MallItemCategoryRoomCard = 3
)

const (
	GameModuleXz 			= 1
)
