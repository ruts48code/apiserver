package main

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	utils "github.com/ruts48code/utils4ruts"
)

func main() {
	if utils.FileExist("/etc/apiserver.hcl") {
		utils.ProcessConfigHCL("/etc/apiserver.hcl", &conf)
		log.Printf("Load /etc/apiserver.hcl sucessfully\n")
	} else if utils.FileExist("/etc/apiserver.yml") {
		utils.ProcessConfig("/etc/apiserver.yml", &conf)
		log.Printf("Load /etc/apiserver.yml sucessfully\n")
	} else {
		log.Printf("Error: cannot load configurationfile\n")
		return
	}

	app := fiber.New(fiber.Config{
		Views:          html.New("./template", ".html"),
		ProxyHeader:    fiber.HeaderXForwardedFor,
		ReadBufferSize: 20480,
		BodyLimit:      104857600,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			log.Printf("E: %d - %s\n", code, ctx.Path())
			switch code {
			case 404:
				return ctx.Render("404", fiber.Map{})
			default:
				ctx.Status(code).SendString(err.Error())
			}

			return nil
		},
	})

	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
	}))

	app.Static("/", "./static")
	app.Get("/", index)
	app.Get("/counter/:name", counterup)
	app.Get("/counter/:name/get", counterget)

	app.Post("/elogin", elogin)
	app.Get("/elogin/delete/:username", eloginDelete)
	app.Get("/elogin/token/:token", eloginToken)
	app.Get("/elogin/clean/:otp", eloginClean)

	app.Get("/personal", PersonalCode)
	app.Get("/personal/academic/:token", PersonalAcademicPrivate)
	app.Post("/personal/academic/:token", PersonalAcademicPrivilege)

	app.Get("/teacher/supervisor/:token", SuperVisor)
	app.Get("/teacher/class/:classid/:token", SuperVisorClass)
	app.Get("/teacher/trace/:classid/:token", SuperVisorTrace)

	app.Get("/student/grade/:id/:token", StudentGrade)
	app.Get("/student/regis/:id/:token", StudentRegis)

	app.Get("/student/report/processalldata/:otp", StudentProcessAllData)
	app.Get("/student/report/getalldata/:token", StudentGetAllData)
	app.Get("/student/report/cleanalldata/:otp", StudentCleanAllData)

	app.Get("/ars/export", ArsExport)
	app.Get("/ars/faculty/:facid", ArsFacExport)
	app.Get("/ars/process/:otp", ArsProcess)
	app.Get("/ars/clean/:otp", ArsClean)
	app.Get("/ars/table", ArsTable)

	app.Post("/mcas/mail/:token", SendMail)

	app.Get("/openathens", OpenAthensLogin)
	app.Get("/404", e404)
	log.Fatal(app.Listen(conf.Listen))
}

func index(c *fiber.Ctx) error {
	return c.Render("index", nil)
}

func e404(c *fiber.Ctx) error {
	return c.Render("404", fiber.Map{})
}
