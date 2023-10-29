package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func Init() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)

	Debug("Initial logger")
}

func Info(args ...interface{}) {
	log.Infoln(args)
}

func Warning(args ...interface{}) {
	log.Warningln(args)
}

func Debug(args ...interface{}) {
	log.Debugln(args)
}

func Error(args ...interface{}) {
	log.Errorln(args)
}
