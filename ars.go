package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	dbs "github.com/ruts48code/dbs4ruts"
)

type (
	ArsUniversity struct {
		Plan             int `json:"plan"`
		Applicant        int `json:"applicant"`
		Confirm          int `json:"confirm"`
		Report           int `json:"report"`
		PlanGoal         int `json:"plangoal"`
		ApplicantGoal    int `json:"applicantgoal"`
		ConfirmGoal      int `json:"confirmgoal"`
		ReportGoal       int `json:"reportgoal"`
		ApplicantPercent int `json:"applicantpercent"`
		ConfirmPercent   int `json:"confirmpercent"`
		ReportPercent    int `json:"reportpercent"`
		ApplicantTCAS    int `json:"applicanttcas"`
		ApplicantTech    int `json:"applicanttech"`
		ApplicantQP      int `json:"applicantqp"`
		ApplicantQPM6    int `json:"applicantqpm6"`
		ApplicantQPTech  int `json:"applicantqptech"`
		ConfirmTCAS      int `json:"confirmtcas"`
		ConfirmTech      int `json:"confirmtech"`
		ConfirmQP        int `json:"confirmqp"`
		ConfirmQPM6      int `json:"confirmqpm6"`
		ConfirmQPTech    int `json:"confirmqptech"`
		ReportTCAS       int `json:"reporttcas"`
		ReportTech       int `json:"reporttech"`
		ReportQP         int `json:"reportqp"`
		ReportQPM6       int `json:"reportqpm6"`
		ReportQPTech     int `json:"reportqptech"`
	}
	ArsCampus struct {
		ID               string `json:"id"`
		Name             string `json:"name"`
		Plan             int    `json:"plan"`
		Applicant        int    `json:"applicant"`
		Confirm          int    `json:"confirm"`
		Report           int    `json:"report"`
		PlanGoal         int    `json:"plangoal"`
		ApplicantGoal    int    `json:"applicantgoal"`
		ConfirmGoal      int    `json:"confirmgoal"`
		ReportGoal       int    `json:"reportgoal"`
		ApplicantPercent int    `json:"applicantpercent"`
		ConfirmPercent   int    `json:"confirmpercent"`
		ReportPercent    int    `json:"reportpercent"`
		ApplicantTCAS    int    `json:"applicanttcas"`
		ApplicantTech    int    `json:"applicanttech"`
		ApplicantQP      int    `json:"applicantqp"`
		ApplicantQPM6    int    `json:"applicantqpm6"`
		ApplicantQPTech  int    `json:"applicantqptech"`
		ConfirmTCAS      int    `json:"confirmtcas"`
		ConfirmTech      int    `json:"confirmtech"`
		ConfirmQP        int    `json:"confirmqp"`
		ConfirmQPM6      int    `json:"confirmqpm6"`
		ConfirmQPTech    int    `json:"confirmqptech"`
		ReportTCAS       int    `json:"reporttcas"`
		ReportTech       int    `json:"reporttech"`
		ReportQP         int    `json:"reportqp"`
		ReportQPM6       int    `json:"reportqpm6"`
		ReportQPTech     int    `json:"reportqptech"`
	}

	ArsFaculty struct {
		ID               string `json:"id"`
		Name             string `json:"name"`
		CampusID         string `json:"campusID"`
		Plan             int    `json:"plan"`
		Applicant        int    `json:"applicant"`
		Confirm          int    `json:"confirm"`
		Report           int    `json:"report"`
		PlanGoal         int    `json:"plangoal"`
		ApplicantGoal    int    `json:"applicantgoal"`
		ConfirmGoal      int    `json:"confirmgoal"`
		ReportGoal       int    `json:"reportgoal"`
		ApplicantPercent int    `json:"applicantpercent"`
		ConfirmPercent   int    `json:"confirmpercent"`
		ReportPercent    int    `json:"reportpercent"`
		ApplicantTCAS    int    `json:"applicanttcas"`
		ApplicantTech    int    `json:"applicanttech"`
		ApplicantQP      int    `json:"applicantqp"`
		ApplicantQPM6    int    `json:"applicantqpm6"`
		ApplicantQPTech  int    `json:"applicantqptech"`
		ConfirmTCAS      int    `json:"confirmtcas"`
		ConfirmTech      int    `json:"confirmtech"`
		ConfirmQP        int    `json:"confirmqp"`
		ConfirmQPM6      int    `json:"confirmqpm6"`
		ConfirmQPTech    int    `json:"confirmqptech"`
		ReportTCAS       int    `json:"reporttcas"`
		ReportTech       int    `json:"reporttech"`
		ReportQP         int    `json:"reportqp"`
		ReportQPM6       int    `json:"reportqpm6"`
		ReportQPTech     int    `json:"reportqptech"`
	}

	ArsProgram struct {
		ID              string   `json:"id"`
		Name            string   `json:"name"`
		FacultyID       string   `json:"facultyID"`
		Period          string   `json:"period"`
		Section         string   `json:"section"`
		ProgramADM      []string `json:"programADM"`
		Plan            int      `json:"plan"`
		Applicant       int      `json:"applicant"`
		Confirm         int      `json:"confirm"`
		Report          int      `json:"report"`
		ApplicantTCAS   int      `json:"applicanttcas"`
		ApplicantTech   int      `json:"applicanttech"`
		ApplicantQP     int      `json:"applicantqp"`
		ApplicantQPM6   int      `json:"applicantqpm6"`
		ApplicantQPTech int      `json:"applicantqptech"`
		ConfirmTCAS     int      `json:"confirmtcas"`
		ConfirmTech     int      `json:"confirmtech"`
		ConfirmQP       int      `json:"confirmqp"`
		ConfirmQPM6     int      `json:"confirmqpm6"`
		ConfirmQPTech   int      `json:"confirmqptech"`
		ReportTCAS      int      `json:"reporttcas"`
		ReportTech      int      `json:"reporttech"`
		ReportQP        int      `json:"reportqp"`
		ReportQPM6      int      `json:"reportqpm6"`
		ReportQPTech    int      `json:"reportqptech"`
	}

	ArsStdObj struct {
		University ArsUniversity `json:"university"`
		Campus     []ArsCampus   `json:"campus"`
		Faculty    []ArsFaculty  `json:"faculty"`
		Program    []ArsProgram  `json:"program"`
	}

	ArsProgramFac struct {
		ID               string       `json:"id"`
		Name             string       `json:"name"`
		Plan             int          `json:"plan"`
		Applicant        int          `json:"applicant"`
		Confirm          int          `json:"confirm"`
		Report           int          `json:"report"`
		PlanGoal         int          `json:"plangoal"`
		ApplicantGoal    int          `json:"applicantgoal"`
		ConfirmGoal      int          `json:"confirmgoal"`
		ReportGoal       int          `json:"reportgoal"`
		ApplicantPercent int          `json:"applicantpercent"`
		ConfirmPercent   int          `json:"confirmpercent"`
		ReportPercent    int          `json:"reportpercent"`
		ApplicantTCAS    int          `json:"applicanttcas"`
		ApplicantTech    int          `json:"applicanttech"`
		ApplicantQP      int          `json:"applicantqp"`
		ApplicantQPM6    int          `json:"applicantqpm6"`
		ApplicantQPTech  int          `json:"applicantqptech"`
		ConfirmTCAS      int          `json:"confirmtcas"`
		ConfirmTech      int          `json:"confirmtech"`
		ConfirmQP        int          `json:"confirmqp"`
		ConfirmQPM6      int          `json:"confirmqpm6"`
		ConfirmQPTech    int          `json:"confirmqptech"`
		ReportTCAS       int          `json:"reporttcas"`
		ReportTech       int          `json:"reporttech"`
		ReportQP         int          `json:"reportqp"`
		ReportQPM6       int          `json:"reportqpm6"`
		ReportQPTech     int          `json:"reportqptech"`
		Program          []ArsProgram `json:"program"`
	}
)

