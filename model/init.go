package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/usermanage")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	db = database
	defer db.Close()
}
