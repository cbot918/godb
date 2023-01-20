package godb

import (
	"database/sql"
	"fmt"
	u "godb/util"

	_ "github.com/lib/pq"
)

type GODB struct {
	db *sql.DB
}

func New(
	db string,
	host string,
	port int,
	user string,
	password string,
	maxOpen int,
	maxIdle int,
) *GODB {

	d := new(GODB)

	// if driver == "mysql" {
	// 	connectString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, use)
	// }

	// postgres string
	var connectString string
	connectString = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, db)

	var err error
	d.db, err = sql.Open("postgres", connectString)
	u.Checke(err, "sql.Open failed")

	return d
}

func (godb *GODB) Ping() interface{} {
	var result interface{}
	err := godb.db.Ping()
	if err != nil {
		result = err
	} else {
		result = "db.Ping success"
	}
	return result
}
