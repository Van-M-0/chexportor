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
	FrontHost 	string
	MaxClient   int

	BackHost 	string
}

type LobbyOption struct {
	GwHost 		string
}

type GameOption struct {
	ClientHost 	string
	GwHost 		string
	Moudles 	[]GameModule
}

type GameCreateor func() IGame
type GameReleaser func(IGame)
type GameModule struct {
	Type 		int
	Creator 	GameCreateor
	Releaser 	GameReleaser
	GameData 	interface{}
	GameConf	[]byte
	PlayerCount int
}

type CommunicatorOption struct {
	Host 			string
	ReadTimeout 	int
	WriteTimeout	int
}

type StartConfigFile struct {
	FrontHost 		string
	BackendHost 	string
	HttpHost 		string
	GameModules 	[]int
	LocalHost		string
	WorldHttp 		string
	WorldHost 		string
}