func (d *ArsStdObj) processData() {
	db, err := dbs.OpenDB(conf.ArsDB.DB)
	switch err {
	case nil:
		d.updateData(db)
		db.Close()
	default:
		log.Printf("Error: ars-processdata - %v\n", err)
	}
}

func (d *ArsStdObj) updateData(db *dbs.DB4ruts) {
	campus := d.getCampus(db)
	faculty := d.getFaculty(db)
	program := d.getProgram(db)
	d.sumFaculty(faculty, program)
	d.sumCampus(campus, faculty)
	university := d.sumUniversity(campus)

	d.University = university
	d.Campus = campus
	d.Faculty = faculty
	d.Program = program
}

func (d *ArsStdObj) getCampus(db *dbs.DB4ruts) []ArsCampus {
	campus := make([]ArsCampus, 0)

	query := "SELECT id, tname FROM campus"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error: ars-getCampus - %v\n", err)
		return campus
	}
	defer rows.Close()

	for rows.Next() {
		id := ""
		tname := ""
		rows.Scan(&id, &tname)
		campus = append(campus, ArsCampus{
			ID:   id,
			Name: tname,
		})
	}

	return campus
}

func (d *ArsStdObj) getFaculty(db *dbs.DB4ruts) []ArsFaculty {
	faculty := make([]ArsFaculty, 0)

	query := "SELECT id, tname, campusid FROM ref_faculty"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error: ars-getFaculty - %v\n", err)
		return faculty
	}
	defer rows.Close()

	for rows.Next() {
		id := ""
		tname := ""
		campusid := ""
		rows.Scan(&id, &tname, &campusid)
		faculty = append(faculty, ArsFaculty{
			ID:       id,
			Name:     tname,
			CampusID: campusid,
		})
	}

	return faculty
}

