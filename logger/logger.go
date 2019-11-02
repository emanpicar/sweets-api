package logger

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

type logger interface {
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Infoln(args ...interface{})
	Warnln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})
}

// Log implements logger interface
var Log logger
var once sync.Once

// Init Initialzes logger module once
// required to be called atleast once to avoid invalid memory address error
func Init(logLevel string) {
	once.Do(func() {
		Log = setUp(logLevel)
	})
}

func setUp(logLevel string) logger {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetLevel(getLoglevel(logLevel))
	log.SetOutput(os.Stdout)

	return log
}

func getLoglevel(logLevel string) logrus.Level {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return logrus.InfoLevel
	}

	return level
}
