package logger

import (
	"regexp"
	"strings"
	"gopkg.in/gookit/color.v1"
	"github.com/pkg/errors"
	"go-util"
	"time"
)
const (
	k_time="time"
	k_level="level"
	k_id="id"
	k_message="message"
)

type Layout struct {
	layout string
	timeFormat string
	level bool
	id bool
}


func NewLayout(layoutString string) (*Layout,error){

	l_layout:=layoutString
	l_timeFormat:=""
	l_level:=false
	l_id:=false

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
		}
	}

	return &Layout{
		layout:l_layout,
		timeFormat:l_timeFormat,
		id:l_id,
		level:l_level,
	},nil
}
func (this *Layout) printf(currtLevel Level, format string, content ...interface{}) (int, error) {

	f:=this.layout
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
