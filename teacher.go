package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type (
	ClassListStructOutput struct {
		Status string            `json:"status"`
		Class  []ClassListStruct `json:"class"`
	}
	ClassListStruct struct {
		ClassID      string `json:"classid"`
		ClassRoom    string `json:"classroom"`
		AdmissYear   int    `json:"admissyear"`
		ClassName    string `json:"classname"`
		CountStudent int    `json:"countstudent"`
	}
	StudentSupervisorStructOutput struct {
		Status  string                    `json:"status"`
		ClassID string                    `json:"classid"`
		Members []StudentSupervisorStruct `json:"members"`
	}
	StudentSupervisorStruct struct {
		ID              string `json:"id"`
		Prefix          string `json:"prefix"`
		Fname           string `json:"fname"`
		Lname           string `json:"lname"`
		Phone           string `json:"phone"`
		GPA             string `json:"gpa"`
		Status          string `json:"status"`
		StatusName      string `json:"statusname"`
		GradSem         int    `json:"gradsem"`
		Facname         string `json:"facname"`
		MajorName       string `json:"majorname"`
		Degree          string `json:"degree"`
		PeriodName      string `json:"periodname"`
		Section         string `json:"section"`
		Lock            string `json:"lock"`
		Email           string `json:"email"`
		MicrosoftID     string `json:"microsoftid"`
		RegisConfirm    int    `json:"regisconfirm"`
		RegisAll        int    `json:"regisall"`
		NumCourse       int    `json:"numcourse"`
		NumPreserv      int    `json:"numpreserv"`
		WithdrawConfirm int    `json:"withdrawconfirm"`
		WithdrawAll     int    `json:"withdrawall"`
		PicStatus       int    `json:"picstatus"`
		Pic             string `json:"pic"`
		FundType        string `json:"fundtype"`
		FundName        string `json:"fundname"`
		SumRegisMoney   int    `json:"sumregismoney"`
		SumAccMoney     int    `json:"sumaccmoney"`
		Plan            int    `json:"plan"`
		Paid            int    `json:"paid"`
		ActUpload       int    `json:"actupload"`
	}
	StudentStatusStructOutput struct {
		Status     string `json:"status"`
		StatusName string `json:"statusname"`
		Count      int    `json:"count"`
		Members    []StudentSupervisorStruct
	}
	StudentTraceStructOutput struct {
		Status                 string                      `json:"status"`
		ClassID                string                      `json:"classid"`
		ConfirmAllNum          int                         `json:"confirmallnum"`
		ConfirmAllMembers      []StudentSupervisorStruct   `json:"confirmallmembers"`
		NotConfirmAllNum       int                         `json:"notconfirmallnum"`
		NotConfirmAllMembers   []StudentSupervisorStruct   `json:"notconfirmallmembers"`
		PreservNum             int                         `json:"preservnum"`
		PreservMembers         []StudentSupervisorStruct   `json:"perservmembers"`
		NotRegisPreservNum     int                         `json:"notregispreservnum"`
		NotRegisPreservMembers []StudentSupervisorStruct   `json:"notregispreservmembers"`
		WithdrawAllNum         int                         `json:"withdrawallnum"`
		WithdrawAllMembers     []StudentSupervisorStruct   `json:"withdrawallmembers"`
		NotWithdrawAllNum      int                         `json:"notwithdrawallnum"`
		NotWithdrawAllMembers  []StudentSupervisorStruct   `json:"notwithdrawallmembers"`
		PaidSuccessNum         int                         `json:"paidsuccessnum"`
		PaidSuccessMembers     []StudentSupervisorStruct   `json:"paidsuccessallmembers"`
		PaidUnSuccessNum       int                         `json:"paidunsuccessnum"`
		PaidUnSuccessMembers   []StudentSupervisorStruct   `json:"paidunsuccessallmembers"`
		ActSuccessNum          int                         `json:"actsuccessnum"`
		ActSuccessMembers      []StudentSupervisorStruct   `json:"actsuccessmembers"`
		ActUnSuccessNum        int                         `json:"actunsuccessnum"`
		ActUnSuccessMembers    []StudentSupervisorStruct   `json:"actunsuccessmembers"`
		FundNum                int                         `json:"fundnum"`
		FundMembers            []StudentSupervisorStruct   `json:"fundmembers"`
		UnFundNum              int                         `json:"unfundnum"`
		UnFundMembers          []StudentSupervisorStruct   `json:"unfundmembers"`
		StudentStatus          []StudentStatusStructOutput `json:"studentstatus"`
	}
)

