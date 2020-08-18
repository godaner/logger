package factory

import "github.com/godaner/logger"

// Options
type Options struct {
	LogPath string        // def /log
	Level   *logger.Level // CRIT,ERRO,WARN,NOTI,INFO,DEBU
}

type Option func(o *Options)

// setter
// SetLogPath
func SetLogPath(logPath string) Option {
	return func(o *Options) {
		o.LogPath = logPath
	}
}

// SetLevel
func SetLevel(level logger.Level) Option {
	return func(o *Options) {
		o.Level = &level
	}
}
