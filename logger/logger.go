package logger

import (
	"github.com/greenac/octopus/config"
	"fmt"
	"time"
	"os"
)

type logger struct {
	config config.Config
	timeFormat string
	isSetup bool
}

func (l *logger)setup() {
	l.config = config.Configuration()

	if l.timeFormat == "" {
		l.timeFormat = time.UnixDate
	}

	l.isSetup = true
}

func (l *logger)log(a ...interface{}) {
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

var l logger;
func Log(a ...interface{}) {
	if !l.isSetup {
		l.setup()
	}

	l.log(a)
}
