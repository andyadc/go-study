package main

/**
Go-MySQL-Driver
https://github.com/go-sql-driver/mysql
go get github.com/go-sql-driver/mysql
username:password@protocol(address)/dbname?param=value

`
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

db, err := sql.Open("mysql", "user:password@/dbname")
`
*/
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:pwd@/idea?charset=utf8")
	checkError(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO spring_transaction SET NAME = ?")
	checkError(err)

	res, err := stmt.Exec("golang")
	checkError(err)

	id, err := res.LastInsertId()
	checkError(err)

	fmt.Println(id)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
