package proto

//
type DbUserLoginArg struct {
	Acc 			string
}

type DbUserLoginReply struct {
	Err 			string
	UserItemList 	[]UserItem
}

//
type DbCreateAccountArg struct {
	Acc			string
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

// master server serivce
type MsServerIdArg struct {
	Type 		string
}

type MsServerIdReply struct {
	Id 			int
}

type MsServerReleaseArg struct {
	Id 			int
}

type MsServerReleaseReply struct {
	ErrCode 	string
}

type MsServerDiscArg struct {
	Id 			int
}

type MsServerDisReply struct {
	ErrCode 	string
}

type MsGsStatusArg struct {

}

type MsGsStatusReply struct {

}

type MsSelectGameServerArg struct {
	Kind 			int
}

type MsSelectGameServerReply struct {
	ServerId 		int
}

// master room service
type MsCreateoomIdArg struct {
	ServerId 		int
}

type MsCreateRoomIdReply struct {
	RoomId 			uint32
}

type MsReleaseRoomArg struct {
	ServerId 		int
	RoomId 			uint32
}

type MsReleaseReply struct {
	ErrCode 		string
}

type MsGetRoomServerIdArg struct {
	RoomId 			uint32
}

type MsGetRoomServerIdReply struct {
	ServerId 		int
}

// master game moudle service
type MsModuleItem struct {
	Kind 			int
	GameConf 		[]byte
	GatewayHost 	string
}

type MsGameMoudleRegisterArg struct {
	ServerId 		int
	ModList 		[]MsModuleItem
}

type MsGameMoudleRegisterReply struct {
	ErrCode 		string
}

// master notice service
type NoticeOperatoin struct {
	Operation 		string
	Notices 		[]*NoticeItem
}

type UserLoadNotice struct {

}

type UserLoadNoticeRet struct {
	Notices 		[]*NoticeItem
}

type UserSendNotice struct {
	Kind 			int
	Content 		string
}

type UserSendNoticeRet struct {
	SendUserName 	string
	Kind 			int
	Content 		string
}

type MsLoadNoticeArg struct {

}

type MsLoadNoticeReply struct {
	Notices 		[]*NoticeItem
}

type MsLoadMallItemListArg struct {

}

type MsLoadMallItemListReply struct {
	Malls 			[]*MallItem
}

type MsLoadUserRankArg struct {
	RankType 		int
	Count 			int
}

type MsLoadItemConfigArg struct {

}

type MsLoadItemConfigReply struct {
	ItemConfigList		[]ItemConfig
}

type MsLoadUserRankReply struct {
	ErrCode 	string
	Users 		[]*UserRankItem
}

type MsLoadGameLibsArg struct {

}

type GameLibItem struct {
	Id 			int
	Name 		string
	Area 		string
	City 		string
	Province 	string
}

type MsLoadGameLibsReply struct {
	ErrCode 	string
	Libs 		[]GameLibItem
}