package logger

import (
	"github.com/Yazzyk/skillfulPenCMS/conf"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

func InitLogger() (err error) {
	lvl, err := logrus.ParseLevel(conf.Config.Logger.Level)
	if err != nil {
		return
	}

	logrus.SetLevel(lvl)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:  true,
		PadLevelText: true,
	})
	logrus.SetReportCaller(true)

	writer, err := rotatelogs.New(
		conf.Config.Logger.Path+"_%Y%m%d",
		rotatelogs.WithLinkName(conf.Config.Logger.Path),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		return
	}
	logrus.SetOutput(writer)

	return
}

func Info(args ...interface{}) {
	logrus.Info(args)
}

func Debug(args ...interface{}) {
	logrus.Debug(args)
}

func Trace(args ...interface{}) {
	logrus.Trace(args)
}

func Warn(args ...interface{}) {
	logrus.Warn(args)
}

func Error(args ...interface{}) {
	logrus.Error(args)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args)
}

func Panic(args ...interface{}) {
	logrus.Panic(args)
}