func SuperVisor(ctx *fiber.Ctx) error {
	username, _, _, status := CheckTKWeb(ctx.Params("token"))
	switch status {
	case "ok":
		return ctx.JSON(GetSupervisorClass(username))
	default:
		output := ClassListStructOutput{
			Status: status,
			Class:  make([]ClassListStruct, 0),
		}
		return ctx.JSON(output)
	}
}

func SuperVisorClass(ctx *fiber.Ctx) error {
	username, _, _, status := CheckTKWeb(ctx.Params("token"))
	switch status {
	case "ok":
		return ctx.JSON(GetSupervisorClassMember(username, ctx.Params("classid")))
	default:
		output := StudentSupervisorStructOutput{
			Status:  status,
			ClassID: "",
			Members: make([]StudentSupervisorStruct, 0),
		}
		return ctx.JSON(output)
	}
}

func SuperVisorTrace(ctx *fiber.Ctx) error {
	username, _, _, status := CheckTKWeb(ctx.Params("token"))
	switch status {
	case "ok":
		return ctx.JSON(StudentTrace(GetSupervisorClassMember(username, ctx.Params("classid"))))
	default:
		output := StudentTraceStructOutput{
			Status:                status,
			ConfirmAllMembers:     make([]StudentSupervisorStruct, 0),
			NotConfirmAllMembers:  make([]StudentSupervisorStruct, 0),
			PreservMembers:        make([]StudentSupervisorStruct, 0),
			WithdrawAllMembers:    make([]StudentSupervisorStruct, 0),
			NotWithdrawAllMembers: make([]StudentSupervisorStruct, 0),
		}
		return ctx.JSON(output)
	}
}

func StudentTrace(student StudentSupervisorStructOutput) (output StudentTraceStructOutput) {
	output.Status = student.Status
	if student.Status != "ok" {
		return
	}
	output.ClassID = student.ClassID
	output.ConfirmAllMembers = make([]StudentSupervisorStruct, 0)
	output.NotConfirmAllMembers = make([]StudentSupervisorStruct, 0)
	output.PreservMembers = make([]StudentSupervisorStruct, 0)
	output.NotRegisPreservMembers = make([]StudentSupervisorStruct, 0)
	output.WithdrawAllMembers = make([]StudentSupervisorStruct, 0)
	output.NotWithdrawAllMembers = make([]StudentSupervisorStruct, 0)
	output.PaidSuccessMembers = make([]StudentSupervisorStruct, 0)
	output.PaidUnSuccessMembers = make([]StudentSupervisorStruct, 0)
	output.ActSuccessMembers = make([]StudentSupervisorStruct, 0)
	output.ActUnSuccessMembers = make([]StudentSupervisorStruct, 0)
	output.FundMembers = make([]StudentSupervisorStruct, 0)
	output.UnFundMembers = make([]StudentSupervisorStruct, 0)

	for _, v := range student.Members {
		switch v.Status {
		case "R", "C1", "P1", "P2", "P3":
			if v.RegisAll != 0 {
				if v.RegisConfirm == v.RegisAll {
					output.ConfirmAllMembers = append(output.ConfirmAllMembers, v)
				} else {
					output.NotConfirmAllMembers = append(output.NotConfirmAllMembers, v)
				}
			}
			if v.NumPreserv != 0 && v.NumCourse == 0 {
				output.PreservMembers = append(output.PreservMembers, v)
			}
			if v.NumPreserv == 0 && v.NumCourse == 0 {
				output.NotRegisPreservMembers = append(output.NotRegisPreservMembers, v)
			}
			if v.WithdrawAll != 0 {
				if v.WithdrawConfirm == v.WithdrawAll {
					output.WithdrawAllMembers = append(output.WithdrawAllMembers, v)
				} else {
					output.NotWithdrawAllMembers = append(output.NotWithdrawAllMembers, v)
				}
			}
			if v.NumCourse != 0 {
				if v.SumRegisMoney == v.SumAccMoney {
					output.PaidSuccessMembers = append(output.PaidSuccessMembers, v)
				} else {
					output.PaidUnSuccessMembers = append(output.PaidUnSuccessMembers, v)
				}
				if v.ActUpload == 0 {
					output.ActUnSuccessMembers = append(output.ActUnSuccessMembers, v)
				} else {
					output.ActSuccessMembers = append(output.ActSuccessMembers, v)
				}
			}

			if v.FundType == "N" || v.FundType == "" {
				output.UnFundMembers = append(output.UnFundMembers, v)
			} else {
				output.FundMembers = append(output.FundMembers, v)
			}
		}
		found := false
		for i := range output.StudentStatus {
			if v.Status == output.StudentStatus[i].Status {
				found = true
				output.StudentStatus[i].Count++
				output.StudentStatus[i].Members = append(output.StudentStatus[i].Members, v)
				break
			}
		}
		if !found {
			studentstatus := StudentStatusStructOutput{
				Status:     v.Status,
				StatusName: v.StatusName,
				Count:      1,
				Members:    make([]StudentSupervisorStruct, 0),
			}
			studentstatus.Members = append(studentstatus.Members, v)
			output.StudentStatus = append(output.StudentStatus, studentstatus)
		}
	}
	output.ConfirmAllNum = len(output.ConfirmAllMembers)
	output.NotConfirmAllNum = len(output.NotConfirmAllMembers)
	output.PreservNum = len(output.PreservMembers)
	output.NotRegisPreservNum = len(output.NotRegisPreservMembers)
	output.WithdrawAllNum = len(output.WithdrawAllMembers)
	output.NotWithdrawAllNum = len(output.NotWithdrawAllMembers)
	output.PaidSuccessNum = len(output.PaidSuccessMembers)
	output.PaidUnSuccessNum = len(output.PaidUnSuccessMembers)
	output.ActSuccessNum = len(output.ActSuccessMembers)
	output.ActUnSuccessNum = len(output.ActUnSuccessMembers)
	output.FundNum = len(output.FundMembers)
	output.UnFundNum = len(output.UnFundMembers)
	return
}

