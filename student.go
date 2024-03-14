package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	dbs "github.com/ruts48code/dbs4ruts"
	utils "github.com/ruts48code/utils4ruts"
)

type (
	StudentGradeStructOutput struct {
		Status            string                       `json:"status"`
		ID                string                       `json:"id"`
		RegisCredit       int                          `json:"regiscredit"`
		EarnCredit        int                          `json:"earncredit"`
		GPA               string                       `json:"gpa"`
		StudentStatus     string                       `json:"studentstatus"`
		StudentStatusName string                       `json:"studentstatusname"`
		Supervisor        []SupervisorForStudentStruct `json:"supervisor"`
		Semester          []Semester                   `json:"semester"`
	}
	Semester struct {
		Sem         string   `json:"semester"`
		Text1       string   `json:"text1"`
		Text2       string   `json:"text2"`
		RegisCredit int      `json:"regiscredit"`
		EarnCredit  int      `json:"earncredit"`
		Status      string   `json:"status"`
		StatusName  string   `json:"statusname"`
		GPS         string   `json:"gps"`
		GPA         string   `json:"gpa"`
		Course      []Course `json:"course"`
	}
	Course struct {
		CourseID     string `json:"courseid"`
		CourseName   string `json:"coursename"`
		TheoryCredit int    `json:"theorycredit"`
		LabCredit    int    `json:"labcredit"`
		Grade        string `json:"grade"`
	}
	CourseRegis struct {
		CourseID       string `json:"courseid"`
		CourseName     string `json:"coursename"`
		TheoryCredit   int    `json:"theorycredit"`
		LabCredit      int    `json:"labcredit"`
		Section        int    `json:"section"`
		CourseStatus   string `json:"coursestatus"`
		TeacherName    string `json:"teachername"`
		AdvisorOK      string `json:"advisorok"`
		AdvisorDate    string `json:"advisordate"`
		MajorOK        string `json:"majorok"`
		MajorName      string `json:"majorname"`
		MajorDate      string `json:"majordate"`
		OfficerOK      string `json:"officerok"`
		OfficerName    string `json:"officername"`
		OfficerDate    string `json:"officerdate"`
		ViceDeanOK     string `json:"vicedeanok"`
		ViceDeanName   string `json:"vicedeanname"`
		ViceDeanDate   string `json:"vicedeandate"`
		DeanOK         string `json:"deanok"`
		DeanName       string `json:"deanname"`
		DeanDate       string `json:"deandate"`
		ViceCampusOK   string `json:"vicecampusok"`
		ViceCampusName string `json:"vicecampusname"`
		ViceCampusDate string `json:"vicecampusdate"`
	}
	CourseWithdraw struct {
		CourseID          string `json:"courseid"`
		CourseName        string `json:"coursename"`
		Section           string `json:"section"`
		CourseType        string `json:"coursetype"`
		DateUpdate        string `json:"dateupdate"`
		TimeUpdate        string `json:"timeupdate"`
		CourseStatus      string `json:"coursestatus"`
		InstructionOK     string `json:"instructionok"`
		InstructionName   string `json:"instructionname"`
		AdvisorOK         string `json:"advisorok"`
		MajorOK           string `json:"majorok"`
		MajorName         string `json:"majorname"`
		MajorNameORG      string `json:"majornameorg"`
		DepartmentOK      string `json:"departmentok"`
		DepartmentName    string `json:"departmentname"`
		DepartmentNameORG string `json:"departmentnameorg"`
	}
	StudentRegisStructOutput struct {
		Status         string                       `json:"status"`
		ID             string                       `json:"id"`
		Sem            string                       `json:"semester"`
		Text1          string                       `json:"text1"`
		Text2          string                       `json:"text2"`
		RegisCredit    int                          `json:"regiscredit"`
		Supervisor     []SupervisorForStudentStruct `json:"supervisor"`
		Course         []CourseRegis                `json:"course"`
		CourseWithdraw []CourseWithdraw             `json:"coursewithdraw"`
	}
	SupervisorForStudentStruct struct {
		Supervisor string `json:"supervisor"`
		Priority   string `json:"priority"`
		Epassport  string `json:"epassport"`
	}
	ClassTraceStruct struct {
		ID      string                        `json:"id"`
		Trace   StudentTraceStructOutput      `json:"trace"`
		Members StudentSupervisorStructOutput `json:"members"`
	}
	SupervisorDataStruct struct {
		Status      string `json:"status"`
		Name        string `json:"name"`
		Epassport   string `json:"epassport"`
		FacultyName string `json:"facultyname"`
		Class       []ClassTraceStruct
	}
)

