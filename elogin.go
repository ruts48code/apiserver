package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	crypt "github.com/ruts48code/crypt4ruts"
	dbs "github.com/ruts48code/dbs4ruts"
	random "github.com/ruts48code/random4ruts"
	utils "github.com/ruts48code/utils4ruts"
)

type (
	LoginStruct struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	TokenStruct struct {
		Token string `json:"token"`
	}

	UserStruct struct {
		Status       string `json:"status"`
		CID          string `json:"cid"`
		Username     string `json:"username"`
		Name         string `json:"name"`
		FirstName    string `json:"firstname"`
		LastName     string `json:"lastname"`
		Type         string `json:"type"`
		FacCode      string `json:"faccode"`
		FacName      string `json:"facname"`
		DepCode      string `json:"depcode"`
		DepName      string `json:"depname"`
		SecCode      string `json:"seccode"`
		SecName      string `json:"secname"`
		Email        string `json:"email"`
		ChiefCode    string `json:"chiefcode"`
		ChiefName    string `json:"chiefname"`
		ChiefFacCode string `json:"chieffaccode"`
		ChiefFacName string `json:"chieffacname"`
		Token        string `json:"token"`
	}

	UserDB struct {
		Username     string `json:"username"`
		Password     string `json:"password"`
		CID          string `json:"cid"`
		Name         string `json:"name"`
		FirstName    string `json:"firstname"`
		LastName     string `json:"lastname"`
		FacCode      string `json:"faccode"`
		FacName      string `json:"facname"`
		DepCode      string `json:"depcode"`
		DepName      string `json:"depname"`
		SecCode      string `json:"seccode"`
		SecName      string `json:"secname"`
		Email        string `json:"email"`
		ChiefCode    string `json:"chiefcode"`
		ChiefName    string `json:"chiefname"`
		ChiefFacCode string `json:"chieffaccode"`
		ChiefFacName string `json:"chieffacname"`
	}
)

func eloginToken(ctx *fiber.Ctx) error {
	token := ctx.Params("token")
	return ctx.JSON(ChkToken(token))
}

func eloginDelete(ctx *fiber.Ctx) error {
	username := utils.NormalizeUsername(ctx.Params("username"))
	go DeleteLoginDatabase(username)
	return ctx.SendString("ok")
}

func eloginClean(ctx *fiber.Ctx) error {
	switch CheckOTP(ctx.Params("otp")) {
	case true:
		go CleanTokenElogin()
		return ctx.JSON(fiber.Map{
			"status": "ok",
		})
	default:
		return ctx.JSON(fiber.Map{
			"status": "otp",
		})
	}
}

func elogin(ctx *fiber.Ctx) error {
	data := LoginStruct{}
	err := json.Unmarshal(ctx.Body(), &data)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status": "json",
		})
	}

	username := utils.NormalizeUsername(data.Username)
	password := data.Password

	if username == "" || password == "" {
		return ctx.JSON(UserStruct{
			Status: "password",
		})
	}

	numlimit := NumLimitElogin(username)
	if numlimit > conf.Elogin.Limit || numlimit == -1 {
		log.Printf("Error: Elogin Limit for %s\n", username)
		return ctx.JSON(fiber.Map{
			"status": "limit",
		})
	}

	tk := SaveLimitLogin(username)
	defer DeleteLimitLogin(tk)

	result := ChkLoginDatabase(username, password)
	if result.Status == "ok" {
		result.Token = getToken(username, result)
		return ctx.JSON(result)
	}

	result = ChkLoginLDAP(username, password)
	if result.Status != "ok" {
		time.Sleep(3 * time.Second)
		return ctx.JSON(result)
	}

	switch result.Type {
	case "staff":
		result = getDataStaff(username, true)
	case "student":
		result = getDataStudent(username, true)
	}
	DeleteLoginDatabase(username)
	CreateLoginDatabase(result, password)
	return ctx.JSON(result)
}

