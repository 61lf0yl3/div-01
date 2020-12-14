package handlers

import (
	"regexp"

	"github.com/61lf0yl3/div-01/forum/internal/models"
)

var (
	regName     = regexp.MustCompile("^([A-Za-z])+$")
	regUsername = regexp.MustCompile("^[a-zA-Z0-9]{3,20}$")
	regEmail    = regexp.MustCompile("^[a-zA-Z0-9.]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$")
	regPwd      = regexp.MustCompile("^[#\\w$&+,:;=?@#|'<>.-^*()%!]{8,20}")
)

func ValidateInput(user *models.User, confirmpwd string) string {

	if !regUsername.MatchString(user.Username) {
		return user.Username
	}
	if !regEmail.MatchString(user.Email) {
		return user.Email
	}
	if !regPwd.MatchString(user.Password) {
		return user.Password
	}

	return "filled correctly"
}
