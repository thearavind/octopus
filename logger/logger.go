package logger

import (
	"github.com/kapitol-app/gologger"
	"github.com/kapitol-app/octopus/config"
)


var goLogger *gologger.GoLogger

func Log(a ...interface{}) {
	if goLogger == nil {
		c := config.Configuration()
		goLogger = &gologger.GoLogger{LogLevel: c.LogInfo.Level, LogPath: c.LogInfo.Path}
		(*goLogger).Setup()
	}

	goLogger.Log(a)
}