func ChkToken(tokenx string) (output UserStruct) {
	token := utils.NormalizedEloginToken(tokenx)
	db, err := dbs.OpenDBS(conf.DBS)
	if err != nil {
		log.Printf("Error: elogin-ChkToken - Check token database for %s - %v\n", token, err)
		output.Status = "database"
		return
	}
	defer db.Close()

	ts := utils.GetTimeStamp(time.Now())
	_, err = db.Exec("UPDATE token SET timestamp=? WHERE token=?;", ts, token)
	if err != nil {
		log.Printf("Error: elogin-ChkToken - Update check token for %s - %v\n", token, err)
		output.Status = "database"
		return
	}

	rows, err := db.Query("SELECT name,firstname,lastname,username,faccode,facname,depcode,depname,seccode,secname,email,cid,chiefcode,chiefname,chieffaccode,chieffacname FROM token WHERE token=?;", token)
	if err != nil {
		log.Printf("Error: elogin-ChkToken - Query check token for %s - %v\n", token, err)
		output.Status = "database"
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&output.Name, &output.FirstName, &output.LastName, &output.Username, &output.FacCode, &output.FacName, &output.DepCode, &output.DepName, &output.SecCode, &output.SecName, &output.Email, &output.CID, &output.ChiefCode, &output.ChiefName, &output.ChiefFacCode, &output.ChiefFacName)
		output.Status = "ok"
		output.Token = token
		output.Type = utils.CheckEpassportType(output.Username)
		break
	}
	if output.Status == "" {
		output.Status = "token"
		log.Printf("Log: No token found for %s\n", token)
	}

	return
}

