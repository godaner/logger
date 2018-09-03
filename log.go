package logger

type Level uint16
const (
	LEVEL_DEBUG Level = iota
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_PANIC
)
var levels = []string{
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
	"PANIC",
}
type Log struct {
	lowestVisibleLevel Level  //最小可视等级
	layout       *Layout //输出格式
}


func NewLog(layoutString string, lowestVisibleLevel Level) (*Log,error) {
	if layoutString == "" {
		layoutString = "${time:2006-01-02 15:04:05.000} > ${level} ${id} ${message}"
	}
	var lay *Layout
	var e error
	if lay,e=NewLayout(layoutString);e!=nil{
		return nil,e
	}
	return &Log{
		lowestVisibleLevel: lowestVisibleLevel,
		layout:       lay,
	},nil
}

func (this *Log) Debug(format string, content ...interface{}) (n int, err error) {
	return this.printf(LEVEL_DEBUG, format, content...)
}
func (this *Log) Info(format string, content ...interface{}) (n int, err error) {
	return this.printf(LEVEL_INFO, format, content...)
}
func (this *Log) Warn(format string, content ...interface{}) (n int, err error) {
	return this.printf(LEVEL_WARN, format, content...)
}
func (this *Log) Error(format string, content ...interface{}) (n int, err error) {
	return this.printf(LEVEL_ERROR, format, content...)
}
func (this *Log) Panic(format string, content ...interface{}) (n int, err error) {
	return this.printf(LEVEL_PANIC, format, content...)
}
func (this *Log) printf(currtLevel Level, format string, content ...interface{}) (n int, err error) {
	if this.lowestVisibleLevel > currtLevel {
		return -1,nil
	}
	return this.layout.printf(currtLevel,format,content...)
}
