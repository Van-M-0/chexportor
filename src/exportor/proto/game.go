package proto

// game proto 5000 - 6000
const (
	CmdGamePlayerLogin 			= 5001
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
	Cmd 		uint32
	Msg 		[]byte
}