package logger

import (
	"regexp"
	"strings"
	"gopkg.in/gookit/color.v1"
	"github.com/pkg/errors"
	"go-util"
	"time"
	"runtime"
	"fmt"
)
const (
	k_time="time"
	k_level="level"
	k_id="id"
	k_message="message"
	k_file="file"
	k_function="function"
	k_linenum="linenum"
)

type Layout struct {
	layout string
	timeFormat string
	level bool
	id bool
	file bool
	function bool
	lineNum bool
}


func NewLayout(layoutString string) (*Layout,error){

	l_layout:=layoutString
	l_timeFormat:=""
	l_level:=false
	l_id:=false
	l_function:=false
	l_file:=false
	l_linenum:=false

	if !strings.Contains(layoutString,k_message) {
		return nil,errors.New("${message} is not exits")
	}
	//"${time:2006-01-02 15:04:05.000} > ${level} ${id} ${message}"
	//get${xxx}
	reg,_:=regexp.Compile(`\$\{[a-zA-Z0-9\-\:\.\ ]*\}`)
	strs:=reg.FindAllString(layoutString,-1)
	//remove ${} , split
	for _,str:=range strs  {
		str=str[2:len(str)-1]
		index:=strings.Index(str,":")
		var k,v string
		if index<=-1{
			k = str
		}else{
			k=str[0:index]
			v=str[index+1:]
		}
		switch k {
		case k_id:
			l_id=true
		case k_level:
			l_level=true
		case k_time:
			l_timeFormat=v
		case k_file:
			l_file=true
		case k_function:
			l_function=true
		case k_linenum:
			l_linenum=true
		}
	}

	return &Layout{
		layout:l_layout,
		timeFormat:l_timeFormat,
		id:l_id,
		level:l_level,
		file:l_file,
		function:l_function,
		lineNum:l_linenum,
	},nil
}
func (this *Layout) printf(currtLevel Level, format string, content ...interface{}) (int, error) {

	f:=this.layout

	pc,file,line,_ := runtime.Caller(2)
	ff := runtime.FuncForPC(pc)
	//file
	if this.file {
		f=strings.Replace(f,"${"+k_file+"}",file,-1)
	}
	//function
	if this.function {
		f=strings.Replace(f,"${"+k_function+"}",ff.Name(),-1)
	}
	//linenum
	if this.lineNum {
		f=strings.Replace(f,"${"+k_linenum+"}",fmt.Sprintf("%d",line),-1)
	}
	//message
	f=strings.Replace(f,"${"+k_message+"}",format,-1)
	//level
	if this.level {
		f=strings.Replace(f,"${"+k_level+"}",levels[currtLevel],-1)
	}
	if this.id {
		id:=go_util.NewUUID().String()[0:8]
		f=strings.Replace(f,"${"+k_id+"}",id,-1)
	}
	reg,_:=regexp.Compile(`\$\{time:[0-9a-zA-Z\:\-\.\ ]*\}`)
	f=reg.ReplaceAllString(f,time.Now().Format(this.timeFormat))

	ensureColor(currtLevel).Printf(f+"\n",content...)
	return -1,nil
}
func ensureColor(level Level) color.Color {
	switch level {
	case LEVEL_DEBUG:
		return color.FgGreen
	case LEVEL_INFO:
		return color.FgBlack
	case LEVEL_WARN:
		return color.FgYellow
	case LEVEL_ERROR:
		return color.FgRed
	case LEVEL_PANIC:
		return color.FgRed
	}
	return color.FgBlack
}
