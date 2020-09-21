package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/golang-migrate/migrate/v4"
	mysql_migrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Client struct {
	db *sql.DB
}

func NewClient(driverName string, host string, port int, user string, password string, dbName string) (*Client, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName)

	db, err := sql.Open(driverName, connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return &Client{db: db}, nil
}

func (cli *Client) Migrate() error {
	if err := cli.db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql_migrate.WithInstance(cli.db, &mysql_migrate.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations/mysql",
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
