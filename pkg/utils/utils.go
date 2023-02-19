package utils

import (
	"os"
	"strings"
)

func GetENV() ENV {
	if len(os.Args) > 1 {
		switch strings.ToLower(os.Args[1]) {
		case "dev":
			return DEV
		case "test":
			return TEST
		case "prod":
			return PROD
		default:
			return PROD
		}
	} else {
		return PROD
	}
}
