package database

import (
	"database/sql"

	"github.com/astaxie/beego"
)

func sqlConnector() (*sql.DB, error) {
	ip := beego.AppConfig.String("mysqlurls")
	dbName := beego.AppConfig.String("mysqldb")
	user := beego.AppConfig.String("mysqluser")
	pass := beego.AppConfig.String("mysqlpass")
	db, err := sql.Open("mysql", ""+user+":"+pass+"@tcp("+ip+")/"+dbName+"")
	return db, err
}

// SQLSingleRowQuery : Single Record Query
func SQLSingleRowQuery(query string, params ...interface{}) *sql.Row {
	db, _ := sqlConnector()
	row := db.QueryRow(query, params...)
	return row
}

// SQLQuery : Multiple Records Query
func SQLQuery(query string, params ...interface{}) (*sql.Rows, error) {
	db, _ := sqlConnector()

	var rows *sql.Rows
	var err error

	if params != nil {
		rows, err = db.Query(query, params...)
	} else {
		rows, err = db.Query(query)
	}
	return rows, err
}

// SQLExec : Exec SQL Statement
func SQLExec(sqlStatement string, params ...interface{}) (sql.Result, error) {
	db, _ := sqlConnector()
	res, err := db.Exec(sqlStatement, params...)
	return res, err
}