func (d *ArsStdObj) getProgram(db *dbs.DB4ruts) []ArsProgram {
	section := make(map[string]string)
	period := make(map[string]string)
	program := make([]ArsProgram, 0)

	querysection := "SELECT id, tname FROM section"
	rowssection, errsection := db.Query(querysection)
	if errsection != nil {
		log.Printf("Error: ars-getProgram 1 - %v\n", errsection)
		return program
	}

	for rowssection.Next() {
		id := ""
		name := ""
		rowssection.Scan(&id, &name)
		//section[id] = name
		section[id] = ArsSectionName(id)
	}
	rowssection.Close()

	queryperiod := "SELECT id, tname FROM period"
	rowsperiod, errperiod := db.Query(queryperiod)
	if errperiod != nil {
		log.Printf("Error: ars-getProgram 2 - %v\n", errperiod)
		return program
	}

	for rowsperiod.Next() {
		id := ""
		name := ""
		rowsperiod.Scan(&id, &name)
		period[id] = ArsPeriodName(id)
	}
	rowsperiod.Close()

	query := "SELECT program_id, program_name_th, faculty_id, major_name_th FROM mainprogram"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error: ars-getProgram 3 - %v\n", err)
		return program
	}
	defer rows.Close()

	for rows.Next() {
		program_name := ""
		program_id := ""
		program_name_th := ""
		faculty_id := ""
		major := ""

		rows.Scan(&program_id, &program_name_th, &faculty_id, &major)
		if major == "0" {
			major = ""
		}

		queryplan := "SELECT total, period, section FROM totalplan WHERE programid like ?"
		rowsplan, errplan := db.Query(queryplan, program_id)
		if errplan != nil {
			log.Printf("Error: ars-getProgram 4 - %v\n", errplan)
			return program
		}
		defer rowsplan.Close()

		for rowsplan.Next() {
			plannum := 0
			periodplan := ""
			sectionplan := ""

			rowsplan.Scan(&plannum, &periodplan, &sectionplan)
			if plannum == 0 {
				continue
			}

			query2 := "SELECT distinct id FROM program WHERE programid like ? and period like ? and section like ?"
			rows2, err2 := db.Query(query2, program_id, periodplan, sectionplan)
			if err2 != nil {
				log.Printf("Error: ars-getProgram 5 - %v\n", err2)
				return program
			}

			programadm := make([]string, 0)
			for rows2.Next() {
				programadmid := ""
				rows2.Scan(&programadmid)

				programadm = append(programadm, programadmid)
			}
			rows2.Close()

			applicant := 0
			applicantqp := 0
			applicanttech := 0
			applicanttcas := 0
			confirm := 0
			confirmqp := 0
			confirmtech := 0
			confirmtcas := 0
			report := 0
			reportqp := 0
			reporttech := 0
			reporttcas := 0
			for i := range programadm {
				count := 0
				queryx1 := "SELECT count(*) FROM applicant_apply WHERE applyprogramid like ? and payment like 'Y'"
				rowsx1, errx1 := db.Query(queryx1, programadm[i])
				if errx1 != nil {
					log.Printf("Error: ars-getProgram 6 - %v\n", errx1)
					return program
				}

				for rowsx1.Next() {
					cc := 0
					rowsx1.Scan(&cc)
					count += cc
					break
				}
				applicant += count
				rowsx1.Close()

				count = 0
				queryz1 := "SELECT count(*) FROM applicant_apply WHERE applyprogramid like ? and roundid like '9' and payment like 'Y'"
				rowsz1, errz1 := db.Query(queryz1, programadm[i])
				if errz1 != nil {
					log.Printf("Error: ars-getProgram 7 - %v\n", errz1)
					return program
				}

				for rowsz1.Next() {
					cc := 0
					rowsz1.Scan(&cc)
					count += cc
					break
				}
				applicantqp += count
				rowsz1.Close()

				count = 0
				queryt1 := "SELECT count(*) FROM applicant_apply WHERE applyprogramid like ? and (roundid like '6' or roundid like '7' or roundid like '8') and payment like 'Y'"
				rowst1, errt1 := db.Query(queryt1, programadm[i])
				if errt1 != nil {
					log.Printf("Error: ars-getProgram 8 - %v\n", errt1)
					return program
				}

				for rowst1.Next() {
					cc := 0
					rowst1.Scan(&cc)
					count += cc
					break
				}
				applicanttech += count
				rowst1.Close()

				count = 0
				queryc1 := "SELECT count(*) FROM applicant_apply WHERE applyprogramid like ? and (roundid like '1' or roundid like '2' or roundid like '3' or roundid like '4') and payment like 'Y'"
				rowsc1, errc1 := db.Query(queryc1, programadm[i])
				if errc1 != nil {
					log.Printf("Error: ars-getProgram 9 - %v\n", errc1)
					return program
				}

				for rowsc1.Next() {
					cc := 0
					rowsc1.Scan(&cc)
					count += cc
					break
				}
				applicanttcas += count
				rowsc1.Close()

				count = 0
				queryx2 := "SELECT count(*) FROM applicant_apply WHERE applyprogramid like ? and confirm like 'Y' and payment like 'Y'"
				rowsx2, errx2 := db.Query(queryx2, programadm[i])
				if errx2 != nil {
					log.Printf("Error: ars-getProgram 10 - %v\n", errx2)
					return program
				}

				for rowsx2.Next() {
					cc := 0
					rowsx2.Scan(&cc)
					count += cc
					break
				}
				confirm += count
				rowsx2.Close()

				count = 0
				queryz2 := "SELECT count(*) FROM applicant_apply WHERE applyprogramid like ? and confirm like 'Y' and roundid like '9' and payment like 'Y'"
				rowsz2, errz2 := db.Query(queryz2, programadm[i])
				if errz2 != nil {
					log.Printf("Error: ars-getProgram 11 - %v\n", errz2)
					return program
				}

				for rowsz2.Next() {
					cc := 0
					rowsz2.Scan(&cc)
					count += cc
					break
				}
				confirmqp += count
				rowsz2.Close()

				count = 0
				queryt2 := "SELECT count(*) FROM applicant_apply WHERE applyprogramid like ? and confirm like 'Y' and (roundid like '6' or roundid like '7' or roundid like '8') and payment like 'Y'"
				rowst2, errt2 := db.Query(queryt2, programadm[i])
				if errt2 != nil {
					log.Printf("Error: ars-getProgram 12 - %v\n", errt2)
					return program
				}

				for rowst2.Next() {
					cc := 0
					rowst2.Scan(&cc)
					count += cc
					break
				}
				confirmtech += count
				rowst2.Close()

				count = 0
				queryc2 := "SELECT count(*) FROM applicant_apply WHERE applyprogramid like ? and confirm like 'Y' and (roundid like '1' or roundid like '2' or roundid like '3' or roundid like '4') and payment like 'Y'"
				rowsc2, errc2 := db.Query(queryc2, programadm[i])
				if errc2 != nil {
					log.Printf("Error: ars-getProgram 13 - %v\n", errc2)
					return program
				}

				for rowsc2.Next() {
					cc := 0
					rowsc2.Scan(&cc)
					count += cc
					break
				}
				confirmtcas += count
				rowsc2.Close()

				count = 0
				queryx3 := "SELECT count(*) FROM applicant_apply WHERE applyprogramid like ? and report like 'Y' and payment like 'Y'"
				rowsx3, errx3 := db.Query(queryx3, programadm[i])
				if errx3 != nil {
					log.Printf("Error: ars-getProgram 14 - %v\n", errx3)
					return program
				}

				for rowsx3.Next() {
					cc := 0
					rowsx3.Scan(&cc)
					count += cc
					break
				}
				report += count
				rowsx3.Close()

				count = 0
				queryz3 := "SELECT count(*) FROM applicant_apply WHERE applyprogramid like ? and report like 'Y' and roundid like '9' and payment like 'Y'"
				rowsz3, errz3 := db.Query(queryz3, programadm[i])
				if errz3 != nil {
					log.Printf("Error: ars-getProgram 15 - %v\n", errz3)
					return program
				}

				for rowsz3.Next() {
					cc := 0
					rowsz3.Scan(&cc)
					count += cc
					break
				}
				reportqp += count
				rowsz3.Close()

				count = 0
				queryt3 := "SELECT count(*) FROM applicant_apply WHERE applyprogramid like ? and report like 'Y' and (roundid like '6' or roundid like '7' or roundid like '8') and payment like 'Y'"
				rowst3, errt3 := db.Query(queryt3, programadm[i])
				if errt3 != nil {
					log.Printf("Error: ars-getProgram 16 - %v\n", errt3)
					return program
				}

				for rowst3.Next() {
					cc := 0
					rowst3.Scan(&cc)
					count += cc
					break
				}
				reporttech += count
				rowst3.Close()

				count = 0
				queryc3 := "SELECT count(*) FROM applicant_apply WHERE applyprogramid like ? and report like 'Y' and (roundid like '1' or roundid like '2' or roundid like '3' or roundid like '4') and payment like 'Y'"
				rowsc3, errc3 := db.Query(queryc3, programadm[i])
				if errc3 != nil {
					log.Printf("Error: ars-getProgram 17 - %v\n", errc3)
					return program
				}

				for rowsc3.Next() {
					cc := 0
					rowsc3.Scan(&cc)
					count += cc
					break
				}
				reporttcas += count
				rowsc3.Close()
			}

			program_name = program_name_th
			if major != "" {
				program_name += " - " + major
			}

			program_name += " (" + period[periodplan] + " " + section[sectionplan] + ") "

			program = append(program, ArsProgram{
				ID:            program_id,
				Name:          program_name,
				FacultyID:     faculty_id,
				Period:        periodplan,
				Section:       sectionplan,
				ProgramADM:    programadm,
				Plan:          plannum,
				Applicant:     applicant,
				Confirm:       confirm,
				Report:        report,
				ApplicantQP:   applicantqp,
				ApplicantTech: applicanttech,
				ApplicantTCAS: applicanttcas,
				ConfirmQP:     confirmqp,
				ConfirmTech:   confirmtech,
				ConfirmTCAS:   confirmtcas,
				ReportQP:      reportqp,
				ReportTech:    reporttech,
				ReportTCAS:    reporttcas,
			})
		}
	}
	return program
}

