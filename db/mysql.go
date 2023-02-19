package db

import (
	"fmt"
	"github.com/Yazzyk/skillfulPenCMS/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Mysql *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Config.Database.User, conf.Config.Database.Password, conf.Config.Database.Host, conf.Config.Database.Port, conf.Config.Database.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}

	Mysql = db

	AutoMigrate()

	return nil
}
