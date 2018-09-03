# Logger

#### 简介：

```
logger 是一个日志打印工具；支持linux、window彩色打印；自定义输出格式；日志文件输出等功能。
```



#### 用法：

```
log, _ := NewLog("${logpath:d:/}${time:2006-01-02 15:04:05.000} ${file} ${function} ${linenum} > [${level}] [${id}] ${message}", LEVEL_DEBUG)
log.Error("今天是个%s天气","好" )
```

#### 参数解释：

参数通过layoutString配置，包括：

```
const (
   k_time     = "time"//显示的输出时间
   k_level    = "level"//是否显示日志等级
   k_id       = "id"//是否显示id
   k_message  = "message"//内容
   k_file     = "file"//是否显示源码所在位置
   k_function = "function"//是否显示源码所在方法
   k_linenum  = "linenum"//是否显示源码所在行数
   k_logpath  = "logpath"//日志路径
   k_logname  = "logname"//日志名
)
```