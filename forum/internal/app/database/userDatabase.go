package database

import (
	"fmt"

	"github.com/61lf0yl3/div-01/forum/internal/models"
)

// AddUser ...
func (d *Database) AddUser(u *models.User) error {
	stmnt, err := d.DB.Prepare("INSERT INTO users (firstname, lastname, username, email, password) VALUES (?, ?, ?, ?, ?)")
	_, err = stmnt.Exec(u.Username, u.Email, u.Password)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
