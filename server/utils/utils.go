package utils

import (
	"os"
	"regexp"
	"time"
)

func PhoneReg(phone string) bool {
	reg := `^1(3|4|5|6|7|8|9)\d{9}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}

// 按天生成日志文件
func todayFilename() string {
	today := time.Now().Format("2006-01-02")
	return "./log/" + today + ".log"
}

// 创建打开文件
func NewLogFile() *os.File {
	filename := todayFilename()
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
