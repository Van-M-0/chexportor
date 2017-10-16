package defines

import (
	"exportor/proto"
)

type PlayerInfo struct {
	Uid 		uint32
	UserId 		uint32
	OpenId 		string
	HeadImg 	string
	Name 		string
	Account		string
	Diamond 	int
	Gold 		int64
	RoomCard 	int
	Score 		int
	Sex 		byte

	RoomId 		uint32

	Items 		[]*proto.UserItem
}

type CreateRoomConf struct {
	Type		int
	RoomId 		uint32
	Conf 		[]byte
}

type TimerHandle interface {
	GetId() int
}

type IGameManager interface {
	ReleaseRoom()
	GetRoomId() uint32
	SendGameMessage(info *PlayerInfo, cmd uint32, data interface{})
	BroadcastMessage(infos []*PlayerInfo, cmd uint32, data interface{})
	SetTimer(t uint32, data interface{}) TimerHandle
	KillTimer(handle TimerHandle) bool
	SaveGameRecord(head, data []byte) int
	SaveUserRecord(userid, id int) error
	UpdateUserInfo(info *PlayerInfo, data *proto.GameUserPpUpdate) bool
	InfoLog(args ...interface{})
	DebugLog(args ...interface{})
	ErrLog(args ...interface{})
}

type IGame interface {
	OnInit(manager IGameManager, config GameModule) error
	OnRelease()
	OnGameCreate(info *PlayerInfo, conf *CreateRoomConf) error
	OnUserEnter(info *PlayerInfo) error
	OnUserLeave(info *PlayerInfo)
	OnUserOffline(info *PlayerInfo)
	OnUserReEnter(info *PlayerInfo)
	OnUserMessage(info *PlayerInfo, cmd uint32, data []byte) error
	OnTimer(id uint32, data interface{})
	GetPlayerCount() int
}

const (
	GameTypeInvalid			= iota
	GameTypeRoomCard
	GameTypeCoin
)