package main

import (
	"encoding/json"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	crypt "github.com/ruts48code/crypt4ruts"
	random "github.com/ruts48code/random4ruts"
	util "github.com/ruts48code/utils4ruts"
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
	username := util.NormalizeUsername(ctx.Params("username"))
	DeleteLoginDatabase(username)
	return ctx.SendString("ok")
}

func elogin(ctx *fiber.Ctx) error {
	data := LoginStruct{}
	err := json.Unmarshal(ctx.Body(), &data)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status": "json",
		})
	}

	username := util.MakeString(util.NormalizeUsername(data.Username))
	password := util.MakeString(data.Password)

	if username == "" || password == "" {
		return ctx.JSON(UserStruct{
			Status: "passowrd",
		})
	}

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
		result = getDataStaff(username)
	case "student":
		result = getDataStudent(username)
	}
	DeleteLoginDatabase(username)
	CreateLoginDatabase(result, password)
	return ctx.JSON(result)
}

func ChkToken(tokenx string) (output UserStruct) {
	token := util.NormalizedEloginToken(tokenx)
	db, err := getDBS()
	if err != nil {
		log.Printf("Error: Check token database for %s - %v\n", token, err)
		output.Status = "database"
		return
	}
	defer db.Close()

	ts := util.GetTimeStamp(time.Now())
	qstring := ""
	switch conf.DBType {
	case "postgres":
		qstring = "UPDATE token SET timestamp=$1 WHERE token=$2;"
	case "mysql":
		qstring = "UPDATE token SET timestamp=? WHERE token=?;"
	}
	_, err = db.Exec(qstring, ts, token)
	if err != nil {
		log.Printf("Error: Update check token for %s - %v\n", token, err)
		output.Status = "database"
		return
	}

	switch conf.DBType {
	case "postgres":
		qstring = "SELECT name,firstname,lastname,username,faccode,facname,depcode,depname,seccode,secname,email,cid,chiefcode,chiefname,chieffaccode,chieffacname FROM token WHERE token=$1;"
	case "mysql":
		qstring = "SELECT name,firstname,lastname,username,faccode,facname,depcode,depname,seccode,secname,email,cid,chiefcode,chiefname,chieffaccode,chieffacname FROM token WHERE token=?;"
	}
	rows, err := db.Query(qstring, token)
	if err != nil {
		log.Printf("Error: Query check token for %s - %v\n", token, err)
		output.Status = "database"
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&output.Name, &output.FirstName, &output.LastName, &output.Username, &output.FacCode, &output.FacName, &output.DepCode, &output.DepName, &output.SecCode, &output.SecName, &output.Email, &output.CID, &output.ChiefCode, &output.ChiefName, &output.ChiefFacCode, &output.ChiefFacName)
		output.Status = "ok"
		output.Token = token
		output.Type = util.CheckEpassportType(output.Username)
		break
	}
	if output.Status == "" {
		output.Status = "token"
		log.Printf("Log: No token found for %s\n", token)
	}

	return
}

func DeleteLoginDatabase(username string) {
	db, err := getDBS()
	if err != nil {
		log.Printf("Error: Delete login database for %s - %v\n", username, err)
		return
	}
	defer db.Close()

	qstring := ""
	switch conf.DBType {
	case "postgres":
		qstring = "DELETE FROM elogin WHERE username=$1;"
	case "mysql":
		qstring = "DELETE FROM elogin WHERE username=?;"
	}
	_, err = db.Exec(qstring, username)
	if err != nil {
		log.Printf("Error: Delete login database for %s - %v\n", username, err)
	}
}

func CreateLoginDatabase(result UserStruct, password string) {
	db, err := getDBS()
	if err != nil {
		log.Printf("Error: Create login database for %s - %v\n", result.Username, err)
		return
	}
	defer db.Close()

	appmutex.Lock()
	salt := random.GetRandomString(4)
	appmutex.Unlock()

	passx := crypt.MooHash([]byte(password), []byte(result.Username), []byte(salt))

	qstring := ""
	switch conf.DBType {
	case "postgres":
		qstring = "INSERT INTO elogin (username, password, name, firstname, lastname, faccode, facname, depcode, depname, seccode, secname, email, cid, chiefcode, chiefname, chieffaccode, chieffacname) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17);"
	case "mysql":
		qstring = "INSERT INTO elogin (username, password, name, firstname, lastname, faccode, facname, depcode, depname, seccode, secname, email, cid, chiefcode, chiefname, chieffaccode, chieffacname) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
	}
	_, err = db.Exec(qstring, result.Username, passx, result.Name, result.FirstName, result.LastName, result.FacCode, result.FacName, result.DepCode, result.DepName, result.SecCode, result.SecName, result.Email, result.CID, result.ChiefCode, result.ChiefName, result.ChiefFacCode, result.ChiefFacName)
	if err != nil {
		log.Printf("Error: Insert login database for %s - %v\n", result.Username, err)
	}
}

