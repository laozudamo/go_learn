package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func main() {
	initData()
	queryData()
}

func initDB() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:root123456@/gp_learn")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\"成功\": %v\n", "成功")
	}
	// See "Important settings" section.
	SqlDB.SetConnMaxLifetime(time.Minute * 3)
	SqlDB.SetMaxOpenConns(10)
	SqlDB.SetMaxIdleConns(10)
}

func init() {
	initDB()
}

func initData() {
	s := "insert into login_info(user,pwd) values(?, ?)"
	r, err := SqlDB.Exec(s, "jackkk", "1212121")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("r: %v\n", r)
}

type User struct {
	user string
	pwd  string
	id   int
}

func queryData() {
	var u User
	s := "select * from login_info where id = ?"
	err := SqlDB.QueryRow(s, 1).Scan(&u.pwd, &u.user, &u.id)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("u: %V\n", u.id, u.pwd, u.user)
	}
}

// func queryManyRow() {

// }

/*
	增删改查
*/
