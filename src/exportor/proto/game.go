package proto

// game proto 5000 - 6000
const (
	CmdGamePlayerLogin 			= 5001
	CmdGameCreateRoom			= 5002
	CmdGameEnterRoom			= 5003
	CmdGamePlayerLeaveRoom		= 5004
	CmdGamePlayerReturn2lobby	= 5005
	CmdGamePlayerMessage 		= 5020
)

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

}