listen = ":80"
otp "secretword" {
  size = 24
  interval = 30
}
dbs = [
  "mysql://user:pass@tcp(10.5.0.1)/database?parseTime=true",
  "mysql://user:pass@tcp(10.5.0.2)/database?parseTime=true",
  "mysql://user:pass@tcp(10.5.0.3)/database?parseTime=true",
  "mysql://user:pass@tcp(10.5.0.4)/database?parseTime=true",
]
elogin {
  ldapdomain = "dc=example,dc=ac,dc=th"
  ldapserver = [
    "10.1.0.1",
    "10.1.0.2",
    "10.1.0.3",
  ]
  expire = 3600
  clean = 60
  tokensize = 64
  limit = 10
}
personal {
  server = "sqlserver://user:password@10.0.0.1/?database=dbname"
  permission {
    readAll = ["someone.x",]
  }
}
student {
  cache {
    update = 3600
    clean = 60
  }
  server "1" "Songkhla" {
    server = "mysql://user:password@tcp(127.0.0.1)/songkhla?parseTime=true"
  }
  server "2" "Rattaphum" {
    server = "mysql://user:password@tcp(127.0.0.1)/rattaphum?parseTime=true"
  }
}
ars {
  db = "mysql://user:pass@tcp(127.0.0.1)/RUTSAdmission?parseTime=true"
  update = 180
  clean = 60
}
openathens {
  connectionid = "12345"
  connectionuri = "https://login.openathens.net/api/v1/{domain}/organisation/{oid}/local-auth/session"
  returnurl = "https://example.com"
  apikey = "12345"
}