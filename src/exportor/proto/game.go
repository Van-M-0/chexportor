package proto

// game proto 5000 - 6000
const (
	CmdGamePlayerLogin 			= 5001
	CmdGameCreateRoom			= 5002
	CmdGameEnterRoom			= 5003
	CmdGamePlayerMessage 		= 5020
)

type PlayerLogin struct {
	Uid 		uint32
	Name 		string
}

type PlayerLoginRet struct {
	UidTest				uint32
	AccountTest			string
	NameTest			string
	UserIdTest			int
	ErrCode 	int
}

type PlayerCreateRoomRet struct {
	ErrCode 	int
}

type PlayerEnterRoom struct {
	RoomId		uint32
}

type PlayerEnterRoomRet struct {
	ErrCode 	int
}

type PlayerLeaveRoom struct {
	RoomId 		uint32
}

type PlayerLeaveRoomRet struct {
	ErrCode 	uint32
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