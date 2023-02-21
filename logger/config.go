package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*Config log配置文件*/
type Config struct {
	defaultLogLevel string          //默认日志记录级别
	stacktraceLevel string          //记录堆栈的级别
	atomicLevel     zap.AtomicLevel //用于动态更改日志记录级别
	projectName     string          //项目名称
	callerSkip      int             //CallerSkip次数
	jsonFormat      bool            //输出json格式
	consoleOut      bool            //是否输出到console
	fileOut         *fileOut
	encoderConfig   *zapcore.EncoderConfig
}

type fileOut struct {
	enable     bool   //是否将日志输出到文件
	fileName   string //日志文件的位置；
	maxSize    int    //在进行切割之前，日志文件的最大大小（以MB为单位）；
	maxBackups int    //保留旧文件的最大个数；
	maxAge     int    //保留旧文件的最大天数；
	compress   bool   //是否压缩/归档旧文件；
}

func newConfig() *Config {
	return &Config{
		defaultLogLevel: "info",
		stacktraceLevel: "panic",
		atomicLevel:     zap.NewAtomicLevel(),
		projectName:     "",
		callerSkip:      1,
		jsonFormat:      true,
		consoleOut:      true,
		fileOut: &fileOut{
			enable:     false,
			fileName:   "",
			maxSize:    10,
			maxBackups: 3,
			maxAge:     3,
			compress:   true,
		},
	}
}

/*SetLevel 设置日志记录级别*/
func (c *Config) SetLevel(level string) {
	c.atomicLevel.SetLevel(getLevel(level))
}

/*SetStacktraceLevel 设置堆栈跟踪的日志级别*/
func (c *Config) SetStacktraceLevel(level string) {
	c.stacktraceLevel = level
}

/*SetProjectName 设置ProjectName*/
func (c *Config) SetProjectName(projectName string) {
	c.projectName = projectName
}

/*SetCallerSkip 设置callerSkip次数*/
func (c *Config) SetCallerSkip(callerSkip int) {
	c.callerSkip = callerSkip
}

/*EnableJSONFormat 开启JSON格式化输出*/
func (c *Config) EnableJSONFormat() {
	c.jsonFormat = true
}

/*DisableJSONFormat 关闭JSON格式化输出*/
func (c *Config) DisableJSONFormat() {
	c.jsonFormat = false
}

/*EnableConsoleOut 开启Console输出*/
func (c *Config) EnableConsoleOut() {
	c.consoleOut = true
}

/*DisableConsoleOut 关闭Console输出*/
func (c *Config) DisableConsoleOut() {
	c.consoleOut = false
}

/*SetEncoderConfig 设置encoderConfig*/
func (c *Config) SetEncoderConfig(encoderConfig *zapcore.EncoderConfig) {
	c.encoderConfig = encoderConfig
}

/*SetFileOut 设置日志输出文件*/
func (c *Config) SetFileOut(fileName string, maxSize, maxBackups, maxAge int, compress bool) {
	c.fileOut.enable = true
	c.fileOut.fileName = fileName
	c.fileOut.maxSize = maxSize
	c.fileOut.maxBackups = maxBackups
	c.fileOut.maxAge = maxAge
	c.fileOut.compress = compress
}
