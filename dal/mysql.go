package dal

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func InitMysql(dns string) {
	var err error
	Db, err = sqlx.Open("mysql", dns)
	if err != nil {
		panic(err)
	}

	err = Db.Ping()
	if err != nil {
		panic(err)
	}
	Db.SetMaxIdleConns(16)
	Db.SetMaxOpenConns(100)
}