func getDataStaff(username string) (output UserStruct) {
	db, err := getSQLServer()
	if err != nil {
		log.Printf("Error: Get data staff for %s - %v\n", username, err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT STF_FNAME, STF_LNAME, FACULTY_CODE, FacName, department_code, depname, section_code, section_tname, CITIZEN_ID FROM vUOC_STAFF_L01 WHERE USERNAME_CISCO=@p1;", username)
	if err != nil {
		log.Printf("Error: Query data staff for %s - %v\n", username, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&output.FirstName, &output.LastName, &output.FacCode, &output.FacName, &output.DepCode, &output.DepName, &output.SecCode, &output.SecName, &output.CID)
		output.Status = "ok"
		output.Username = username
		output.Name = output.FirstName + " " + output.LastName
		output.Type = util.CheckEpassportType(username)
		output.Email = username + "@rmutsv.ac.th"
		break
	}

	rows2, err2 := db.Query("SELECT CHIEF_CODE, CHIEF_NAME, FACULTY_CODE, Facname from vADMIN WHERE USERNAME_CISCO=@p1;", username)
	if err2 != nil {
		log.Printf("Error: Query data staff admin for %s - %v\n", username, err)
		return
	}
	defer rows2.Close()

	for rows2.Next() {
		rows2.Scan(&output.ChiefCode, &output.ChiefName, &output.ChiefFacCode, &output.ChiefFacName)
		break
	}
	output.Token = getToken(username, output)
	return
}

func ChkLoginLDAP(username string, password string) (output UserStruct) {
	result := ldapLogin(username, password)
	if result == "none" {
		output.Status = "password"
		return
	}
	output.Status = "ok"
	output.Type = util.CheckEpassportType(username)
	return
}

func ChkLoginDatabase(username string, password string) (output UserStruct) {
	output.Status = "fail"
	data := getUsernameDatabase(username)
	if data.Username == "" {
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
		output.Type = util.CheckEpassportType(username)
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
	return
}

func getUsernameDatabase(username string) (output UserDB) {
	db, err := getDBS()
	if err != nil {
		log.Printf("Error: Get username %s - %v\n", username, err)
		return
	}
	defer db.Close()

	qstring := ""
	switch conf.DBType {
	case "postgres":
		qstring = "SELECT username, password, name, firstname, lastname, faccode, facname, depcode, depname, seccode, secname, email, cid, chiefcode, chiefname, chieffaccode, chieffacname FROM elogin WHERE username=$1"
	case "mysql":
		qstring = "SELECT username, password, name, firstname, lastname, faccode, facname, depcode, depname, seccode, secname, email, cid, chiefcode, chiefname, chieffaccode, chieffacname FROM elogin WHERE username=?"
	}
	rows, err := db.Query(qstring, username)
	if err != nil {
		log.Printf("Error: Query for %s - %v\n", username, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&output.Username, &output.Password, &output.Name, &output.FirstName, &output.LastName, &output.FacCode, &output.FacName, &output.DepCode, &output.DepName, &output.SecCode, &output.SecName, &output.Email, &output.CID, &output.ChiefCode, &output.ChiefName, &output.ChiefFacCode, &output.ChiefFacName)
		break
	}

	log.Printf("Log: Get username database for %s\n", username)
	return
}

func getToken(username string, u UserStruct) (output string) {
	appmutex.Lock()
	output = username + ":" + random.GetRandomString(conf.Elogin.TokenSize)
	appmutex.Unlock()
	db, err := getDBS()
	if err != nil {
		log.Printf("Error: Get token for %s - %v\n", username, err)
		return
	}
	defer db.Close()

	qstring := ""
	switch conf.DBType {
	case "postgres":
		qstring = "INSERT INTO token (token,timestamp,name,firstname,lastname,faccode,facname,depcode,depname,seccode,secname,email,username,cid,chiefcode,chiefname,chieffaccode,chieffacname) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18);"
	case "mysql":
		qstring = "INSERT INTO token (token,timestamp,name,firstname,lastname,faccode,facname,depcode,depname,seccode,secname,email,username,cid,chiefcode,chiefname,chieffaccode,chieffacname) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
	}
	ts := util.GetTimeStamp(time.Now())
	_, err = db.Exec(qstring, output, ts, u.Name, u.FirstName, u.LastName, u.FacCode, u.FacName, u.DepCode, u.DepName, u.SecCode, u.SecName, u.Email, username, u.CID, u.ChiefCode, u.ChiefName, u.ChiefFacCode, u.ChiefFacName)
	if err != nil {
		log.Printf("Error: Insert token for %s - %v\n", username, err)
	}

	log.Printf("Log: Save token for %s\n", username)
	return
}
