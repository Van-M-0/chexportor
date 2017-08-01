package proto

const (
	CmdBaseUpsePropUpdate = 100
)

type UserProp struct {
	Ppkey 		int
	PpVal 		interface{}
}

type SyncUserProps struct {
	Props 		UserProp
}
