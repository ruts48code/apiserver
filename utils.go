package main

import (
	"log"
	"strings"
	"time"

	dbs "github.com/ruts48code/dbs4ruts"
	otp "github.com/ruts48code/otp4ruts"
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

func CheckOTP(otptxt string) bool {
	return otp.ChkTimeOTPxHex([]byte(conf.OTP.Key), conf.OTP.Size, otptxt, conf.OTP.Interval)
}

func SaveCache(domain string, data string) {
	db, err := dbs.OpenDBS(conf.DBS)
	if err != nil {
		return
	}
	defer db.Close()

	_, err = db.Exec("Insert into cache (domain,data,timestamp) values (?,?,?);", domain, data, utils.GetTimeStamp(time.Now()))
	if err != nil {
		log.Printf("Error: cannot write cache domain %s : %v\n", domain, err)
	}
}
