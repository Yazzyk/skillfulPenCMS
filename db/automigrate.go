package db

import (
	"github.com/Yazzyk/skillfulPenCMS/models"
	"github.com/Yazzyk/skillfulPenCMS/pkg/logger"
)

func AutoMigrate() {
	if err := Mysql.AutoMigrate(&models.Account{}); err != nil {
		logger.Error(err)
		return
	}
}
