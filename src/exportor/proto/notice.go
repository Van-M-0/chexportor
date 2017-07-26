package proto

import "time"

type NoticeItem struct {
	Id 			int
	Content 	string
	StartTime 	time.Time
	FinishTime 	time.Time
	Counts 		int
	PlayTime 	int
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