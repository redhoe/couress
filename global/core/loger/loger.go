package loger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/redhoe/couress/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"time"
)

var errorLoggerMap = make(map[LogTagName]*zap.Logger)

var errorSugaredLogger *zap.SugaredLogger
var errorSugaredLoggerMap = make(map[LogTagName]*zap.SugaredLogger)

type LogConf struct {
	Path string        `json:"path" yaml:"path"`
	Logs []logConfTags `json:"logs" yaml:"logs"`
}

type logConfTags struct {
	Tag         LogTagName `json:"tag" yaml:"tag"`
	FileName    string     `json:"filename" yaml:"filename"`
	ErrFileName string     `json:"errFileName" yaml:"errFileName"`
}

type LogTagName string

var logConf LogConf

// 日志类型定义
const (
	Task       LogTagName = "task"
	App        LogTagName = "app"
	Cli        LogTagName = "cli"
	DefaultLog LogTagName = "default"
)

func LogerInit() {
	// 初始化 日志服务Tags
	logConf = LogConf{
		Path: "./logs/",
		Logs: []logConfTags{
			{Task, "task.log", "taskErr.log"},
			{App, "app.log", "appErr.log"},
			{Cli, "cli.log", "cliErr.log"},
			{DefaultLog, "DefaultLog.log", "DefaultLogErr.log"},
		},
	}
	for _, logc := range logConf.Logs {
		encoder := getEncoder()
		infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= global.GbCONFIG.Zap.TransportLevel()
		})
		errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})
		// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
		infoWriter := getWriter(logConf.Path + logc.FileName)
		errorWriter := getWriter(logConf.Path + logc.ErrFileName)
		cores := make([]zapcore.Core, 0)
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel))
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel))
		if global.GbCONFIG.Zap.LogInConsole {
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), infoLevel))
		}
		// 最后创建具体的Logger
		core := zapcore.NewTee(cores...)
		// 开启开发模式，堆栈跟踪
		caller := zap.AddCaller()
		development := zap.Development()
		filed := zap.Fields(zap.String("Tag", string(logc.Tag)))
		stackTrace := zap.AddStacktrace(zap.ErrorLevel) // 当错误等级error时 触发堆栈跟踪
		log := zap.New(core, caller, stackTrace, development, filed)
		// 是否显示行
		if global.GbCONFIG.Zap.ShowLine {
			log = log.WithOptions(zap.AddCaller())
		}
		errorLoggerMap[logc.Tag] = log                // 日志糖
		errorSugaredLoggerMap[logc.Tag] = log.Sugar() // 日志糖
	}
	errorSugaredLogger = errorSugaredLoggerMap[DefaultLog]
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs 的Logger 实际生成的文件名 demo.log.YYmmddHH
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".log", "", -1)+"-%Y%m%d.log",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*2),    // 保存2天日志
		rotatelogs.WithRotationTime(time.Hour*2), // 分割日志周期：1小时
	)
	if err != nil {
		panic(err)
	}
	return hook
}

func getEncoder() zapcore.Encoder {
	if global.GbCONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.GbCONFIG.Zap.StacktraceKey, // "error"
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    global.GbCONFIG.Zap.ZapEncodeLevel(),
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.GbCONFIG.Zap.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

// SetSugaredLoggerCli 设置日志糖服务名
func SetSugaredLoggerCli(tag LogTagName) {
	errorSugaredLogger = errorSugaredLoggerMap[tag]
}

func NewSugaredLogger(tag LogTagName) *zap.SugaredLogger {
	return errorSugaredLoggerMap[tag]
}

func NewLogger(tag LogTagName) *zap.Logger {
	return errorLoggerMap[tag]
}

func Debug(args ...interface{}) {
	errorSugaredLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	errorSugaredLogger.Debugf(template, args...)
}
func Info(args ...interface{}) {
	errorSugaredLogger.Info(args...)
}
func Infof(template string, args ...interface{}) {
	errorSugaredLogger.Infof(template, args...)
}
func Warn(args ...interface{}) {
	errorSugaredLogger.Warn(args...)
}
func Warnf(template string, args ...interface{}) {
	errorSugaredLogger.Warnf(template, args...)
}
func Error(args ...interface{}) {
	errorSugaredLogger.Error(args...)
}
func Errorf(template string, args ...interface{}) {
	errorSugaredLogger.Errorf(template, args...)
}
func DPanic(args ...interface{}) {
	errorSugaredLogger.DPanic(args...)
}
func DPanicf(template string, args ...interface{}) {
	errorSugaredLogger.DPanicf(template, args...)
}
func Panic(args ...interface{}) {
	errorSugaredLogger.Panic(args...)
}
func Panicf(template string, args ...interface{}) {
	errorSugaredLogger.Panicf(template, args...)
}
func Fatal(args ...interface{}) {
	errorSugaredLogger.Fatal(args...)
}
func Fatalf(template string, args ...interface{}) {
	errorSugaredLogger.Fatalf(template, args...)
}
