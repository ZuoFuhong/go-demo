package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"log"
	"time"
)

/**
 * xorm是一个简单而强大的Go语言ORM库. 通过它可以使数据库操作非常简便
 *
 * 文档地址：https://www.kancloud.cn/xormplus/xorm/167077
 * Github地址：https://github.com/xormplus/xorm
 */

const (
	dataSourceName = "root:123456@tcp(127.0.0.1:3306)/mooc?charset=utf8"
)

type ScUser struct {
	Id         uint32    `xorm:"id"`
	Username   string    `xorm:"username"`
	Password   string    `xorm:"password"`
	CreateTime time.Time `xorm:"create_time"`
	UpdateTime time.Time `xorm:"update_time"`
}

var db *xorm.Engine

func init() {
	var err error
	db, err = xorm.NewMySQL("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
}

func testAutoTransaction() {
	_, _ = db.Transaction(func(session *xorm.Session) (interface{}, error) {
		user := ScUser{Username: "marszuo", Password: "123456", CreateTime: time.Now(), UpdateTime: time.Now()}
		if _, err := session.Where("id = ?", 2).Update(&user); err != nil {
			return nil, err
		}
		return nil, nil
	})
}

func testManualTransaction() {
	session := db.NewSession()
	user := ScUser{Username: "marszuo", Password: "12345", CreateTime: time.Now(), UpdateTime: time.Now()}
	if _, err := session.Where("id = ?", 2).Update(&user); err != nil {
		_ = session.Rollback()
	}
	_ = session.Commit()
}

func testGet() {
	var scUser ScUser
	_, err := db.Table("sc_user").Where("id = 1").Get(&scUser)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(scUser)
}

func testFind() {
	userList := make([]ScUser, 0)
	err := db.Table("sc_user").Where("id <= 2").Find(&userList)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(userList)
}

func main() {
	testManualTransaction()
}
