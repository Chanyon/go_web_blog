package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	filepath := "log/log"
	src, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0755)

	if err != nil {
		fmt.Println("err:", err)
	}

	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)

	logWriter, _ := retalog.New(
		filepath+"%Y%m%d.log",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)

	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		stopTime := time.Since(startTime).Microseconds()
		spendTime := fmt.Sprintf("%d ms", stopTime)
		clientIp := ctx.ClientIP()
		userAgent := ctx.Request.UserAgent()
		dataSize := ctx.Writer.Size()

		if dataSize < 0 {
			dataSize = 0
		}

		method := ctx.Request.Method
		path := ctx.Request.RequestURI
		hostName, err := os.Hostname()
		statusCode := ctx.Writer.Status()
		if err != nil {
			hostName = "unknown"
		}
		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"Status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})

		if len(ctx.Errors) > 0 {
			entry.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
