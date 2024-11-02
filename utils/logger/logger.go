package logger

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	debugLogger *logrus.Logger
	infoLogger  *logrus.Logger
	warnLogger  *logrus.Logger
	errorLogger *logrus.Logger
)

func InitLogger() {
	debugLogger = newLogger("logs/debug", logrus.DebugLevel)
	infoLogger = newLogger("logs/info", logrus.InfoLevel)
	warnLogger = newLogger("logs/warn", logrus.WarnLevel)
	errorLogger = newLogger("logs/error", logrus.ErrorLevel)
}
func GinErrorLogger() *logrus.Logger {
	return errorLogger
}

func GinInfoLogger() *logrus.Logger {
	return infoLogger
}

func newLogger(path string, level logrus.Level) *logrus.Logger {
	logger := logrus.New()

	// 创建日志目录
	if err := os.MkdirAll(path, 0755); err != nil {
		fmt.Printf("could not create log directory: %v\n", err)
	}

	// 设置输出文件
	logger.SetOutput(&lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/log_%s.log", path, time.Now().Format("2006-01-02")),
		MaxSize:    100, // 100 MB
		MaxBackups: 7,
		MaxAge:     30, // 保留30天
		Compress:   true,
	})

	// 使用自定义格式化器
	logger.SetFormatter(&SimpleFormatter{})

	logger.SetLevel(level)

	return logger
}

type SimpleFormatter struct{}

func (f *SimpleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//Custom log format: Time log level content
	log := fmt.Sprintf("%s [%s] %s\n", entry.Time.Format("2006-01-02 15:04:05"), entry.Level.String(), entry.Message)

	//Check for additional fields, such as stack information
	if stack, ok := entry.Data["stack"]; ok {
		log += fmt.Sprintf("%s\n", stack)
	}
	return []byte(log), nil
}

func Debug(args ...interface{}) {
	debugLogger.Debug(fmt.Sprint(args...))
}

func Info(args ...interface{}) {
	infoLogger.Info(fmt.Sprint(args...))
}

func Warn(args ...interface{}) {
	warnLogger.Warn(fmt.Sprint(args...))
}

func Error(args ...interface{}) {
	//Record error stack information
	errMsg := fmt.Sprint(args...)
	errorLogger.WithField("stack", string(debug.Stack())).Error(errMsg)
}
