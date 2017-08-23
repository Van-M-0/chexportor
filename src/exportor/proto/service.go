package proto

import "time"

//
type DbUserLoginArg struct {
	LoginType 		int
	Acc 			string
	Name 			string
	Headimg 		string
	Sex 			int
}

type UserData struct {
	Activity 		[]byte
	Quest 			[]byte
}

type DbUserLoginReply struct {
	Err 			string
	UserItemList 	[]UserItem
	Ud 				[]byte
	Identify 		SynceIdentifyInfo
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
	Conf 			[]byte
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
	Conf 			[]byte
}

// master skd service
type MsSdkWechatLoginArg struct {
	Code 			string
	Device			string
}

type MsSdkWechatLoginReply struct {
	ErrCode 		string
	OpenId 			string
	Token 			string
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

type MsLoadActivitysArg struct {

}

type ActivityItem struct {
	Id 			int
	Desc 		string
	Actype 		string
	Starttime 	time.Time
	Finishtime 	time.Time
	Rewardids 	string
}

type ActivityRewardItem struct {
	Id 			int
	RewardType 	string
	ItemId 		int
	Num 		int
}

type MsLoadActivitysReply struct {
	ErrCode 			string
	Activitys 			[]*ActivityItem
	ActivityRewards 	[]*ActivityRewardItem
}

type MsLoadQuestArg struct {

}

type MsLoadQuestReply struct {
	ErrCode 			string
	Quests 				[]QuestItem
	Rewards 			[]QuestRewardItem
}


// user data
type MsSaveUserDataArg struct {
	UserId 				uint32
	UserData 			[]byte
}

type MsSaveUserDataReply struct {
	ErrCode 			string
}

type MsSaveIdentifyInfoArg struct {
	Userid 				uint32
	Phone 				string
	Name 				string
	Idcard 				string
}

type MsSaveIdentifyInfoReply struct {
	ErrCode 			string
}

// world service

type ModuleInfo struct {
	Province 		string
	City 			string
	Name 			string
	Kind 			int
	Conf 			interface{}
	GateIp 			string
}

type GameModule struct {
	ModuleConf 		interface{}
	GatewayHost 	string
}