package libs

import (
	"database/sql"
	"moviedemo/config"

	_ "github.com/go-sql-driver/mysql"
	xerrors "github.com/pkg/errors"
)

func GetMySQLDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.MYSQL_CONN_STR)
	return db, xerrors.Wrapf(err, "open db connection failed: %v", config.MYSQL_CONN_STR)
}
