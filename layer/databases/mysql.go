package databases

import (
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MySQL struct {
	user     string
	password string
	dbName   string
	host     string
	port     int
	database *gorm.DB
}

func (m *MySQL) OpenConnection() error {
	var connectionError error

	m.database, connectionError = gorm.Open("mysql", m.getDataSourceName())
	return connectionError
}

func (m *MySQL) CloseConnection() error {
	return m.database.Close()
}

func (m *MySQL) MigrateTables(tables ...interface{}) {
	m.database.AutoMigrate(tables...)
}

func (m *MySQL) AddForeignKey(model interface{}, field string, destination string) {
	m.database.Model(model).AddForeignKey(field, destination, "RESTRICT", "RESTRICT")
}

func (m *MySQL) DropColumn(model interface{}, columnName string) {
	m.database.Model(model).DropColumn(columnName)
}

func (m *MySQL) Insert(data interface{}) error {
	return m.database.Create(data).Error
}

func (m *MySQL) Find(model interface{}, filter interface{}) {
	m.database.Where(filter).Find(model)
}

func (m *MySQL) FindOne(model interface{}, filter interface{}) {
	m.database.Order("updated_at desc").Where(filter).Limit(1).First(model)
}

func (m *MySQL) Update(model interface{}, where interface{}, data interface{}) error {
	return m.database.Model(model).Where(where).Update(data).Error
}

func (m *MySQL) Delete(model interface{}, where interface{}) error {
	return m.database.Unscoped().Delete(model, where).Error
}

func (m *MySQL) SetUser(_user string) {
	m.user = _user
}

func (m *MySQL) GetUser() string {
	return m.user
}

func (m *MySQL) SetPassword(_password string) {
	m.password = _password
}

func (m *MySQL) GetPassword() string {
	return m.password
}

func (m *MySQL) SetDatabase(_database string) {
	m.dbName = _database
}

func (m *MySQL) GetDatabase() string {
	return m.dbName
}

func (m *MySQL) SetHost(_host string) {
	m.host = _host
}

func (m *MySQL) GetHost() string {
	return m.host
}

func (m *MySQL) SetPort(_port int) {
	m.port = _port
}

func (m *MySQL) GetPort() int {
	return m.port
}

func (m *MySQL) getDataSourceName() string {
	var concatenateString strings.Builder

	concatenateString.WriteString(m.user)
	concatenateString.WriteString(":")
	concatenateString.WriteString(m.password)
	concatenateString.WriteString("@(")
	concatenateString.WriteString(m.host)
	concatenateString.WriteString(":")
	concatenateString.WriteString(strconv.Itoa(m.port))
	concatenateString.WriteString(")/")
	concatenateString.WriteString(m.dbName)
	concatenateString.WriteString("?charset=utf8&parseTime=True&loc=Local")

	return concatenateString.String()
}
