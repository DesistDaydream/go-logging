package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		// 这里可以修改 logrus 中默认的字段名称
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "customTimeName",
			logrus.FieldKeyLevel: "customLevelName",
			logrus.FieldKeyMsg:   "customMessageName",
			logrus.FieldKeyFunc:  "customCallerName",
		},
		// CallerPrettyfier: func(*runtime.Frame) (string, string) {
		// },
		PrettyPrint: false,
	})
	logrus.Infof("Hello World")
}
