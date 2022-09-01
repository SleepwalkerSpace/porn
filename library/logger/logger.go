package logger

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func New(logPath, logName string) (*logrus.Logger, error) {
	log := logrus.New()
	log.SetReportCaller(true)

	r, err := rotatelogs.New(
		logName+".%Y%m%d%H",
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithMaxAge(time.Hour*24*30),
	)
	if err != nil {
		return nil, err
	}

	log.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: r,
			logrus.InfoLevel:  r,
			logrus.WarnLevel:  r,
			logrus.ErrorLevel: r,
			logrus.PanicLevel: r,
			logrus.FatalLevel: r,
			logrus.TraceLevel: r,
		},
		&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			PrettyPrint:     false,
		},
	))

	return log, nil
}
