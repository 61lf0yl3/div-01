package database

import "database/sql"

// InitDB ..
func (d *Database) InitDB(dbName string) error {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return err
	}
	d.db = db
	return nil
}
