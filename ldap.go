package main

import (
	"log"

	util "github.com/ruts48code/utils4ruts"

	ldapserv "github.com/mavricknz/ldap"
)

func getLDAP() []string {
	return util.RandomArrayString(conf.Elogin.LDAPServer)
}

func ldapLogin(username, password string) string {
	if username == "" {
		log.Printf("Error: LDAP nil username\n")
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
			log.Printf("Error: LDAP [%s] for %s - %v \n", ldapN[i], username, err)
			l.Close()
			continue
		}
		ldapNum = i
		log.Printf("Log: Connect to LDAP: %s for %s\n", ldapN[i], username)
		break
	}
	if ldapNum == -1 {
		return "none"
	}
	defer l.Close()

	err := l.Bind("", "")
	if err != nil {
		log.Printf("Error: LDAP [%s] for %s - %v\n", ldapN, username, err)
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
		log.Printf("Error: LDAP [%s] for %s - %v\n", ldapN, username, err)
		return "none"
	}

	if sr.Entries == nil {
		log.Printf("Error: LDAP [%s] for %s - %v\n", ldapN, username, err)
		return "none"
	}

	err = l.Bind(sr.Entries[0].DN, password)
	if err != nil {
		log.Printf("Error: LDAP [%s] for %s - %v\n", ldapN, username, err)
		return "none"
	}

	return sr.Entries[0].GetAttributeValues("cn")[0]
}
