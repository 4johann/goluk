package http

type Client interface {
	SendRequest(RequestMessage) error
	GetHost() string
	GetRequestTimeoutInSeconds() int
	SetHost(string)
	SetRequestTimeoutInSeconds(int)
}

type Server interface {
	CloseServer() error
	RunServer() error
}
