package proto

type PMLoadUser struct {
	Acc 		string
	Create 		bool
	Guest 		bool
}

type PMLoadUserFinish struct {
	Acc 		string
	Err 		error
	Code 		int
}

type PMCreateAccount struct {
	Name 		string
	Sex 		byte
	Pwd 		string
}

type PMCreateAccountFinish struct {
	Err 		int
	Account 	string
	Pwd 		string
}

type PMUserCreateRoom struct {
	ServerId 	int
	Uid 		uint32
	Message 	UserCreateRoomReq
}

type PMUserCreateRoomRet struct {
	ErrCode 	int
}

type PMUserEnterRoom struct {
	ServerId 	int
	RoomId 		uint32
	Uid 		uint32
}

type PMUserEnterRoomRet struct {
	ErrCode 	int
}

type PMMallItemUdpate struct {
	Items 		[]MallItem
}

type PMMallItemLoaded struct {
	Items 		[]MallItem
}