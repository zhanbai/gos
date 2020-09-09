package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zhanbai/gos/library/config"
)

var Db *sql.DB

func init() {
	var err error
	config := config.Get()
	Db, err = sql.Open("mysql", config.Database.Link)

	if err != nil {
		log.Fatal(err)
	}

	return
}
