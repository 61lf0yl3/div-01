package database

// InsertData ...
func (d *Database) InsertData() {
	d.DB.Exec()
	//d.db.Exec()
}

// GetUserBySession ...
func (d *Database) GetUserBySession(session string) error {
	return nil
}