func StudentRegis(ctx *fiber.Ctx) error {
	_, _, _, status := CheckTKWeb(ctx.Params("token"))
	switch status {
	case "ok":
		return ctx.JSON(GetStudentRegis(ctx.Params("id")))
	default:
		output := StudentRegisStructOutput{
			Status: status,
		}
		return ctx.JSON(output)
	}
}

func StudentGrade(ctx *fiber.Ctx) error {
	_, _, _, status := CheckTKWeb(ctx.Params("token"))
	switch status {
	case "ok":
		return ctx.JSON(GetStudentGrade(ctx.Params("id")))
	default:
		output := StudentGradeStructOutput{
			Status:   status,
			Semester: make([]Semester, 0),
		}
		return ctx.JSON(output)
	}
}

func StudentAllData(ctx *fiber.Ctx) error {
	_, _, _, status := CheckTKWeb(ctx.Params("token"))
	switch status {
	case "ok":
		return ctx.JSON(GetAllStudentData())
	default:
		return ctx.JSON(make([]SupervisorDataStruct, 0))
	}
}

func GetStudentRegis(id string) (output StudentRegisStructOutput) {
	idx := strings.TrimSpace(id)
	output = StudentRegisStructOutput{
		Status:         "ok",
		ID:             idx,
		Supervisor:     GetStudentSupervisor(idx),
		Course:         make([]CourseRegis, 0),
		CourseWithdraw: make([]CourseWithdraw, 0),
	}
	dbname := GetStudentDBNameFromID(string(idx[0]))
	db, err := dbs.OpenDB(dbname)
	if err != nil {
		log.Printf("Error: Cannot connect to MySQL %s for %s - %v\n", dbname, idx, err)
		output.Status = "databaseconnect"
		return output
	} else {
		log.Printf("Log: Connect to MySQL %s for %s\n", dbname, idx)
	}

	rows, err := db.Query("select r.semester,sem.semestertext,sem.semestertext2,r.student,r.course as courseid,c.tname as coursename,c.th_cr,c.lb_cr,r.section,r.status as courseStatus,fInstructorName(cfl.instructor) as teacherName,r.advisorok,o.uploadDate as advisorDate,o.majorok,o.majorDate,fInstructorName(m.head) as majorname,o.officeok,o.officeDate,f.officer_user,o.vice_deanok,o.vice_deanDate,fInstructorName(f.vice_dean) as vice_deanname,o.deanok,o.deanDate,fInstructorName(f.dean) as deanname,o.vice_campusok,o.vice_campusDate,fInstructorName(cp.vice_campus) as vice_campusname from basketregis r,basketregisok o,semester sem,course c,course_offer_limit cfl,login_web s,advisor_classroom adv,majorregis m,department d,facultyofcourse f,campus cp where r.student=o.student and r.semester=o.semester and r.semester=sem.semester and r.course=c.id and r.semester=cfl.semester and r.course=cfl.course and r.section=cfl.section and r.student=s.id and s.classroom=adv.classroom and s.admiss_year=adv.admiss_year and adv.majorregis=m.id and m.depid=d.id and d.faculty=f.id and sem.regis_status='Y' and r.student=?;", idx)
	if err != nil {
		log.Printf("Error: Query get data student for %s - %v\n", idx, err)
		output.Status = "databasequery"
		return output
	}
	output.RegisCredit = 0
	first := true
	for rows.Next() {
		course := CourseRegis{}
		sem := ""
		text1 := ""
		text2 := ""
		id := ""
		rows.Scan(&sem, &text1, &text2, &id, &course.CourseID, &course.CourseName, &course.TheoryCredit, &course.LabCredit, &course.Section, &course.CourseStatus, &course.TeacherName, &course.AdvisorOK, &course.AdvisorDate, &course.MajorOK, &course.MajorDate, &course.MajorName, &course.OfficerOK, &course.OfficerDate, &course.OfficerName, &course.ViceDeanOK, &course.ViceDeanDate, &course.ViceDeanName, &course.DeanOK, &course.DeanDate, &course.DeanName, &course.ViceCampusOK, &course.ViceCampusDate, &course.ViceCampusName)
		if first {
			output.Sem = sem
			output.Text1 = text1
			output.Text2 = text2
		}
		output.RegisCredit += course.TheoryCredit + course.LabCredit
		output.Course = append(output.Course, course)
	}

	rows2, err2 := db.Query("select b.course,c.tname as coursename,b.section,'วิชาในสาขา' as coursetype,b.dateUpdate,b.timeUpdate,r.status as courseStatus,b.instructorok,cfl.instructor,fInstructorName(cfl.instructor) as instructorname,b.advisorok,'อาจารย์ที่ปรึกษา' as advisor,b.majorok,m.head as mhead,fInstructorName(m.head) as mheadname,m.tname as majorname,b.departmentok,d.head as dhead,fInstructorName(d.head) as dheadname,d.tname as departmentname from basketwithdraw b,course_offer_limit cfl,course c,instrfac i,majorregis m,department d,basketregis r where b.semester=cfl.semester and b.course=cfl.course and b.section=cfl.section and b.course=c.id and b.student=r.student and b.semester=r.semester and b.course=r.course and b.section=r.section and fAdvisorMain(b.student)=i.instructor and i.majorid=m.id and m.depid=d.id and c.coursetype <> 1 and b.semester in (select semester from semester where regis_status='Y') and b.student=? UNION select b.course,c.tname as coursename,b.section,'วิชาศึกษาทั่วไป' as coursetype,b.dateUpdate,b.timeUpdate,r.status as courseStatus,b.instructorok,cfl.instructor,fInstructorName(cfl.instructor) as instructorname,b.advisorok,'อาจารย์ที่ปรึกษา' as advisor,b.majorok,m.head as mhead,fInstructorName(m.head) as mheadname,m.tname as majorname,b.departmentok,d.head as dhead,fInstructorName(d.head) as dheadname,d.tname as departmentname from basketwithdraw b,course_offer_limit cfl,course c,instrfac i,majorregis m,department d,basketregis r where b.semester=cfl.semester and b.course=cfl.course and b.section=cfl.section and b.course=c.id and b.student=r.student and b.semester=r.semester and b.course=r.course and b.section=r.section and cfl.instructor=i.instructor and i.majorid=m.id and m.depid=d.id and c.coursetype=1 and b.semester in (select semester from semester where regis_status='Y') and b.student=?;", idx, idx)
	if err2 != nil {
		log.Printf("Error: Query get data student for %s - %v\n", idx, err)
		output.Status = "databasequery"
		return output
	}
	for rows2.Next() {
		course := CourseWithdraw{}
		id1 := 0
		id2 := ""
		id3 := 0
		id4 := 0
		rows2.Scan(&course.CourseID, &course.CourseName, &course.Section, &course.CourseType, &course.DateUpdate, &course.TimeUpdate, &course.CourseStatus, &course.InstructionOK, &id1, &course.InstructionName, &course.AdvisorOK, &id2, &course.MajorOK, &id3, &course.MajorName, &course.MajorNameORG, &course.DepartmentOK, &id4, &course.DepartmentName, &course.DepartmentNameORG)
		output.CourseWithdraw = append(output.CourseWithdraw, course)
	}
	return
}