func GetSupervisorClassMember(username string, classid string) (output StudentSupervisorStructOutput) {
	serverID := GetServerStudentIDForTeacher(username)
	if len(serverID) == 0 {
		output.Status = "error"
	}
	members := make([]StudentSupervisorStruct, 0)
	for _, serverIDx := range serverID {
		memberx := GetMemberClassForTeacherServer(username, classid, serverIDx)
		members = append(members, memberx...)
	}
	output.Status = "ok"
	output.ClassID = classid
	output.Members = members
	return
}

func GetMemberClassForTeacherServer(username string, classid string, serverIDx string) (output []StudentSupervisorStruct) {
	output = make([]StudentSupervisorStruct, 0)
	dbname := GetStudentDBNameFromID(serverIDx)
	db, err := sql.Open("mysql", dbname)
	if err != nil {
		log.Printf("Error: Cannot connect to MySQL %s for %s - %v\n", dbname, username, err)
	} else {
		log.Printf("Log: Connect to MySQL %s for %s\n", dbname, username)
	}
	class, year := ExtractClassID(classid)
	rows, err := db.Query("select s.id, prefix, tfirst, tlast, c_tel, gpa, status, fStatusName(s.status) as statusname, grad_sem, faculty, major, degree, fPeriodName(s.period) as periodname, s.section, fLockStudent(s.id) as stdlock, fStudentEmail(s.id) as studentEmail, fSumCr(s.id,sem.semester) as regisConfirm, fSumCrAll(s.id,sem.semester) as regisAll, fNumCourse(s.id,sem.semester) as numCourse, fStudentPreserv(s.id,sem.semester) as numPreserv, fCountCourseWithdraw(s.id,sem.semester) as wdAll, fCountCourseWithdrawOk(s.id,sem.semester) as wdConfirm, fCountImageRStatus(s.citizen) as imageStatus, s.citizen, fFundStatus(sem.semester,s.id) as fundtype, fFundName(sem.semester,s.id) as fundname, fSumRegisMoney(sem.semester,s.id) as sumRegisMoney ,fSumAccMoney(sem.semester,s.id) as sumAccMoney,fCountPaymentBillPeriod(s.id,sem.semester) as plan,fCountPaymentUpload(s.id,sem.semester) as paid,fCountPaymentActUpload(s.id,concat(substring(sem.semester,1,2),'1')) as actUpload from login_web s,semester sem where sem.regis_status='Y' and s.classroom=? and s.admiss_year=? and s.status in (select id from status where used_status='Y') order by gpa asc;", class, year)
	if err != nil {
		log.Printf("Error: Query get data student for %s - %v\n", username, err)
	}
	for rows.Next() {
		d := StudentSupervisorStruct{}
		citizen := ""
		rows.Scan(&d.ID, &d.Prefix, &d.Fname, &d.Lname, &d.Phone, &d.GPA, &d.Status, &d.StatusName, &d.GradSem, &d.Facname, &d.MajorName, &d.Degree, &d.PeriodName, &d.Section, &d.Lock, &d.Email, &d.RegisConfirm, &d.RegisAll, &d.NumCourse, &d.NumPreserv, &d.WithdrawAll, &d.WithdrawConfirm, &d.PicStatus, &citizen, &d.FundType, &d.FundName, &d.SumRegisMoney, &d.SumAccMoney, &d.Plan, &d.Paid, &d.ActUpload)
		location := ""
		picpath := ""
		d.MicrosoftID = "s" + d.ID + "@ms.rmutsv.ac.th"
		for _, v := range conf.Student.Server {
			if v.ID == string(d.ID[0]) {
				location = strings.ToLower(v.Name)
				break
			}
		}
		if location != "" {
			picpath = "https://reg.rmutsv.ac.th/regis2file/" + location + "/stdImage/" + citizen + ".jpg"
		}
		d.Pic = picpath
		output = append(output, d)
	}
	return
}

