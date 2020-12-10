package database

import (
	"database/sql"
	"fmt"
)

// InitDB ..
func (d *Database) InitDB(dbName string) error {

	db, err := sql.Open("sqlite3", dbName)
	fmt.Println(d.DB)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	d.DB = db
	return nil
}
