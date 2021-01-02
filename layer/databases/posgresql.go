package databases

import (
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgreSQL struct {
	user     string
	password string
	dbName   string
	host     string
	port     int
	database *gorm.DB
}

func (p *PostgreSQL) OpenConnection() error {
	var connectionError error

	p.database, connectionError = gorm.Open("postgres", p.getDataSourceName())
	return connectionError
}

func (p *PostgreSQL) CloseConnection() error {
	return p.database.Close()
}

func (p *PostgreSQL) MigrateTables(tables ...interface{}) {
	p.database.AutoMigrate(tables...)
}

func (p *PostgreSQL) AddForeignKey(model interface{}, field string, destination string) {
	p.database.Model(model).AddForeignKey(field, destination, "RESTRICT", "RESTRICT")
}

func (p *PostgreSQL) DropColumn(model interface{}, columnName string) {
	p.database.Model(model).DropColumn(columnName)
}

func (p *PostgreSQL) Insert(data interface{}) error {
	return p.database.Create(data).Error
}

func (p *PostgreSQL) Find(model interface{}, filter interface{}) {
	p.database.Where(filter).Find(model)
}

func (p *PostgreSQL) FindOne(model interface{}, filter interface{}) {
	p.database.Order("updated_at desc").Where(filter).Limit(1).First(model)
}

func (p *PostgreSQL) Update(model interface{}, where interface{}, data interface{}) error {
	return p.database.Model(model).Where(where).Update(data).Error
}

func (p *PostgreSQL) Delete(model interface{}, where interface{}) error {
	return p.database.Unscoped().Delete(model, where).Error
}

func (p *PostgreSQL) SetUser(_user string) {
	p.user = _user
}

func (p *PostgreSQL) GetUser() string {
	return p.user
}

func (p *PostgreSQL) SetPassword(_password string) {
	p.password = _password
}

func (p *PostgreSQL) GetPassword() string {
	return p.password
}

func (p *PostgreSQL) SetDatabase(_database string) {
	p.dbName = _database
}

func (p *PostgreSQL) GetDatabase() string {
	return p.dbName
}

func (p *PostgreSQL) SetHost(_host string) {
	p.host = _host
}

func (p *PostgreSQL) GetHost() string {
	return p.host
}

func (p *PostgreSQL) SetPort(_port int) {
	p.port = _port
}

func (p *PostgreSQL) GetPort() int {
	return p.port
}

func (p *PostgreSQL) getDataSourceName() string {
	var connectionRoute strings.Builder

	connectionRoute.WriteString("host=")
	connectionRoute.WriteString(p.host)
	connectionRoute.WriteString(" port=")
	connectionRoute.WriteString(strconv.Itoa(p.port))
	connectionRoute.WriteString(" user=")
	connectionRoute.WriteString(p.user)
	connectionRoute.WriteString(" dbname=")
	connectionRoute.WriteString(p.dbName)
	connectionRoute.WriteString(" password=")
	connectionRoute.WriteString(p.password)
	connectionRoute.WriteString(" sslmode=disable")
	return connectionRoute.String()
}
