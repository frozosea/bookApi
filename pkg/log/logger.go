package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

type IBaseExceptionLogger interface {
	ExceptionLog(err error)
}

func New(logFileName string, logMaxAge, logRotationTime int) *logrus.Logger {
	logger := logrus.New()
	path := fmt.Sprintf("/var/log/%s.log", logFileName)
	writer, err := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(logMaxAge)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(logRotationTime)*time.Second),
	)
	if err != nil {
		fmt.Printf(`logger was nor create success error: %s`, err.Error())
		return logger
	}
	logHook := lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.ErrorLevel: writer,
		}, &logrus.TextFormatter{})
	logger.AddHook(logHook)
	return logger
}
