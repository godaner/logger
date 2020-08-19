package go_logging

import (
	logger "github.com/godaner/logger"
	logging "github.com/op/go-logging"
	"os"
	"sync"
)

type Logger struct {
	sync.Once
	log     *logging.Logger
	L       logger.Level
	Project string
	Tag     string
	LogPath string
}

func (l *Logger) init() {
	l.Do(func() {
		if l.Project == "" {
			panic("pls set the project")
		}
		// log
		log, err := logging.GetLogger(l.Project)
		if err != nil {
			panic(err)
		}
		if log == nil {
			panic("nil log")
		}
		if l.Tag != "" {
			l.Tag = " " + l.Tag
		}
		var format = logging.MustStringFormatter(
			"%{color}%{time:2006-01-02 15:04:05.000} " + l.Project + l.Tag + " > %{level:.4s} %{color:reset} %{message}",
		)
		bs := make([]logging.Backend, 0)
		// std
		stdLog := logging.NewLogBackend(os.Stdout, "", 0)
		stdLogF := logging.NewBackendFormatter(stdLog, format)
		bs = append(bs, stdLogF)
		if l.LogPath != "" {
			// file
			fileLog := logging.NewLogBackend(NewLogWrite(l.LogPath+"/"+l.Project, l.Project), "", 0)
			fileLogF := logging.NewBackendFormatter(fileLog, format)
			bs = append(bs, fileLogF)
		}

		// set lev
		le := logging.MultiLogger(bs...)
		le.SetLevel(logging.Level(l.L), "")
		log.SetBackend(le)
		l.log = log
	})
}

func (l *Logger) Noticef(fms string, arg ...interface{}) {
	l.init()
	l.log.Noticef(fms, arg...)
}

func (l *Logger) Notice(arg ...interface{}) {
	l.init()
	l.log.Notice(arg...)
}

func (l *Logger) Criticalf(fms string, arg ...interface{}) {
	l.init()
	l.log.Criticalf(fms, arg...)
}

func (l *Logger) Critical(arg ...interface{}) {
	l.init()
	l.log.Critical(arg...)
}

func (l *Logger) Debugf(fms string, arg ...interface{}) {
	l.init()
	l.log.Debugf(fms, arg...)
}
func (l *Logger) Debug(arg ...interface{}) {
	l.init()
	l.log.Debug(arg...)
}

func (l *Logger) Infof(fms string, arg ...interface{}) {
	l.init()
	l.log.Infof(fms, arg...)
}

func (l *Logger) Info(arg ...interface{}) {
	l.init()
	l.log.Info(arg...)
}

func (l *Logger) Warningf(fms string, arg ...interface{}) {
	l.init()
	l.log.Warningf(fms, arg...)
}

func (l *Logger) Warning(arg ...interface{}) {
	l.init()
	l.log.Warning(arg...)
}

func (l *Logger) Errorf(fms string, arg ...interface{}) {
	l.init()
	l.log.Errorf(fms, arg...)
}

func (l *Logger) Error(arg ...interface{}) {
	l.init()
	l.log.Error(arg...)
}
