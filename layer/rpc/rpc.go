package rpc

type RPC interface {
	CreateServer() error
	OpenConnection() error
	SetHost(_host_ string)
	SetPort(_port_ int)
	GetHost() string
	GetPort() int
}
