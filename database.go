package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"net"
)

type Database struct {
	connection *sql.DB
}

func OpenDatabase() (*Database, error) {
	config, err := GetConfig()

	if err != nil {
		return nil, err
	}

	var dsnConfig = mysql.Config{}

	dsnConfig.Addr = net.JoinHostPort(config.Database.Username, config.Database.Port)

	db, err := sql.Open("mysql", dsnConfig.FormatDSN())

	if err != nil {
		return nil, err
	}
	return &Database{connection: db}, nil
}

func (d *Database) GetVersion() (string, error) {
	stmt, err := d.connection.Prepare("SELECT metaValue FROM _meta WHERE metaKey='version'")
	if err != nil {
		return "", err
	}

	defer stmt.Close()

	var version string

	err = stmt.QueryRow().Scan(&version)
	if err == sql.ErrNoRows {
		version = "Not Installed"
		err = nil
	}

	return version, err
}
