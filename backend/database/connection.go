package database

import (
	"database/sql"
	"go-auth-token/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() *sql.DB {

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/db_go_auth")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db

}