func GetSupervisorClass(username string) (output ClassListStructOutput) {
	serverID := GetServerStudentIDForTeacher(username)
	if len(serverID) == 0 {
		output.Status = "error"
	}
	classlist := make([]ClassListStruct, 0)
	for _, serverIDx := range serverID {
		classlistx := GetClassForTeacherServer(username, serverIDx)
		classlist = append(classlist, classlistx...)
	}
	output.Status = "ok"
	output.Class = classlist
	return
}

func GetClassForTeacherServer(username string, serverIDx string) (output []ClassListStruct) {
	output = make([]ClassListStruct, 0)
	dbname := GetStudentDBNameFromID(serverIDx)
	db, err := sql.Open("mysql", dbname)
	if err != nil {
		log.Printf("Error: Cannot connect to MySQL %s for %s - %v\n", dbname, username, err)
	} else {
		log.Printf("Log: Connect to MySQL %s for %s\n", dbname, username)
	}

	rows, err := db.Query("select classroom,admiss_year,fClassNameNew(advs.student) as classname,count(advs.student) as cntStd from advisor_student advs,login_web s,instructorLogin i where advs.advisor=i.instructor and advs.student=s.id and i.esearch=? and i.loginstatus='epassport' and s.status in (select id from status where in_status='Y') group by classroom,admiss_year;", username)
	if err != nil {
		log.Printf("Error: Query get data student for %s - %v\n", username, err)
	}
	for rows.Next() {
		classroom := ""
		admiss_year := 0
		classname := ""
		countStudent := 0
		rows.Scan(&classroom, &admiss_year, &classname, &countStudent)
		class := ClassListStruct{
			ClassID:      fmt.Sprintf("%s_%d", classroom, admiss_year),
			ClassRoom:    classroom,
			ClassName:    classname,
			AdmissYear:   admiss_year,
			CountStudent: countStudent,
		}
		output = append(output, class)
	}
	return
}

func GetServerStudentIDForTeacher(username string) (output []string) {
	dataUser := getAcademicInfo(username)
	output = make([]string, 0)
	switch dataUser.Data[0].CampusCode {
	case "S":
		output = append(output, "1", "2")
	case "K", "N":
		output = append(output, "3", "4")
	case "H":
		output = append(output, "5")
	case "T":
		output = append(output, "6")
	}
	return
}

func ExtractClassID(classid string) (class string, year int) {
	data := strings.Split(classid, "_")
	class = data[0]
	if len(data) != 2 {
		return
	}
	year, err := strconv.Atoi(data[1])
	if err != nil {
		return
	}
	return
}
