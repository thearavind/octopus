package logger

import (
	"github.com/kapitol-app/gologger"
	"github.com/kapitol-app/octopus/config"
)

var goLogger *gologger.GoLogger

//Log - Global logger function
func Log(a ...interface{}) {
	if goLogger == nil {
		goLogger = &gologger.GoLogger{LogLevel: config.C.LogInfo.Level, LogPath: config.C.LogInfo.Path}
		(*goLogger).Setup()
	}

	goLogger.Log(a)
}
