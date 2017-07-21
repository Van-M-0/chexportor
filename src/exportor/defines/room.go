package defines

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
	Sex 		byte

	RoomId 		uint32
}

type IRoomManager interface {

}

type IRoom interface {

}

type IPlayer interface {
	GetPlayerInfo()	PlayerInfo
}
