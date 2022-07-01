package app

import (
	"database/sql"
	"fmt"

	"github.com/go-seidon/local/internal/repository"
	repository_mysql "github.com/go-seidon/local/internal/repository-mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_PROVIDER_MYSQL = "mysql"
	DB_PROVIDER_MONGO = "mongo"
)

func NewRepository(o RepositoryOption) (*NewRepositoryResult, error) {
	if o == nil {
		return nil, fmt.Errorf("invalid repository option")
	}

	var p NewRepositoryOption
	o.Apply(&p)

	if p.Provider == DB_PROVIDER_MYSQL {
		return newMySQLRepository(p)
	}

	return nil, fmt.Errorf("db provider is not supported")
}

func newMySQLRepository(p NewRepositoryOption) (*NewRepositoryResult, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		p.MySQLUser, p.MySQLPassword,
		p.MySQLHost, p.MySQLPort, p.MySQLDBName,
	)
	client, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	fileRepo, err := repository_mysql.NewFileRepository(client)
	if err != nil {
		return nil, err
	}

	r := &NewRepositoryResult{
		FileRepo: fileRepo,
	}
	return r, nil
}

type RepositoryOption interface {
	Apply(*NewRepositoryOption)
}

type NewRepositoryOption struct {
	Provider string

	MySQLHost     string
	MySQLPort     int
	MySQLUser     string
	MySQLPassword string
	MySQLDBName   string
}

type NewRepositoryResult struct {
	FileRepo repository.FileRepository
}

type mysqlRepositoryOption struct {
	host     string
	port     int
	username string
	password string
	dbName   string
}

func (o *mysqlRepositoryOption) Apply(p *NewRepositoryOption) {
	p.MySQLHost = o.host
	p.MySQLPort = o.port
	p.MySQLDBName = o.dbName
	p.MySQLUser = o.username
	p.MySQLPassword = o.password
	p.Provider = DB_PROVIDER_MYSQL
}

func WithMySQLRepository(username string, password string, dbName string, host string, port int) *mysqlRepositoryOption {
	return &mysqlRepositoryOption{
		username: username,
		password: password,
		dbName:   dbName,
		host:     host,
		port:     port,
	}
}
