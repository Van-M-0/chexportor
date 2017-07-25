package proto

// game proto 5000 - 6000
const (
	CmdGamePlayerLogin 			= 5001
	CmdGamePlayerEnterRoom		= 5002
	CmdGamePlayerLeaveRoom		= 5003

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

type PlayerCreateRoom struct {
	//common
	RoomId 		uint32
	Kind 		int
	Enter 		bool

	//special
	Conf 		[]byte
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