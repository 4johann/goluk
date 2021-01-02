package databases

type Database interface {
	OpenConnection() error
	CloseConnection() error
	SetUser(_user string)
	SetPassword(_password string)
	SetDatabase(_database string)
	SetHost(_host string)
	SetPort(_port int)
	GetUser() string
	GetPassword() string
	GetDatabase() string
	GetHost() string
	GetPort() int
}
