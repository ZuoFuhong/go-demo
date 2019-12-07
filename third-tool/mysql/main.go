package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, e := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mooc")
	if e != nil {
		log.Fatal("Connect to mysql error：", e)
		return
	}
	stmt, e := db.Prepare("SELECT * FROM sc_user WHERE id = ?")
	if e != nil {
		log.Fatal("Prepare sql error：", e)
		return
	}
	rows, e := stmt.Query(1)

	if e != nil {
		log.Fatal("Query error：", e)
		return
	}

	for rows.Next() {
		var (
			id         int
			username   string
			password   string
			createTime string
			updateTime string
		)
		e = rows.Scan(&id, &username, &password, &createTime, &updateTime)
		if e != nil {
			panic(e)
		}
		fmt.Printf("id = %d, username = %s createTime = %s\n", id, username, createTime)
	}
	_ = db.Close()
}
