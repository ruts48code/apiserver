package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	dbs "github.com/ruts48code/dbs4ruts"
)

type (
	UserAcademicInfoArrayStruct struct {
		Status string                   `json:"status"`
		Data   []UserAcademicInfoStruct `json:"data"`
	}
	UserAcademicInfoStruct struct {
		CID            string            `json:"cid"`
		Prefix         string            `json:"prefix"`
		Fname          string            `json:"fname"`
		Lname          string            `json:"lname"`
		FnameE         string            `json:"fnamee"`
		LnameE         string            `json:"lnamee"`
		Aposition      string            `json:"aposition"`
		CampusCode     string            `json:"campuscode"`
		CampusName     string            `json:"campusname"`
		FacultyCode    string            `json:"facultycode"`
		FacultyName    string            `json:"facultyname"`
		DepartmentCode string            `json:"departmentcode"`
		DepartmentName string            `json:"departmentname"`
		SectionCode    string            `json:"sectioncode"`
		SectionName    string            `json:"sectionname"`
		Email          string            `json:"email"`
		Gender         string            `json:"gender"`
		Epassport      string            `json:"epassport"`
		Education      []EducationStruct `json:"education"`
	}
	EducationStruct struct {
		Graduate   string `json:"graduate"`
		Level      string `json:"level"`
		Degree     string `json:"degree"`
		Program    string `json:"program"`
		University string `json:"univeristy"`
		Nation     string `json:"nation"`
		NationE    string `json:"natione"`
	}
	UserPersonalRequestStruct struct {
		Username []string `json:"username"`
	}
	PersonalCodeStruct struct {
		Gender                 []string
		Position               []string
		Campus                 []CodeStruct
		Faculty                []CodeStruct
		Department             []CodeStruct
		Section                []CodeStruct
		EducationLevel         []string
		EducationDegreeProgram []CodeStruct
	}
	CodeStruct struct {
		Code string
		Name string
	}
)

func PersonalAcademicPrivate(ctx *fiber.Ctx) error {
	username, _, _, status := CheckTKWeb(ctx.Params("token"))
	switch status {
	case "ok":
		return ctx.JSON(getAcademicInfo(username))
	default:
		output := UserAcademicInfoArrayStruct{
			Status: status,
		}
		return ctx.JSON(output)
	}
}

func PersonalAcademicPrivilege(ctx *fiber.Ctx) error {
	username, _, _, status := CheckTKWeb(ctx.Params("token"))
	switch status {
	case "ok":
		body := UserPersonalRequestStruct{}
		err := json.Unmarshal(ctx.Body(), &body)
		if err != nil {
			return ctx.JSON(fiber.Map{
				"status": "json",
			})
		}

		output := UserAcademicInfoArrayStruct{
			Status: status,
			Data:   make([]UserAcademicInfoStruct, 0),
		}

		privilege := false
		for _, v := range conf.Personal.Permission.ReadAll {
			if username == v {
				privilege = true
				break
			}
		}
		if !privilege {
			output.Status = "permission"
			return ctx.JSON(output)
		}
		for _, v := range body.Username {
			userDat := getAcademicInfo(v)
			if userDat.Status != "ok" {
				continue
			}
			output.Data = append(output.Data, userDat.Data...)
		}
		output.Status = "ok"
		return ctx.JSON(output)

	default:
		output := UserAcademicInfoArrayStruct{
			Status: status,
		}
		return ctx.JSON(output)
	}
}

func PersonalCode(ctx *fiber.Ctx) error {
	return ctx.Render("personalcode", pcode())
}

func getAcademicInfo(username string) (output UserAcademicInfoArrayStruct) {
	output.Status = "ok"
	db, err := dbs.OpenDB(conf.Personal.Server)
	if err != nil {
		log.Printf("Error: Get academic info databse for %s - %v\n", username, err)
		output.Status = "database"
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT CITIZEN_ID,PREFIX_NAME,STF_FNAME,STF_LNAME,STF_FNAME_EN,STF_LNAME_EN,POSITION_TNAME,CAMPUS_CODE,CAMPUS_TNAME,FACULTY_CODE,FacName,DEPARTMENT_CODE,DepName,SECTION_CODE,SECTION_TNAME,UOC_STAFF_EMAIL,GENDER_NAME,HGRAD_OUTDATE,GRAD_LEV_TNAME,HGRAD_DEGREE,PROGRAM_NAME,HGRAD_UNIV,NATION_TNAME,NATION_ENAME FROM vRISS WHERE USERNAME_CISCO=?;", username)
	if err != nil {
		log.Printf("Error: Query get academic info for %s - %v\n", username, err)
		output.Status = "database"
		return
	}
	defer rows.Close()

	output.Status = "not found"
	person := UserAcademicInfoStruct{}
	person.Education = make([]EducationStruct, 0)
	for rows.Next() {
		output.Status = "ok"
		education := EducationStruct{}
		var grad time.Time
		rows.Scan(&person.CID, &person.Prefix, &person.Fname, &person.Lname, &person.FnameE, &person.LnameE, &person.Aposition, &person.CampusCode, &person.CampusName, &person.FacultyCode, &person.FacultyName, &person.DepartmentCode, &person.DepartmentName, &person.SectionCode, &person.SectionName, &person.Email, &person.Gender, &grad, &education.Level, &education.Degree, &education.Program, &education.University, &education.Nation, &education.NationE)
		education.Graduate = grad.Format("2006-01-02")
		person.Epassport = username
		person.Education = append(person.Education, education)
	}
	output.Data = make([]UserAcademicInfoStruct, 0)
	output.Data = append(output.Data, person)
	return
}

func pcode() (output PersonalCodeStruct) {
	output.Gender = getDataString("gender_name")
	output.Position = getDataString("position_tname")
	output.Campus = getPairDataString("campus_code", "campus_tname")
	output.Faculty = getPairDataString("faculty_code", "facname")
	output.Department = getPairDataString("department_code", "depname")
	output.Section = getPairDataString("section_code", "section_tname")
	output.EducationLevel = getDataString("grad_lev_tname")
	output.EducationDegreeProgram = getPairDataString("hgrad_degree", "program_name")
	return
}

func getDataString(data string) (output []string) {
	db, err := dbs.OpenDB(conf.Personal.Server)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT DISTINCT " + data + " FROM vRISS;")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	for rows.Next() {
		data := ""
		rows.Scan(&data)
		output = append(output, data)
	}
	return
}

func getPairDataString(data1 string, data2 string) (output []CodeStruct) {
	db, err := dbs.OpenDB(conf.Personal.Server)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT DISTINCT " + data1 + "," + data2 + " FROM vRISS;")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	for rows.Next() {
		datax := CodeStruct{}
		rows.Scan(&datax.Code, &datax.Name)
		output = append(output, datax)
	}
	return
}