func (d *ArsStdObj) sumFaculty(faculty []ArsFaculty, program []ArsProgram) {
	fac := make(map[string]*ArsFaculty)
	for i := range faculty {
		fac[faculty[i].ID] = &faculty[i]
	}
	for i := range program {
		fac[program[i].FacultyID].Plan += program[i].Plan
		fac[program[i].FacultyID].Applicant += program[i].Applicant
		fac[program[i].FacultyID].ApplicantQP += program[i].ApplicantQP
		fac[program[i].FacultyID].ApplicantTech += program[i].ApplicantTech
		fac[program[i].FacultyID].ApplicantTCAS += program[i].ApplicantTCAS
		fac[program[i].FacultyID].Confirm += program[i].Confirm
		fac[program[i].FacultyID].ConfirmQP += program[i].ConfirmQP
		fac[program[i].FacultyID].ConfirmTech += program[i].ConfirmTech
		fac[program[i].FacultyID].ConfirmTCAS += program[i].ConfirmTCAS
		fac[program[i].FacultyID].Report += program[i].Report
		fac[program[i].FacultyID].ReportQP += program[i].ReportQP
		fac[program[i].FacultyID].ReportTech += program[i].ReportTech
		fac[program[i].FacultyID].ReportTCAS += program[i].ReportTCAS
		fac[program[i].FacultyID].PlanGoal++
		if program[i].Plan <= program[i].Applicant {
			fac[program[i].FacultyID].ApplicantGoal++
		}
		if program[i].Plan <= program[i].Confirm {
			fac[program[i].FacultyID].ConfirmGoal++
		}
		if program[i].Plan <= program[i].Report {
			fac[program[i].FacultyID].ReportGoal++
		}
	}
	for i := range faculty {
		faculty[i].ApplicantPercent = faculty[i].ApplicantGoal * 100 / faculty[i].PlanGoal
		faculty[i].ConfirmPercent = faculty[i].ConfirmGoal * 100 / faculty[i].PlanGoal
		faculty[i].ReportPercent = faculty[i].ReportGoal * 100 / faculty[i].PlanGoal
	}
}

