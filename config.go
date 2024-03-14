package main

import (
	util "github.com/ruts48code/utils4ruts"
	"gopkg.in/yaml.v3"
)

type (
	Conf struct {
		Listen     string           `yaml:"listen"`
		DBUsername string           `yaml:"dbusername"`
		DBPassword string           `yaml:"dbpassword"`
		DBName     string           `yaml:"dbname"`
		DBParam    string           `yaml:"dbparam"`
		DBType     string           `yaml:"dbtype"`
		DBS        []string         `yaml:"dbs"`
		Elogin     EloginStruct     `yaml:"elogin"`
		Personal   PersonalStruct   `yaml:"personal"`
		Student    StudentStruct    `yaml:"student"`
		OpenAthens OpenAthensStruct `yaml:"openathens"`
	}

	EloginStruct struct {
		LDAPDomain string   `yaml:"ldapdomain"`
		LDAPServer []string `yaml:"ldapserver"`
		Expire     int      `yaml:"expire"`
		Clean      int      `yaml:"clean"`
		TokenSize  int      `yaml:"tokensize"`
	}

	PersonalStruct struct {
		Server     string           `yaml:"server"`
		Permission PermissionStruct `yaml:"permission"`
	}

	StudentStruct struct {
		Server []SisServerStruct `yaml:"server"`
	}

	SisServerStruct struct {
		Name   string `yaml:"name"`
		ID     string `yaml:"id"`
		Server string `yaml:"server"`
	}

	PermissionStruct struct {
		ReadAll []string `yaml:"readAll"`
	}

	OpenAthensStruct struct {
		ConnectionID  string `yaml:"connectionid"`
		ConnectionURI string `yaml:"connectionuri"`
		ReturnURL     string `yaml:"returnurl"`
		APIKey        string `yaml:"apikey"`
	}
)

var (
	conf Conf
)

func processConfig() {
	confdata := util.ReadFile("/etc/apiserver.yml")
	yaml.Unmarshal(confdata, &conf)
}
