package cmd

import (
	"github.com/Yazzyk/skillfulPenCMS/api"
	"github.com/Yazzyk/skillfulPenCMS/conf"
	"github.com/Yazzyk/skillfulPenCMS/db"
	"github.com/Yazzyk/skillfulPenCMS/pkg/logger"
)

func Exec() {
	// 配置文件
	if err := conf.InitConf(); err != nil {
		panic("读取配置文件失败:" + err.Error())
	}

	// 日志
	if err := logger.InitLogger(); err != nil {
		panic("初始化logger失败:" + err.Error())
	}

	// 数据库
	if err := db.InitDB(); err != nil {
		panic("连接数据库失败:" + err.Error())
	}

	// api 路由
	api.InitAPI()
}