func (d *ArsStdObj) sumCampus(campus []ArsCampus, faculty []ArsFaculty) {
	cam := make(map[string]*ArsCampus)
	for i := range campus {
		cam[campus[i].ID] = &campus[i]
	}
	for i := range faculty {
		cam[faculty[i].CampusID].Plan += faculty[i].Plan
		cam[faculty[i].CampusID].Applicant += faculty[i].Applicant
		cam[faculty[i].CampusID].ApplicantQP += faculty[i].ApplicantQP
		cam[faculty[i].CampusID].ApplicantTech += faculty[i].ApplicantTech
		cam[faculty[i].CampusID].ApplicantTCAS += faculty[i].ApplicantTCAS
		cam[faculty[i].CampusID].Confirm += faculty[i].Confirm
		cam[faculty[i].CampusID].ConfirmQP += faculty[i].ConfirmQP
		cam[faculty[i].CampusID].ConfirmTech += faculty[i].ConfirmTech
		cam[faculty[i].CampusID].ConfirmTCAS += faculty[i].ConfirmTCAS
		cam[faculty[i].CampusID].Report += faculty[i].Report
		cam[faculty[i].CampusID].ReportQP += faculty[i].ReportQP
		cam[faculty[i].CampusID].ReportTech += faculty[i].ReportTech
		cam[faculty[i].CampusID].ReportTCAS += faculty[i].ReportTCAS
		cam[faculty[i].CampusID].PlanGoal += faculty[i].PlanGoal
		cam[faculty[i].CampusID].ApplicantGoal += faculty[i].ApplicantGoal
		cam[faculty[i].CampusID].ConfirmGoal += faculty[i].ConfirmGoal
		cam[faculty[i].CampusID].ReportGoal += faculty[i].ReportGoal
	}
	for i := range campus {
		campus[i].ApplicantPercent = campus[i].ApplicantGoal * 100 / campus[i].PlanGoal
		campus[i].ConfirmPercent = campus[i].ConfirmGoal * 100 / campus[i].PlanGoal
		campus[i].ReportPercent = campus[i].ReportGoal * 100 / campus[i].PlanGoal
	}
}

