package logger

import (
	"fmt"
	"github.com/godaner/go-util"
	"github.com/pkg/errors"
	"gopkg.in/gookit/color.v1"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
)

const (
	LOG_FILE_TIME_FORMAT = "2006-01-02"
)
const (
	k_time     = "time"
	k_level    = "level"
	k_id       = "id"
	k_message  = "message"
	k_file     = "file"
	k_function = "function"
	k_linenum  = "linenum"
	k_logpath  = "logpath"
	k_logname  = "logname"
)

type Layout struct {
	layout     string
	timeFormat string
	level      bool
	id         bool
	file       bool
	function   bool
	lineNum    bool
	logPath    string
	logName    string
}

func NewLayout(layoutString string) (*Layout, error) {

	l_layout := layoutString
	l_logpath := ""
	l_logname := ""
	l_timeFormat := ""
	l_level := false
	l_id := false
	l_function := false
	l_file := false
	l_linenum := false

	if !strings.Contains(layoutString, k_message) {
		return nil, errors.New("${message} is not exits")
	}
	//"${time:2006-01-02 15:04:05.000} > ${level} ${id} ${message}"
	//get${xxx}
	reg, _ := regexp.Compile(`\$\{[a-zA-Z0-9\-\:\.\ \/\\]*\}`)
	strs := reg.FindAllString(layoutString, -1)
	//remove ${} , split
	for _, str := range strs {
		str = str[2 : len(str)-1]
		index := strings.Index(str, ":")
		var k, v string
		if index <= -1 {
			k = str
		} else {
			k = str[0:index]
			v = str[index+1:]
		}
		switch k {
		case k_id:
			l_id = true
		case k_level:
			l_level = true
		case k_file:
			l_file = true
		case k_function:
			l_function = true
		case k_linenum:
			l_linenum = true
		case k_time:
			l_timeFormat = v
		case k_logpath:
			l_logpath = v
		case k_logname:
			l_logname = v
		}
	}

	return &Layout{
		layout:     l_layout,
		logPath:    l_logpath,
		logName:    l_logname,
		timeFormat: l_timeFormat,
		id:         l_id,
		level:      l_level,
		file:       l_file,
		function:   l_function,
		lineNum:    l_linenum,
	}, nil
}
func (this *Layout) printf(currtLevel Level, format string, content ...interface{}) (int, error) {

	f := this.layout

	pc, file, line, _ := runtime.Caller(2)
	ff := runtime.FuncForPC(pc)
	//file
	if this.file {
		f = strings.Replace(f, "${"+k_file+"}", file, -1)
	}
	//function
	if this.function {
		f = strings.Replace(f, "${"+k_function+"}", ff.Name(), -1)
	}
	//linenum
	if this.lineNum {
		f = strings.Replace(f, "${"+k_linenum+"}", fmt.Sprintf("%d", line), -1)
	}
	//message
	f = strings.Replace(f, "${"+k_message+"}", format, -1)
	//level
	if this.level {
		f = strings.Replace(f, "${"+k_level+"}", levels[currtLevel], -1)
	}
	//id
	if this.id {
		id := go_util.NewUUID().String()[0:8]
		f = strings.Replace(f, "${"+k_id+"}", id, -1)
	}
	//time
	reg, _ := regexp.Compile(`\$\{` + k_time + `:[0-9a-zA-Z\:\-\.\ ]*\}`)
	f = reg.ReplaceAllString(f, time.Now().Format(this.timeFormat))

	//remove logpath
	reg, _ = regexp.Compile(`\$\{` + k_logpath + `:[0-9a-zA-Z\:\-\.\_\ \/\\]*\}`)
	f = reg.ReplaceAllString(f, "")

	//console
	ensureColor(currtLevel).Printf(f+"\r\n", content...)
	//log path
	if this.logPath == "" {
		return -1, nil
	}
	if this.logName == ""{
		this.logName = "unKnow"
	}


	date := time.Now().Format(LOG_FILE_TIME_FORMAT)
	todayLogFilePathName := this.logPath + "/" + this.logName + "-" + date + ".log"
	totalLogFilePathName := this.logPath + "/" + this.logName + ".log"
	e := appendLogFile(todayLogFilePathName, ensureColor(currtLevel).Sprintf(f+"\r\n", content...))
	if e != nil {
		return -1, e
	}
	e = appendLogFile(totalLogFilePathName, ensureColor(currtLevel).Sprintf(f+"\r\n", content...))

	if e != nil {
		return -1, e
	}
	return -1, nil
}

func appendLogFile(logFilePathName string, appendContent string) error {
	logFile, err := os.OpenFile(logFilePathName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return errors.New("write log err")
	}
	fd_content := strings.Join([]string{appendContent}, "")
	buf := []byte(fd_content)
	logFile.Write(buf)
	logFile.Close()
	return nil
}
func ensureColor(level Level) color.Color {
	switch level {
	case LEVEL_DEBUG:
		return color.FgGreen
	case LEVEL_INFO:
		return color.FgWhite
	case LEVEL_WARN:
		return color.FgYellow
	case LEVEL_ERROR:
		return color.FgRed
	case LEVEL_PANIC:
		return color.FgRed
	}
	return color.FgBlack
}
