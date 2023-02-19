package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/Yazzyk/skillfulPenCMS/models"
	"github.com/Yazzyk/skillfulPenCMS/pkg/utils"
)

var Config models.Conf

func InitConf() error {
	var confName string
	switch utils.GetENV() {
	case utils.DEV:
		confName = "conf_dev.toml"
		fmt.Println("=== DEV模式 ===")
	case utils.TEST:
		confName = "conf_test.toml"
		fmt.Println("=== TEST模式 ===")
	default:
		confName = "conf.toml"
		fmt.Println("=== PROD模式 ===")
	}

	_, err := toml.DecodeFile(confName, &Config)
	if err != nil {
		return err
	}
	return nil
}
