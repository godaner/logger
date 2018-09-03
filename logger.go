package logger

//定义接口
type Logger interface {
	Debug(string,...interface{})
	Info(string,...interface{})
	Warn(string,...interface{})
	Error(string,...interface{})
	Panic(string,...interface{})
}
