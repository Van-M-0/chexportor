package proto

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

type ClientLoadMallItem struct {

}

type ClientLoadMallItemRet struct {
	Items 		[]*MallItem
}