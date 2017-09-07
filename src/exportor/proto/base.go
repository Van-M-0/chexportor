package proto

const (
	CmdBaseUpsePropUpdate = 100
	CmdBaseSynceLoginItems = 101
	CmdBaseSynceIdentifyInfo = 102
	CmdBaseSynceClubInfo = 103
)

type UserProp struct {
	Ppkey 		int
	PpVal 		interface{}
}

type ItemProp struct {
	Flag 		int			//1 update(add/del) 2 remove 3 new
	ItemId 		uint32
	Count 		int
}

type SyncUserProps struct {
	Props 		*UserProp
	Items 		*ItemProp
}

type SynceIdentifyInfo struct {
	Name 		string
	Phone 		string
	Card 		string
}

type SyncClubInfo struct {
	ClubId 		int
	CreateorId 	int
	CreatorName string
}

type GameUserPpUpdate struct {
	Diamond		*int
	Score 		*int
	Gold 		*int64
	Item 		*ItemProp
}
