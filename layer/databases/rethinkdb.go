package databases

import (
	"strconv"
	"strings"

	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type RethinkDB struct {
	user     string
	password string
	dbName   string
	host     string
	port     int
	database *r.Session
}

func (e *RethinkDB) OpenConnection() error {
	var connectionError error

	clientOptions := r.ConnectOpts{
		Address:  e.getDataSourceName(),
		Database: e.dbName,
		Username: e.user,
		Password: e.password,
	}
	e.database, connectionError = r.Connect(clientOptions)
	return connectionError
}

func (e *RethinkDB) CloseConnection() error {
	return e.database.Close()
}

func (e *RethinkDB) SetUser(_user string) {
	e.user = _user
}

func (e *RethinkDB) GetUser() string {
	return e.user
}

func (e *RethinkDB) SetPassword(_password string) {
	e.password = _password
}

func (e *RethinkDB) GetPassword() string {
	return e.password
}

func (e *RethinkDB) SetDatabase(_database string) {
	e.dbName = _database
}

func (e *RethinkDB) GetDatabase() string {
	return e.dbName
}

func (e *RethinkDB) SetHost(_host string) {
	e.host = _host
}

func (e *RethinkDB) GetHost() string {
	return e.host
}

func (e *RethinkDB) SetPort(_port int) {
	e.port = _port
}

func (e *RethinkDB) GetPort() int {
	return e.port
}

func (e *RethinkDB) getDataSourceName() string {
	var connectionRoute strings.Builder

	connectionRoute.WriteString(e.host)
	connectionRoute.WriteString(":")
	connectionRoute.WriteString(strconv.Itoa(e.port))
	return connectionRoute.String()
}
