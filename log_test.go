package logger

import (
	"testing"
)

func TestDebug(t *testing.T) {
	testCases := []struct {
		format   string
		content   interface{}
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
	for _, test := range testCases {
		log,_:=NewLog("${time:2006-01-02 15:04:05.000} ${file} ${function} ${linenum} > [${level}] [${id}] ${message}",LEVEL_DEBUG)
		log.Panic(test.format,test.content)
	}
}
