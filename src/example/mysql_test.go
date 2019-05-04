// go get github.com/go-sql-driver/mysql
package example

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func TestMysql (t *testing.T) {
	db, e := sql.Open("mysql", "dazuo:123456@tcp(47.98.199.80:3306)/mooc")
	if e != nil {
		log.Fatal("Connect to mysql error：", e)
		return
	}
	stmt, e := db.Prepare("SELECT * FROM `user` WHERE id = ?")
	if e != nil  {
		log.Fatal("Prepare sql error：", e)
		return
	}
	rows, e := stmt.Query(1)

	if e != nil {
		log.Fatal("Query error：", e)
		return
	}

	for rows.Next() {
		strings, _ := rows.Columns()
		t.Log("strings: ", strings)

		var (
			id int
			username string
			password string
			email 	 string
			phone	 string
			create_time string
			update_time string
		)
		e = rows.Scan(&id, &username, &password, &email, &phone, &create_time, &update_time)
		if e != nil {
			log.Fatal("row scan error: ", e)
			return
		}
		fmt.Printf("id = %d, username = %s\n", id, username)
	}
	t.Log("query successful")
	_ = db.Close()
}