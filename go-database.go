package main

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
