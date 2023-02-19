package main

import (
	"github.com/Yazzyk/skillfulPenCMS/cmd"
	_ "github.com/Yazzyk/skillfulPenCMS/docs"
)

// @title 妙笔CMS
// @version 0.1
// @description 妙笔CMS API
// @contact.name Author
// @contact.email go_010@qq.com
// @host localhost:8080
// @BasePath /
func main() {
	cmd.Exec()
}
