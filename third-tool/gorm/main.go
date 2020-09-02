package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

/**
 * GORM ORM框架，文档地址：https://v1.gorm.io/
 */
func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mooc?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic(err)
	}
	// 日志模式
	db.LogMode(true)
}

type User struct {
	Id        int        `gorm:"column:id;primary_key" json:"id"`
	Username  string     `gorm:"column:username" json:"username"`
	Password  string     `gorm:"column:password" json:"password"`
	CreatedAt *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdatedAt *time.Time `gorm:"column:update_time" json:"update_time"`
	DeletedAt *time.Time `gorm:"column:delete_time" json:"delete_time"`
}

// TableName 设置表名
func (u *User) TableName() string {
	return "t_user"
}

func Table() {
	ret := db.HasTable(&User{})
	fmt.Println(ret)

	ret = db.HasTable("t_user")
	fmt.Println(ret)
}

func Insert() {
	user := User{Username: "dazuo", Password: "12345"}
	if err := db.Create(&user).Error; err != nil {
		log.Panic(err)
	}
	// 返回自增ID
	fmt.Printf("autoId:%d", user.Id)
}

func Query() {
	// 获取第一记录
	u := new(User)
	if err := db.First(u).Error; err != nil {
		log.Panic(err)
	}

	// 获取全部记录
	var users []User
	if err := db.Find(&users).Error; err != nil {
		log.Panic(err)
	}

	// 通过主键获取记录
	u2 := new(User)
	if err := db.First(&u2, 2).Error; err != nil {
		log.Panic(err)
	}

	// Where查询
	u3 := new(User)
	if err := db.Where("username = ?", "dazuo").First(u3).Error; err != nil {
		log.Panic(err)
	}

	u4 := new(User)
	if err := db.Where(&User{Username: "dazuo"}).First(u4).Error; err != nil {
		log.Panic(err)
	}

	u5 := new(User)
	if err := db.Find(u5, User{Username: "dazuo"}).Error; err != nil {
		log.Panic(err)
	}
	fmt.Println(u5)
}

func Update() {
	if err := db.Model(User{Id: 1}).Update("username", "dz").Error; err != nil {
		log.Panic(err)
	}
	if err := db.Model(User{Id: 1}).Updates(User{Username: "dazuo"}).Error; err != nil {
		log.Panic(err)
	}
}

func Delete() {
	// 含DeletedAt 则是软删除
	if err := db.Delete(&User{Id: 1}).Error; err != nil {
		log.Panic(err)
	}
}

func TX() {
	db.Begin()
	db.Rollback()
	db.Commit()
}

func ExecSQL() {
	if err := db.Exec("UPDATE t_user SET username = ? WHERE id = ?", "DAZUO", 1).Error; err != nil {
		log.Panic(err)
	}

	rows, err := db.Raw("SELECT * FROM t_user WHERE id > ?", 0).Rows()
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			log.Panic(err)
		}
		users = append(users, user)
	}
	fmt.Println(users)
}

func main() {
	ExecSQL()
}