func GetStudentGrade(id string) (output StudentGradeStructOutput) {
	idx := strings.TrimSpace(id)
	output = StudentGradeStructOutput{
		Status:     "ok",
		ID:         idx,
		Supervisor: GetStudentSupervisor(idx),
		Semester:   make([]Semester, 0),
	}
	dbname := GetStudentDBNameFromID(string(idx[0]))
	db, err := dbs.OpenDB(dbname)
	if err != nil {
		log.Printf("Error: Cannot connect to MySQL %s for %s - %v\n", dbname, idx, err)
		output.Status = "databaseconnect"
		return output
	} else {
		log.Printf("Log: Connect to MySQL %s for %s\n", dbname, idx)
	}

	rows, err := db.Query("select t.student,sem.semester,sem.semestertext,sem.semestertext2,g.regis_cr,g.earn_cr,g.gps,g.all_regis_cr,g.all_earn_cr,g.gpa,p.status as proStatus,fStatusName(p.status) as proStatusName,fStudentStatus(t.student) as std_status,fStatusName(fStudentStatus(t.student)) as std_statusName,c.id as courseid,c.tname as coursename,c.th_cr,c.lb_cr,t.grade from transcript t,gpa g,semester sem,course c,pro_status p where t.student=g.student and t.semester=g.semester and g.semester=sem.semester and t.course=c.id and g.student=p.student and g.semester=p.semester and t.student=? order by sem.semester;", idx)
	if err != nil {
		log.Printf("Error: Query get data student for %s - %v\n", idx, err)
		output.Status = "databasequery"
		return output
	}
	semester := Semester{
		Course: make([]Course, 0),
	}
	for rows.Next() {
		course := Course{}
		idd := ""
		sem := ""
		text1 := ""
		text2 := ""
		regiscredit := 0
		earncredit := 0
		aregiscredit := 0
		aearncredit := 0
		gpa := ""
		gps := ""
		proStatus := ""
		proStatusName := ""
		studentStatus := ""
		studentStatusName := ""
		rows.Scan(&idd, &sem, &text1, &text2, &regiscredit, &earncredit, &gps, &aregiscredit, &aearncredit, &gpa, &proStatus, &proStatusName, &studentStatus, &studentStatusName, &course.CourseID, &course.CourseName, &course.TheoryCredit, &course.LabCredit, &course.Grade)
		output.RegisCredit = aregiscredit
		output.EarnCredit = aearncredit
		output.GPA = gpa
		output.StudentStatus = studentStatus
		output.StudentStatusName = studentStatusName
		if sem != semester.Sem {
			if semester.Sem != "" {
				output.Semester = append(output.Semester, semester)
			}

			semester = Semester{
				Sem:         sem,
				Text1:       text1,
				Text2:       text2,
				RegisCredit: regiscredit,
				EarnCredit:  earncredit,
				GPS:         gps,
				GPA:         gpa,
				Status:      proStatus,
				StatusName:  proStatusName,
				Course:      make([]Course, 0),
			}
			semester.Course = append(semester.Course, course)
		} else {
			semester.Course = append(semester.Course, course)
		}
	}
	if semester.Sem != "" {
		output.Semester = append(output.Semester, semester)
	}
	return
}

