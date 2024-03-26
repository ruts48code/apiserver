package main

type (
	Conf struct {
		Listen     string           `yaml:"listen" hcl:"listen"`
		OTP        OTPStruct        `yaml:"otp" hcl:"otp,block"`
		DBS        []string         `yaml:"dbs" hcl:"dbs"`
		Elogin     EloginStruct     `yaml:"elogin" hcl:"elogin,block"`
		Personal   PersonalStruct   `yaml:"personal" hcl:"personal,block"`
		Student    StudentStruct    `yaml:"student" hcl:"student,block"`
		OpenAthens OpenAthensStruct `yaml:"openathens" hcl:"openathens,block"`
		ArsDB      ARSStruct        `yaml:"ars" hcl:"ars,block"`
	}

	OTPStruct struct {
		Key      string `yaml:"key" hcl:"key,label"`
		Size     int    `yaml:"size" hcl:"size"`
		Interval int    `yaml:"interval" hcl:"interval"`
	}

	EloginStruct struct {
		LDAPDomain string   `yaml:"ldapdomain" hcl:"ldapdomain"`
		LDAPServer []string `yaml:"ldapserver" hcl:"ldapserver"`
		Expire     int      `yaml:"expire" hcl:"expire"`
		Clean      int      `yaml:"clean" hcl:"clean"`
		TokenSize  int      `yaml:"tokensize" hcl:"tokensize"`
		Limit      int      `yaml:"limit" hcl:"limit"`
	}

	PersonalStruct struct {
		Server     string           `yaml:"server" hcl:"server"`
		Permission PermissionStruct `yaml:"permission" hcl:"permission,block"`
	}

	StudentStruct struct {
		Cache  StudentCacheStruct `yaml:"cache" hcl:"cache,block"`
		Server []SisServerStruct  `yaml:"server" hcl:"server,block"`
	}

	StudentCacheStruct struct {
		Update int `yaml:"update" hcl:"update"`
		Clean  int `yaml:"clean" hcl:"clean"`
	}

	SisServerStruct struct {
		ID     string `yaml:"id" hcl:"id,label"`
		Name   string `yaml:"name" hcl:"name,label"`
		Server string `yaml:"server" hcl:"server"`
	}

	PermissionStruct struct {
		ReadAll []string `yaml:"readAll" hcl:"readAll"`
	}

	OpenAthensStruct struct {
		ConnectionID  string `yaml:"connectionid" hcl:"connectionid"`
		ConnectionURI string `yaml:"connectionuri" hcl:"connectionuri"`
		ReturnURL     string `yaml:"returnurl" hcl:"returnurl"`
		APIKey        string `yaml:"apikey" hcl:"apikey"`
	}

	ARSStruct struct {
		DB     string `yaml:"db" hcl:"db"`
		Update int    `yaml:"update" hcl:"update"`
		Clean  int    `yaml:"clean" hcl:"clean"`
	}
)

var (
	conf Conf
)
