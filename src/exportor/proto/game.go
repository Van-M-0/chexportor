package proto

// game proto 5000 - 6000
const (
	CmdGamePlayerLogin 			= 5001
	CmdGamePlayerCreateRoom		= 5002
	CmdGamePlayerEnterRoom		= 5003
	CmdGamePlayerLeaveRoom		= 5004

	CmdGamePlayerMessage 		= 5020
)

type PlayerLogin struct {
	Uid 		uint32
	Name 		string
}

type PlayerLoginRet struct {
	ErrCode 	int
}

type PlayerCreateRoom struct {
	//common
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