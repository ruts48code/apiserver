package main

import (
	"log"

	utils "github.com/ruts48code/utils4ruts"

	ldapserv "github.com/mavricknz/ldap"
)

func getLDAP() []string {
	return utils.RandomArrayString(conf.Elogin.LDAPServer)
}

func ldapLogin(username, password string) string {
	if username == "" {
		log.Printf("Log: ldap-ldapLogin 1 - LDAP username is empy string\n")
		return "none"
	}

	var attributes []string = []string{"uid", "cn"}
	ldapN := getLDAP()
	var l *ldapserv.LDAPConnection
	ldapNum := -1

	for i := range ldapN {
		l = ldapserv.NewLDAPConnection(ldapN[i], 389)
		err := l.Connect()
		if err != nil {
			log.Printf("Error: ldap-ldapLogin 2 - fail to connect LDAP server [%s] for %s - %v \n", ldapN[i], username, err)
			l.Close()
			continue
		}
		ldapNum = i
		break
	}
	if ldapNum == -1 {
		log.Printf("Log: ldap-ldapLogin - Cannot connect to any ldap server\n")
		return "none"
	}
	defer l.Close()

	err := l.Bind("", "")
	if err != nil {
		log.Printf("Error: ldap-ldapLogin 3 - fail to anonymous bind LDAP server [%s] for %s - %v\n", ldapN, username, err)
		return "none"
	}

	search_request := ldapserv.NewSearchRequest(
		conf.Elogin.LDAPDomain,
		ldapserv.ScopeWholeSubtree, ldapserv.DerefAlways, 0, 0, false,
		"(uid="+username+")",
		attributes,
		nil)

	sr, err := l.SearchWithPaging(search_request, 5)
	if err != nil {
		log.Printf("Error: ldap-ldapLogin 4 - fail to search in LDAP server [%s] for %s - %v\n", ldapN, username, err)
		return "none"
	}

	if sr.Entries == nil {
		log.Printf("Error: ldap-ldapLogin 5 - no entry in LDAP server [%s] for %s - %v\n", ldapN, username, err)
		return "none"
	}

	err = l.Bind(sr.Entries[0].DN, password)
	if err != nil {
		log.Printf("Error: ldap-ldapLogin 6 - fail to direct bind LDAP server [%s] for %s - %v\n", ldapN, username, err)
		return "none"
	}

	return sr.Entries[0].GetAttributeValues("cn")[0]
}
