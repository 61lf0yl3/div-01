package models

// User ...
type User struct {
	ID       int64
	Username string
	Email    string
	Password string
}

// NewUser ...
func NewUser() *User {
	return &User{}
}
