package databases

import (
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	user     string
	password string
	dbName   string
	host     string
	port     int
	database *mongo.Client
}

func (m *MongoDB) OpenConnection() error {
	var connectionError error

	authenticationOptions := options.Credential{
		Username:   m.user,
		Password:   m.password,
		AuthSource: m.dbName,
	}
	clientOptions := options.Client().ApplyURI(m.getDataSourceName()).SetAuth(authenticationOptions)
	m.database, connectionError = mongo.NewClient(clientOptions)
	return connectionError
}

func (m *MongoDB) CloseConnection() error {
	return m.database.Disconnect()
}

func (m *MongoDB) SetUser(_user string) {
	m.user = _user
}

func (m *MongoDB) GetUser() string {
	return m.user
}

func (m *MongoDB) SetPassword(_password string) {
	m.password = _password
}

func (m *MongoDB) GetPassword() string {
	return m.password
}

func (m *MongoDB) SetDatabase(_database string) {
	m.dbName = _database
}

func (m *MongoDB) GetDatabase() string {
	return m.dbName
}

func (m *MongoDB) SetHost(_host string) {
	m.host = _host
}

func (m *MongoDB) GetHost() string {
	return m.host
}

func (m *MongoDB) SetPort(_port int) {
	m.port = _port
}

func (m *MongoDB) GetPort() int {
	return m.port
}

func (m *MongoDB) getDataSourceName() string {
	var connectionRoute strings.Builder

	connectionRoute.WriteString("mongodb://")
	connectionRoute.WriteString(m.host)
	connectionRoute.WriteString(":")
	connectionRoute.WriteString(strconv.Itoa(m.port))
	return connectionRoute.String()
}
