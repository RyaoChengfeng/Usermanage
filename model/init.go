package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var cli redis.Conn

func init() {
	database, err := sqlx.Open("mysql", "root:liaocfe@tcp(127.0.0.1:3306)/usermanage")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	db = database
	fmt.Println("mysql conn success")
	//defer db.Close()

	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	cli = c
	fmt.Println("redis conn success")
	//defer cli.Close()
}
