package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var sqlDriver = "sqlite3"

var dbCon = "./wol.db"

func InitDB() (db *sql.DB, err error) {
	driver, err := sql.Open(sqlDriver, dbCon)
	if err != nil {
		return nil, err
	}
	err = ExecStatement(driver, "CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT NOT NULL UNIQUE, password TEXT NOT NULL)")
	if err != nil {
		return nil, err
	}
	err = ExecStatement(driver, "CREATE TABLE IF NOT EXISTS device (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, mac_address TEXT NOT NULL UNIQUE, last_online DATETIME, user_id INTEGER, FOREIGN KEY (user_id) REFERENCES user(id))")
	if err != nil {
		return nil, err
	}
	return driver, nil

}

func ExecStatement(driver *sql.DB, query string, args ...interface{}) error {
	statement, err := driver.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(args...)
	if err != nil {
		return err
	}
	return nil
}

func Query(driver *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := driver.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func QueryOne(driver *sql.DB, query string, args ...interface{}) (*sql.Row, error) {
	row := driver.QueryRow(query, args...)
	return row, nil
}
