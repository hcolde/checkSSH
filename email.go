package main

import (
	"flag"
	"gopkg.in/gomail.v2"
)

func main() {
	msg := ""
	flag.StringVar(&msg, "msg", "message")
	flag.Parse()

	m := gomail.NewMessage()
	m.SetHeader("From", "12345@qq.com") // 发送方邮箱
	m.SetHeader("To", "hcolde@xx.com") // 接收方邮箱
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan") // 抄送
	m.SetHeader("Subject", "Warn") // 标题
	m.SetBody("text/html", msg) // 内容

	// 发件服务器 & 端口 & 账号 & 密码(授权码)
	d := gomail.NewDialer("smtp.qq.com", 465, "12345@qq.com", "abcdefgxxx666")

	_ = d.DialAndSend(m)
}