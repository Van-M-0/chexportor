package defines

// lobby
const (
	ErrCommonSuccess			= 1
	ErrCommonCache 				= 2
	ErrCommonWait				= 3
	ErrCoomonSystem 			= 4

	ErrClientLoginWait 	  		= 100
	ErrClientLoginNeedCreate	= 1001

	ErrCreateAccountErr			= 100
	ErrCreateAccountWait 		= 101
	ErrCreateAccountExists 		= 102

	ErrClientBuyItemNotExists	= 101
	ErrClientBuyItemMoneyNotEnough = 102
	ErrClientBuyItemTime 		= 103

	ErrUserHornMessageUserErr 	= 101
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

	ErrEnterRoomSuccess			= 101
	ErrEnterRoomUserNotIn		= 102
	ErrEnterRoomNotExists 		= 103
	ErrEnterRoomMoudle			= 104
	ErrEnterRoomNotSame 		= 105

	ErrLeaveRoomSuccess			= 101
	ErrLeaveRoomUserNotIn		= 102
	ErrLeaveRoomNotExists		= 103
)