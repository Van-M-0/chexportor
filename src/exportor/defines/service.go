package defines

//
type DbUserLoginArg struct {
	Acc 			string
}

type DbUserLoginReply struct {
	Err 			string
}

//
type DbCreateAccountArg struct {
	UserName		string
}

type DbCreateAccountReply struct {
	Err 			string
	Acc 			string
}

//
type LbGetRoomIdArg struct {

}

type LbGetRoomIdReply struct {
	Err 		string
	RoomId 		uint32
}

//
type LbReportRoomInfoArg struct {
	Uid 		uint32
	Kind 		int	//1 create 2 destory 3 enter
	ServerId 	int
	RoomId 		uint32
}

type LbReportRoomInfoReply struct {
	Err 		string
}
