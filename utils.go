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
		log.Printf("Error: utils-SaveCache 1 - cannot connect to database - %v\n", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("Insert into cache (domain,data,timestamp) values (?,?,?);", domain, data, utils.GetTimeStamp(time.Now()))
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
	log.Printf("Domain = %s - Timestamp = %s\n", domain, utils.GetTimeStamp(t))
	_, err = db.Exec("delete from cache where domain=? and timestamp < ?", domain, utils.GetTimeStamp(t))
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
