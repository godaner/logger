package logger

import (
	"testing"
)

func TestDebug(t *testing.T) {
	testCases := []struct {
		format  string
		content interface{}
	}{
		{
			"a%s a",
			"6666",
		},
		{
			"a%d a",
			10,
		},
	}
	go func() {
		for {
			log, _ := NewLog("${logpath:d:/}${time:2006-01-02 15:04:05.000} ${file} ${function} ${linenum} > [${level}] [${id}] ${message}", LEVEL_DEBUG)
			log.Debug(testCases[0].format, testCases[0].content)
		}
	}()
	go func() {
		for {
			log, _ := NewLog("${logpath:d:/}${time:2006-01-02 15:04:05.000} ${file} ${function} ${linenum} > [${level}] [${id}] ${message}", LEVEL_DEBUG)
			log.Info(testCases[0].format, testCases[0].content)
		}

	}()
	go func() {
		for {
			log, _ := NewLog("${logpath:d:/}${time:2006-01-02 15:04:05.000} ${file} ${function} ${linenum} > [${level}] [${id}] ${message}", LEVEL_DEBUG)
			log.Warn(testCases[0].format, testCases[0].content)
		}
	}()
	go func() {
		for {
			log, _ := NewLog("${logpath:d:/}${time:2006-01-02 15:04:05.000} ${file} ${function} ${linenum} > [${level}] [${id}] ${message}", LEVEL_DEBUG)
			log.Error(testCases[0].format, testCases[0].content)
		}

	}()

	forever:=make(chan bool,1)
	<-forever

}
