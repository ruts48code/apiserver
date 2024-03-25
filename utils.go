package main

import (
	"log"
	"strings"
	"time"

	dbs "github.com/ruts48code/dbs4ruts"
	otp "github.com/ruts48code/otp4ruts"
	random "github.com/ruts48code/random4ruts"
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
		log.Printf("Error: utils-SaveCache 1 - cannot connect to database - %v\n", err)
		return
	}
	defer db.Close()

	ts := utils.GetTimeStamp(time.Now())
	_, err = db.Exec("Insert into cache (domain,data,timestamp) values (?,?,?);", domain, data, ts)
	if err != nil {
		log.Printf("Error: utils-SaveCache 2 - cannot write cache domain %s : %v\n", domain, err)
	}
}

func CleanCache(domain string) {
	db, err := dbs.OpenDBS(conf.DBS)
	if err != nil {
		log.Printf("Error: utils-CleanCache 1 - cannot connect to database - %v\n", err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select timestamp from cache where domain=? order by timestamp desc limit 1;", domain)
	if err != nil {
		log.Printf("Error: utils-CleanCache 2 - cannot query cache domain %s : %v\n", domain, err)
		return
	}
	defer rows.Close()

	t := time.Now()
	for rows.Next() {
		rows.Scan(&t)
	}
	ts := utils.GetTimeStamp(t)
	_, err = db.Exec("delete from cache where domain=? and timestamp < ?", domain, ts)
	if err != nil {
		log.Printf("Error: utils-CleanCache 3 - cannot delete cache domain %s : %v\n", domain, err)
	}
}

func ReadCache(domain string) (output string) {
	db, err := dbs.OpenDBS(conf.DBS)
	if err != nil {
		log.Printf("Error: utils-ReadCache 1 - cannot connect database : %v\n", err)
		return ""
	}
	defer db.Close()

	rows, err := db.Query("select data from cache where domain=? order by timestamp desc limit 1;", domain)
	if err != nil {
		log.Printf("Error: utils-ReadCache 2 - cannot query cache domain %s : %v\n", domain, err)
		return ""
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&output)
	}
	return
}

func NumLimitElogin(username string) (out int) {
	db, err := dbs.OpenDBS(conf.DBS)
	if err != nil {
		log.Printf("Error: utils-NumLimitElogi 1 - cannot connect database : %v\n", err)
		return -1
	}
	defer db.Close()

	rows, err := db.Query("select count(*) from elogin_limit where username=?;", username)
	if err != nil {
		log.Printf("Error: utils-NumLimitElogi 2 - cannot query NumLimitElogin for user %s : %v\n", username, err)
		return -1
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&out)
	}
	return
}

func SaveLimitLogin(username string) (out string) {
	db, err := dbs.OpenDBS(conf.DBS)
	if err != nil {
		log.Printf("Error: utils-SaveLimitLogin 1 - cannot connect to database - %v\n", err)
		return
	}
	defer db.Close()
	out = random.GetRandomString(conf.Elogin.TokenSize)

	_, err = db.Exec("Insert into elogin_limit (username,token) values (?,?);", username, out)
	if err != nil {
		log.Printf("Error: utils-SaveLimitLogin 2 - cannot insert NumLimitElogin for user %s : %v\n", username, err)
	}
	return
}

func DeleteLimitLogin(token string) {
	db, err := dbs.OpenDBS(conf.DBS)
	if err != nil {
		log.Printf("Error: utils-DeleteLimitLogin 1 - cannot connect to database - %v\n", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("Delete from elogin_limit where token=?;", token)
	if err != nil {
		log.Printf("Error: utils-DeleteLimitLogin 2 - cannot delete NumLimitElogin for token %s : %v\n", token, err)
	}
}
