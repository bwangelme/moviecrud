package config

import (
	"fmt"
	"moviedemo/config/env"
)

const (
	USERNAME = "root"
	PASSWORD = "passwd"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
)

var (
	MYSQL_CONN_STR = ""
	DATABASE       = ""
)

func init() {
	if env.IsTest() {
		DATABASE = "test_subject"
	} else {
		DATABASE = "subject"
	}
	MYSQL_CONN_STR = fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=true", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
}
