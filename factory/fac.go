package factory

import (
	"github.com/godaner/logger"
	logging "github.com/godaner/logger/go-logging"
	"os"
	"sync"
)

var loggerFactory Factory
var levs = map[string]logger.Level{
	"CRIT": logger.CRITICAL,
	"ERRO": logger.ERROR,
	"WARN": logger.WARNING,
	"NOTI": logger.NOTICE,
	"INFO": logger.INFO,
	"DEBU": logger.DEBUG,
}

// Init
func Init(project string, options ...Option) {
	// options
	opts := &Options{}
	for _, v := range options {
		if v == nil {
			continue
		}
		v(opts)
	}
	loggerFactory = &defFactory{
		project: project,
		logPath: opts.LogPath,
		lev:     opts.Level,
	}
}

// GetLogger
func GetLogger(tag string) (logger logger.Logger) {
	if loggerFactory == nil {
		panic("pls init first")
	}
	return loggerFactory.GetLogger(tag)
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
	f.init()
	return &logging.Logger{
		L:       *f.lev,
		Project: f.project,
		Tag:     tag,
		LogPath: f.logPath,
	}
}

func (f *defFactory) GetProject() (pj string) {
	f.init()
	return f.project
}
func (f *defFactory) GetLogPath() (logPath string) {
	f.init()
	return f.logPath
}
func (f *defFactory) GetLevel() (level *logger.Level) {
	f.init()
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