func (d *ArsStdObj) sumUniversity(campus []ArsCampus) ArsUniversity {
	university := ArsUniversity{}
	for i := range campus {
		university.Plan += campus[i].Plan
		university.Applicant += campus[i].Applicant
		university.ApplicantQP += campus[i].ApplicantQP
		university.ApplicantTech += campus[i].ApplicantTech
		university.ApplicantTCAS += campus[i].ApplicantTCAS
		university.Confirm += campus[i].Confirm
		university.ConfirmQP += campus[i].ConfirmQP
		university.ConfirmTech += campus[i].ConfirmTech
		university.ConfirmTCAS += campus[i].ConfirmTCAS
		university.Report += campus[i].Report
		university.ReportQP += campus[i].ReportQP
		university.ReportTech += campus[i].ReportTech
		university.ReportTCAS += campus[i].ReportTCAS
		university.PlanGoal += campus[i].PlanGoal
		university.ApplicantGoal += campus[i].ApplicantGoal
		university.ConfirmGoal += campus[i].ConfirmGoal
		university.ReportGoal += campus[i].ReportGoal

		university.ApplicantPercent = campus[i].ApplicantGoal / campus[i].PlanGoal
		university.ConfirmPercent = campus[i].ConfirmGoal / campus[i].PlanGoal
		university.Report = campus[i].ReportGoal / campus[i].PlanGoal
	}
	return university
}

