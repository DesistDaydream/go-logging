package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func logInit() error {
	// 设置日志级别,
	level := "info"
	// 设置日志输出的文件路径
	file := ""
	// 设置日志输出的格式
	format := "text"

	// 设置日志格式
	switch format {
	case "text":
		// 文本格式
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	case "json":
		// JSON 格式
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}

	// 设置日志级别
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(l)

	if file != "" {
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		logrus.SetOutput(f)
	}

	return nil
}

func main() {
	logInit()
}
