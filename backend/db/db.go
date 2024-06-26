package db

import (
	"database/sql"
	"log"
)

func InitDB() {
	err := ExecStatement("CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT NOT NULL UNIQUE, password TEXT NOT NULL)")
	if err != nil {
		log.Fatal(err)
	}
	err = ExecStatement("CREATE TABLE IF NOT EXISTS device (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, mac_address TEXT NOT NULL UNIQUE, last_online DATETIME, user_id INTEGER, FOREIGN KEY (user_id) REFERENCES user(id))")
	if err != nil {
		log.Fatal(err)
	}

}

func ExecStatement(query string, args ...interface{}) error {
	driver, err := sql.Open("sqlite3", "./wol.db")
	if err != nil {
		return err
	}
	statement, err := driver.Prepare(query)
	if err != nil {
		defer driver.Close()
		return err
	}
	_, err = statement.Exec(args...)
	if err != nil {
		defer driver.Close()
		return err
	}
	defer driver.Close()
	return nil
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	driver, err := sql.Open("sqlite3", "./wol.db")
	if err != nil {
		defer driver.Close()
		return nil, err
	}
	rows, err := driver.Query(query, args...)
	if err != nil {
		defer driver.Close()
		return nil, err
	}
	defer rows.Close()
	return rows, nil
}
