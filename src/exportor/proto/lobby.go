package proto

// lobby proto 1000 - 2000

const (
	CmdClientLogin   	= 1001
	CmdGuestLogin    	= 1002
	CmdCreateAccount 	= 1003
	CmdWechatLogin	 	= 1004

	CmdCreateRoom 		= 1005
	CmdEnterRoom 		= 1006
	CmdLeaveRoom 		= 1007

	CmdClientLoadMallItem = 1008
	CmdClientBuyItem 	= 1009

	CmdUserLoadNotice	= 1010
	CmdUserSendMessage	= 1011

)

type ClientLogin struct {
	Account 	string
}

type GuestLogin struct {
	Account 	string
}

type ClientLoginRet struct {
	ErrCode 	int
	Uid 		uint32
	Account 	string
	Name 		string
	UserId 		uint32
}

type CreateAccount struct {
	Name 		string
	Sex 		byte
}

type CreateAccountRet struct {
	ErrCode 	int
	Account 	string
	Pwd 		string
}

type WechatLoginReq struct {
	Code 		string
	Device 		string
}

type WechatLoginRet struct {
	Account 	string
}

type UserCreateRoomReq struct {
	RoomId 		uint32
	Kind 		int
	Enter 		bool

	//special
	Conf 		[]byte
}

type UserCreateRoomRet struct {
	ErrCode 	int
	RoomId		uint32
}

type UserEnterRoomReq struct {
	RoomId 		uint32
}

type UserEnterRoomRet struct {
	ErrCode 	int
}
