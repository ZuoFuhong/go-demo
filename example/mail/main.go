package main

import "gopkg.in/gomail.v2"

// 发送邮件
func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "863329112@qq.com")
	m.SetHeader("To", "2601481148@qq.com")
	m.SetHeader("Subject", "测试邮件")
	m.SetBody("text/plain", "hello world")
	m.Attach("/Users/dazuo/workplace/go_learning_notes/README.md")

	d := gomail.NewDialer("smtp.qq.com", 465, "863329112@qq.com", "xxx")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