func DeleteLoginDatabase(username string) {
	db, err := dbs.OpenDBS(conf.DBS)
	if err != nil {
		log.Printf("Error: elogin-DeleteLoginDatabase - Delete login database for %s - %v\n", username, err)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM elogin WHERE username=?;", username)
	if err != nil {
		log.Printf("Error: elogin-DeleteLoginDatabase - Delete login database for %s - %v\n", username, err)
	}
}

func CreateLoginDatabase(result UserStruct, password string) {
	db, err := dbs.OpenDBS(conf.DBS)
	if err != nil {
		log.Printf("Error: elogin-CreateLoginDatabase 1 - Create login database for %s - %v\n", result.Username, err)
		return
	}
	defer db.Close()

	salt := random.GetRandomString(4)
	passx := crypt.MooHash([]byte(password), []byte(result.Username), []byte(salt))

	_, err = db.Exec("INSERT INTO elogin (username, password, name, firstname, lastname, faccode, facname, depcode, depname, seccode, secname, email, cid, chiefcode, chiefname, chieffaccode, chieffacname) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);", result.Username, passx, result.Name, result.FirstName, result.LastName, result.FacCode, result.FacName, result.DepCode, result.DepName, result.SecCode, result.SecName, result.Email, result.CID, result.ChiefCode, result.ChiefName, result.ChiefFacCode, result.ChiefFacName)
	if err != nil {
		log.Printf("Error: elogin-CreateLoginDatabase 2 - Insert login database for %s - %v\n", result.Username, err)
	}
}

func getDataStaff(username string, token bool) (output UserStruct) {
	db, err := dbs.OpenDB(conf.Personal.Server)
	if err != nil {
		log.Printf("Error: elogin-getDataStaff 1 - Get data staff for %s - %v\n", username, err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT STF_FNAME, STF_LNAME, FACULTY_CODE, FacName, department_code, depname, section_code, section_tname, CITIZEN_ID FROM vUOC_STAFF_L01 WHERE USERNAME_CISCO=?", username)
	if err != nil {
		log.Printf("Error: elogin-getDataStaff 2 - Query data staff for %s - %v\n", username, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&output.FirstName, &output.LastName, &output.FacCode, &output.FacName, &output.DepCode, &output.DepName, &output.SecCode, &output.SecName, &output.CID)
		output.Status = "ok"
		output.Username = username
		output.Name = output.FirstName + " " + output.LastName
		output.Type = utils.CheckEpassportType(username)
		output.Email = username + "@rmutsv.ac.th"
		break
	}

	rows2, err2 := db.Query("SELECT CHIEF_CODE, CHIEF_NAME, FACULTY_CODE, Facname from vADMIN WHERE USERNAME_CISCO=?;", username)
	if err2 != nil {
		log.Printf("Error: elogin-getDataStaff 3 - Query data staff admin for %s - %v\n", username, err)
		return
	}
	defer rows2.Close()

	for rows2.Next() {
		rows2.Scan(&output.ChiefCode, &output.ChiefName, &output.ChiefFacCode, &output.ChiefFacName)
		break
	}
	if token {
		output.Token = getToken(username, output)
	}
	return
}

func ChkLoginLDAP(username string, password string) (output UserStruct) {
	result := ldapLogin(username, password)
	if result == "none" {
		log.Printf("Log: elogin-ChkLoginLDAP - password fail in LDAP for %s\n", username)
		output.Status = "password"
		return
	}
	output.Status = "ok"
	output.Type = utils.CheckEpassportType(username)
	return
}

func ChkLoginDatabase(username string, password string) (output UserStruct) {
	output.Status = "fail"
	data := getUsernameDatabase(username)
	if data.Username == "" {
		log.Printf("Log: elogin-ChkLoginDatabase - username is empy string\n")
		return
	}

	salt := data.Password[:4]
	passx := crypt.MooHash([]byte(password), []byte(username), []byte(salt))
	if passx == data.Password {
		output.Status = "ok"
		output.Username = username
		output.Name = data.Name
		output.FirstName = data.FirstName
		output.LastName = data.LastName
		output.Type = utils.CheckEpassportType(username)
		output.FacCode = data.FacCode
		output.FacName = data.FacName
		output.DepCode = data.DepCode
		output.DepName = data.DepName
		output.SecCode = data.SecCode
		output.SecName = data.SecName
		output.Email = data.Email
		output.CID = data.CID
		output.ChiefCode = data.ChiefCode
		output.ChiefName = data.ChiefName
		output.ChiefFacCode = data.ChiefFacCode
		output.ChiefFacName = data.ChiefFacName
		return
	}
	output.Status = "password"
	log.Printf("Log: elogin-ChkLoginDatabase - fail password in elogin database for %s\n", username)
	return
}

func getUsernameDatabase(username string) (output UserDB) {
	db, err := dbs.OpenDBS(conf.DBS)
	if err != nil {
		log.Printf("Error: elogin-getUsernameDatabase 1 - Get username %s - %v\n", username, err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT username, password, name, firstname, lastname, faccode, facname, depcode, depname, seccode, secname, email, cid, chiefcode, chiefname, chieffaccode, chieffacname FROM elogin WHERE username=?", username)
	if err != nil {
		log.Printf("Error: elogin-getUsernameDatabase 2 - Query for %s - %v\n", username, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&output.Username, &output.Password, &output.Name, &output.FirstName, &output.LastName, &output.FacCode, &output.FacName, &output.DepCode, &output.DepName, &output.SecCode, &output.SecName, &output.Email, &output.CID, &output.ChiefCode, &output.ChiefName, &output.ChiefFacCode, &output.ChiefFacName)
		break
	}
	return
}

func getToken(username string, u UserStruct) (output string) {
	output = username + ":" + random.GetRandomString(conf.Elogin.TokenSize)
	db, err := dbs.OpenDBS(conf.DBS)
	if err != nil {
		log.Printf("Error: elogin-getToken 1 - Get token for %s - %v\n", username, err)
		return
	}
	defer db.Close()

	ts := utils.GetTimeStamp(time.Now())
	_, err = db.Exec("INSERT INTO token (token,timestamp,name,firstname,lastname,faccode,facname,depcode,depname,seccode,secname,email,username,cid,chiefcode,chiefname,chieffaccode,chieffacname) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);", output, ts, u.Name, u.FirstName, u.LastName, u.FacCode, u.FacName, u.DepCode, u.DepName, u.SecCode, u.SecName, u.Email, username, u.CID, u.ChiefCode, u.ChiefName, u.ChiefFacCode, u.ChiefFacName)
	if err != nil {
		log.Printf("Error: elogin-getToken 2 - Insert token for %s - %v\n", username, err)
	}
	return
}

func CleanTokenElogin() {
	db, err := dbs.OpenDBS(conf.DBS)
	if err != nil {
		log.Printf("Error: elogin-CleanTokenElogin 1 - %v\n", err)
		return
	}
	defer db.Close()

	ts := utils.GetTimeStamp(time.Now().Add(time.Duration(conf.Elogin.Expire) * time.Second * -1))
	_, err = db.Exec("DELETE FROM token WHERE timestamp < ?;", ts)
	if err != nil {
		log.Printf("Error: elogin-CleanTokenElogin 2 - %v\n", err)
	}
}
