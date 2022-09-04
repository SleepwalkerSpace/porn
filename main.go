package main

import (
	_ "porn/common"
	"porn/common/email"
	"time"
)

func main() {
	email.Email.Send("TEST", "TEST", "", "410520827@qq.com")
	time.Sleep(time.Hour)
}
