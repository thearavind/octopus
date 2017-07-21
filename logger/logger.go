package logger

import (
	"github.com/kapitol-app/gologger"
	"github.com/kapitol-app/octopus/config"
)

var goLogger *gologger.GoLogger

func setup() {
	if goLogger == nil {
		goLogger = &gologger.GoLogger{LogLevel: config.C.LogInfo.Level, LogPath: config.C.LogInfo.Path}
		(*goLogger).Setup()
	}
}

func Log(a ...interface{}) {
	setup()
	goLogger.Log(a)
}

func Error(a ...interface{}) {
	setup()
	goLogger.Error(a)
}

func Warn(a ...interface{}) {
	setup()
	goLogger.Warn(a)
}
