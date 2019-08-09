package main

import (
	"fmt"
	"encoding/xml"
)

var x string = `
<?xml version="1.0" encoding="UTF-8"?>
<config>
<smtpServer>smtp.163.com</smtpServer>
<smtpPort>25</smtpPort>
<sender>user@163.com</sender>
<senderPasswd>123456</senderPasswd>
<receivers flag="true">
<user>Mike_Zhang@live.com</user>
<user>test1@qq.com</user>
</receivers>
</config>`

type Child struct {
	XMLName xml.Name `xml:"receivers"`
	User []string `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"config"`
	SmtpServer string `xml:"smtpServer"`
	SmtpPort string `xml:"smtpPort"`
	Receivers Child
}

func test() {
	var res User
	err := xml.Unmarshal([]byte(x), &res)
	if err != nil {
		panic("line on : 31")
	}
	fmt.Println(res)
	fmt.Println(res.Receivers.XMLName.Local)
}
 
func main() {
	test()
}