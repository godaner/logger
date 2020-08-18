package factory

import (
	"github.com/godaner/logger"
	logging "github.com/godaner/logger/go-logging"
	"os"
	"sync"
)

var LoggerFactory Factory
var levs = map[string]logger.Level{
	"CRIT": logger.CRITICAL,
	"ERRO": logger.ERROR,
	"WARN": logger.WARNING,
	"NOTI": logger.NOTICE,
	"INFO": logger.INFO,
	"DEBU": logger.DEBUG,
}
// InitLoggerFactory
func InitLoggerFactory(project string, options ...logger.Option) {
	// options
	opts := &logger.Options{}
	for _, v := range options {
		if v == nil {
			continue
		}
		v(opts)
	}
	LoggerFactory = &defFactory{
		project: project,
		logPath: opts.LogPath,
		lev:     opts.Level,
	}
}

type Factory interface {
	GetLogger(tag string) (logger logger.Logger)
	GetProject() (pj string)
	GetLogPath() (logPath string)
	GetLevel() (level *logger.Level)
}

type defFactory struct {
	project string
	logPath string
	lev     *logger.Level
	sync.Once
}

func (f *defFactory) init() {
	f.Do(func() {
		if f.lev == nil {
			f.lev = getEnvLev()
		}
		if f.logPath == "" {
			f.logPath = getEnvLogPath()
		}
	})
}
func (f *defFactory) GetLogger(tag string) (logger logger.Logger) {
	return &logging.Logger{
		L:       *f.lev,
		Project: f.project,
		Tag:     tag,
		LogPath: f.logPath,
	}
}

func (f *defFactory) GetProject() (pj string) {
	return f.project
}
func (f *defFactory) GetLogPath() (logPath string) {
	return f.logPath
}
func (f *defFactory) GetLevel() (level *logger.Level) {
	return f.lev
}

// getEnvLogPath
func getEnvLogPath() string {
	lp := os.Getenv("LOG_PATH")
	if lp != "" {
		return lp
	}
	return ""
}

// getEnvLev
func getEnvLev() *logger.Level {
	ls := os.Getenv("LOG_LEV")
	l, ok := levs[ls]
	if !ok {
		l = logger.CRITICAL
	}
	return &l
}
