package libs

import (
	"html/template"
	"time"
)

// html 使用的函数
//时间戳转换为日期
func TimeToDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func StrToHtml(content string) template.HTML {
	return template.HTML(content)
}

func AddKey(content int) int {
	return content + 1
}

func TimeYear() string {
	timeNow := time.Now()
	return timeNow.Format("2006")
}
