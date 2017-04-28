package logger

import (
	"github.com/greenac/octopus/config"
	"fmt"
	"time"
	"os"
)

type Logger struct {
	config config.Config
	TimeFormat string
}

func (l *Logger)Setup() {
	l.config = config.Configuration()

	if l.TimeFormat == "" {
		l.TimeFormat = time.UnixDate
	}
}

func (l *Logger)Log(a ...interface{}) {
	args := fmt.Sprint(a)
	msg := time.Now().Format(time.UnixDate) + " :: " + args[1: len(args) - 1]
	fmt.Println(msg)

	if l.config.LogInfo.Level == 1 {
		go func(message string) {
			f, err:= os.OpenFile(l.config.LogInfo.Path, os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0644)
			if err != nil {
				fmt.Println("Error: could not write to octopus log file:", l.config.LogInfo.Path)
				return
			}
			defer f.Close()

			message += "\n"
			_, err = f.WriteString(message)
			if err != nil {
				fmt.Println("Error: failed to write message to log file:", l.config.LogInfo.Path)
			}
		}(msg)
	}
}