func getStudentDB(username string) (db *dbs.DB4ruts, err error) {
	dbname := GetStudentDBNameFromID(string(username[1]))
	db, err = dbs.OpenDB(dbname)
	if err != nil {
		log.Printf("Error: Cannot connect to MySQL %s for %s - %v\n", dbname, username, err)
	} else {
		log.Printf("Log: Connect to MySQL %s for %s\n", dbname, username)
	}
	return
}

func getDataStudent(username string) (output UserStruct) {
	db, err := getStudentDB(username)
	if err != nil {
		log.Printf("Error: Get data student for %s - %v\n", username, err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select s.citizen as citizen, s.tfirst as tfirst, s.tlast as tlast, f.id as faculty_id, f.tname as faculty_name, d.id as department_id, d.tname as department_name, m.id as major_id, m.tname as major_name, n.id as minor_id, n.tname as minor_name, fStudentEmail(s.id) as email, fStudentEmailUser(s.id) as emailUser from login_web s,advisor_classroom adv,minorregis n,majorregis m,department d,facultyofcourse f where s.classroom=adv.classroom and s.admiss_year=adv.admiss_year and adv.majorregis=m.id and adv.minorregis=n.id and m.depid=d.id and d.faculty=f.id and s.id=?;", username[1:])
	if err != nil {
		log.Printf("Error: Query get data student for %s - %v\n", username, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		minorcode := ""
		minorname := ""
		emailuser := ""
		rows.Scan(&output.CID, &output.FirstName, &output.LastName, &output.FacCode, &output.FacName, &output.DepCode, &output.DepName, &output.SecCode, &output.SecName, &minorcode, &minorname, &output.Email, &emailuser)
		output.Status = "ok"
		output.Username = username
		output.Name = output.FirstName + " " + output.LastName
		output.Type = utils.CheckEpassportType(username)
		output.Token = getToken(username, output)
		return
	}
	return
}

func GetStudentDBNameFromID(id string) (dbname string) {
	dbname = ""
	for _, v := range conf.Student.Server {
		if v.ID == id {
			dbname = v.Server
			break
		}
	}
	return
}

func GetStudentSupervisor(id string) (output []SupervisorForStudentStruct) {
	output = make([]SupervisorForStudentStruct, 0)
	dbname := GetStudentDBNameFromID(string(id[0]))
	db, err := dbs.OpenDB(dbname)
	if err != nil {
		log.Printf("Error: Cannot connect to MySQL %s for %s - %v\n", dbname, id, err)
		return
	} else {
		log.Printf("Log: Connect to MySQL %s for %s\n", dbname, id)
	}

	rows, err := db.Query("select fInstructorName(a.advisor) as advisorName,priority,l.esearch,l.loginstatus from advisor_student a,instructorLogin l where a.advisor=l.instructor and a.student=?;", id)
	if err != nil {
		log.Printf("Error: Query get data student for %s - %v\n", id, err)
		return
	}
	for rows.Next() {
		a := ""
		supervisor := SupervisorForStudentStruct{}
		rows.Scan(&supervisor.Supervisor, &supervisor.Priority, &supervisor.Epassport, &a)
		supervisor.Epassport = supervisor.Epassport + "@rmutsv.ac.th"
		output = append(output, supervisor)
	}
	return
}

func GetAllStudentData() (output []SupervisorDataStruct) {
	output = GetAllSupervisor()
	FillAllClassToSupervisor(output)
	FillTraceToClassOfSupervisor(output)
	return
}

func GetAllSupervisor() (output []SupervisorDataStruct) {
	output = make([]SupervisorDataStruct, 0)
	for i := range conf.Student.Server {
		data := GetSupervisorFromServer(conf.Student.Server[i].Server)
		output = append(output, data...)
	}
	return
}

func GetSupervisorFromServer(server string) (output []SupervisorDataStruct) {
	output = make([]SupervisorDataStruct, 0)
	db, err := dbs.OpenDB(server)
	if err != nil {
		log.Printf("Error: Cannot connect to SiS MySQL %s - %v\n", server, err)
		return
	} else {
		log.Printf("Log: Connect to SiS MySQL %s\n", server)
	}
	defer db.Close()

	rows, err := db.Query("select fInstructorName(a.advisor),i.esearch,f.tname as facultyname from login_web s,advisor_student a,instructorLogin i,instrfac t,majorregis m,department d,facultyofcourse f where s.id=a.student and a.advisor=i.instructor and a.advisor=t.instructor and t.majorid=m.id and m.depid=d.id and d.faculty=f.id and a.priority='M' and i.loginstatus='epassport' and s.status in (select id from status where in_status='Y') and s.admiss_year <= (select academicyear from campus) group by a.advisor,i.esearch,f.tname")
	if err != nil {
		log.Printf("Error: Query get supervisor from server %s - %v\n", server, err)
		return
	}
	for rows.Next() {
		data := SupervisorDataStruct{}
		data.Class = make([]ClassTraceStruct, 0)
		rows.Scan(&data.Name, &data.Epassport, &data.FacultyName)
		output = append(output, data)
	}
	return
}

func FillAllClassToSupervisor(supervisor []SupervisorDataStruct) {
	for i := range supervisor {
		FillClassToSupervisor(i, supervisor)
	}
}

func FillClassToSupervisor(i int, supervisor []SupervisorDataStruct) {
	classroom := GetSupervisorClass(supervisor[i].Epassport)
	for j := range classroom.Class {
		classx := ClassTraceStruct{}
		classx.ID = classroom.Class[j].ClassID
		supervisor[i].Class = append(supervisor[i].Class, classx)
	}
}

func FillTraceToClassOfSupervisor(supervisor []SupervisorDataStruct) {
	for i := range supervisor {
		for j := range supervisor[i].Class {
			supervisor[i].Class[j].Members = GetSupervisorClassMember(supervisor[i].Epassport, supervisor[i].Class[j].ID)
			supervisor[i].Class[j].Trace = StudentTrace(supervisor[i].Class[j].Members)
		}
	}
}
