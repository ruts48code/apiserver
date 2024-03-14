package main

import (
	"strings"

	utils "github.com/ruts48code/utils4ruts"
)

func CheckTKWeb(token string) (username, name, email, status string) {
	data := ChkToken(token)
	if data.Status != "ok" {
		status = "token"
		return
	}
	status = "ok"
	usernamex := strings.Split(utils.NormalizedEloginToken(token), ":")
	username = usernamex[0]
	name = data.Name
	email = data.Email
	return
}