func (d *ArsStdObj) exportfaculty(fac string) ArsProgramFac {
	program := ArsProgramFac{
		Program: make([]ArsProgram, 0),
		ID:      fac,
	}

	for i := range d.Faculty {
		if d.Faculty[i].ID == fac {
			program.Name = d.Faculty[i].Name
			program.Plan = d.Faculty[i].Plan
			program.PlanGoal = d.Faculty[i].PlanGoal
			program.Applicant = d.Faculty[i].Applicant
			program.ApplicantQP = d.Faculty[i].ApplicantQP
			program.ApplicantTech = d.Faculty[i].ApplicantTech
			program.ApplicantTCAS = d.Faculty[i].ApplicantTCAS
			program.ApplicantGoal = d.Faculty[i].ApplicantGoal
			program.ApplicantPercent = d.Faculty[i].ApplicantPercent
			program.Confirm = d.Faculty[i].Confirm
			program.ConfirmQP = d.Faculty[i].ConfirmQP
			program.ConfirmTech = d.Faculty[i].ConfirmTech
			program.ConfirmTCAS = d.Faculty[i].ConfirmTCAS
			program.ConfirmGoal = d.Faculty[i].ConfirmGoal
			program.ConfirmPercent = d.Faculty[i].ConfirmPercent
			program.Report = d.Faculty[i].Report
			program.ReportQP = d.Faculty[i].ReportQP
			program.ReportTech = d.Faculty[i].ReportTech
			program.ReportTCAS = d.Faculty[i].ReportTCAS
			program.ReportGoal = d.Faculty[i].ReportGoal
			program.ReportPercent = d.Faculty[i].ReportPercent
			break
		}
	}

	for i := range d.Program {
		if d.Program[i].FacultyID == fac {
			program.Program = append(program.Program, d.Program[i])
		}
	}

	return program
}

func ArsSectionName(n string) string {
	switch n {
	case "A":
		return "สมทบ"
	case "N":
		return "ปกติ"
	default:
		return ""
	}
}

func ArsPeriodName(n string) string {
	switch n {
	case "1":
		return "ปวส."
	case "2":
		return "4 ปี"
	case "3":
		return "5 ปี"
	case "4":
		return "6 ปี"
	case "5":
		return "เทียบโอน"
	case "6":
		return "ต่อเนื่อง"
	case "7":
		return "ป.โท"
	default:
		return ""
	}
}

func ArsExport(ctx *fiber.Ctx) error {
	return ctx.JSON(ArsReadFromDB())
}

func ArsFacExport(ctx *fiber.Ctx) error {
	return ctx.JSON(ArsReadFacFromDB(ctx.Params("facid")))
}

func ArsReadFromDB() (output ArsStdObj) {
	data := ReadCache("ARS")
	if data == "" {
		return
	}
	err := json.Unmarshal([]byte(data), &output)
	if err != nil {
		return
	}
	return
}

func ArsReadFacFromDB(fac string) (output ArsProgramFac) {
	data := ReadCache("ARS")
	if data == "" {
		return
	}
	dataUniver := ArsStdObj{}
	err := json.Unmarshal([]byte(data), &dataUniver)
	if err != nil {
		return
	}
	output = dataUniver.exportfaculty(fac)
	return
}

func ArsProcess(ctx *fiber.Ctx) error {
	switch CheckOTP(ctx.Params("otp")) {
	case true:
		data := ArsStdObj{}
		data.processData()
		jdata, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error: ars-ArsProcess - %v\n", err)
			return ctx.JSON(fiber.Map{
				"status": "json",
			})
		}
		SaveCache("ARS", string(jdata))
		return ctx.JSON(fiber.Map{
			"status": "ok",
		})
	default:
		return ctx.JSON(fiber.Map{
			"status": "otp",
		})
	}
}

func ArsClean(ctx *fiber.Ctx) error {
	switch CheckOTP(ctx.Params("otp")) {
	case true:
		CleanCache("ARS")
		return ctx.JSON(fiber.Map{
			"status": "ok",
		})
	default:
		return ctx.JSON(fiber.Map{
			"status": "otp",
		})
	}
}
