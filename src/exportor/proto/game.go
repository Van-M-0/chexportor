package proto

// game proto 5000 - 6000
const (
	CmdGamePlayerLogin 			= 5001
	CmdGameCreateRoom			= 5002
	CmdGameEnterRoom			= 5003
	CmdGamePlayerLeaveRoom		= 5004
	CmdGamePlayerReturn2lobby	= 5005
	CmdGamePlayerReleaseRoom 	= 5006
	CmdGamePlayerReleaseRoomResponse = 5007
	CmdGamePlayerRoomChat 		= 5010
	CmdGameEnterCoinRoom		= 5011
	CmdGamePlayerMessage 		= 5020
)


type PlayerGameCommonError struct {
	ErrCode 		int
}

type PlayerLogin struct {
	Uid 		uint32
}

type PlayerLoginRet struct {
	UidTest				uint32
	AccountTest			string
	NameTest			string
	UserIdTest			int
	ReEnter 			bool
	ErrCode 			int
}

type PlayerCreateRoom struct {
	Kind 			int
	Enter 			bool
	Conf 			[]byte
}

type PlayerCreateRoomRet struct {
	ErrCode 			int
	RoomId 				uint32
	ServerId 			int
	Conf 				[]byte
}

type PlayerEnterRoom struct {
	ServerId 			uint32
	RoomId				uint32
}

type PlayerEnterRoomRet struct {
	ErrCode 			int
	ServerId 			int
	Conf 				[]byte
}

type PlayerLeaveRoom struct {
	RoomId 				uint32
}

type PlayerLeaveRoomRet struct {
	ErrCode 			uint32
}

type PlayerGameMessage struct {
	A 		uint32
	B 		[]byte
}

type PlayerGameMessageRet struct {
	Cmd 		uint32
	Msg 		interface{}
}

type PlayerSubGameMessageRet struct {
	Cmd 		uint32
	Msg 		interface{}
}

type PlayerReturn2Lobby struct {
	ErrCode 	int
}

type PlayerRoomChat struct {
	Content 		string
}

type PlayerRoomChatRet struct {
	ErrCode 		int
	Userid 			uint32
	Content 		string
}

type PlayerGameReleaseRoom struct {
}

type PlayerGameReleaseRoomRet struct {
	ErrCode 		int
	Sponser	 		uint32
	Released 		bool
}

type PlayerGameReleaseRoomResponse struct {
	Agree 			bool
}

type PlayerGameReleaseRoomResponseRet struct {
	ErrCode 		int
	Released 		bool
	Agree 			bool
	Voter			uint32
}

type PlayerGameEnterCoinRoom struct {
/*
	EnterType :

	1. 换房，由一个房间换到另一个房间
		RoomId 不为0， 且必须是和服务器保持一致才能换

	2. 再次进入房间
		RoomId 不为0， 且玩家身上原先有这个roomid

	3. 创建房间
		Kind 不为0， 且必须正确, Conf不为空

	4. 加入金币房间
		RoomId 不为0， 且玩家身上原先没有房间

	5. 匹配房间
		Kind 不为0， 且必须正确

*/
	EnterType 		int
	Kind 			int
	RoomId 			int
	Conf 			[]byte
}

type PlayerGameEnterCoinRoomRet struct {
	ErrCode 		int
}