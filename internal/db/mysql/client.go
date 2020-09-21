package mysql

import (
	"fmt"
	"log"

	"database/sql"

	migrate "github.com/golang-migrate/migrate/v4"
	mysql_migrate "github.com/golang-migrate/migrate/v4/database/mysql"

	// needed by migrate
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// driver for mysql
	_ "github.com/go-sql-driver/mysql"
)

type Client struct {
	db *sql.DB
}

func NewClient(driverName, host string, port int, user, password, dbName string) (*Client, error) {
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
