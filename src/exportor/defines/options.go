package defines

import (
	"net"
)

type CodecCreator func() ICodec
type NetClientOption struct {
	Host       	string
	ConnectCb  	ClientConnectCb
	CloseCb    	ClientCloseCb
	MsgCb      	ClientMessageCb
	AuthCb     	AuthCb
	SendChSize 	int
	SendActor 	int
}

type NetServerOption struct {
	GwHost		string
	CmHost 		string
	Host 		string
	SendChSize 	int
	SendActor 	int
	RecvNum		int

	MaxRecvSize int
	EncryptCode string
	ConnectCb   ClientConnectCb
	CloseCb     ClientCloseCb
	MsgCb      	ClientMessageCb
	AuthCb      AuthCb
	listenConn  net.Conn
}

type GatewayOption struct {
	ClientHost 	string
	FrontHost 	string
	MaxClient   int

	BackHost 	string
}

type LobbyOption struct {
	GwHost 		string
}

type GameOption struct {
	GwHost 		string
	Moudles 	[]GameModule
}

type DbProxyOption struct {
	Name 		string
	User 		string
	Pwd 		string
}

type WorldOptoin struct {
	Name 		string
	User 		string
	Pwd 		string
}

type GameCreateor func() IGame
type GameReleaser func(IGame)
type GameModule struct {
	Type 		int			//游戏id
	GameType 	int			//游戏种类
	Creator 	GameCreateor
	Releaser 	GameReleaser
	GameData 	interface{}
	GameConf	[]byte
	PlayerCount int
}

type GameModuleConfig struct {
	Enabled 	bool 		//是否开启
	Desc 		string		//描述
	Kind 		int			//游戏id
	GameType 	int			//游戏种类
	Cfg 		interface{}
}

type CommunicatorOption struct {
	Host 			string
	ReadTimeout 	int
	WriteTimeout	int
}

type StartConfigFile struct {

	ClusterId		int

	WorldHost		string

	ClientVisitIp 	string
	LocalIp 		string

	//ClientHost 		string

	FrontHost 		string
	BackendHost 	string
	HttpHost 		string

	DbName 			string
	DbUser 			string
	DbPwd 			string

	WorldService 	string
	DBService 		string
	MSservice 		string
	LobbyService 	string
}
