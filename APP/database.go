package app

import (
	"database/sql"
	"restful_api/helper"
	"time"
)

func NewDB() *sql.DB {

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belejar_golang_restful_api")

	helper.ErrorT(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Hour)
	db.SetConnMaxIdleTime(10 * time.Hour)
	return db
}
