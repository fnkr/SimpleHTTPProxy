package main

import (
	"flag"
	"strings"
	"regexp"
	"log"
	"encoding/json"
)

type Options struct {
	Verbose bool
	Bind string
	Users map[string]string
	Hosts []*regexp.Regexp
}

type userFlag map[string]string

type hostFlag []*regexp.Regexp

type hostRegexFlag []*regexp.Regexp

func (u *userFlag) String() string {
	return "userFlag.String()"
}

func (u *userFlag) Set(value string) error {
	if *u == nil {
		*u = userFlag{}
	}
	user_pass := strings.SplitN(value, ":", 2)
	if len(user_pass) > 1 {
		(*u)[user_pass[0]] = user_pass[1]
	} else {
		(*u)[user_pass[0]] = ""
	}
	return nil
}

func (d *hostFlag) String() string {
	// Not actually in use, but needs to be implemented
	return "hostFlag.String()"
}

func (d *hostFlag) Set(value string) error {
	*d = append(*d, regexp.MustCompile("^" + regexp.QuoteMeta(value) + "$"))
	return nil
}

func (d *hostRegexFlag) String() string {
	// Not actually in use, but needs to be implemented
	return "hostRegexFlag.String()"
}

func (d *hostRegexFlag) Set(value string) error {
	regex, err := regexp.Compile(value)
	if err != nil {
		return err
	}
	*d = append(*d, regex)
	return nil
}

func parseArgs() *Options {
	verbose := flag.Bool("verbose", false, "")
	bind := flag.String("bind", ":8080", "")

	var users userFlag
	flag.Var(&users, "user", "e.g. \"admin:password\", can be used multiple times")

	var hosts hostFlag
	flag.Var(&hosts, "host", "e.g. \"example.com\", can be used multiple times")

	var hostregex hostRegexFlag
	flag.Var(&hostregex, "host-match", "regular expression, e.g. \"^*\\.example\\.com$\", can be used multiple times")

	flag.Parse()

	for _, host := range hostregex {
		hosts = append(hosts, host)
	}

	opts := Options{
		Verbose: *verbose,
		Bind: *bind,
		Users: users,
		Hosts: hosts,
	}

	if *verbose {
		js, _ := json.Marshal(opts)
		log.Println("Options: " + string(js))
	}

	return &opts
}
