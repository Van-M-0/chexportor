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
)

type ClientLogin struct {
	Account 	string
}

type GuestLogin struct {
	Account 	string
}

type ClientLoginRet struct {
	ErrCode 	int
	Uid 		uint32
	Account 	string
	Name 		string
	UserId 		uint32
	Diamond 	int
	Gold 		int64
	RoomCard 	int
	Score 		int
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
	Account 	string
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
	ItemId 			int
	BuyType			int
}

type ClientBuyMallItemRet struct {
	ErrCode 		int
}

type ClientLoadMallList struct {

}

type ClientLoadMallListRet struct {
	Items 		[]MallItem
}