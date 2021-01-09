package config

import "fmt"

const (
	USERNAME = "root"
	PASSWORD = "passwd"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "subject"
)

var (
	MYSQL_CONN_STR = ""
)

func init() {
	MYSQL_CONN_STR = fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=true", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
}
