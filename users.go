package main

import (
	"encoding/base64"
	"strings"
)

type ProxyAuthorization struct {
	Username string
	Password string
}

func parseProxyAuthorization(s string) ProxyAuthorization {
	a := strings.SplitN(s, " ", 2)

	switch m := strings.ToLower(a[0]); m {
	case "basic":
		if len(a) != 2 {
			return ProxyAuthorization{}
		}

		db, err := base64.StdEncoding.DecodeString(a[1])
		if err != nil {
			return ProxyAuthorization{}
		}

		da := strings.SplitN(string(db), ":", 2)
		if len(da) != 2 {
			return ProxyAuthorization{}
		}

		return ProxyAuthorization{
			Username: da[0],
			Password: da[1],
		}
	}

	return ProxyAuthorization{}
}

func checkUser(auth ProxyAuthorization) bool {
	if x, ok := config.Users[auth.Username]; ok && x == auth.Password {
		return true
	}
	return false
}
