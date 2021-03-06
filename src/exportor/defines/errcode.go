package defines

// lobby
const (
	ErrCommonSuccess			= 1
	ErrCommonCache 				= 2
	ErrCommonWait				= 3
	ErrCoomonSystem 			= 4
	ErrComononUserNotIn			= 5
	ErrCommonInvalidReq 		= 6

	ErrClientLoginWait 	  		= 100
	ErrClientLoginNeedCreate	= 1001
	ErrClientLoginRelogin 		= 103

	ErrCreateAccountErr			= 100
	ErrCreateAccountWait 		= 101
	ErrCreateAccountExists 		= 102

	ErrClientBuyItemNotExists	= 101
	ErrClientBuyItemMoneyNotEnough = 102
	ErrClientBuyItemTime 		= 103
	ErrClientBuyInvalid			= 104
	ErrClientBuyConsumeErr 		= 105

	ErrUserHornMessageUserErr 	= 101


	ErrQuestProcessNotHave		= 101
	ErrQuestProcessMustWait		= 102
	ErrQuestProcessFinished 	= 103
	ErrQuestPorcessNotFinish 	= 104
	ErrQuestProcessCompletioned = 105

	ErrClubCreateHaveClub 		= 101
	ErrClubCreateTryAgain		= 102
	ErrClubCreateWait 			= 103
	ErrClubNotAgent 			= 104

	ErrClubJoinNotExists		= 101
	ErrClubJoinHaveClub			= 102

	ErrClubLeaveNotExists 		= 101
	ErrClubLeaveNotAllow		= 102
	ErrClubLeaveMoneyNotEnough  = 103
	ErrClubLeaveError			= 104
	ErrClubLeaveNotJoind 		= 105

	ErrReleaseRoomNotExists 	= 101
	ErrReleaseRoomIng 			= 102
)

// game
const (
	ErrPlayerLoginSuccess 		= 101
	ErrPlayerLoginErr     		= 102
	ErrPlayerLoginCache   		= 103

	ErrCreateRoomSuccess		= 101
	ErrCreateRoomUserNotIn		= 102
	ErrCreateRoomWait			= 103
	ErrCreateRoomKind 			= 104
	ErrCreateRoomGameMoudele	= 105
	ErrCreateRoomSystme			= 106
	ErrCreateRoomHaveRoom 		= 107
	ErrCreateRoomRoomId 		= 108
	ErrCreateRoomCmd			= 109
	ErrCreateRoomCreate			= 110
	ErrCreateRoomType 			= 111
	ErrCoinCreateRoomMaxCount 	= 112

	ErrEnterRoomSuccess			= 101
	ErrEnterRoomUserNotIn		= 102
	ErrEnterRoomNotExists 		= 103
	ErrEnterRoomMoudle			= 104
	ErrEnterRoomNotSame 		= 105
	ErrEnterRoomQueryConf 		= 106

	ErrLeaveRoomSuccess			= 101
	ErrLeaveRoomUserNotIn		= 102
	ErrLeaveRoomNotExists		= 103

	ErrEnterCoinRoomChangeNotSame = 101
	ErrEnterCoinRoomEnterNotSame = 102
	ErrEnterCoinRoomInvalidReq 	= 103
	ErrEnterCoinRoomHaveRoom	= 104
	ErrEnterCoinOldRoomNotExist	= 105
	ErrEnterCoinCategory		= 106
	ErrEnterCoinRoomChangeInvalid = 107
)