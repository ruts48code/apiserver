package main

type (
	Conf struct {
		Listen     string           `yaml:"listen"`
		OTP        OTPStruct        `yaml:"otp"`
		DBS        []string         `yaml:"dbs"`
		Elogin     EloginStruct     `yaml:"elogin"`
		Personal   PersonalStruct   `yaml:"personal"`
		Student    StudentStruct    `yaml:"student"`
		OpenAthens OpenAthensStruct `yaml:"openathens"`
	}

	OTPStruct struct {
		Key      string `yaml:"key"`
		Size     int    `yaml:"size"`
		Interval int    `yaml:"interval"`
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
		Cache  StudentCacheStruct `yaml:"cache"`
		Server []SisServerStruct  `yaml:"server"`
	}

	StudentCacheStruct struct {
		Interval int `yaml:"interval"`
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
