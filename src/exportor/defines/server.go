package defines

import (
	"exportor/proto"
	"time"
)

type IServer interface {
	Start() error
	Stop() error
}

type IGateway interface {
	IServer
}

type ILobby interface {
	IServer
}

type IGameServer interface {
	IServer
}

type IDbProxy interface {
	IServer
}

type CommunicatorCb func([]byte)
type ICommunicatorClient interface {
	Notify(chanel string, v interface{})	error
	JoinChanel(chanel string, reg bool, time int, cb CommunicatorCb) error
	WaitChannel(channel string, time int) ([] byte, error)
}

type IMsgPublisher interface {
	IServer
	WaitPublish(channel string, key string, data interface{}) error
	SendPublish(channel string, data interface{}) error
}

type IMsgConsumer interface {
	IServer
	WaitMessage(channel string, key string, t time.Duration) interface{}
	GetMessage(channel string, key string) interface{}
}

type ICommunicator interface {
	WaitPublish(channel string, key string, data interface{}) error
	SendPublish(channel string, data interface{}) error
	WaitMessage(channel string, key string, t time.Duration) interface{}
	GetMessage(channel string, key string) interface{}
}

type ICacheClient interface {
	IServer
	GetUserInfo(name string, user *proto.CacheUser) error
	GetUserInfoById(uid uint32, user *proto.CacheUser) error
	SetUserInfo(d interface{}, dbRet bool) error
	SetUserCidUserId(uid uint32, userId int) error
	GetUserCidUserId(uid uint32) int
	DelUserCidUserId(uid uint32)
	UpdateUserInfo(uid uint32, prop int, value interface{}) bool
	SetServer(server *proto.CacheServer) error
	GetServers() ([]*proto.CacheServer, error)
	UpdateServer(server *proto.CacheServer) error
	NoticeOperation(notice *[]*proto.CacheNotice, op string) error
	UpdateUserItems(userid uint32, items []proto.UserItem) error
	UpdateSingleItem(userid uint32, flag int, id uint32, count int) error
	GetUserItems(userid uint32) ([]*proto.UserItem, error)
	FlushAll()
	Scripts(args ...interface{})

	GetAllUsers() ([]*proto.CacheUser, error)
	GetAllUserItem() ([]*proto.CacheUserItem, error)

	SaveGameRecord(head, content []byte) int
	SaveUserRecord(userId, recordId int) error
	GetGameRecordHead(userId int) (map[int][]byte, error)
	GetGameRecordContent(id int) ([]byte, error)
}

type ICacheLoader interface {
	LoadUsers(name string)
}
