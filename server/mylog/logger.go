package mylog

import (
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"server/config"
)

var Logger *logrus.Logger

func InitMyLogger() {
	Logger = logrus.New()
	parseLevel, err := logrus.ParseLevel(config.Config.Log.Level)
	if err != nil {
		log.Error(err)
		parseLevel = logrus.DebugLevel
	}

	Logger.SetLevel(parseLevel)

	//todo 日志格式先放一下
}
