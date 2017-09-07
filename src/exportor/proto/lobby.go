package proto

import "time"
// lobby proto 1000 - 2000

const (
	CmdClientLogin   	= 1001
	CmdGuestLogin    	= 1002
	CmdCreateAccount 	= 1003
	CmdWechatLogin	 	= 1004

	CmdCreateRoom 		= 1005
	CmdEnterRoom 		= 1006
	CmdLeaveRoom 		= 1007
	CmdDestoyRooom 		= 1008

	CmdClientLoadMallList = 1008
	CmdClientBuyItem 	= 1009

	CmdUserLoadNotice	= 1010
	CmdHornMessage		= 1011
	CmdNoticeUpdate 	= 1012

	CmdUserLoadRank 		= 1013

	CmdUserGetRecordList 	= 1015
	CmdUserGetRecord 		= 1016

	CmdUserLoadActivityList = 1020

	CmdUserLoadQuest 		= 1030
	CmdUserProcessQuest 	= 1031
	CmdUserCompleteQuest	= 1032

	CmdUserIdentify 		= 1040
	CmdUserLoadIdentify 	= 1041

	CmdUserCreatClub 		= 1050
	CmdUserJoinClub 		= 1051
	CmdUserLeaveClub		= 1052

	CmdSystemSyncItem 		= 1200

	CmdLobbyPerformance		= 1999
)

type LobbyPerformance struct {
	T 				time.Time
	SubCmd 			int
}

type LobbyPerformanceRet struct {
	SubCmd 			int
	T1				time.Time
	T2				time.Time
	T3 				time.Time
}

type ClientLogin struct {
	LoginType 	int
	Account 	string
	Name 		string
	Sex 		int
	Headimg 	string
}

type GuestLogin struct {
	Account 	string
}

type ClientLoginRet struct {
	ErrCode 	int
	Uid 		uint32		//动态id
	UserId 		uint32		//用户id
	Account 	string		//用户账号，或者openid
	Sex 		uint8
	Name 		string
	HeadImg 	string
	Diamond 	int
	Gold 		int64
	Score 		int
	RoomId 		int
}

type CreateAccount struct {
	Name 		string
	Sex 		byte
}

type CreateAccountRet struct {
	ErrCode 	int
	Account 	string
	Pwd 		string
}

type WechatLoginReq struct {
	Code 		string
	Device 		string
}

type WechatLoginRet struct {
	ErrCode 	string
	Code 		string
	OpenId 		string
	Token		string
}

//--------------------------------------------
// user item
//--------------------------------------------
type UserItem struct {
	ItemId 		uint32
	Count 		int
}

type SystemSyncItems struct {
	Items 		[]*UserItem
}

//--------------------------------------------
// notice
//--------------------------------------------
type NoticeItem struct {
	Id 			int
	Kind 		string		//'notice', 'horn'
	Content 	string
	StartTime 	time.Time
	FinishTime 	time.Time
	Counts 		int
	PlayTime 	int
}

type LoadNoticeListReq struct {
}

type LoadNoticeListRet struct {
	List 		[]NoticeItem
}

type NoticeUpdateItem struct {
	Operation 		int		//1 add 2 update 3 delete
	Item 			NoticeItem
}

type NoticeUpdate struct {
	List 			[]NoticeUpdateItem
}

type UserHornMessageReq struct {
	Uid 		uint32
	Content 	string
}

type UserHornMessgaeRet struct {
	ErrCode 	int
	UserName 	string
	Item 		NoticeItem
}

//--------------------------------------------
// mall
//--------------------------------------------
type MallItem struct {
	Id 			int
	Name 		string
	Category 	int
	BuyValue 	int
	Nums 		int
	BuyLimt 	int
	StartTime, FinishTime int
}

type ClientBuyReq struct {
	Item 			int	// 购买的物品
	BuyType			int	// 购买类型
}

type ClientBuyMallItemRet struct {
	ErrCode 		int
	Item 			int	// 购买的物品
}

type ClientLoadMallList struct {

}

type ClientLoadMallListRet struct {
	Items 		[]MallItem
}

//--------------------------------------------
// item config
//--------------------------------------------
type ItemConfig struct {
	Itemid		uint32
	Itemname 	string
	Category 	int
	Nums 		int
	Sell 		int
	Buyvalue 	int
	GameKind 	int
	Description string
}

type ItemArea struct {
	Area 		int
	Gamelib 	int
}

//--------------------------------------------
// rank
//--------------------------------------------
type UserRankItem struct {
	Rank 		int
	Name 		string
	UserId 		int
	HeadImg 	string
	Value 		int64
}

type ClientLoadUserRank struct {
	RankType 	int
}

type ClientLoadUserRankRet struct {
	RankType 	int
	Users 		[]UserRankItem
}

//--------------------------------------------
// record
//--------------------------------------------
type RecordItem struct {
	RecordId 		int
	RecordData 		[]byte
}

type ClientGetRecordList struct {

}

type ClientGetRecordListRet struct {
	ErrCode 		int
	Records 		[]RecordItem
}

type ClientGetRecord struct {
	RecordId 		int
}

type ClientGetRecordRet struct {
	ErrCode 		int
	Content 		[]byte
}

//--------------------------------------------
// quest
//--------------------------------------------
type QuestItem struct {
	Id 			int
	Title 		string
	Content 	string
	Type 		string
	MaxCount 	int
	RewardIds 	string
}

type QuestData struct {
	Id 			int
	CurCount 	int
	TolCount 	int
	Status 		int 	// 1 正在进行  2 已完成  3 已领取
}

type QuestRewardItem struct {
	Id 			int
	ItemId 		int
	Num 		int
}

type ClientLoadQuest struct {

}

type ClientLoadQuestRet struct {
	ErrCode 		int
	Process 		[]*QuestData
}

type ClientProcessQuest struct {
	Id 				int
	CurCount 		int
}

type ClientProcessQuestRet struct {
	ErrCode 		int
	Id 				int
	CurCount 		int
	Status 			int		// 1 正在进行  2 已完成  3 已领取
}

type ClientCompleteQuest struct {
	Id 				int
}

type ClientCompleteQuestRet struct {
	ErrCode 		int
	Id 				int
	RewardId 		string
}

//--------------------------------------------
// activity
//--------------------------------------------
type ClientLoadActitity struct {

}

type ClientLoadActitityRet struct {
	Activities 		[]*ActivityItem
	Rewards 		[]*ActivityRewardItem
	OpenIds 		[]int
}

//--------------------------------------------
// identify
//--------------------------------------------
type ClientIdentify struct {
	Id 				string
	Phone 			string
	Name 			string
}

type ClientIdentifyRet struct {
	ErrCode 		int
}

type ClientLoadIdentify struct {

}

type ClientLoadIdentifyRet struct {

}

//--------------------------------------------
// club
//--------------------------------------------
type ClientCreateClub struct {

}

type ClientCreateClubRet struct {
	ErrCode 		int
}

type ClientJoinClub struct {
	ClubId 			int
}

type ClientJoinClubRet struct {
	ErrCode 		int
}

type ClientLeaveClub struct {
	ClubId 			int
}

type ClientLeaveClubRet struct {
	ErrCode 		int
}