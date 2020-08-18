package factory

import (
	"github.com/godaner/logger"
	logging "github.com/godaner/logger/go-logging"
	"os"
	"sync"
)

var levs = map[string]logger.Level{
	"CRIT": logger.CRITICAL,
	"ERRO": logger.ERROR,
	"WARN": logger.WARNING,
	"NOTI": logger.NOTICE,
	"INFO": logger.INFO,
	"DEBU": logger.DEBUG,
}

type Factory struct {
	Project string
	LogPath string
	Lev     *logger.Level
	sync.Once
}

func (f *Factory) init() {
	f.Do(func() {
		if f.Lev == nil {
			f.Lev = getEnvLev()
		}
		if f.LogPath == "" {
			f.LogPath = getEnvLogPath()
		}
	})
}
func (f *Factory) GetLogger(tag string) (logger logger.Logger) {
	return &logging.Logger{
		L:       *f.Lev,
		Project: f.Project,
		Tag:     tag,
		LogPath: f.LogPath,
	}
}

func (f *Factory) GetProject() (pj string) {
	return f.Project
}
func (f *Factory) GetLogPath() (logPath string) {
	return f.LogPath
}
func (f *Factory) GetLevel() (level *logger.Level) {
	return f.Lev
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
