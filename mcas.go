package main

import (
	"crypto/tls"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	gomail "gopkg.in/mail.v2"
)

type (
	SendMailOutStruct struct {
		Status string `json:"status"`
	}
	SendMailStruct struct {
		To      string `json:"to"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}
)

func SendMail(ctx *fiber.Ctx) error {
	username, name, email, status := CheckTKWeb(ctx.Params("token"))
	switch status {
	case "ok":
		data := SendMailStruct{}
		err := json.Unmarshal(ctx.Body(), &data)
		if err != nil {
			log.Printf("Error: mcas-SendMail - json error api\n")
			return ctx.JSON(SendMailOutStruct{
				Status: "json",
			})
		}
		return ctx.JSON(SendMailAPI(username, name, email, data))
	default:
		output := SendMailOutStruct{
			Status: status,
		}
		return ctx.JSON(output)
	}
}

func SendMailAPI(username string, name string, email string, data SendMailStruct) (output SendMailOutStruct) {
	output = SendMailOutStruct{
		Status: "ok",
	}

	originsend := "ข้อความจาก : " + name + " (" + email + ")\n\n"
	m := gomail.NewMessage()
	m.SetHeader("From", "mcas@rmutsv.ac.th")
	m.SetHeader("To", data.To)
	m.SetHeader("CC", email)
	m.SetHeader("Reply-to", email)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/plain", originsend+data.Body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "mcas@rmutsv.ac.th", "rmut$v2s48")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error: mcas-SendMailAPI - send mail error : %v\n", err)
		output.Status = "mail"
	}
	return
}
