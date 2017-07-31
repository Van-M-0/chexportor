package proto

// account -> userid
type CacheUserId struct {
	Uid 		int
}

// userid -> cacheuser
type CacheUser struct {
	Account 	string		`redis:"account"`
	Openid 		string		`redis:"openid"`
	Pwd 		string		`redis:"pwd"`
	Uid 		int			`redis:"uid"`
	Name 		string		`redis:"name"`
	Sex 		byte		`redis:"sex"`
	HeadImg 	string		`redis:"headimg"`
	Diamond 	int			`redis:"diamond"`
	RoomCard 	int			`redis:"roomcard"`
	Gold 		int64		`redis:"gold"`
	RoomId 		int 		`redis:"roomid"`
}

type CacheServer struct {
	Type 			string		`redis:"type,omitempty"`
	Id 				int 		`redis:"id,omitempty"`
	OnlineCount 	int			`redis:"onlinecount,omitempty"`
}

type CacheNotice struct {
	Id 				int 		`redis:"id,omitempty"`
	StartTime 		string		`redis:"starttime,omitempty"`
	FinishTime 		string	 	`redis:"finishtime,omitempty"`
	Kind 			string 		`redis:"kind,omitempty"`
	Content 		string 		`redis:"content,omitempty"`
	PlayTime 		int			`redis:"playtime,omitempty"`
	PlayCount 		int			`redis:"playcount,omitempty"`
}